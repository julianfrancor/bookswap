package application

import (
	"github.com/julianfrancor/bookswap/internal/domain"
	"github.com/julianfrancor/bookswap/internal/infrastructure/persistence"
)

type CreateBookRequest struct {
	UserID int    `json:"userID"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Genre  string `json:"genre"`
	Status string `json:"status"`
}

type UpdateBookRequest struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Genre  string `json:"genre"`
	Status string `json:"status"`
}

var bookIDCounter int

type BookService struct {
	repository persistence.BookRepository
}

func NewBookService(repository persistence.BookRepository) *BookService {
	return &BookService{
		repository: repository,
	}
}

func (s *BookService) CreateBook(request CreateBookRequest) {
	bookIDCounter++
	book := domain.Book{
		UserID: request.UserID,
		Title:  request.Title,
		Author: request.Author,
		Genre:  request.Genre,
		Status: request.Status,
	}
	s.repository.Create(book)
}

func (s *BookService) UpdateBook(request UpdateBookRequest) {
	book := domain.Book{
		ID:     request.ID,
		Title:  request.Title,
		Author: request.Author,
		Genre:  request.Genre,
		Status: request.Status,
	}
	s.repository.Update(book)
}

func (s *BookService) GetBookByID(id int) (domain.Book, error) {
	return s.repository.GetByID(id)
}

func (s *BookService) DeleteBook(id int) {
	s.repository.Delete(id)
}
