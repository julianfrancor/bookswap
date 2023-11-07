package domain

import "errors"

var ErrUserNotFound = errors.New("book not found")

type User struct {
	ID       int
	Username string
	Email    string
	Books    []Book
}
