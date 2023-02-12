package db

import (
	"api-books/entities"
	"database/sql"
)

type BookRepository struct {
	DB *sql.DB
}

// Create inserts a new book into the database
func (repo BookRepository) Create(book *entities.Book) (string, error) {
	sql := "INSERT INTO books (name, isbn, price) VALUES (?, ?, ?) RETURNING id"

	var id string
	err := repo.DB.QueryRow(sql, book.Name, book.ISBN, book.Price).Scan(&id)
	if err != nil {
		return "", err
	}

	return id, nil
}

// Get returns all books into the database
func (repo BookRepository) Get() ([]entities.Book, error) {
	sql := "SELECT id, name, isbn, price FROM books"

	res, err := repo.DB.Query(sql)
	if err != nil {
		return nil, err
	}

	var books []entities.Book

	for res.Next() {
		var book entities.Book
		err := res.Scan(&book.Id, &book.Name, &book.ISBN, &book.Price)

		if err != nil {
			return nil, err
		}

		books = append(books, book)
	}

	return books, nil
}

// Update updates the book in the database
func (repo BookRepository) Update(book *entities.Book) error {
	sql := "UPDATE books SET name = ?, isbn = ?, price = ? WHERE id = ?"
	_, err := repo.DB.Exec(sql, book.Name, book.ISBN, book.Price, book.Id)

	if err != nil {
		return err
	}

	return nil
}

// RemoveBook removes a book from the database
func (repo BookRepository) Remove(id string) error {
	sql := "DELETE FROM books WHERE id = ?"
	_, err := repo.DB.Exec(sql, id)

	if err != nil {
		return err
	}

	return nil
}
