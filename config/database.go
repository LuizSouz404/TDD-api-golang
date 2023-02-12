package config

import (
	"database/sql"
	"fmt"
	"os"
)

var conn *sql.DB

var (
	database = os.Getenv("DATABASE")
	user     = os.Getenv("USER")
	password = os.Getenv("PASSWORD")
	host     = os.Getenv("HOST")
	port     = os.Getenv("PORT")
)

// GetDatabaseInstance instantiates a new database singleton
func GetDatabaseInstance() (conn *sql.DB, err error) {
	if conn == nil {
		uri := fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			host, port, user, password, database)

		db, err := sql.Open("postgres", uri)

		if err != nil {
			return nil, err
		}

		conn = db

		return conn, nil
	}

	return conn, nil
}
