package main

import (
	"log"
	"net/http"

	"github.com/julianfrancor/bookswap/internal/application"
	"github.com/julianfrancor/bookswap/internal/domain"
	"github.com/julianfrancor/bookswap/internal/infrastructure/persistence"
	"github.com/gorilla/mux"
)

func main() {
	// Crear el repositorio y el servicio
	bookRepository := persistence.NewBookRepository()
	bookService := application.NewBookService(bookRepository)

	// Crear el enrutador
	r := mux.NewRouter()

	// Definir las rutas
	r.HandleFunc("/books", CreateBookHandler(bookService)).Methods("POST")
	r.HandleFunc("/books/{id}", GetBookHandler(bookService)).Methods("GET")
	r.HandleFunc("/books/{id}", UpdateBookHandler(bookService)).Methods("PUT")
	r.HandleFunc("/books/{id}", DeleteBookHandler(bookService)).Methods("DELETE")

	// Iniciar el servidor
	log.Println("Starting server on :8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
