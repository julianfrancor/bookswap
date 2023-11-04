package domain

import "errors"

var ErrBookNotFound = errors.New("book not found")

type Book struct {
    ID     int
    Title  string
    Author string
    Genre  string
    Status string
}
