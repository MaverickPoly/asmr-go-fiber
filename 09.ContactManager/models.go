package main

type Contact struct {
	ID    string `json:"id,omitempty" bson:"_id,omitempty"`
	Name  string `json:"name" bson:"name"`
	Email string `json:"email" bson:"email"`
	Phone string `json:"phone" bson:"phone"` // Phone Number
}
