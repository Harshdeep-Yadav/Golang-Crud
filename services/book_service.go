package services

import (
	"github.com/harsh/book_crud/models"
	"github.com/harsh/book_crud/repository"
)

func CreateBookService(book *models.Book) (interface{}, error) {
    return repository.CreateBook(book)
}

func GetAllBooksService() ([]models.Book, error) {
    return repository.GetAllBooks()
}

func GetBookByIDService(id string) (*models.Book, error) {
    return repository.GetBookByID(id)
}

func UpdateBookService(id string, book *models.Book) error {
    return repository.UpdateBook(id, book)
}

func DeleteBookService(id string) error {
    return repository.DeleteBook(id)
}
