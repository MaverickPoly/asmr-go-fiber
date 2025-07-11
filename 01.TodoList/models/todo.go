package models

type Todo struct {
	ID        int    `db:"id" json:"id"`
	Title     string `db:"title" json:"title"`
	Completed bool   `db:"completed" json:"completed"`
}
