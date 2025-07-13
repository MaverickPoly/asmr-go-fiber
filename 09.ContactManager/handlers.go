package main

import (
	"context"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetContacts(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := ContactCollection.Find(ctx, bson.M{})

	contacts := make([]Contact, 0)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("Error getting contacts: %v", err.Error()),
		})
	}

	if err := cursor.All(ctx, &contacts); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("Error getting contacts: %v", err.Error()),
		})
	}

	fmt.Println("Get contacts:", contacts)

	return c.JSON(contacts)
}

func CreateContact(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	contact := new(Contact)

	if err := c.BodyParser(&contact); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid body!",
		})
	}

	if contact.Name == "" || contact.Email == "" || contact.Phone == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Some fields are missing!",
		})
	}

	res, err := ContactCollection.InsertOne(ctx, contact)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error creating contact!",
		})
	}
	fmt.Println("Added")
	contact.ID = res.InsertedID.(primitive.ObjectID).Hex()

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Contact created successfully!",
		"data":    contact,
	})
}

func GetContact(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	contactId := c.Params("contactId")
	objectID, err := primitive.ObjectIDFromHex(contactId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid contact ID",
		})
	}

	var contact Contact
	filter := bson.M{"_id": objectID}
	err = ContactCollection.FindOne(ctx, filter).Decode(&contact)

	fmt.Println("Get contact:", contact)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Contact not found!",
		})
	}

	return c.JSON(contact)
}

func UpdateContact(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Check If ID valid
	contactId := c.Params("contactId")
	objectID, err := primitive.ObjectIDFromHex(contactId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid contact ID",
		})
	}

	// Parse Body
	var contact Contact
	if err := c.BodyParser(&contact); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid Body!",
		})
	}

	// Prepare update fields
	update := bson.M{}
	if contact.Name != "" {
		update["name"] = contact.Name
	}
	if contact.Phone != "" {
		update["phone"] = contact.Phone
	}
	if contact.Email != "" {
		update["email"] = contact.Email
	}

	if len(update) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "No valid fields provided for update!",
		})
	}

	// Perform Update
	res, err := ContactCollection.UpdateOne(ctx,
		bson.M{"_id": objectID},
		bson.M{"$set": update},
	)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update contact",
		})
	}
	if res.MatchedCount == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Contact not found!",
		})
	}

	// Return Updated data
	var updatedContact Contact
	_ = ContactCollection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&updatedContact)
	return c.JSON(fiber.Map{
		"message": "Contact updated successfully!",
		"data":    updatedContact,
	})
}

func DeleteContact(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	contactId := c.Params("contactId")
	objectID, err := primitive.ObjectIDFromHex(contactId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid contact ID!",
		})
	}

	var deletedContact Contact
	if err := ContactCollection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&deletedContact); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Contact not Found!",
		})
	}

	result, err := ContactCollection.DeleteOne(ctx, bson.M{"_id": objectID})
	if result.DeletedCount == 0 || err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Contact not Found!",
		})
	}

	return c.JSON(fiber.Map{"message": "Contact Deleted successfully!", "data": deletedContact})
}
