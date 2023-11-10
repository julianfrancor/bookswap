package domain

import "errors"

var ErrUserNotFound = errors.New("user not found")

// User represents a user entity.
// swagger:model
type User struct {
	// ID is the unique identifier for the user.
	// Example: 1
	// required: true
	ID int `json:"id"`

	// Username is the username of the user.
	// Example: "johndoe"
	// required: true
	Username string `json:"username"`

	// Email is the email address of the user.
	// Example: "john@example.com"
	// required: true
	Email string `json:"email"`

	// Books is the list of books associated with the user.
	// Example: [{"id": 1, "title": "Book 1", "author": "Author 1", "genre": "Genre 1", "status": "Available"}]
	Books []Book `json:"books"`
}
