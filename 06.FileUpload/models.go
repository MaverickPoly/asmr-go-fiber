package main

import "gorm.io/gorm"

type File struct {
	gorm.Model

	Filename string
	Path     string
	Size     float64
}
