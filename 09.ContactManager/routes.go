package main

import "github.com/gofiber/fiber/v2"

func SetupRoutes(app *fiber.App) {
	/*
		Get all Contacts
		Create Contact
		Get Contact
		Update Contact
		Delete Contact
	*/
	contactRoute := app.Group("/contacts")

	contactRoute.Get("/", GetContacts)
	contactRoute.Post("/", CreateContact)
	contactRoute.Get("/:contactId", GetContact)
	contactRoute.Put("/:contactId", UpdateContact)
	contactRoute.Delete("/:contactId", DeleteContact)
}
