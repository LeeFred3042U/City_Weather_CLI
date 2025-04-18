package db

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)

// Conn is a global connection handle.
// Only use global like this in CLI tools — never in web servers.
var Conn *pgx.Conn

//Establishing connection to db
func ConnectDB(connStr string) {
	var err error

	Conn, err = pgx.Connect(context.Background(), connStr)
	if err != nil {
		log.Fatal("Could not connect to DB:", err)
	}
	fmt.Println("Connected to PostgreSQL (Neon)")
}

//Closing the DB
func CloseDB() {
	if Conn != nil {
		if err := Conn.Close(context.Background()); err != nil {
			log.Printf("Error closing DB connection: %v", err)
		} else {
			fmt.Println("PostgreSQL connection closed")
		}
	} else {
		log.Println("No active DB connection to close.")
	}
}
