package persistence

import "github.com/julianfrancor/bookswap/internal/domain"

type BookRepository struct {
	books []domain.Book
}

func NewBookRepository() *BookRepository {
	return &BookRepository{
		books: []domain.Book{},
	}
}

func (r *BookRepository) Create(book domain.Book) {
	r.books = append(r.books, book)
}

func (r *BookRepository) Update(book domain.Book) {
	for i, b := range r.books {
		if b.ID == book.ID {
			r.books[i] = book
			break
		}
	}
}

func (r *BookRepository) GetByID(id int) (domain.Book, error) {
	for _, b := range r.books {
		if b.ID == id {
			return b, nil
		}
	}
	return domain.Book{}, domain.ErrBookNotFound
}

func (r *BookRepository) GetAll() []domain.Book {
	return r.books
}

func (r *BookRepository) Delete(id int) {
	for i, b := range r.books {
		if b.ID == id {
			r.books = append(r.books[:i], r.books[i+1:]...)
			break
		}
	}
}
