package models

import (
	"errors"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Book represents a book in the system
type Book struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Title       string             `json:"title" bson:"title" validate:"required,min=1,max=200"`
	Author      string             `json:"author" bson:"author" validate:"required,min=1,max=100"`
	Description string             `json:"description" bson:"description" validate:"max=1000"`
	ISBN        string             `json:"isbn" bson:"isbn" validate:"required,isbn"`
	CreatedAt   time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at" bson:"updated_at"`
}

// Validate performs validation on the Book struct
func (b *Book) Validate() error {
	if strings.TrimSpace(b.Title) == "" {
		return errors.New("title is required")
	}
	if len(b.Title) > 200 {
		return errors.New("title must be less than 200 characters")
	}

	if strings.TrimSpace(b.Author) == "" {
		return errors.New("author is required")
	}
	if len(b.Author) > 100 {
		return errors.New("author name must be less than 100 characters")
	}

	if len(b.Description) > 1000 {
		return errors.New("description must be less than 1000 characters")
	}

	if strings.TrimSpace(b.ISBN) == "" {
		return errors.New("ISBN is required")
	}
	if !isValidISBN(b.ISBN) {
		return errors.New("invalid ISBN format")
	}

	return nil
}

// isValidISBN checks if the ISBN is valid
func isValidISBN(isbn string) bool {
	// Remove any spaces or hyphens
	isbn = strings.ReplaceAll(isbn, " ", "")
	isbn = strings.ReplaceAll(isbn, "-", "")

	// Check length (ISBN-10 or ISBN-13)
	if len(isbn) != 10 && len(isbn) != 13 {
		return false
	}

	// Basic format check
	for _, c := range isbn {
		if !((c >= '0' && c <= '9') || c == 'X') {
			return false
		}
	}

	return true
}

// BeforeCreate sets the CreatedAt and UpdatedAt timestamps
func (b *Book) BeforeCreate() {
	now := time.Now()
	b.CreatedAt = now
	b.UpdatedAt = now
}

// BeforeUpdate sets the UpdatedAt timestamp
func (b *Book) BeforeUpdate() {
	b.UpdatedAt = time.Now()
}
