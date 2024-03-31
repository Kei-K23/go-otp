package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func NewDB(connectionString string) (*sql.DB, error) {
	db, err := sql.Open("mysql", connectionString)

	if err != nil {
		log.Fatalf("failed to open database : %s", err.Error())
	}

	return db, nil
}

func InitDB(db *sql.DB) {
	err := db.Ping()

	if err != nil {
		log.Fatalf("failed to connect to database : %s", err.Error())
	}

	fmt.Println("successfully connected to database")
}
