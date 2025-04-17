package db

import (
	"context"
	"fmt"
	"log"
	"github.com/jackc/pgx/v5"
)
var Conn *pgx.Connect
func connectDB(connStr string){
	var err error 
	Conn, err = pgx.Connect(context.Background(), connStr)
	if err !=nil {
		
	}

}
func ConnectDB(connStr string) {
	var err error
	Conn, err = pgx.Connect(context.Background(), connStr)
	if err != nil {
		log.Fatal("❌ Could not connect to DB:", err)
	}
	fmt.Println("✅ Connected to PostgreSQL (Neon)")
}

func CloseDB() {
	Conn.Close(context.Background())
}
