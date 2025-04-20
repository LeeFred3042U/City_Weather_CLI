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
	//Loading variables
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//Connect to db
	conn, err := db.Connect()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer conn.Close(context.Background())

	//If table does'nt exist create it
	err = db.CreateTableIfNotExists(conn)
	if err != nil {
		log.Fatal("Error creating table:", err)
	}

	//Fetching Data
	city := "Lucknow" // Example city
	weather, err := api.GetWeather(city)
	if err != nil {
		log.Fatal("Error fetching weather data:", err)
	}


	//Display
	ui.DisplayWeather(weather)

	fmt.Println("Weather data saved and displayed successfully.")


	//Saving data in db
	err = db.SaveWeatherData(conn, weather)
	if err != nil {
		log.Fatal("Error saving weather data:", err)
	}

	//15 row limit
	err = db.EnforceRowLimit(conn)
	if err != nil {
		log.Fatal("Error enforcing row limit:", err)
	}

}

