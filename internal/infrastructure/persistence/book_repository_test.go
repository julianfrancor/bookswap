package persistence

import (
	"testing"

	"github.com/julianfrancor/bookswap/internal/domain"
)

func TestBookRepository_Create(t *testing.T) {
	repo := NewBookRepository()

	book := domain.Book{
		ID:     1,
		Title:  "The Catcher in the Rye",
		Author: "J.D. Salinger",
		Genre:  "Fiction",
		Status: "Available",
	}

	repo.Create(book)

	// Verify if the book has been created successfully
	createdBook, err := repo.GetByID(1)
	if err != nil {
		t.Errorf("Error fetching the created book: %v", err)
	}
	if createdBook != book {
		t.Errorf("Created book does not match the expected one")
	}
}

func TestBookRepository_Update(t *testing.T) {
	repo := NewBookRepository()

	originalBook := domain.Book{
		ID:     1,
		Title:  "The Catcher in the Rye",
		Author: "J.D. Salinger",
		Genre:  "Fiction",
		Status: "Available",
	}

	repo.Create(originalBook)

	updatedBook := domain.Book{
		ID:     1,
		Title:  "Updated Title",
		Author: "Updated Author",
		Genre:  "Updated Genre",
		Status: "Updated Status",
	}

	repo.Update(updatedBook)

	// Verify if the book has been updated successfully
	result, err := repo.GetByID(1)
	if err != nil {
		t.Errorf("Error fetching the updated book: %v", err)
	}
	if result != updatedBook {
		t.Errorf("Updated book does not match the expected one")
	}
}

func TestBookRepository_GetByID(t *testing.T) {
	repo := NewBookRepository()

	book := domain.Book{
		ID:     1,
		Title:  "The Catcher in the Rye",
		Author: "J.D. Salinger",
		Genre:  "Fiction",
		Status: "Available",
	}

	repo.Create(book)

	// Verify if the correct book is fetched by ID
	result, err := repo.GetByID(1)
	if err != nil {
		t.Errorf("Error fetching the book by ID: %v", err)
	}
	if result != book {
		t.Errorf("Fetched book does not match the expected one")
	}
}

func TestBookRepository_GetAll(t *testing.T) {
	repo := NewBookRepository()

	book1 := domain.Book{
		ID:     1,
		Title:  "Book 1",
		Author: "Author 1",
		Genre:  "Genre 1",
		Status: "Available",
	}

	book2 := domain.Book{
		ID:     2,
		Title:  "Book 2",
		Author: "Author 2",
		Genre:  "Genre 2",
		Status: "Not Available",
	}

	repo.Create(book1)
	repo.Create(book2)

	// Verify if all books are fetched
	result := repo.GetAll()
	if len(result) != 2 {
		t.Errorf("Expected 2 books, got %d", len(result))
	}
	if result[0] != book1 || result[1] != book2 {
		t.Errorf("Fetched books do not match the expected ones")
	}
}

func TestBookRepository_Delete(t *testing.T) {
	repo := NewBookRepository()

	book := domain.Book{
		ID:     1,
		Title:  "The Catcher in the Rye",
		Author: "J.D. Salinger",
		Genre:  "Fiction",
		Status: "Available",
	}

	repo.Create(book)
	repo.Delete(1)

	// Verify if the book has been deleted successfully
	_, err := repo.GetByID(1)
	if err != domain.ErrBookNotFound {
		t.Errorf("Expected an error when trying to fetch a deleted book")
	}
}
