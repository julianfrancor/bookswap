// internal/application/exchange_service.go

package application

import (
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
	repository     persistence.ExchangeRepository
	bookRepository persistence.BookRepository
}

// NewExchangeService creates a new ExchangeService instance.
func NewExchangeService(repository persistence.ExchangeRepository, bookRepository persistence.BookRepository) *ExchangeService {
	return &ExchangeService{
		repository:     repository,
		bookRepository: bookRepository,
	}
}

// ExchangeBooks performs the exchange of a book from one user to another.
func (s *ExchangeService) ExchangeBooks(user1ID, user2ID, book1ID int, book2ID int) error {
	// Check if the book1 exists
	book1, err := s.bookRepository.GetByID(book1ID)
	if err != nil {
		return err
	}

	// Check if the book1 exists
	book2, err := s.bookRepository.GetByID(book2ID)
	if err != nil {
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
	s.bookRepository.Update(book1)

	// Transfer the book2 to user1
	book2.UserID = user1ID
	s.bookRepository.Update(book2)

	// Save the exchange
	s.repository.Create(exchange)

	return nil
}

// GetExchangeByID retrieves an exchange by its ID.
func (s *ExchangeService) GetExchangeByID(id int) (domain.Exchange, error) {
	return s.repository.GetByID(id)
}
