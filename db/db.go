package db

import (
	"context"   // for context.Background()
	"fmt"       // for fmt.Println
	"log"       // for log.Fatal
	"github.com/jackc/pgx/v5"
)

var Conn *pgx.Conn

func ConnectDB(connStr string) {
	var err error
	Conn, err = pgx.Connect(context.Background(), connStr)
	if err != nil {
		log.Fatal("Could not connect to DB:", err)
	}
	fmt.Println("Connected to PostgreSQL (Neon)")
}

func CloseDB() {
	Conn.Close(context.Background())
}
