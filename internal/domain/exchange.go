package domain

import (
	"time"
)

// Exchange represents an exchange or trade between users.
type Exchange struct {
	ID        int       `json:"id"`
	User1ID   int       `json:"user1_id"`
	User2ID   int       `json:"user2_id"`
	Book1ID   int       `json:"book1_id"`
	Book2ID   int       `json:"book2_id"`
	CreatedAt time.Time `json:"created_at"`
}
