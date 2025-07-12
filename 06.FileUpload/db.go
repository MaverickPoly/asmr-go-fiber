package main

import (
	"github.com/gofiber/fiber/v2/log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	log.Info("Connecting DB...")

	DB, err = gorm.Open(sqlite.Open("./sqlite.db"), &gorm.Config{})

	if err != nil {
		log.Fatalf("Error connecting DB: %v", err.Error())
	}

	DB.AutoMigrate(&File{})
	log.Info("Connected to DB successfully!")
}
