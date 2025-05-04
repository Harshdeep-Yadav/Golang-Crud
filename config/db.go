// package config

// import (
// 	"context"
// 	"fmt"
// 	"log"

// 	// "os"
// 	"time"

// 	// "github.com/joho/godotenv"
// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// )

// var DB *mongo.Database

// // func ConnectDB() {
// // 	// Load environment variables from .env file
// // 	// This is optional, but it's a good practice to keep sensitive data out of your codebase
// // 	// err := godotenv.Load()
// // 	// if err != nil {
// // 	// 	log.Fatal("Error loading .env file")
// // 	// }

// // 	// Get MongoDB connection string and database name from environment variables
// // 	mongoURI := "mongodb+srv://harsh:harshy@eccomerceapi.h0iaweq.mongodb.net/bookstore?retryWrites=true&w=majority"
// // 	dbName := "bookstore"

// // 	fmt.Println("Mongo URI:", mongoURI)
// // 	fmt.Println("DB Name:", dbName)

// // 	// Create a new MongoDB client and connect to the database
// // 	// Use the connection string and database name from the environment variables
// // 	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
// // 	if err != nil {
// // 		log.Fatal("Error creating MongoDB client:", err)
// // 	}

// // 	// Set a timeout for the connection
// // 	// This is important to avoid hanging indefinitely if the database is unreachable
// // 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// // 	defer cancel()

// // 	// Connect to the MongoDB server
// // 	// This is where the actual connection happens
// // 	err = client.Connect(ctx)
// // 	if err != nil {
// // 		log.Fatal("Error connecting to MongoDB:", err)
// // 	}

// // 	// Ping the MongoDB server to check if the connection is successful
// // 	// This is a good way to verify that the connection is working as expected
// // 	err = client.Ping(ctx, nil)
// // 	if err != nil {
// // 		log.Fatal("Error pinging MongoDB:", err)
// // 	}
// // 	fmt.Println("Connected to MongoDB")
// // 	// Set the database to use
// // 	// This is where you specify which database you want to work with
// // 	DB = client.Database(dbName)
// // }


package config

import (
    "context"
    "fmt"
    "log"
    "time"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func ConnectDB() {
    // MongoDB connection string and database name
    mongoURI := "mongodb+srv://harsh:harshy@eccomerceapi.h0iaweq.mongodb.net/bookstore?retryWrites=true&w=majority"
    dbName := "bookstore"

    fmt.Println("Mongo URI:", mongoURI)
    fmt.Println("DB Name:", dbName)

    // Set a timeout for the connection
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    // Connect to MongoDB server
    client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
    if err != nil {
        log.Fatal("Error connecting to MongoDB:", err)
    }

    // Ping the MongoDB server to verify the connection
    err = client.Ping(ctx, nil)
    if err != nil {
        log.Fatal("Error pinging MongoDB:", err)
    }
    fmt.Println("Connected to MongoDB")

    // Set the database to use
    DB = client.Database(dbName)
}