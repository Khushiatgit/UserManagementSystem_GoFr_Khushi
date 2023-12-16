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

func main() {
	app := gofr.New()

	app.Start()

}
