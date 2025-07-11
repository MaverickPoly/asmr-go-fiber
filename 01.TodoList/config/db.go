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
	todoTable := `CREATE TABLE IF NOT EXISTS todos (
		id SERIAL PRIMARY KEY,
		title TEXT NOT NULL,
		completed BOOLEAN NOT NULL DEFAULT false
	);`

	DB.Exec(todoTable)

	fmt.Println("Created Tables!")
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
	fmt.Println("DBUrl: ", dbURL)

	var err error
	DB, err = sqlx.Connect("postgres", dbURL)
	if err != nil {
		log.Fatalln("Unable to connect to database:", err)
	}
	CreateTables()
	fmt.Println("Database connected!")
}
