package main

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client

// Collections
var ContactCollection *mongo.Collection

func ConnectDB() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatalln("MongoDB connection error:", err)
	}

	MongoClient = client
	ContactCollection = client.Database("go-contact-manager").Collection("users")
	log.Println("Connected to MongoDB!")
}
