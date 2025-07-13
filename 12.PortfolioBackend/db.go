package main

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})

	if err != nil {
		log.Fatalf("Error connecting db: %v", err.Error())
	}

	DB.AutoMigrate(&Project{})
	fmt.Println("Connected DB successfully!")
}
