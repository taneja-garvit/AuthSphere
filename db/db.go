package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client             // Global MongoDB client (like a db connection pool)
var UserCollection *mongo.Collection // Global reference to "users" collection

func InitDB() error {
	clientOptions := options.Client().ApplyURI("Mongo connection string")

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return err
		}
		Client = client
		UserCollection = client.Database("authdb").Collection("users")
		log.Println("Connected to MongoDB")
		return nil

}
