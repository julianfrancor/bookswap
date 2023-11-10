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

// CreateBookHandler Create a Book Request
// @Summary			Create a new book
// @Description		Create a new book and associate it with a user
// @Tags			books
// @Accept			json
// @Produce			json
// @Param			request			body		application.CreateBookRequest	true	"Create Book Request"
// @Param			Authorization	header		string							true	"Bearer {token}"
// @Success			201				{object}	domain.Book
// @Router			/books [post]
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

// GetAllBooksHandler
//
// @Summary Get all books
// @Description	Get a list of all books
// @Tags		books
// @Produce		json
// @Success		200	{array}	domain.Book
// @Router		/books [get]
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

// CreateUserHandler @Summary Create a new user
//
// @Description		Create a new user
// @Tags			users
// @Accept			json
// @Produce			json
// @Param			request		body		application.CreateUserRequest	true	"CreateUserRequest"
// @Success			201			{object}	domain.User
// @Router			/users [post]
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

// GetUserHandler
// @Summary Get a user by ID
// @Description	Get a user by their ID
// @Tags		users
// @Produce		json
// @Param		id	path		int	true	"User ID"
// @Success		200	{object}	domain.User
// @Router		/users/{id} [get]
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

// GetAllUsersHandler
// @Summary Get all users
// @Description		Get a list of all users
// @Tags			users
// @Produce			json
// @Success			200	{array}	domain.User
// @Router			/users [get]
func GetAllUsersHandler(userService *application.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users := userService.GetAllUsers()

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(users)
	}
}

// UpdateUserHandler @Summary Update a user
//
// @Description		Update a user's information
// @Tags			users
// @Accept			json
// @Produce			json
// @Param			id		path		int					true	"User ID"
// @Param			request	body		application.UpdateUserRequest	true	"UpdateUserRequest"
// @Success			200		{object}	domain.User
// @Router			/users/{id} [put]
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

// DeleteUserHandler
// @Summary Delete a user
// @Description	Delete a user by their ID
// @Tags			users
// @Param			id				path	int	true	"User ID"
// @Success			204				"No Content"
// @Router			/users/{id}		[delete]
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
