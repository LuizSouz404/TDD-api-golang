package main

import (
	"api-books/config"

	"github.com/gorilla/mux"
)

func main() {
	db, err := config.GetDatabaseInstance()
	if err != nil {
		panic(err)
	}

	defer db.Close()

	r := mux.NewRouter()
	server := config.Server{
		DB:     db,
		Router: r,
	}

	server.Run(":3333")
}
