package config

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	DB     *sql.DB
	Router *mux.Router
}

func (s Server) Run(port string) {
	http.ListenAndServe(port, s.Router)
}
