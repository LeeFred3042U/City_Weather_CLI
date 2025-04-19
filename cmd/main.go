package main

import (
	"City_Weather_CLI/config"
	"City_Weather_CLI/db"
	"fmt"
)

func main() {
	// Loading from .env file
	config.LoadEnv()

	//Getting DB connection string
	connStr := config.GetDBURL()

	//connecting DB
	db.ConnectDB(connStr)
	defer db.CloseDB()

	fmt.Println("App started successfully.")
}
