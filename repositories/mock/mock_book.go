package mock

import "api-books/entities"

type MockBookRepository struct{}

// MOCK Create inserts the new book into database
func (m MockBookRepository) Create(book *entities.Book) (string, error) {
	return book.Id, nil
}

// MOCK Get returns every book in database
func (m MockBookRepository) Get() ([]entities.Book, error) {
	return []entities.Book{}, nil
}

// Update changes the book in database
func (m MockBookRepository) Update(book *entities.Book) error {
	return nil
}

// MOCK Remove delete the book from database
func (m MockBookRepository) Remove(Id string) error {
	return nil
}
