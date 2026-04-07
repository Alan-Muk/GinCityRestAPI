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