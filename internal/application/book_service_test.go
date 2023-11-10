package application

import (
	"errors"
	"testing"

	"github.com/julianfrancor/bookswap/internal/domain"
	"github.com/julianfrancor/bookswap/internal/infrastructure/persistence"
)

func TestBookService(t *testing.T) {
	// Create the repositories and services
	bookRepository := persistence.NewBookRepository()
	userRepository := persistence.NewUserRepository()
	bookService := NewBookService(*bookRepository)
	userService := NewUserService(*userRepository)

	// Create a User
	userRequest := CreateUserRequest{
		Username: "julian",
		Email:    "julian@bookswap.com",
	}
	userService.CreateUser(userRequest)
	userID := userService.GetAllUsers()[0].ID

	// Test Create a book request
	bookRequest := CreateBookRequest{
		UserID: userID,
		Title:  "Book 1",
		Author: "Author 1",
		Genre:  "Genre 1",
		Status: "Available",
	}
	bookService.CreateBook(bookRequest, userService)

	// Verify if the book has been created successfully
	createdBook, err := bookService.GetBookByID(1)
	if err != nil {
		t.Errorf("Error fetching the book: %v", err)
	}
	if createdBook != (domain.Book{
		ID:     1,
		UserID: userID,
		Title:  "Book 1",
		Author: "Author 1",
		Genre:  "Genre 1",
		Status: "Available",
	}) {
		t.Errorf("Created book does not match the expected one")
	}

	// Test Update a book
	updatedBookRequest := UpdateBookRequest{
		ID:     1,
		UserID: userID,
		Title:  "Modified_Book1",
		Author: "Author1",
		Genre:  "Genre1",
		Status: "Exchanged",
	}
	bookService.UpdateBook(updatedBookRequest)

	// Verify if the book has been updated successfully
	updatedBookResult, err := bookService.GetBookByID(1)

	if err != nil {
		t.Errorf("Error fetching the updated book: %v", err)
	}

	if updatedBookResult != (domain.Book{
		ID:     1,
		UserID: userID,
		Title:  "Modified_Book1",
		Author: "Author1",
		Genre:  "Genre1",
		Status: "Exchanged",
	}) {
		t.Errorf("Updated book does not match the expected one")
	}

	// Test Delete a book
	bookService.DeleteBook(1)

	// Verify if the book has been deleted successfully
	_, err = bookService.GetBookByID(1)
	if !errors.Is(err, domain.ErrBookNotFound) {
		t.Errorf("Expected an error when trying to fetch a deleted book")
	}
}
