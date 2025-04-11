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

func getEnv(key string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return ""
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	connString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", getEnv("DB_HOST"), getEnv("DB_PORT"), getEnv("DB_USER"), getEnv("DB_PASSWORD"), getEnv("DB_NAME"))
	db, err := sql.Open("postgres", connString)
	if err != nil {
		log.Fatal("Error opening database")
	}
	defer db.Close()

	accountRepository := repository.NewAccountRepository(db)
	accountService := service.NewAccountService(accountRepository)

	server := server.NewServer(getEnv("HTTP_PORT"), accountService)
	if err := server.Start(); err != nil {
		log.Fatal("Error starting server")
	}

}
