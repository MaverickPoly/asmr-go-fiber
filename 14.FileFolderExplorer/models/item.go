package models

import (
	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	Name     string `json:"name" gorm:"not null"`
	IsFolder bool   `json:"is_folder" gorm:"not null;default:false"`
	ParentID *uint  `json:"parent_id"` // nullable for root items
}
