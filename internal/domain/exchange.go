package domain

import (
	"errors"
	"time"
)

var ErrExchangeNotFound = errors.New("exchange not found")

// ExchangeStatus represents the status of an exchange.
type ExchangeStatus string

const (
	// ExchangeStatusPending represents the pending status of an exchange.
	ExchangeStatusPending ExchangeStatus = "Pending"
	// ExchangeStatusCompleted represents the completed status of an exchange.
	ExchangeStatusCompleted ExchangeStatus = "Completed"
)

// Exchange represents an exchange or trade between users.
type Exchange struct {
	ID          int            `json:"id"`
	User1ID     int            `json:"user1_id"`
	User2ID     int            `json:"user2_id"`
	Book1ID     int            `json:"book1_id"`
	Book2ID     int            `json:"book2_id"`
	Status      ExchangeStatus `json:"status"`
	CreatedAt   time.Time      `json:"created_at"`
	CompletedAt time.Time      `json:"completed_at"`
}

// ExchangeRepository defines the interface for interacting with exchanges.
type ExchangeRepository interface {
	Create(exchange Exchange)
	Update(exchange Exchange)
	GetByID(id int) (Exchange, error)
	GetAll() []Exchange
	Delete(id int)
}
