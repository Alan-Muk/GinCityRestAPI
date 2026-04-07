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