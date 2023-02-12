package mock

import (
	"api-books/entities"
	"time"
)

type MockUserRepository struct {
	Expectation map[string]interface{}
}

// MOCK Create inserts a new user into the repository
func (m MockUserRepository) Create(user *entities.User) (string, error) {
	result, ok := m.Expectation["Create"]
	if ok {
		id, _ := result.(string)

		return id, nil
	}

	return "d8577c4c-96e7-43f1-a546-e89187dc2096", nil
}

// MOCK SearchEmail returns a row with specified email
func (m MockUserRepository) SearchEmail(email string) (*entities.User, error) {
	if email == "john.doe@dev.io" {
		return &entities.User{
			Id:        "d8577c4c-96e7-43f1-a546-e89187dc2096",
			Email:     "john.doe@dev.io",
			Username:  "John Doe",
			Password:  "123456",
			CreatedAt: time.Now(),
		}, nil
	}

	return nil, nil
}
