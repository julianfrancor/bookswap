package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julianfrancor/bookswap/internal/application"
	"github.com/gorilla/mux"
)

func CreateBookHandler(bookService *application.BookService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book application.CreateBookRequest
		err := json.NewDecoder(r.Body).Decode(&book)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		bookService.CreateBook(book)
		w.WriteHeader(http.StatusCreated)
	}
}

func GetBookHandler(bookService *application.BookService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		book, err := bookService.GetBookByID(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		json.NewEncoder(w).Encode(book)
	}
}

func UpdateBookHandler(bookService *application.BookService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		var book application.UpdateBookRequest
		err = json.NewDecoder(r.Body).Decode(&book)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		book.ID = id
		bookService.UpdateBook(book)

		w.WriteHeader(http.StatusOK)
	}
}

func DeleteBookHandler(bookService *application.BookService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		bookService.DeleteBook(id)

		w.WriteHeader(http.StatusOK)
	}
}
