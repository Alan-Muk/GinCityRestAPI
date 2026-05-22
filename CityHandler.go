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
    file, err := os.ReadFile("data/cities.json")
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

/*

1. CityService Struct
- Stores a slice of cities loaded from the JSON data file.
- Acts as the service layer for city-related operations.

2. NewCityService()
- Reads the "data/cities.json" file.
- Parses the JSON into Go structs using json.Unmarshal.
- Loads all cities into memory and returns a new CityService instance.
- Uses panic() to stop execution if file reading or JSON parsing fails.

3. GetAll()
- Returns all cities stored in the service.

4. GetByName(name string)
- Searches for a city by its name.
- Uses strings.EqualFold() for case-insensitive comparison.
- Returns a pointer to the matching city if found.
- Returns nil if no matching city exists.

5. GetByCountry(country string)
- Filters cities by country name.
- Uses case-insensitive comparison.
- Returns a slice containing all matching cities.
*/
