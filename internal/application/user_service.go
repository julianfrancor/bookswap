// internal/application/user_service.go
package application

import (
	"github.com/julianfrancor/bookswap/internal/domain"
	"github.com/julianfrancor/bookswap/internal/infrastructure/persistence"
)

type CreateUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

type UpdateUserRequest struct {
	ID       int           `json:"id"`
	Username string        `json:"username"`
	Email    string        `json:"email"`
	Books    []domain.Book `json:"books,omitempty"`
}

type UserService struct {
	repository persistence.UserRepository
}

func NewUserService(repository persistence.UserRepository) *UserService {
	return &UserService{
		repository: repository,
	}
}

func (s *UserService) CreateUser(request CreateUserRequest) {
	user := domain.User{
		ID:       len(s.repository.GetAll()) + 1,
		Username: request.Username,
		Email:    request.Email,
		Books:    []domain.Book{},
	}
	s.repository.Create(user)
}

func (s *UserService) UpdateUser(request UpdateUserRequest) {
	// Obtener el usuario existente
	user, err := s.repository.GetByID(request.ID)
	if err != nil {
		// Manejar el error si no se encuentra al usuario
		return
	}

	// Actualizar campos
	user.Username = request.Username
	user.Email = request.Email

	// Si se proporciona una lista de libros en la solicitud, actualizarla
	if request.Books != nil {
		user.Books = request.Books
	}

	// Guardar el usuario actualizado
	s.repository.Update(user)
}

func (s *UserService) GetUserByID(id int) (domain.User, error) {
	return s.repository.GetByID(id)
}

func (s *UserService) GetAllUsers() []domain.User {
	return s.repository.GetAll()
}

func (s *UserService) DeleteUser(id int) {
	s.repository.Delete(id)
}
