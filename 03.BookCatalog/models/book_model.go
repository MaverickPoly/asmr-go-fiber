package models

type Book struct {
	ID          int    `db:"id" json:"id"`
	Title       string `db:"title" json:"title"`
	Author      string `db:"author" json:"author"`
	Description string `db:"description" json:"description"`
	Rating      uint   `db:"rating" json:"rating"`
}
