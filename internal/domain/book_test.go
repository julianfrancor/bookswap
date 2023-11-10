package domain

import (
	"testing"
)

func TestNewBook(t *testing.T) {
	book := Book{
		ID:     1,
		UserID: 1,
		Title:  "The Catcher in the Rye",
		Author: "J.D. Salinger",
		Genre:  "Fiction",
		Status: "Available",
	}

	if book.ID != 1 {
		t.Errorf("Expected book ID to be 1, got %d", book.ID)
	}

	if book.UserID != 1 {
		t.Errorf("Expected book UserID to be 1, got %d", book.UserID)
	}

	if book.Title != "The Catcher in the Rye" {
		t.Errorf("Expected book Title to be 'The Catcher in the Rye', got %s", book.Title)
	}

	if book.Author != "J.D. Salinger" {
		t.Errorf("Expected book Author to be 'J.D. Salinger', got %s", book.Author)
	}

	if book.Genre != "Fiction" {
		t.Errorf("Expected book Genre to be 'Fiction', got %s", book.Genre)
	}

	if book.Status != "Available" {
		t.Errorf("Expected book Status to be 'Available', got %s", book.Status)
	}
}

func TestErrBookNotFound(t *testing.T) {
	expectedError := "book not found"

	if ErrBookNotFound.Error() != expectedError {
		t.Errorf("Expected error message '%s', got '%s'", expectedError, ErrBookNotFound.Error())
	}
}
