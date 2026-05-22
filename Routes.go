package routes

import (
    "github.com/gin-gonic/gin"
    "yourmodule/handlers"
)

func RegisterRoutes(r *gin.Engine, h *handlers.CityHandler) {
    api := r.Group("/api/v1")
    {
        api.GET("/cities", h.GetCities)
        api.GET("/cities/:name", h.GetCity)
        api.GET("/countries/:country", h.GetByCountry)
    }
}

/*

1. RegisterRoutes()
- Registers all API routes for the application.
- Connects HTTP endpoints to handler functions.

2. Router Parameter
- r *gin.Engine represents the main Gin router instance.
- Used to define application routes.

3. Handler Parameter
- h *handlers.CityHandler contains the controller logic
  for handling city-related requests.

4. API Route Group
- Creates a grouped route prefix: "/api/v1".
- Helps organize and version the API.

5. Routes Defined

- GET /api/v1/cities
  Calls h.GetCities
  Returns all cities.

- GET /api/v1/cities/:name
  Calls h.GetCity
  Returns a city by name using a URL parameter.

- GET /api/v1/countries/:country
  Calls h.GetByCountry
  Returns all cities belonging to a specific country.

6. Route Group Benefits
- Keeps routes organized.
- Makes API versioning easier.
- Reduces repeated route prefixes.
*/
