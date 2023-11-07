package main

import (
	"encoding/json"
	"errors"
	"github.com/julianfrancor/bookswap/internal/domain"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/julianfrancor/bookswap/internal/application"
)

func CreateBookHandler(bookService *application.BookService, userService *application.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book application.CreateBookRequest
		err := json.NewDecoder(r.Body).Decode(&book)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = bookService.CreateBook(book, userService)
		if err != nil {
			if errors.Is(err, domain.ErrUserNotFound) {
				http.Error(w, "User not found", http.StatusNotFound)
			} else {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}
		w.WriteHeader(http.StatusCreated)
	}
}

func GetAllBooksHandler(bookService *application.BookService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		books := bookService.GetAllBooks()

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(books)
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

func GetAllUsersHandler(userService *application.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users := userService.GetAllUsers()

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(users)
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
