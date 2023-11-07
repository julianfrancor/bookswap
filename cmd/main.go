package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/julianfrancor/bookswap/internal/application"
	_ "github.com/julianfrancor/bookswap/internal/domain"
	"github.com/julianfrancor/bookswap/internal/infrastructure/persistence"
)

func main() {
	// Create repositories and services
	bookRepository := persistence.NewBookRepository()
	bookService := application.NewBookService(*bookRepository)

	userRepository := persistence.NewUserRepository()
	userService := application.NewUserService(*userRepository)

	// Create the router
	router := mux.NewRouter()

	// Define the routes
	router.HandleFunc("/books", CreateBookHandler(bookService)).Methods("POST")
	router.HandleFunc("/books/{id}", GetBookHandler(bookService)).Methods("GET")
	router.HandleFunc("/books/{id}", UpdateBookHandler(bookService)).Methods("PUT")
	router.HandleFunc("/books/{id}", DeleteBookHandler(bookService)).Methods("DELETE")

	router.HandleFunc("/users", CreateUserHandler(userService)).Methods("POST")
	router.HandleFunc("/users/{id}", GetUserHandler(userService)).Methods("GET")
	router.HandleFunc("/users/{id}", UpdateUserHandler(userService)).Methods("PUT")
	router.HandleFunc("/users/{id}", DeleteUserHandler(userService)).Methods("DELETE")

	// Init the server
	log.Println("Starting server on :8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
