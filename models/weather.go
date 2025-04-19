type apiResponse struct {
	Name string `json:"name"`
	Main struct {
		Temp     float64 `json:"temp"`
		Humidity int     `json:"humidity"` // <-- new
	} `json:"main"`
	Wind struct {
		Speed float64 `json:"speed"` // <-- new
	} `json:"wind"`
	Weather []struct {
		Description string `json:"description"`
	} `json:"weather"`
}

func GetWeather(city string) (*models.Weather, error) {
	// ... existing code ...

	weather := &models.Weather{
		City:        result.Name,
		Temperature: result.Main.Temp,
		Description: result.Weather[0].Description,
		Humidity:    result.Main.Humidity,   // <-- new
		WindSpeed:   result.Wind.Speed,      // <-- new
	}

	return weather, nil
}
