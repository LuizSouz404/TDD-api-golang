package db_test

import (
	"api-books/entities"
	dbRepo "api-books/repositories/db"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestCreateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	repo := dbRepo.UserRepository{
		DB: db,
	}

	user := &entities.User{
		Id:       "ca37b54f-dd4d-46a9-b64b-8c5b40c6ed35",
		Email:    "johndoe@dev.com",
		Username: "john doe",
		Password: "123456",
	}

	rows := sqlmock.NewRows([]string{"id"}).AddRow(user.Id)

	mock.ExpectQuery("^INSERT INTO users").WithArgs(user.Username, user.Email, user.Password).WillReturnRows(rows)

	got, err := repo.Create(user)
	if err != nil {
		t.Error(err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}

	if got != user.Id {
		t.Errorf("Create user expected %v, got %v, err %v", user.Id, got, err)
	}
}

func TestSearchEmail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	repo := dbRepo.UserRepository{
		DB: db,
	}

	user := &entities.User{
		Id:       "ca37b54f-dd4d-46a9-b64b-8c5b40c6ed35",
		Email:    "johndoe@dev.com",
		Username: "john doe",
		Password: "123456",
	}

	rows := sqlmock.NewRows([]string{"id", "username", "email", "password"}).
		AddRow(user.Id, user.Username, user.Email, user.Password)

	mock.ExpectQuery("^SELECT id, username, email, password FROM users").
		WithArgs(user.Email).
		WillReturnRows(rows)

	got, err := repo.SearchEmail(user.Email)
	if err != nil {
		t.Error(err)
	}

	if got.Email != user.Email {
		t.Errorf("SearchEmail expected %v, got %v, err %v", user.Email, got.Id, err)
	}
}
