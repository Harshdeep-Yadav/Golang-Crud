package repository

import (
	"context"
	"errors"
	"time"

	"github.com/harsh/book_crud/config"
	"github.com/harsh/book_crud/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var bookCollection *mongo.Collection = config.DB.Collection("bookstore") // Initialize the book collection

// CreateBook inserts a new book into the database
// It takes a pointer to a Book struct and returns the result of the insertion and any error that occurred
func CreateBook(book *models.Book) (*mongo.InsertOneResult, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second) // Set a timeout for the context to avoid hanging indefinitely
    defer cancel()

	// Generate a new ObjectID for the book
    result, err := bookCollection.InsertOne(ctx, book)
    return result, err
}

// GetAllBooks retrieves all books from the database
func GetAllBooks() ([]models.Book, error) {
	// Set a timeout for the context to avoid hanging indefinitely
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    cursor, err := bookCollection.Find(ctx, bson.M{}) // Find all books in the collection
    if err != nil {
        return nil, err
    }
    defer cursor.Close(ctx) // Close the cursor when done

    var books []models.Book
    for cursor.Next(ctx) {
        var book models.Book
        if err := cursor.Decode(&book); err != nil {
            return nil, err
        }
        books = append(books, book)
    }

    return books, nil
}

// GetBookByID retrieves a book by its ID from the database
func GetBookByID(id string) (*models.Book, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    objID, err := primitive.ObjectIDFromHex(id) // Convert the string ID to an ObjectID
    if err != nil {
        return nil, err
    }

    var book models.Book
    err = bookCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&book) // Find the book by ID and decode it into the book variable
    if err != nil {
        return nil, err
    }

    return &book, nil
}

func UpdateBook(id string, book *models.Book) error {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    objID, err := primitive.ObjectIDFromHex(id) // Convert the string ID to an ObjectID
    if err != nil {
        return err
    }

    result, err := bookCollection.ReplaceOne(ctx, bson.M{"_id": objID}, book) // Replace the book with the new data
    if err != nil {
        return err
    }

    if result.MatchedCount == 0 {
        return errors.New("no book found to update")
    }

    return nil
}

func DeleteBook(id string) error {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    objID, err := primitive.ObjectIDFromHex(id) // Convert the string ID to an ObjectID 
    if err != nil {
        return err
    }

    result, err := bookCollection.DeleteOne(ctx, bson.M{"_id": objID}) // Delete the book by ID
    if err != nil {
        return err
    }

    if result.DeletedCount == 0 {
        return errors.New("no book found to delete")
    }

    return nil
}
