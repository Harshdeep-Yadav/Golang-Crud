package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
    DB     *mongo.Database
    client *mongo.Client
)

// ConnectDB establishes a connection to MongoDB
func ConnectDB() error {
    // Get MongoDB connection string and database name from environment variables
    mongoURI := os.Getenv("MONGODB_URI")
    if mongoURI == "" {
        mongoURI = "mongodb+srv://harsh:harshy@eccomerceapi.h0iaweq.mongodb.net/bookstore?retryWrites=true&w=majority"
    }

    dbName := os.Getenv("MONGODB_DB_NAME")
    if dbName == "" {
        dbName = "bookstore"
    }

    // Configure client options
    clientOptions := options.Client().
        ApplyURI(mongoURI).
        SetMaxPoolSize(100).
        SetMinPoolSize(5).
        SetMaxConnIdleTime(5 * time.Minute).
        SetConnectTimeout(10 * time.Second)

    // Set a timeout for the connection
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    // Create and connect the MongoDB client
    var err error
    client, err = mongo.Connect(ctx, clientOptions)
    if err != nil {
        return fmt.Errorf("error connecting to MongoDB: %v", err)
    }

    // Ping the MongoDB server
    if err = client.Ping(ctx, nil); err != nil {
        return fmt.Errorf("error pinging MongoDB: %v", err)
    }

    // Set the database
    DB = client.Database(dbName)

    log.Printf("Successfully connected to MongoDB database: %s", dbName)
    return nil
}

// DisconnectDB closes the MongoDB connection
func DisconnectDB() error {
    if client == nil {
        return nil
    }

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    if err := client.Disconnect(ctx); err != nil {
        return fmt.Errorf("error disconnecting from MongoDB: %v", err)
    }

    log.Println("Successfully disconnected from MongoDB")
    return nil
}

// GetCollection returns a MongoDB collection
func GetCollection(collectionName string) *mongo.Collection {
    return DB.Collection(collectionName)
}