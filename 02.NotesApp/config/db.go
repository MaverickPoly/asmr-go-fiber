package config

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "modernc.org/sqlite"
)

var DB *sqlx.DB

func CreateTables() {
	todoQuery := `
		CREATE TABLE IF NOT EXISTS notes (
			id INTEGER PRIMARY KEY,
			title TEXT NOT NULL,
			created_at TEXT NOT NULL DEFAULT (datetime('now', 'localtime'))
		)
	`

	_, err := DB.Exec(todoQuery)

	if err != nil {
		log.Fatalf("Error creating tables: %s", err.Error())
	}

	fmt.Println("Created tables successfully!")
}

func ConnectDB() {
	dbURL := "./sqlite.db"

	// db, err := sqlx.Open("sqlite3", "file:./mydb.db?_foreign_keys=on")
	var err error
	DB, err = sqlx.Open("sqlite", fmt.Sprintf("file:%s", dbURL))
	if err != nil {
		log.Fatalln("Unable to connect to database:", err)
	}
	CreateTables()
	fmt.Println("Database connected!")
}
