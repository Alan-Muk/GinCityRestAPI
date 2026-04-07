package services

import (
    "encoding/json"
    "os"
    "strings"

    "yourmodule/models"
)

type CityService struct {
    Cities []models.City
}

func NewCityService() *CityService {
    file, err := os.ReadFile("world-cities.json")
    if err != nil {
        panic(err)
    }

    var data models.Data
    if err := json.Unmarshal(file, &data); err != nil {
        panic(err)
    }

    return &CityService{Cities: data.Cities}
}

// Get all
func (s *CityService) GetAll() []models.City {
    return s.Cities
}

// Get by name
func (s *CityService) GetByName(name string) *models.City {
    for _, c := range s.Cities {
        if strings.EqualFold(c.Name, name) {
            return &c
        }
    }
    return nil
}

// Filter by country
func (s *CityService) GetByCountry(country string) []models.City {
    var result []models.City
    for _, c := range s.Cities {
        if strings.EqualFold(c.Country, country) {
            result = append(result, c)
        }
    }
    return result
}