package application

import "github.com/julianfrancor/bookswap/internal/domain"

type CreateBookRequest struct {
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

type BookService struct {
	repository domain.BookRepository
}

func NewBookService(repository domain.BookRepository) *BookService {
	return &BookService{
		repository: repository,
	}
}

func (s *BookService) CreateBook(request CreateBookRequest) {
	book := domain.Book{
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
