package main

import "gorm.io/gorm"

type Project struct {
	gorm.Model

	Title       string `json:"title" gorm:"title"`
	Description string `json:"description" gorm:"description"`
	Language    string `json:"language" gorm:"language"`
	Link        string `json:"link" gorm:"link"`
}
