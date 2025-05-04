package routes

import (
	"github.com/harsh/book_crud/controllers"

	"github.com/gorilla/mux"
)

// SetupRoutes sets up the routes for the application
func RegisterBookRoutes(router *mux.Router) {
	// Define the routes for the book resource
	router.HandleFunc("/books", controllers.CreateBook).Methods("POST") // Create a new book
	router.HandleFunc("/books", controllers.GetBooks).Methods("GET")    // Get all books
	router.HandleFunc("/books/{id}", controllers.GetBook).Methods("GET") // Get a book by ID
	router.HandleFunc("/books/{id}", controllers.UpdateBook).Methods("PUT") // Update a book by ID
	router.HandleFunc("/books/{id}", controllers.DeleteBook).Methods("DELETE") // Delete a book by ID
}
