package usecases_test

import (
	"api-books/entities"
	"api-books/repositories/mock"
	"api-books/usecases"
	"testing"
)

var mockedUserRepository *mock.MockUserRepository
var userService *usecases.UserService

func setupUser() {
	// Initialize what is necessary to test usecase user
	mockedUserRepository = new(mock.MockUserRepository)
	userService = &usecases.UserService{
		Repository: mockedUserRepository,
	}
}

func TestCreateUser(t *testing.T) {
	setupUser()
	tableStruct := []struct {
		Description   string
		User          entities.User
		ExpectedError error
		ExpectedId    string
	}{
		{
			Description: "Should create a user",
			User: entities.User{
				Id:       "d8577c4c-96e7-43f1-a546-e89187dc2096",
				Email:    "john.doe@dev.com",
				Username: "John Doe",
				Password: "123456",
			},
			ExpectedError: nil,
			ExpectedId:    "d8577c4c-96e7-43f1-a546-e89187dc2096",
		},
		{
			Description: "Should require email user",
			User: entities.User{
				Id:       "d8577c4c-96e7-43f1-a546-e89187dc2096",
				Email:    "",
				Username: "John Doe",
				Password: "123456",
			},
			ExpectedError: usecases.ErrorEmailRequired,
			ExpectedId:    "",
		},
		{
			Description: "Should require password user",
			User: entities.User{
				Id:       "d8577c4c-96e7-43f1-a546-e89187dc2096",
				Email:    "john.doe@dev.com",
				Username: "John Doe",
				Password: "",
			},
			ExpectedError: usecases.ErrorPasswordRequired,
			ExpectedId:    "",
		},
		{
			Description: "Should return error already email used",
			User: entities.User{
				Id:       "d8577c4c-96e7-43f1-a546-e89187dc2096",
				Email:    "john.doe@dev.io",
				Username: "John Doe",
				Password: "123456",
			},
			ExpectedError: usecases.ErrorEmailAlreadyExists,
			ExpectedId:    "",
		},
	}

	for _, test := range tableStruct {
		t.Run(test.Description, func(t *testing.T) {
			id, err := userService.Create(&test.User)
			if err != test.ExpectedError {
				t.Errorf("Expected error %v, got %v", err, test.ExpectedError)
			}

			if id != test.ExpectedId {
				t.Errorf("Expected id %v, got %v", id, test.ExpectedId)
			}
		})
	}
}

func TestAuthenticateUser(t *testing.T) {
	setupUser()
	tableStruct := []struct {
		Description   string
		User          entities.User
		ExpectedError error
	}{
		{
			Description: "Should authenticate user",
			User: entities.User{
				Email:    "john.doe@dev.io",
				Password: "123456",
			},
			ExpectedError: nil,
		},
		{
			Description: "Should require email",
			User: entities.User{
				Email:    "",
				Password: "123456",
			},
			ExpectedError: usecases.ErrorEmailRequired,
		},
		{
			Description: "Should require password",
			User: entities.User{
				Email:    "john.doe@dev.io",
				Password: "",
			},
			ExpectedError: usecases.ErrorPasswordRequired,
		},
		{
			Description: "Password is invalid",
			User: entities.User{
				Email:    "john.doe@dev.io",
				Password: "987654",
			},
			ExpectedError: usecases.ErrorUserNotFound,
		},
		{
			Description: "Email nout found",
			User: entities.User{
				Email:    "john.doe@dev.com",
				Password: "123456",
			},
			ExpectedError: usecases.ErrorUserNotFound,
		},
	}

	for _, test := range tableStruct {
		err := userService.Authenticate(test.User.Email, test.User.Password)
		if err != test.ExpectedError {
			t.Errorf("Expected error %v, got %v", err, test.ExpectedError)
		}
	}
}
