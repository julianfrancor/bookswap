package domain

import (
	"testing"
)

func TestNewUser(t *testing.T) {
	user := User{
		ID:       1,
		Username: "john_doe",
		Email:    "john@example.com",
		Books:    nil,
	}

	if user.ID != 1 {
		t.Errorf("Expected user ID to be 1, got %d", user.ID)
	}

	if user.Username != "john_doe" {
		t.Errorf("Expected username to be 'john_doe', got %s", user.Username)
	}

	if user.Email != "john@example.com" {
		t.Errorf("Expected email to be 'john@example.com', got %s", user.Email)
	}

	if user.Books != nil {
		t.Error("Expected user Books to be nil")
	}
}

func TestErrUserNotFound(t *testing.T) {
	expectedError := "user not found"

	if ErrUserNotFound.Error() != expectedError {
		t.Errorf("Expected error message '%s', got '%s'", expectedError, ErrUserNotFound.Error())
	}
}
