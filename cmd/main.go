package main

import (
	httpSwagger "github.com/swaggo/http-swagger"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/julianfrancor/bookswap/cmd/docs"
	"github.com/julianfrancor/bookswap/internal/application"
	_ "github.com/julianfrancor/bookswap/internal/domain"
	"github.com/julianfrancor/bookswap/internal/infrastructure/persistence"
)

// @title         BookSwap API
// @description   BookSwap is a RESTful API for managing books and users within the BookSwap application.
//
//	This API provides endpoints to perform CRUD operations on books and users, allowing users
//	to interact with the BookSwap system.
//
// @version       1.0
// @host          localhost:8081
// @BasePath      /
// @contact       name:"BookSwap Support" email:"support@bookswap.com"
func main() {
	// Create repositories and services
	userRepository := persistence.NewUserRepository()
	userService := application.NewUserService(*userRepository)

	bookRepository := persistence.NewBookRepository()
	bookService := application.NewBookService(*bookRepository)

	exchangeRepository := persistence.NewExchangeRepository()
	exchangeService := application.NewExchangeService(*exchangeRepository)

	// Create the router
	router := mux.NewRouter()

	// Define the routes
	router.HandleFunc("/books", CreateBookHandler(bookService, userService)).Methods("POST")
	router.HandleFunc("/books", GetAllBooksHandler(bookService)).Methods("GET")
	router.HandleFunc("/books/{id}", GetBookHandler(bookService)).Methods("GET")
	router.HandleFunc("/books/{id}", UpdateBookHandler(bookService)).Methods("PUT")
	router.HandleFunc("/books/{id}", DeleteBookHandler(bookService)).Methods("DELETE")

	router.HandleFunc("/users", CreateUserHandler(userService)).Methods("POST")
	router.HandleFunc("/users", GetAllUsersHandler(userService)).Methods("GET")
	router.HandleFunc("/users/{id}", GetUserHandler(userService)).Methods("GET")
	router.HandleFunc("/users/{id}", UpdateUserHandler(userService)).Methods("PUT")
	router.HandleFunc("/users/{id}", DeleteUserHandler(userService)).Methods("DELETE")

	// Exchange handlers
	router.HandleFunc("/exchanges", ExchangeBooksHandler(exchangeService, bookService)).Methods("POST")
	router.HandleFunc("/exchanges/{id:[0-9]+}", GetExchangeHandler(exchangeService)).Methods("GET")
	// Add more exchange handlers as needed

	// docs route
	//router.PathPrefix("/documentation/").Handler(httpSwagger.WrapHandler)
	router.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8081/swagger/doc.json"),
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("swagger-ui"),
	)).Methods(http.MethodGet)

	// Init the server
	log.Println("Starting server on :8081...")
	log.Fatal(http.ListenAndServe(":8081", router))
}
