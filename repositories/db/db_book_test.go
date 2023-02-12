package db_test

import (
	"api-books/entities"
	dbRepo "api-books/repositories/db"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestCreateBook(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()
	book := &entities.Book{
		Id:    "0de20dd1-4b6d-4240-b3fd-8de433b663f6",
		Name:  "Insgnia",
		ISBN:  "8576835088",
		Price: 3400,
	}

	repo := dbRepo.BookRepository{
		DB: db,
	}

	rows := sqlmock.NewRows([]string{"id"}).AddRow("0de20dd1-4b6d-4240-b3fd-8de433b663f6")

	mock.ExpectQuery("^INSERT INTO books").WithArgs(book.Name, book.ISBN, book.Price).WillReturnRows(rows)

	got, err := repo.Create(book)
	if err != nil {
		t.Error(err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}

	if got != book.Id {
		t.Errorf("create expect: %s, got: %s, err: %v", book.Id, got, err)
	}
}

func TestUpdateBook(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	repo := dbRepo.BookRepository{
		DB: db,
	}

	changedBook := &entities.Book{
		Id:    "0de20dd1-4b6d-4240-b3fd-8de433b663f6",
		Name:  "Insgnia",
		ISBN:  "8576835088",
		Price: 4500,
	}

	mock.ExpectExec("^UPDATE books").
		WithArgs(changedBook.Name, changedBook.ISBN, changedBook.Price, changedBook.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.Update(changedBook)
	if err != nil {
		t.Error(err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}
}

func TestGetBook(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	repo := dbRepo.BookRepository{
		DB: db,
	}

	book := &entities.Book{
		Id:    "0de20dd1-4b6d-4240-b3fd-8de433b663f6",
		Name:  "Insgnia",
		ISBN:  "8576835088",
		Price: 4500,
	}

	rows := sqlmock.NewRows([]string{"id", "name", "isbn", "price"}).
		AddRow(book.Id, book.Name, book.ISBN, book.Price)

	mock.ExpectQuery("^SELECT id, name, isbn, price FROM books").WillReturnRows(rows)

	_, err = repo.Get()
	if err != nil {
		t.Error(err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}
}

func TestRemoveBook(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	repo := dbRepo.BookRepository{
		DB: db,
	}

	book := &entities.Book{
		Id: "ca37b54f-dd4d-46a9-b64b-8c5b40c6ed35",
	}

	mock.ExpectExec("^DELETE FROM books").WithArgs(book.Id).WillReturnResult(sqlmock.NewResult(0, 1))

	err = repo.Remove(book.Id)
	if err != nil {
		t.Error(err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}
}
