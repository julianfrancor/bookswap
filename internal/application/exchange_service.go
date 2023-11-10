// Package application provides the application services for the BookSwap application.
package application

import (
	"time"

	"github.com/julianfrancor/bookswap/internal/domain"
)

// ExchangeService provides the business logic for managing exchanges.
type ExchangeService struct {
	repository domain.ExchangeRepository
}

// NewExchangeService creates a new ExchangeService instance.
func NewExchangeService(repository domain.ExchangeRepository) *ExchangeService {
	return &ExchangeService{
		repository: repository,
	}
}

// CreateExchangeRequest represents the request structure for creating an exchange.
type CreateExchangeRequest struct {
	User1ID int `json:"user1_id"`
	User2ID int `json:"user2_id"`
	Book1ID int `json:"book1_id"`
	Book2ID int `json:"book2_id"`
}

// UpdateExchangeRequest represents the request structure for updating an exchange.
type UpdateExchangeRequest struct {
	ID        int    `json:"id"`
	Status    string `json:"status"`
	Completed bool   `json:"completed"`
}

// CreateExchange creates a new exchange.
func (s *ExchangeService) CreateExchange(request CreateExchangeRequest) {
	exchange := domain.Exchange{
		User1ID:   request.User1ID,
		User2ID:   request.User2ID,
		Book1ID:   request.Book1ID,
		Book2ID:   request.Book2ID,
		Status:    domain.ExchangeStatusPending,
		CreatedAt: time.Now(),
	}

	s.repository.Create(exchange)
}

// UpdateExchange updates an existing exchange.
func (s *ExchangeService) UpdateExchange(request UpdateExchangeRequest) {
	existingExchange, err := s.repository.GetByID(request.ID)
	if err != nil {
		// Handle error, exchange not found
		return
	}

	existingExchange.Status = domain.ExchangeStatus(request.Status)
	if request.Completed {
		existingExchange.CompletedAt = time.Now()
	}

	s.repository.Update(existingExchange)
}

// GetExchangeByID retrieves details of an exchange by ID.
func (s *ExchangeService) GetExchangeByID(id int) (domain.Exchange, error) {
	return s.repository.GetByID(id)
}

// DeleteExchange deletes an exchange by ID.
func (s *ExchangeService) DeleteExchange(id int) {
	s.repository.Delete(id)
}
