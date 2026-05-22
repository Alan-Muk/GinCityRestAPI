package models

type City struct {
    Name       string  `json:"name"`
    Country    string  `json:"country"`
    Population int     `json:"population"`
    Lat        float64 `json:"lat"`
    Lon        float64 `json:"lon"`
    Timezone   string  `json:"timezone"`
}

type Data struct {
    Cities []City `json:"cities"`
}

/*

1. City Struct
- Represents a city object in the application.
- Each field maps to data from the JSON file using struct tags.

2. City Fields
- Name: Stores the city name.
- Country: Stores the country where the city is located.
- Population: Stores the total population of the city.
- Lat: Stores the latitude coordinate.
- Lon: Stores the longitude coordinate.
- Timezone: Stores the city's timezone.

3. JSON Tags
- The `json:"..."` tags define how JSON keys map to struct fields.
- Example:
  JSON key "name" maps to the Name field.

4. Data Struct
- Represents the root structure of the JSON file.
- Contains a slice of City objects.

5. Cities Field
- Holds all city records loaded from the JSON data.
- Maps to the "cities" array in the JSON file.
*/
