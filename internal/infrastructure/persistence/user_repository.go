package persistence

import "github.com/julianfrancor/bookswap/internal/domain"

type UserRepository struct {
	users []domain.User
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		users: []domain.User{},
	}
}

func (r *UserRepository) Create(user domain.User) {
	r.users = append(r.users, user)
}

func (r *UserRepository) Update(user domain.User) {
	for i, u := range r.users {
		if u.ID == user.ID {
			r.users[i] = user
			break
		}
	}
}

func (r *UserRepository) GetByID(id int) (domain.User, error) {
	for _, u := range r.users {
		if u.ID == id {
			return u, nil
		}
	}
	return domain.User{}, domain.ErrUserNotFound
}

func (r *UserRepository) Delete(id int) {
	for i, u := range r.users {
		if u.ID == id {
			r.users = append(r.users[:i], r.users[i+1:]...)
			break
		}
	}
}

func (r *UserRepository) GetAll() []domain.User {
	return r.users
}
