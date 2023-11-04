package application

import (
    "testing"

    "github.com/julianfrancor/bookswap/internal/domain"
)

func TestBookService(t *testing.T) {
    repository := domain.NewBookRepository()
    service := NewBookService(repository)

    // Test Crear un libro
    book := domain.Book{ID: 1, Title: "Libro 1", Author: "Autor 1", Genre: "Género 1", Status: "Disponible"}
    service.CreateBook(book)

    // Verificar si el libro se ha creado correctamente
    createdBook, err := service.GetBookByID(1)
    if err != nil {
        t.Errorf("Error al obtener el libro: %v", err)
    }
    if createdBook != book {
        t.Errorf("El libro creado no coincide con el esperado")
    }

    // Test Actualizar un libro
    updatedBook := domain.Book{ID: 1, Title: "Libro 1 Modificado", Author: "Autor 1", Genre: "Género 1", Status: "Intercambiado"}
    service.UpdateBook(updatedBook)

    // Verificar si el libro se ha actualizado correctamente
    updatedBook, err = service.GetBookByID(1)
    if err != nil {
        t.Errorf("Error al obtener el libro actualizado: %v", err)
    }
    if updatedBook != updatedBook {
        t.Errorf("El libro actualizado no coincide con el esperado")
    }

    // Test Eliminar un libro
    service.DeleteBook(1)

    // Verificar si el libro se ha eliminado correctamente
    _, err = service.GetBookByID(1)
    if err != domain.ErrBookNotFound {
        t.Errorf("Se esperaba un error al intentar obtener un libro eliminado")
    }
}
