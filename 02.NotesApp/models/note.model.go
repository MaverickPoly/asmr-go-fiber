package models

type Note struct {
	ID        int    `db:"id" json:"id"`
	Title     string `db:"title" json:"title"`
	CreatedAt string `db:"created_at" json:"created_at"`
}
