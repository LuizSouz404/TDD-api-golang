package usecases

import "api-books/entities"

// UserUsecase represents the bussiness logic for user
type UserUsecase interface {
	Create(*entities.User) (string, error)
	Authenticate(email string, password string) error
}

// BookUsecase represents the bussiness logic for book
type BookUsecase interface {
	Create(*entities.Book) (string, error)
	Get() ([]entities.Book, error)
	Update(*entities.Book) error
	Remove(string) error
}
