package domain

import "errors"

var ErrBookNotFound = errors.New("book not found")

// Book represents a book entity.
type Book struct {
	// ID is the unique identifier for the book.
	// Example: 1
	ID int `json:"id"`

	// UserID is the ID of the user associated with the book.
	// Example: 1
	UserID int `json:"user_id"`

	// Title is the title of the book.
	// Example: "The Catcher in the Rye"
	Title string `json:"title"`

	// Author is the author of the book.
	// Example: "J.D. Salinger"
	Author string `json:"author"`

	// Genre is the genre of the book.
	// Example: "Fiction"
	Genre string `json:"genre"`

	// Status represents the availability status of the book.
	// Example: "Available"
	Status string `json:"status"`
}
