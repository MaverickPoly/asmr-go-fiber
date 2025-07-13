package main

import "gorm.io/gorm"

type Recipe struct {
	gorm.Model

	Name        string `json:"name" gorm:"name"`
	Description string `json:"description" gorm:"description"`
	Ingredients string `json:"ingredients" gorm:"ingredients"`
	Steps       string `json:"steps" gorm:"steps"`
}
