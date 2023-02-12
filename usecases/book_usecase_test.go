package usecases_test

import (
	"api-books/entities"
	"api-books/repositories/mock"
	"api-books/usecases"
	"testing"
)

var mockedBookRepository *mock.MockBookRepository
var bookService *usecases.BookService

func setupBook() {
	// Initialize what is necessary to test usacase book
	mockedBookRepository = new(mock.MockBookRepository)
	bookService = &usecases.BookService{
		Repository: mockedBookRepository,
	}
}

func TestCreateBook(t *testing.T) {
	setupBook()

	testTable := []struct {
		Description   string
		Book          entities.Book
		ExpectedError error
		ExpectedId    string
	}{
		{
			Description: "Should Create book",
			Book: entities.Book{
				Id:    "d8577c4c-96e7-43f1-a546-e89187dc2096",
				Name:  "example",
				ISBN:  "ABSFEFE-QEXFOGKGK-MWIEFJFE",
				Price: 1200,
			},
			ExpectedError: nil,
			ExpectedId:    "d8577c4c-96e7-43f1-a546-e89187dc2096",
		},
		{
			Description: "Should require name book",
			Book: entities.Book{
				Id:    "d8577c4c-96e7-43f1-a546-e89187dc2096",
				Name:  "",
				ISBN:  "ABSFEFE-QEXFOGKGK-MWIEFJFE",
				Price: 1200,
			},
			ExpectedError: usecases.ErrorBookNameRequired,
			ExpectedId:    "",
		},
		{
			Description: "Should require price book greater than zero",
			Book: entities.Book{
				Id:    "d8577c4c-96e7-43f1-a546-e89187dc2096",
				Name:  "example",
				ISBN:  "ABSFEFE-QEXFOGKGK-MWIEFJFE",
				Price: 0,
			},
			ExpectedError: usecases.ErrorBookPriceRequired,
			ExpectedId:    "",
		},
	}

	for _, test := range testTable {
		t.Run(test.Description, func(t *testing.T) {
			id, err := bookService.Create(&test.Book)
			if err != test.ExpectedError {
				t.Errorf("Expected error: %v, got: %v", err, test.ExpectedError)
			}

			if id != test.ExpectedId {
				t.Errorf("Expected id: %v, got: %v", id, test.ExpectedId)
			}
		})
	}
}

func TestGetBook(t *testing.T) {
	setupBook()
	_, err := bookService.Get()
	if err != nil {
		t.Error(err)
	}
}

func TestUpdateBook(t *testing.T) {
	setupBook()
	testTable := []struct {
		Description   string
		Book          entities.Book
		ExpectedError error
	}{
		{
			Description: "Should Update book",
			Book: entities.Book{
				Id:    "d8577c4c-96e7-43f1-a546-e89187dc2096",
				Name:  "example",
				ISBN:  "ABSFEFE-QEXFOGKGK-MWIEFJFE",
				Price: 1200,
			},
			ExpectedError: nil,
		},
		{
			Description: "Should require name book",
			Book: entities.Book{
				Id:    "d8577c4c-96e7-43f1-a546-e89187dc2096",
				Name:  "",
				ISBN:  "ABSFEFE-QEXFOGKGK-MWIEFJFE",
				Price: 1200,
			},
			ExpectedError: usecases.ErrorBookNameRequired,
		},
		{
			Description: "Should require price book greater than zero",
			Book: entities.Book{
				Id:    "d8577c4c-96e7-43f1-a546-e89187dc2096",
				Name:  "example",
				ISBN:  "ABSFEFE-QEXFOGKGK-MWIEFJFE",
				Price: 0,
			},
			ExpectedError: usecases.ErrorBookPriceRequired,
		},
	}

	for _, test := range testTable {
		t.Run(test.Description, func(t *testing.T) {
			err := bookService.Update(&test.Book)
			if err != test.ExpectedError {
				t.Errorf("Expected error: %v, got: %v", err, test.ExpectedError)
			}
		})
	}
}

func TestRemoveBook(t *testing.T) {
	setupBook()
	err := bookService.Remove("d8577c4c-96e7-43f1-a546-e89187dc2096")
	if err != nil {
		t.Error(err)
	}
}
