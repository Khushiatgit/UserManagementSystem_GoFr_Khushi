// main.go
// USER MANAGEMENT SYSTEM
package main

import (
	"context"
	"fmt"
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

func main() {
	app := gofr.New()

	app.POST("/users", CreateUserEndpoint)

	app.Start()

}
