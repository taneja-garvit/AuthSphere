package controllers

import (
	"context"
	"login-signup-gin/db"
	"login-signup-gin/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func Signup(c *gin.Context) {
	var user models.User // Empty user struct to hold request data

	// Bind JSON from request body to user struct (like req.body in Express)
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": "Invalid req body"}) // Gin’s way of sending JSON response
		return
	}
	// Check if user already exists
	var existingUser models.User

	err := db.UserCollection.FindOne(context.Background(), 
        bson.M{"email": user.Email}).Decode(&existingUser) // bson.M is a map representing the MongoDB query, where email is the key and user.Email is the value you're searching for. So, the query is looking for a document where the email field matches the email provided by the user object.
	//After executing the query, the result is decoded into the existingUser variable. The Decode() function is used to convert the BSON document returned by MongoDB into a Go struct.
	if err == nil {
		c.JSON(400, gin.H{"error": "User already exists"}) // Gin’s way
		return
	}
	result,err := db.UserCollection.InsertOne(context.Background(), user) // Inserting a new user into the collection
	if err != nil {
		c.JSON(400, gin.H{"error": "Failed to create user"}) // Gin’s
		return
		}
	c.JSON(201,gin.H{
		"message":"User created",
		"id":result.InsertedID,
	})
}

func Login(c *gin.Context){
	var user models.User;
	if err := c.ShouldBindJSON(&user); err != nil {          // similar to req.body all data came from json
		c.JSON(400, gin.H{"error": "Invalid req body"}) // Gin’s way
		return
		}
		var foundUser models.User

		err:= db.UserCollection.FindOne(context.Background(),
		bson.M{"email": user.Email}).Decode(&foundUser) // similar to the previous
		if err != nil {
			c.JSON(400, gin.H{"error": "User not found"}) // Gin’s way
			return
			}
			if foundUser.Password != user.Password {
				c.JSON(400, gin.H{"error": "Invalid password"}) // Gin’s way
				return
				}
			c.JSON(200,gin.H{
				"message": "User logged in",
				"email":foundUser.Email,
			})
}
