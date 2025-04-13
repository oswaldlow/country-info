package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type Country struct {
	Name        string `json:"name"`
	Alpha2      string `json:"alpha2"`
	Alpha3      string `json:"alpha3"`
	Numeric     string `json:"numeric"`
	IsoCode     string `json:"iso_code"`
	Independent bool   `json:"independent"`
	ChineseName string `json:"chinese_name"`
}

// Whitelist country/region code
var whiteList = map[string]string{
	"中华人民共和国":   "CN",
	"台湾":        "TW",
	"台湾省":       "TW",
	"中国台湾":      "TW",
	"中国香港":      "HK",
	"香港特别行政区":   "HK",
	"中国香港特别行政区": "HK",
	"澳门特别行政区":   "MO",
	"中国澳门特别行政区": "MO",
	"中国，香港":     "HK",
	"中国,香港":     "HK",
	"中国,澳门":     "MO",
	"中国，澳门":     "MO",
	"中华民国":      "TW",
	"中国,台湾":     "TW",
	"中国，台湾":     "TW",
}

// Read and parse json
func loadCountries() ([]Country, error) {
	data, err := ioutil.ReadFile("countries.json")
	if err != nil {
		return nil, err
	}

	var countries []Country
	if err := json.Unmarshal(data, &countries); err != nil {
		return nil, err
	}

	return countries, nil
}

// Whitelist module
func handleWhiteList(searchTerm string) string {
	if replacement, found := whiteList[searchTerm]; found {
		return replacement
	}
	return searchTerm
}

// Search country module
func searchCountry(countries []Country, searchTerm string) []Country {
	var result []Country

	searchTerm = handleWhiteList(searchTerm)

	searchTermLower := strings.ToLower(searchTerm)

	for _, country := range countries {
		if strings.ToLower(country.Name) == searchTermLower ||
			strings.ToLower(country.Alpha2) == searchTermLower ||
			strings.ToLower(country.Alpha3) == searchTermLower ||
			strings.ToLower(country.IsoCode) == searchTermLower ||
			strings.ToLower(country.ChineseName) == searchTermLower {
			result = append(result, country)
		}
	}
	return result
}

// Handle search request
func countrySearchHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// 添加 CORS 支持
	w.Header().Set("Access-Control-Allow-Origin", "*")                   
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS") 
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")       

	if r.Method == http.MethodOptions {
		return
	}

	// Get user query info
	searchTerm := r.URL.Path[len("/country/"):]

	countries, err := loadCountries()
	if err != nil {
		http.Error(w, fmt.Sprintf("Unable to load countries data: %v", err), http.StatusInternalServerError)
		return
	}

	// Search for the result
	results := searchCountry(countries, searchTerm)

	// If no result, then return 404
	if len(results) == 0 {
		http.Error(w, "No country found", http.StatusNotFound)
		return
	}

	// Respone result
	if err := json.NewEncoder(w).Encode(results); err != nil {
		http.Error(w, fmt.Sprintf("Unable to encode data: %v", err), http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/country/", countrySearchHandler)

	// 启动 HTTP 服务器
	port := ":8080"
	log.Fatal(http.ListenAndServe(port, nil))
}
