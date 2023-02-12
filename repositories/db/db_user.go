package db

import (
	"api-books/entities"
	"database/sql"
)

type UserRepository struct {
	DB *sql.DB
}

// Create inserts a new user into the database
func (repo UserRepository) Create(user *entities.User) (string, error) {
	sql := `INSERT INTO users (username, email, password) VALUES (?, ?, ?) RETURNING id`
	var id string
	err := repo.DB.QueryRow(sql, user.Username, user.Email, user.Password).Scan(&id)

	if err != nil {
		return id, err
	}

	return id, nil
}

// SearchEmail returns the user by email
func (repo UserRepository) SearchEmail(email string) (*entities.User, error) {
	sql := "SELECT id, username, email, password FROM users WHERE email = ?"

	var result entities.User

	err := repo.DB.QueryRow(sql, email).
		Scan(&result.Id, &result.Username, &result.Email, &result.Password)

	if err != nil {
		return nil, err
	}

	return &result, nil
}
