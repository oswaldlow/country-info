# Country Lookup API

A lightweight and fast country info API built with **Golang**.  
You can query by **ISO code**, **English name**, or **Chinese name** to get detailed information including ISO codes, numeric codes, independence status, and names in multiple languages.

The service is **precompiled to run on port `8080`**.  
To change the port, please **compile from source**.

**Data Source**: All the country information is sourced from [Wikipedia - ISO 3166-1 Alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2).

**中文翻译**: 本文中的中文翻译由 ChatGPT 提供。

---

## Features

- 🔁 Convert between ISO code, English name, and Chinese name
- 🌐 Supports ISO 3166-1 alpha-2 / alpha-3 / numeric codes
- 🏳️ Provides independence status and full ISO identifier
- 🈶 Built-in Simplified Chinese support
- ⚡ Fast and lightweight, built with Go

---

## 📘 Example Usage

### 🔍 ISO Code Query Example

**Request:**

### 🌏 ISO 3166-1 alpha-2 codes Query Example

```http
GET https://country-api.demonsword.top/country/jp
```

**Response:**

```json
[
  {
    "name": "Japan",
    "alpha2": "JP",
    "alpha3": "JPN",
    "numeric": "392",
    "iso_code": "ISO 3166-2:JP",
    "independent": true,
    "chinese_name": "日本"
  }
]
```

---

### 🌏 Chinese Name Query Example

**Request:**

```http
GET https://country-api.demonsword.top/country/澳门
```

**Response:**

```json
[
  {
    "name": "Macao",
    "alpha2": "MO",
    "alpha3": "MAC",
    "numeric": "446",
    "iso_code": "ISO 3166-2:MO",
    "independent": false,
    "chinese_name": "澳门"
  }
]
```

---

### 🇬🇧 English Name Query Example

**Request:**

```http
GET https://country-api.demonsword.top/country/Poland
```

**Response:**

```json
[
  {
    "name": "Poland",
    "alpha2": "PL",
    "alpha3": "POL",
    "numeric": "616",
    "iso_code": "ISO 3166-2:PL",
    "independent": true,
    "chinese_name": "波兰"
  }
]
```

---

## API Endpoints

| Method | Endpoint                   | Description                                        |
|--------|----------------------------|----------------------------------------------------|
| GET    | `/country/{query}`         | Query by ISO code, English name, or Chinese name   |

> Example queries: `jp`, `Japan`, `日本`

---

## Build & Run

The project is written in **Go** and precompiled to run on **port `8080`**.

### 🔧 Run the precompiled binary

```bash
./country-lookup
```

> Server starts on: `http://localhost:8080`

---

### 🛠️ Dev mode

If you're developing locally:

```bash
go run main.go
```

---

## License

MIT License
