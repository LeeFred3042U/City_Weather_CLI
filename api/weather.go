package api

import (
	"City_Weather_CLI/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type apiResponse struct {
	Name string `json:"name"`
	Main struct {
		Temp     float64 `json:"temp"`
		Humidity int     `json:"humidity"`
	} `json:"main"`
	Wind struct {
		Speed float64 `json:"speed"`
	} `json:"wind"`
	Weather []struct {
		Description string `json:"description"`
	} `json:"weather"`
}

// GetWeather fetches weather data from OpenWeather API
func GetWeather(city string) (*models.Weather, error) {
	apiKey := os.Getenv("OPENWEATHER_API_KEY")
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", city, apiKey)


	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get data: %s", resp.Status)
	}

	var result apiResponse
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}



	TempInCelsius := int(result.Main.Temp - 273.15)

	weather := &models.Weather{
		City:        result.Name,
		Temperature: TempInCelsius,
		Description: result.Weather[0].Description,
		Humidity:    result.Main.Humidity,
		WindSpeed:   result.Wind.Speed,
	}

	return weather, nil
}
