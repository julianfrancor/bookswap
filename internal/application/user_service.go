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
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
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
	user := domain.User{
		ID:       request.ID,
		Username: request.Username,
		Email:    request.Email,
	}
	s.repository.Update(user)
}

func (s *UserService) GetUserByID(id int) (domain.User, error) {
	return s.repository.GetByID(id)
}

func (s *UserService) DeleteUser(id int) {
	s.repository.Delete(id)
}
