package controller

import (
	"api-books/entities"
	"api-books/usecases"
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var (
	jwtToken = os.Getenv("JWT_TOKEN")
)

// UserHandler represents a controller of a user
type UserHandler struct {
	Service usecases.UserUsecase
}

// Register inserts a new user
func (handler UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var user entities.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := handler.Service.Create(&user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user.Id = id

	response := make(map[string]interface{})
	response["data"] = user

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

// Login authenticates the user
func (handler UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	credentials := make(map[string]string)
	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	email := credentials["email"]
	password := credentials["password"]

	err := handler.Service.Authenticate(email, password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Duration(24) * time.Hour),
		"iat":   time.Now().Unix(),
	})

	tokenString, err := token.SignedString([]byte(jwtToken))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := map[string]interface{}{
		"token": tokenString,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
