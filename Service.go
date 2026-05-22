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

/*

1. CityService Struct
- Stores all city data loaded from the JSON file.
- Provides methods for accessing and filtering city information.

2. NewCityService()
- Reads data from the "world-cities.json" file.
- Converts JSON data into Go structs using json.Unmarshal().
- Stores the loaded cities inside a CityService instance.
- Uses panic() to stop execution if file loading or JSON parsing fails.

3. JSON Processing
- os.ReadFile() reads the contents of the JSON file.
- json.Unmarshal() converts JSON data into Go objects.

4. GetAll()
- Returns the complete list of cities stored in memory.

5. GetByName(name string)
- Searches for a city using its name.
- Uses strings.EqualFold() for case-insensitive comparison.
- Returns a pointer to the matching city if found.
- Returns nil if no city matches the provided name.

6. GetByCountry(country string)
- Filters cities by country name.
- Uses case-insensitive comparison for matching.
- Returns a slice containing all cities from the specified country.

7. Overall Purpose
- Acts as the service layer between the data source
  and the application handlers/routes.
- Centralizes city-related business logic.
*/
