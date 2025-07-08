package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var DB *sql.DB

func InitDB() {
	godotenv.Load()
	url := os.Getenv("DATABASE_URL")

	var err error
	DB, err = sql.Open("pgx", url)
	if err != nil {
		log.Fatal("Error abriendo DB: ", err)
	}

	if err := DB.Ping(); err != nil {
		log.Fatal("Error conectando a DB: ", err)
	}
}
