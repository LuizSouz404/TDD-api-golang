package repositories

import "api-books/entities"

// SQLUserRepository is a contract of user repositories
type SQLUserRepository interface {
	Create(*entities.User) (string, error)
	SearchEmail(string) (*entities.User, error)
}

// SQLBookRepository is a contract of the data provider
type SQLBookRepository interface {
	Create(*entities.Book) (string, error)
	Get() ([]entities.Book, error)
	Update(*entities.Book) error
	Remove(string) error
}
