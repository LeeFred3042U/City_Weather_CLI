package ui

import (
	"City_Weather_CLI/models"
	"fmt"
)

// DisplayWeather prints the weather information to the console
func DisplayWeather(weather *models.Weather) {
	fmt.Println("Weather in", weather.City)
	fmt.Println("Temperature:", weather.Temperature, "°C")
	fmt.Println("Description:", weather.Description)
	fmt.Println("Humidity:", weather.Humidity, "%")
	fmt.Println("Wind Speed:", weather.WindSpeed, "m/s")
}
