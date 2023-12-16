// main.go
// USER MANAGEMENT SYSTEM
package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gofr.dev/pkg/gofr"
	"log"
	"time"
)

// User struct represents the user model
type User struct {
	ID        string    `json:"id,omitempty" bson:"_id,omitempty"`
	Username  string    `json:"username,omitempty" bson:"username,omitempty"`
	Email     string    `json:"email,omitempty" bson:"email,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
}

var client *mongo.Client

// Connect to MongoDB
func init() {
	// Replace the connection string with your MongoDB cluster connection string
	connectionString := "mongodb+srv://khushi:khushi123@cluster0.fg03hwb.mongodb.net/?retryWrites=true&w=majority"
	clientOptions := options.Client().ApplyURI(connectionString)
	// Assume 'client' is your MongoDB client
	//databaseName := "UMS"
	//collectionName := "users"

	var err error
	client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
}

// CreateUserEndpoint creates a new user.
func CreateUserEndpoint(ctx *gofr.Context) (interface{}, error) {
	// Set content type to JSON
	//header := ctx.Header("Content-Type", "application/json")

	// Initialize user
	var user User

	// Bind request body to user variable
	if err := ctx.Bind(&user); err != nil {
		response := "Error decoding user data"
		return response, err
	}

	// Set user creation time
	user.CreatedAt = time.Now()

	// Get database collection
	collection := client.Database("UMS").Collection("users")

	// Insert user into the collection
	result, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		response := "Error creating user"
		return response, err
	}

	// Return the result as JSON
	return result, nil
}

// Get user details
func GetUserEndpoint(ctx *gofr.Context) (interface{}, error) {
	// Get user ID from path parameter
	id := ctx.PathParam("id")

	// Get database collection
	collection := client.Database("UMS").Collection("users")

	// If an ID is provided, retrieve a single user by ID
	if id != "" {
		// Convert the string ID to ObjectId
		objectID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			response := "Invalid ID format"
			return response, err
		}

		// Find user by ID
		var user User
		err = collection.FindOne(context.TODO(), bson.M{"_id": objectID}).Decode(&user)
		if err != nil {
			response := "Could not fetch user details"
			return response, err
		}

		// Return the user details
		return user, nil
	}

	// If no ID is provided, retrieve all users
	var users []User

	// Find all users in the collection
	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		response := "Error fetching users"
		return response, err
	}
	defer cursor.Close(context.TODO())

	// Iterate over the cursor and decode each document into the users slice
	for cursor.Next(context.TODO()) {
		var user User
		err := cursor.Decode(&user)
		if err != nil {
			response := "Error decoding user data"
			return response, err
		}
		users = append(users, user)
	}

	// Check for errors from iterating over the cursor
	if err := cursor.Err(); err != nil {
		response := "Error iterating over users"
		return response, err
	}

	// Return the list of users
	return users, nil
}

func main() {
	app := gofr.New()

	app.POST("/users", CreateUserEndpoint)
	app.GET("/users", GetUserEndpoint)
	app.GET("/users/{id}", GetUserEndpoint)

	app.Start()

}
