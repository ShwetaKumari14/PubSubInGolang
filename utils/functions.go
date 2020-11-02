package utils

import (
	"log"
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func GenerateLogs(err error, msg string){
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func GetDBConnection() (*sql.DB, error) {
	database, err := sql.Open("sqlite3", "DB.db")
	if err != nil {
		return nil, err
	}

	return database, nil
}

