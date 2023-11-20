package application

import (
	"github.com/julianfrancor/bookswap/internal/domain"
	"github.com/julianfrancor/bookswap/internal/infrastructure/persistence"
)

// CreateBookRequest represents the request body for creating a new book.
type CreateBookRequest struct {
	// UserID is the ID of the user associated with the book.
	// Example: 1
	UserID int `json:"userID"`

	// Title of the book.
	// Example: "The Catcher in the Rye"
	Title string `json:"title"`

	// Author of the book.
	// Example: "J.D. Salinger"
	Author string `json:"author"`

	// Genre of the book.
	// Example: "Fiction"
	Genre string `json:"genre"`

	// Status of the book.
	// Example: "Available"
	Status string `json:"status"`
}

type UpdateBookRequest struct {
	ID     int    `json:"id"`
	UserID int    `json:"userID"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Genre  string `json:"genre"`
	Status string `json:"status"`
}

type BookService struct {
	repository persistence.BookRepository
}

func NewBookService(repository persistence.BookRepository) *BookService {
	return &BookService{
		repository: repository,
	}
}

func (s *BookService) CreateBook(request CreateBookRequest, userService *UserService) error {
	// We get the user with the matching ID
	user, err := userService.GetUserByID(request.UserID)
	if err != nil {
		// Handle when user is not found
		return domain.ErrUserNotFound
	}

	book := domain.Book{
		ID:     len(s.repository.GetAll()) + 1,
		UserID: request.UserID,
		Title:  request.Title,
		Author: request.Author,
		Genre:  request.Genre,
		Status: request.Status,
	}
	s.repository.Create(book)

	// We add the book to the user's book list
	user.Books = append(user.Books, book)

	updateUserRequest := UpdateUserRequest{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Books:    user.Books,
	}

	userService.UpdateUser(updateUserRequest)
	return nil
}

func (s *BookService) UpdateBook(request UpdateBookRequest) {
	book := domain.Book{
		ID:     request.ID,
		UserID: request.UserID,
		Title:  request.Title,
		Author: request.Author,
		Genre:  request.Genre,
		Status: request.Status,
	}
	s.repository.Update(book)
}

func (s *BookService) GetBookByID(id int) (domain.Book, error) {
	return s.repository.GetByID(id)
}

func (s *BookService) GetAllBooks() []domain.Book {
	return s.repository.GetAll()
}

func (s *BookService) DeleteBook(id int) {
	s.repository.Delete(id)
}

// ConvertToUpdateBookRequest converts a domain.Book to an UpdateBookRequest.
func (s *BookService) ConvertToUpdateBookRequest(book domain.Book) UpdateBookRequest {
	return UpdateBookRequest{
		ID:     book.ID,
		UserID: book.UserID,
		Title:  book.Title,
		Author: book.Author,
		Genre:  book.Genre,
		Status: book.Status,
	}
}
