package controller_test

import (
	controller "api-books/controllers"
	"api-books/repositories/mock"
	"api-books/usecases"
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

var mockedUserRepository *mock.MockUserRepository
var userService *usecases.UserService
var userHandler *controller.UserHandler

func setupUserController() {
	// Initilize the controller user
	mockedUserRepository = new(mock.MockUserRepository)

	userService = &usecases.UserService{
		Repository: mockedUserRepository,
	}

	userHandler = &controller.UserHandler{
		Service: userService,
	}
}

func TestRegisterHandler(t *testing.T) {
	setupUserController()
	tableStruct := []struct {
		Description    string
		Body           []byte
		ExpectedStatus int
	}{
		{
			Description: "Should return 201 created",
			Body: []byte(`
				{
					"username":"John Doe",
					"email":"john.doe@dev.com",
					"password":"123456"
				}
			`),
			ExpectedStatus: http.StatusCreated,
		},
		{
			Description: "Should return 400 with email already used",
			Body: []byte(`
				{
					"username":"John Doe",
					"email":"john.doe@dev.io",
					"password":"123456"
				}
			`),
			ExpectedStatus: http.StatusBadRequest,
		},
		{
			Description: "Should return 400 with email empty",
			Body: []byte(`
				{
					"username":"John Doe",
					"email":"",
					"password":"123456"
				}
			`),
			ExpectedStatus: http.StatusBadRequest,
		},
		{
			Description: "Should return 400 with password empty",
			Body: []byte(`
				{
					"username":"John Doe",
					"email":"john.doe@dev.io",
					"password":""
				}
			`),
			ExpectedStatus: http.StatusBadRequest,
		},
		{
			Description: "Should return 400 with body invalid",
			Body: []byte(`
				{
					"username":"John Doe",
					"email":"john.doe@dev.io",
					"password":""
				
			`),
			ExpectedStatus: http.StatusBadRequest,
		},
	}

	for _, test := range tableStruct {
		rec := httptest.NewRecorder()
		res := httptest.NewRequest(http.MethodPost, "/register", bytes.NewReader(test.Body))
		userHandler.Register(rec, res)

		response := rec.Result()
		if response.StatusCode != test.ExpectedStatus {
			t.Errorf("Expected status code %v, got %v", response.StatusCode, test.ExpectedStatus)
		}
	}
}

func TestAuthenticateHandler(t *testing.T) {
	setupUserController()
	tableStruct := []struct {
		Description    string
		Body           []byte
		ExpectedStatus int
	}{
		{
			Description: "Should login sucessully",
			Body: []byte(`
				{
					"email":"john.doe@dev.io",
					"password":"123456"
				}
			`),
			ExpectedStatus: http.StatusOK,
		},
		{
			Description: "Should return failed login, email not founded",
			Body: []byte(`
				{
					"email":"john.doe@dev.com",
					"password":"123456"
				}
			`),
			ExpectedStatus: http.StatusBadRequest,
		},
		{
			Description: "Should return failed login, email empty",
			Body: []byte(`
				{
					"email":"",
					"password":"123456"
				}
			`),
			ExpectedStatus: http.StatusBadRequest,
		},
		{
			Description: "Should return failed login, password empty",
			Body: []byte(`
				{
					"email":"john.doe@dev.io",
					"password":""
				}
			`),
			ExpectedStatus: http.StatusBadRequest,
		},
		{
			Description: "Should return failed login, password empty",
			Body: []byte(`
				{
					"email":"john.doe@dev.io",
					"password":"987654"
				}
			`),
			ExpectedStatus: http.StatusBadRequest,
		},
		{
			Description: "Should return error request body invalid",
			Body: []byte(`
				{
					"email":"john.doe@dev.io",
					"password":"987654"				
			`),
			ExpectedStatus: http.StatusBadRequest,
		},
	}

	for _, test := range tableStruct {
		rec := httptest.NewRecorder()
		res := httptest.NewRequest(http.MethodPost, "/login", bytes.NewReader(test.Body))
		userHandler.Login(rec, res)

		response := rec.Result()
		if response.StatusCode != test.ExpectedStatus {
			t.Errorf("Expected status code %v, got %v", response.StatusCode, test.ExpectedStatus)
		}
	}
}
