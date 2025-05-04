package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/harsh/book_crud/models"
	"github.com/harsh/book_crud/services"
	"github.com/harsh/book_crud/utils"

	"github.com/gorilla/mux"
)

// CreateBook handles the creation of a new book
// It reads the book data from the request body, validates it, and calls the service to create the book
// It returns a JSON response with the created book or an error message
func CreateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	// Decode the request body into the book struct
	// This is where the actual decoding happens
	// If there's an error during decoding, send a bad request response
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	result, err := services.CreateBookService(&book) // Call the service to create the book
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to create book")
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, result) // Send a JSON response with the created book
}

// GetAllBooks handles the retrieval of all books
// It calls the service to get all books and returns them as a JSON response
// If there's an error, it sends an internal server error response
func GetBooks(w http.ResponseWriter, r *http.Request) {
	books, err := services.GetAllBooksService() // Call the service to get all books
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to retrieve books")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, books) // Send a JSON response with the list of books
}

// GetBookByID handles the retrieval of a book by its ID
// It extracts the ID from the URL, calls the service to get the book, and returns it as a JSON response
// If there's an error, it sends an internal server error response
// If the book is not found, it sends a not found response
// If the ID is invalid, it sends a bad request response
func GetBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r) // Extract the URL parameters
	id := params["id"]    // Get the book ID from the URL parameters

	book, err := services.GetBookByIDService(id) // Call the service to get the book by ID
	if err != nil {
		utils.RespondWithError(w, http.StatusNotFound, "Book not found")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, book) // Send a JSON response with the book details
}


// UpdateBook handles the update of a book by its ID
// It extracts the ID from the URL, decodes the request body into a book struct, and calls the service to update the book
// It returns a JSON response with the updated book or an error message
// If there's an error, it sends an internal server error response
// If the book is not found, it sends a not found response
// If the ID is invalid, it sends a bad request response
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r) // Extract the URL parameters
	id := params["id"]    // Get the book ID from the URL parameters

	var book models.Book
	// Decode the request body into the book struct
	// This is where the actual decoding happens
	// If there's an error during decoding, send a bad request response
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	err := services.UpdateBookService(id, &book) // Call the service to update the book
	if err != nil {
		utils.RespondWithError(w, http.StatusNotFound, "Book not found")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, book) // Send a JSON response with the updated book details

}


// DeleteBook handles the deletion of a book by its ID
// It extracts the ID from the URL, calls the service to delete the book, and returns a success message or an error
// If there's an error, it sends an internal server error response
// If the book is not found, it sends a not found response
// If the ID is invalid, it sends a bad request response
// If the deletion is successful, it sends a no content response
// If the book is not found, it sends a not found response
// If the ID is invalid, it sends a bad request response
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r) // Extract the URL parameters
	id := params["id"]    // Get the book ID from the URL parameters

	err := services.DeleteBookService(id) // Call the service to delete the book
	if err != nil {
		utils.RespondWithError(w, http.StatusNotFound, "Book not found")
		return
	}

	utils.RespondWithJSON(w, http.StatusNoContent, nil) // Send a no content response
	// This indicates that the request was successful, but there's no content to return
}