# City REST API

![Go](https://img.shields.io/badge/Go-1.x-00ADD8?logo=go)
![Gin](https://img.shields.io/badge/Gin-Web_Framework-008ECF)
![REST API](https://img.shields.io/badge/API-REST-green)
![JSON](https://img.shields.io/badge/Data-JSON-lightgrey)
![License](https://img.shields.io/badge/License-MIT-green)

A lightweight REST API built with Go and the Gin web framework.

The service provides structured access to city data through a clean backend architecture focused on maintainability, separation of concerns, and scalable API design.

---

# Overview

City REST API demonstrates how to build a backend service using Go with a layered architecture approach.

The application workflow:

```text
Client Request

      ↓

Gin Router

      ↓

HTTP Handlers

      ↓

Service Layer

      ↓

Data Repository

      ↓

JSON Storage
```

The project focuses on backend fundamentals:

- REST API design
- request handling
- service abstraction
- data modeling
- clean project organization

---

# Features

## City Retrieval

Provides endpoints for:

- retrieving all cities
- searching cities by name
- filtering cities by country

---

## RESTful API Design

The API follows standard HTTP principles:

- resource-based routes
- structured responses
- clear endpoint responsibilities

---

## Layered Architecture

The application separates:

```text
Routes

↓

Handlers

↓

Services

↓

Data Layer
```

Benefits:

- easier maintenance
- improved testability
- clearer responsibilities
- future database migration support

---

# Architecture

## Router Layer

Responsible for:

- defining API routes
- mapping requests to handlers

Example:

```text
GET /cities

GET /cities/:name

GET /cities?country=value
```

---

## Handler Layer

Responsible for:

- processing HTTP requests
- validating parameters
- returning API responses

---

## Service Layer

Contains application logic:

- filtering cities
- searching records
- transforming data

Keeping logic outside handlers makes the API easier to extend.

---

## Data Layer

Currently uses JSON storage.

Responsibilities:

- loading city records
- providing structured data access

The design allows future replacement with:

- PostgreSQL
- MongoDB
- external APIs

---

# Project Structure

```text
city-rest-api/

├── main.go
│
├── data/
│   └── cities.json
│
├── models/
│
├── services/
│
├── handlers/
│
├── routes/
│
└── README.md
```

---

# API Documentation

## Get All Cities

```http
GET /cities
```

Returns:

```json
[
  {
    "name": "Amsterdam",
    "country": "Netherlands"
  }
]
```

---

## Get City By Name

```http
GET /cities/:name
```

Example:

```http
GET /cities/Amsterdam
```

---

## Filter Cities By Country

```http
GET /cities?country=Netherlands
```

Example response:

```json
[
  {
    "name": "Amsterdam",
    "country": "Netherlands"
  }
]
```

---

# Tech Stack

## Backend

- Go
- Gin Web Framework

## Data

- JSON storage

## API Style

- REST architecture

---

# Engineering Highlights

This project demonstrates:

- Go backend development
- REST API implementation
- layered architecture
- HTTP request handling
- service abstraction
- data modeling
- maintainable project structure

---

# Design Decisions

## Separation of Responsibilities

Business logic is separated from HTTP concerns.

Example:

```text
Handler

"Receive request"

        ↓

Service

"Process data"

        ↓

Repository

"Access storage"
```

This improves:

- code organization
- testing capability
- scalability

---

## JSON-Based Storage

The project starts with a lightweight data source.

Advantages:

- simple development setup
- easy testing
- portable dataset

The architecture supports migration to a database without rewriting the API layer.

---

# Getting Started

## Clone Repository

```bash
git clone https://github.com/your-username/city-rest-api.git

cd city-rest-api
```

---

## Install Dependencies

```bash
go mod tidy
```

---

## Run Application

```bash
go run main.go
```

Server starts:

```text
http://localhost:8080
```

---

# Future Improvements

## Database Integration

- PostgreSQL support
- repository abstraction
- database migrations

---

## API Improvements

- pagination
- sorting
- advanced filtering
- request validation
- OpenAPI documentation

---

## Production Features

- environment configuration
- structured logging
- middleware support
- authentication
- rate limiting
- Docker deployment

---

## Testing

- unit tests
- handler tests
- integration tests
- API contract testing

---

# Learning Context

This project represents a progression into Go backend development.

It explores:

- building HTTP services
- designing backend layers
- structuring maintainable Go applications
- applying software engineering principles

---

# What This Project Demonstrates

City REST API demonstrates:

- REST API development in Go
- Gin framework usage
- backend architecture design
- clean separation of concerns
- scalable service organization

---

# License

MIT License
