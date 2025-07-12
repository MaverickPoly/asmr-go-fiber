package main

import (
	"github.com/gofiber/fiber/v2/log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	log.Info("Connecting DB...")
	DB, err = gorm.Open(sqlite.Open("./sqlite.db"), &gorm.Config{})

	if err != nil {
		log.Fatalf("Error connecting db: %v", err.Error())
	}

	DB.AutoMigrate(&User{})
	log.Info("Connected to DB successfully!")
}
