package main

import (
	"City_Weather_CLI/api"
	"City_Weather_CLI/config"
	"fmt"
)

func main() {
	config.LoadEnv()

	weather, err := api.GetWeather("Lucknow")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("🌆 City: %s\n🌡️ Temp: %.2f°C\n🌤️ Description: %s\n",
		weather.City, weather.Temperature, weather.Description)
}
