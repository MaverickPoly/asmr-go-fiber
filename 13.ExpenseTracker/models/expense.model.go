package models

import "gorm.io/gorm"

type Expense struct {
	gorm.Model

	Title  string  `json:"title" gorm:"not null"`
	Amount float64 `json:"amount" gorm:"not null;default:0"`
	Type   string  `json:"type" gorm:"not null;default:'expenditure'"` // expenditure, earning

	UserID uint `json:"user_id" gorm:"not null"` // foreign key
}
