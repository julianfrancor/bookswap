package persistence

import (
	"testing"

	"github.com/julianfrancor/bookswap/internal/domain"
)

func TestUserRepository_Create(t *testing.T) {
	repo := NewUserRepository()

	user := domain.User{
		ID:       1,
		Username: "johndoe",
		Email:    "john@example.com",
		Books:    []domain.Book{},
	}

	repo.Create(user)

	// Verify if the user has been created successfully
	createdUser, err := repo.GetByID(1)
	if err != nil {
		t.Errorf("Error fetching the created user: %v", err)
	}

	if !compareUsers(createdUser, user) {
		t.Errorf("Created user does not match the expected one")
	}
}

func TestUserRepository_Update(t *testing.T) {
	repo := NewUserRepository()

	originalUser := domain.User{
		ID:       1,
		Username: "johndoe",
		Email:    "john@example.com",
		Books:    []domain.Book{},
	}

	repo.Create(originalUser)

	updatedUser := domain.User{
		ID:       1,
		Username: "updateduser",
		Email:    "updated@example.com",
		Books:    []domain.Book{},
	}

	repo.Update(updatedUser)

	// Verify if the user has been updated successfully
	result, err := repo.GetByID(1)
	if err != nil {
		t.Errorf("Error fetching the updated user: %v", err)
	}
	if !compareUsers(updatedUser, result) {
		t.Errorf("Updated user does not match the expected one")
	}
}

func TestUserRepository_GetByID(t *testing.T) {
	repo := NewUserRepository()

	user := domain.User{
		ID:       1,
		Username: "johndoe",
		Email:    "john@example.com",
		Books:    []domain.Book{},
	}

	repo.Create(user)

	// Verify if the correct user is fetched by ID
	result, err := repo.GetByID(1)
	if err != nil {
		t.Errorf("Error fetching the user by ID: %v", err)
	}
	if !compareUsers(result, user) {
		t.Errorf("Fetched user does not match the expected one")
	}
}

func TestUserRepository_GetAll(t *testing.T) {
	repo := NewUserRepository()

	user1 := domain.User{
		ID:       1,
		Username: "user1",
		Email:    "user1@example.com",
		Books:    []domain.Book{},
	}

	user2 := domain.User{
		ID:       2,
		Username: "user2",
		Email:    "user2@example.com",
		Books:    []domain.Book{},
	}

	repo.Create(user1)
	repo.Create(user2)

	// Verify if all users are fetched
	result := repo.GetAll()
	if len(result) != 2 {
		t.Errorf("Expected 2 users, got %d", len(result))
	}
	if !compareUsers(result[0], user1) || !compareUsers(result[1], user2) {
		t.Errorf("Fetched users do not match the expected ones")
	}
}

func TestUserRepository_Delete(t *testing.T) {
	repo := NewUserRepository()

	user := domain.User{
		ID:       1,
		Username: "johndoe",
		Email:    "john@example.com",
		Books:    []domain.Book{},
	}

	repo.Create(user)
	repo.Delete(1)

	// Verify if the user has been deleted successfully
	_, err := repo.GetByID(1)
	if err != domain.ErrUserNotFound {
		t.Errorf("Expected an error when trying to fetch a deleted user")
	}
}

// Helper function to compare users including the Books field
func compareUsers(u1, u2 domain.User) bool {
	return u1.ID == u2.ID &&
		u1.Username == u2.Username &&
		u1.Email == u2.Email &&
		compareBooks(u1.Books, u2.Books)
}

// Helper function to compare books
func compareBooks(b1, b2 []domain.Book) bool {
	if len(b1) != len(b2) {
		return false
	}

	for i := range b1 {
		if b1[i] != b2[i] {
			return false
		}
	}

	return true
}
