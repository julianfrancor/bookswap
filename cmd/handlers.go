package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/julianfrancor/bookswap/internal/application"
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

func CreateUserHandler(userService *application.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var userRequest application.CreateUserRequest
		err := json.NewDecoder(r.Body).Decode(&userRequest)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		userService.CreateUser(userRequest)
		w.WriteHeader(http.StatusCreated)
	}
}

func GetUserHandler(userService *application.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		user, err := userService.GetUserByID(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		json.NewEncoder(w).Encode(user)
	}
}

func UpdateUserHandler(userService *application.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		var userRequest application.UpdateUserRequest
		err = json.NewDecoder(r.Body).Decode(&userRequest)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		userRequest.ID = id
		userService.UpdateUser(userRequest)

		w.WriteHeader(http.StatusOK)
	}
}

func DeleteUserHandler(userService *application.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		userService.DeleteUser(id)

		w.WriteHeader(http.StatusOK)
	}
}
