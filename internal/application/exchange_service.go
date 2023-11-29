package application

import (
	"fmt"
	"github.com/julianfrancor/bookswap/internal/infrastructure/persistence"
	"time"

	"github.com/julianfrancor/bookswap/internal/domain"
)

// CreateExchangeRequest represents the request structure for creating an exchange.
type CreateExchangeRequest struct {
	User1ID int `json:"user1_id"`
	User2ID int `json:"user2_id"`
	Book1ID int `json:"book1_id"`
	Book2ID int `json:"book2_id"`
}

// ExchangeService provides the business logic for managing exchanges.
type ExchangeService struct {
	repository persistence.ExchangeRepository
}

// NewExchangeService creates a new ExchangeService instance.
func NewExchangeService(repository persistence.ExchangeRepository) *ExchangeService {
	return &ExchangeService{
		repository: repository,
	}
}

// ExchangeBooks performs the exchange of a book from one user to another.
func (s *ExchangeService) ExchangeBooks(exchangeRequest CreateExchangeRequest, bookService *BookService, userService *UserService) error {
	user1ID := exchangeRequest.User1ID
	user2ID := exchangeRequest.User2ID
	book1ID := exchangeRequest.Book1ID
	book2ID := exchangeRequest.Book2ID

	// Check if the book1 exists
	book1, err := bookService.GetBookByID(book1ID)
	if err != nil {
		fmt.Println("Book1 is not in DB")
		return err
	}

	// Check if the book1 exists
	book2, err := bookService.GetBookByID(book2ID)
	if err != nil {
		fmt.Println("Book2 is not in DB")
		return err
	}

	// Check if the books belongs to its users
	if book1.UserID != user1ID || book2.UserID != user2ID {
		return domain.ErrBookNotOwnedByUser
	}

	// Create the exchange
	exchange := domain.Exchange{
		User1ID:   user1ID,
		User2ID:   user2ID,
		Book1ID:   book1ID,
		Book2ID:   book2ID,
		CreatedAt: time.Now(),
	}

	// Transfer the book1 to user2
	book1.UserID = user2ID
	updateBook1Request := bookService.ConvertToUpdateBookRequest(book1)
	bookService.UpdateBook(updateBook1Request)

	// Transfer the book2 to user1
	book2.UserID = user1ID
	updateBook2Request := bookService.ConvertToUpdateBookRequest(book2)
	bookService.UpdateBook(updateBook2Request)

	// Update users shelves
	userService.UpdateUserBooksOnExchange(user1ID, book1ID, book2ID, *bookService)
	userService.UpdateUserBooksOnExchange(user2ID, book2ID, book1ID, *bookService)

	// Save the exchange
	s.repository.Create(exchange)

	return nil
}

// GetExchangeByID retrieves an exchange by its ID.
func (s *ExchangeService) GetExchangeByID(id int) (domain.Exchange, error) {
	return s.repository.GetByID(id)
}
