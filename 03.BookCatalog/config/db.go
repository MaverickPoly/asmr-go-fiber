package config

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB

func CreateTables() {
	bookTable := `
		CREATE TABLE IF NOT EXISTS books (
			id SERIAL PRIMARY KEY,
			title TEXT NOT NULL,
			author TEXT NOT NULL,
			description TEXT NOT NULL,
			rating INTEGER NOT NULL
		)
	`

	if _, err := DB.Query(bookTable); err != nil {
		log.Fatalln("Error creating book table!")
	}
}

func ConnectDB() {
	_ = godotenv.Load()
	dbURL := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	var err error
	DB, err = sqlx.Connect("postgres", dbURL)
	if err != nil {
		log.Fatalln("Unable to connect to database:", err)
	}
	CreateTables()
	fmt.Println("Database connected!")
}
