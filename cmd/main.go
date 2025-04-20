package main

import (
	"fmt"
	"log"
	"context"

	"City_Weather_CLI/db"
	"City_Weather_CLI/ui"
	"City_Weather_CLI/api"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connect to database
	conn, err := db.Connect()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer conn.Close(context.Background())

	// Create table if not exists
	err = db.CreateTableIfNotExists(conn)
	if err != nil {
		log.Fatal("Error creating table:", err)
	}

	// Fetch weather data
	city := "Lucknow" // Example city
	weather, err := api.GetWeather(city)
	if err != nil {
		log.Fatal("Error fetching weather data:", err)
	}

	// Save weather data to the database
	err = db.SaveWeatherData(conn, weather)
	if err != nil {
		log.Fatal("Error saving weather data:", err)
	}

	// Display weather data
	ui.DisplayWeather(weather)

	fmt.Println("Weather data saved and displayed successfully.")
}
