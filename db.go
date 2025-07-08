package main

import (
	"log"
	"os"

	"database/sql"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var db *sql.DB

func InitDB() error {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("La variable DATABASE_URL no está definida")
	}

	var err error
	db, err = sql.Open("pgx", dsn)
	if err != nil {
		return err
	}
	err = db.Ping()
	if err != nil {
		return err
	}
	log.Println("Conectado a la base de datos")
	return nil
}
