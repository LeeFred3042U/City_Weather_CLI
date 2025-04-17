package models

import (
	"context"
	"fmt"
	"City_Weather_CLI/db"
)

func InitSchema() {
	query := `
	CREATE TABLE IF NOT EXISTS city_weather (
		id SERIAL PRIMARY KEY,
		city VARCHAR(100) UNIQUE NOT NULL,
		temperature DOUBLE PRECISION NOT NULL
	);
	`
	_, err := db.Conn.Exec(context.Background(), query)
	if err != nil {
		panic(err)
	}
}

func AddOrUpdateCity(city string, temp float64) {
	_, err := db.Conn.Exec(context.Background(), `
		INSERT INTO city_weather (city, temperature)
		VALUES ($1, $2)
		ON CONFLICT (city) DO UPDATE SET temperature = EXCLUDED.temperature
	`, city, temp)

	if err != nil {
		fmt.Println("Failed:", err)
		return
	}
	fmt.Println("City saved.")
}

func ListCities() {
	rows, err := db.Conn.Query(context.Background(), `SELECT city, temperature FROM city_weather`)
	if err != nil {
		fmt.Println("Query failed:", err)
		return
	}
	defer rows.Close()

	fmt.Println("\nAll Cities:")
	for rows.Next() {
		var city string
		var temp float64
		rows.Scan(&city, &temp)
		fmt.Printf("- %s: %.2f°C\n", city, temp)
	}
}

func ListSorted(order string) {
	var query string
	if order == "asc" {
		query = `SELECT city, temperature FROM city_weather ORDER BY temperature ASC`
	} else {
		query = `SELECT city, temperature FROM city_weather ORDER BY temperature DESC`
	}

	rows, err := db.Conn.Query(context.Background(), query)
	if err != nil {
		fmt.Println("Query failed:", err)
		return
	}
	defer rows.Close()

	fmt.Printf("\nSorted by Temperature (%s):\n", order)
	for rows.Next() {
		var city string
		var temp float64
		rows.Scan(&city, &temp)
		fmt.Printf("- %s: %.2f°C\n", city, temp)
	}
}
