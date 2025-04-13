package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/devfullcycle/imersao22/go-gateway/internal/repository"
	"github.com/devfullcycle/imersao22/go-gateway/internal/service"
	"github.com/devfullcycle/imersao22/go-gateway/internal/web/server"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func getEnv(key, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	connString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		getEnv("DB_HOST", "localhost"),
		getEnv("DB_PORT", "5432"),
		getEnv("DB_USER", "postgres"),
		getEnv("DB_PASSWORD", "postgres"),
		getEnv("DB_NAME", "go_gateway"))

	log.Printf("Connecting to database: %s", connString)
	db, err := sql.Open("postgres", connString)
	if err != nil {
		log.Fatal("Error opening database: ", err)
	}
	defer db.Close()

	// Test database connection
	if err := db.Ping(); err != nil {
		log.Fatal("Error connecting to database: ", err)
	}
	log.Println("Successfully connected to database")

	accountRepository := repository.NewAccountRepository(db)
	accountService := service.NewAccountService(*accountRepository)

	port := getEnv("HTTP_PORT", "8080")
	log.Printf("Starting server on port %s", port)
	server := server.NewServer(port, *accountService)
	if err := server.Start(); err != nil {
		log.Fatal("Error starting server: ", err)
	}

}
