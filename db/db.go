package db

import (
	"context"
	"fmt"
	"os"
	"City_Weather_CLI/models" 
	"github.com/jackc/pgx/v5"
)

//Connects to the database using the connection URL in the .env file.
func Connect() (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}
	return conn, nil
}

//Creates Table != exist
func CreateTableIfNotExists(conn *pgx.Conn) error {
	query := `
	CREATE TABLE IF NOT EXISTS weather (
		id SERIAL PRIMARY KEY,
		city TEXT,
		temperature FLOAT,
		description TEXT,
		humidity INT,
		wind_speed FLOAT,
		created_at TIMESTAMPTZ DEFAULT now()
	);
	`

	_, err := conn.Exec(context.Background(), query)
	if err != nil {
		return fmt.Errorf("failed to create table: %w", err)
	}
	return nil
}
//Inserts data to table in db
func SaveWeatherData(conn *pgx.Conn, weather *models.Weather) error {
	query := `
		INSERT INTO weather (city, temperature, description, humidity, wind_speed)
		VALUES ($1, $2, $3, $4, $5)
	`

	_, err := conn.Exec(
		context.Background(),
		query,
		weather.City,
		weather.Temperature,
		weather.Description,
		weather.Humidity,
		weather.WindSpeed,
	)

	if err != nil {
		return fmt.Errorf("failed to insert weather data: %w", err)
	}
	return nil
}

func EnforceRowLimit(conn *pgx.Conn) error {
	query := `
		DELETE FROM weather
		WHERE id IN (
			SELECT id FROM weather
			ORDER BY created_at ASC
			OFFSET 15
		);
	`
	_, err := conn.Exec(context.Background(), query)
	if err != nil {
		return fmt.Errorf("failed to enforce row limit: %w", err)
	}
	return nil
}
