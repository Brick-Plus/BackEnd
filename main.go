package main

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
        log.Println(".env file not found, using system environment variables")
    }

    user := "postgres"
    password := os.Getenv("DB_PASSWORD")
    host := os.Getenv("DB_HOST")
    dbname := os.Getenv("DB_NAME")
    port := os.Getenv("DB_PORT")

	encodedPassword := url.QueryEscape(password)


    databaseURL := fmt.Sprintf(
        "postgresql://%s:%s@%s:%s/%s?sslmode=require",
        user,
        encodedPassword,
        host,
        port,
        dbname,
    )



    fmt.Println("DATABASE_URL:", databaseURL)
	conn, err := pgx.Connect(context.Background(), databaseURL)

	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer conn.Close(context.Background())

	var version string
	if err := conn.QueryRow(context.Background(), "SELECT version()").Scan(&version); err != nil {
		log.Fatalf("Query failed: %v", err)
	}

	log.Println("Connected to:", version)
}