# City REST API (Go + Gin)

A clean and lightweight REST API built with **Go (Golang)** and the **Gin web framework**.

This project provides structured access to city data stored in a JSON file and demonstrates how to design a scalable backend using layered architecture principles.

---

#  Features

* Retrieve all cities
* Get a city by name
* Filter cities by country
* Clean layered architecture (handlers, services, models, routes)
* Fast and lightweight HTTP API using Gin
* JSON-based data source

---

#  Tech Stack

* Go (Golang)
* Gin Web Framework
* JSON (data storage format)

---

#  Project Structure

```bash id="structure1"
.
├── main.go
├── data/
│   └── cities.json
├── models/
├── services/
├── handlers/
└── routes/
```

---

#  Architecture Overview

```text id="arch1"
Client Request
      ↓
Gin Router
      ↓
Route Layer
      ↓
Handler Layer (HTTP Controllers)
      ↓
Service Layer (Business Logic)
      ↓
JSON Data Source
```

This layered design ensures:

* Separation of concerns
* Easy maintainability
* Scalable project structure

---

#  API Features

## Cities

### Get all cities

```http id="api1"
GET /cities
```

### Get city by name

```http id="api2"
GET /cities/:name
```

### Filter by country

```http id="api3"
GET /cities?country=Netherlands
```

---

#  Getting Started

## 1. Clone the repository

```bash id="setup1"
git clone https://github.com/your-username/city-rest-api.git
cd city-rest-api
```

---

## 2. Install dependencies

```bash id="setup2"
go mod tidy
```

---

## 3. Run the server

```bash id="run1"
go run main.go
```

Server runs at:

```text id="run2"
http://localhost:8080
```

---

#  Future Improvements

* Add pagination and sorting
* Integrate PostgreSQL database
* Add environment configuration (.env support)
* Implement structured logging and middleware
* Add unit and integration tests
* Dockerize the application
* Add Swagger/OpenAPI documentation

---

#  Learning Context

This project is part of my Go backend learning journey.

It started as a simple API focused on:

* JSON parsing
* Basic Gin routing
* HTTP request handling

and evolved into a structured backend using layered architecture principles.

---

#  What This Project Demonstrates

* REST API development in Go
* Gin framework usage
* Clean architecture principles
* Layered backend design
* Basic data modeling and filtering logic
* Scalable API structure design
