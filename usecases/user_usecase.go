package usecases

import (
	"api-books/entities"
	"api-books/repositories"
	"errors"
)

var (
	// ErrorPasswordRequired the field password must be filled
	ErrorPasswordRequired = errors.New("the password is required")
	// ErrorEmailRequired the field email must be filled
	ErrorEmailRequired = errors.New("the email is required")
	// ErrorEmailAlreadyExists the field email must be unique
	ErrorEmailAlreadyExists = errors.New("the email must be unique")
	// ErrorPasswordInvalid the password is invalid
	ErrorUserNotFound = errors.New("email or password is invalid")
)

type UserService struct {
	Repository repositories.SQLUserRepository
}

func (service UserService) Create(user *entities.User) (string, error) {
	if len(user.Password) == 0 {
		return "", ErrorPasswordRequired
	}

	if len(user.Email) == 0 {
		return "", ErrorEmailRequired
	}

	userFounded, err := service.Repository.SearchEmail(user.Email)
	if err != nil {
		return "", err
	}

	if userFounded != nil {
		return "", ErrorEmailAlreadyExists
	}

	return service.Repository.Create(user)
}

func (service UserService) Authenticate(email string, password string) error {
	if len(email) == 0 {
		return ErrorEmailRequired
	}

	if len(password) == 0 {
		return ErrorPasswordRequired
	}

	userFounded, err := service.Repository.SearchEmail(email)
	if err != nil {
		return err
	}

	if userFounded == nil {
		return ErrorUserNotFound
	}

	if userFounded.Password != password {
		return ErrorUserNotFound
	}

	return nil
}
