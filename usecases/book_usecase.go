package usecases

import (
	"api-books/entities"
	"api-books/repositories"
	"errors"
)

var (
	// ErrorBookNameRequired must be used to validation
	ErrorBookNameRequired = errors.New("the name of book is required")
	// ErrorBookPriceRequired must be used to validation
	ErrorBookPriceRequired = errors.New("the price of book must be greater than zero")
)

// BookService represents the book business logic
type BookService struct {
	Repository repositories.SQLBookRepository
}

// Create inserts a new book into the database
func (service BookService) Create(book *entities.Book) (string, error) {
	if book.Name == "" {
		return "", ErrorBookNameRequired
	}

	if book.Price <= 0 {
		return "", ErrorBookPriceRequired
	}

	return service.Repository.Create(book)
}

// Get returns all the books in the database
func (service BookService) Get() ([]entities.Book, error) {
	return service.Repository.Get()
}

// Update changes the book in the database
func (service BookService) Update(book *entities.Book) error {
	if book.Name == "" {
		return ErrorBookNameRequired
	}

	if book.Price <= 0 {
		return ErrorBookPriceRequired
	}

	return service.Repository.Update(book)
}

// Remove removes a book
func (service BookService) Remove(id string) error {
	return service.Repository.Remove(id)
}
