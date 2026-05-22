package main

import (
    "github.com/gin-gonic/gin"
    "yourmodule/handlers"
    "yourmodule/routes"
    "yourmodule/services"
)

func main() {
    r := gin.Default()

    cityService := services.NewCityService()
    cityHandler := handlers.NewCityHandler(cityService)

    routes.RegisterRoutes(r, cityHandler)

    r.Run(":8080")
}

/*

1. Main Function
- Entry point of the application.
- Initializes the web server and application components.

2. Gin Router
- gin.Default() creates a Gin router with
  default middleware such as logging and recovery.

3. Service Initialization
- services.NewCityService() creates a new CityService.
- Loads city data from the JSON file into memory.

4. Handler Initialization
- handlers.NewCityHandler(cityService) creates
  a new handler and injects the service dependency.
- Allows handlers to access city-related operations.

5. Route Registration
- routes.RegisterRoutes() connects API endpoints
  to their corresponding handler functions.

6. Run Server
- r.Run(":8080") starts the HTTP server on port 8080.
- The API becomes accessible through:
  http://localhost:8080

7. Overall Flow
- Main initializes:
  Router -> Service -> Handler -> Routes -> Server
- This structure follows a layered application design.
*/
