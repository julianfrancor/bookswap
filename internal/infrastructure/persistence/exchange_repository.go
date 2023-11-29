// Package persistence provides the persistence layer for the BookSwap application.
package persistence

import "github.com/julianfrancor/bookswap/internal/domain"

// ExchangeRepository implements the ExchangeRepository interface using an in-memory storage.
type ExchangeRepository struct {
	exchanges []domain.Exchange
}

// NewExchangeRepository creates a new ExchangeRepository instance.
func NewExchangeRepository() *ExchangeRepository {
	return &ExchangeRepository{
		exchanges: []domain.Exchange{},
	}
}

// Create adds a new exchange to the repository.
func (r *ExchangeRepository) Create(exchange domain.Exchange) {
	r.exchanges = append(r.exchanges, exchange)
}

// Update updates an existing exchange in the repository.
func (r *ExchangeRepository) Update(exchange domain.Exchange) {
	for i, e := range r.exchanges {
		if e.ID == exchange.ID {
			r.exchanges[i] = exchange
			break
		}
	}
}

// GetByID retrieves an exchange by its ID.
func (r *ExchangeRepository) GetByID(id int) (domain.Exchange, error) {
	for _, e := range r.exchanges {
		if e.ID == id {
			return e, nil
		}
	}
	return domain.Exchange{}, domain.ErrExchangeNotFound
}

// GetAll returns all exchanges in the repository.
func (r *ExchangeRepository) GetAll() []domain.Exchange {
	return r.exchanges
}

// Delete deletes an exchange by ID.
func (r *ExchangeRepository) Delete(id int) {
	for i, e := range r.exchanges {
		if e.ID == id {
			r.exchanges = append(r.exchanges[:i], r.exchanges[i+1:]...)
			break
		}
	}
}
