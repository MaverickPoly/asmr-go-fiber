package main

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "modernc.org/sqlite"
)

var DB *sqlx.DB

func ConnectDB() {
	dbURL := "./sqlite.db"

	var err error
	DB, err = sqlx.Open("sqlite", fmt.Sprintf("file:%s", dbURL))
	if err != nil {
		log.Fatalln("Unable to connect to database:", err)
	}
	fmt.Println("Database connected!")

	query := `CREATE TABLE IF NOT EXISTS urls (
		id INTEGER PRIMARY KEY,
		path TEXT NOT NULL,
		url TEXT NOT NULL
	)`

	if _, err := DB.Exec(query); err != nil {
		log.Fatalf("Error creating table: %s", err.Error())
	}
	fmt.Println("Created table successfully!")
}
