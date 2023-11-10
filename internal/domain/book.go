package domain

import "errors"

var ErrBookNotFound = errors.New("book not found")

// Book represents a book entity.
type Book struct {
	ID     int    `json:"id"`
	UserID int    `json:"user_id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Genre  string `json:"genre"`
	Status string `json:"status"`
}
