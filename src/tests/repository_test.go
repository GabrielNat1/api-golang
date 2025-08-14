package tests

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type MockDB struct {
	users []User
}
type User struct {
	ID        int
	Name      string
	Email     string
	CreatedAt time.Time
}

// UserRepository provides methods to interact with users in MockDB.
type UserRepository struct {
	db *MockDB
}

// NewUserRepository creates a new UserRepository.
func NewUserRepository(db *MockDB) *UserRepository {
	return &UserRepository{db: db}
}

// Create adds a new user to the MockDB, checking for duplicate emails.
func (r *UserRepository) Create(ctx context.Context, user *User) error {
	for _, u := range r.db.users {
		if u.Email == user.Email {
			return fmt.Errorf("duplicate email")
		}
	}
	user.ID = len(r.db.users) + 1
	user.CreatedAt = time.Now()
	r.db.users = append(r.db.users, *user)
	return nil
}

// Get retrieves a user by ID from the MockDB.
func (r *UserRepository) Get(ctx context.Context, id int) (*User, error) {
	for _, u := range r.db.users {
		if u.ID == id {
			return &u, nil
		}
	}
	return nil, fmt.Errorf("user not found")
}

func TestUserRepository_Create(t *testing.T) {
	mockDB := &MockDB{users: []User{}}
	repo := NewUserRepository(mockDB)

	tests := []struct {
		name    string
		user    User
		wantErr bool
	}{
		{
			name: "valid user",
			user: User{
				Name:  "John Doe",
				Email: "john@example.com",
			},
			wantErr: false,
		},
		{
			name: "duplicate email",
			user: User{
				Name:  "Jane Doe",
				Email: "john@example.com",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := repo.Create(context.Background(), &tt.user)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotZero(t, tt.user.ID)
			}
		})
	}
}

func TestUserRepository_Get(t *testing.T) {
	mockDB := &MockDB{users: []User{
		{ID: 1, Name: "John Doe", Email: "john@example.com"},
	}}
	repo := NewUserRepository(mockDB)

	t.Run("existing user", func(t *testing.T) {
		user, err := repo.Get(context.Background(), 1)
		assert.NoError(t, err)
		assert.Equal(t, "John Doe", user.Name)
	})

	t.Run("non-existing user", func(t *testing.T) {
		_, err := repo.Get(context.Background(), 999)
		assert.Error(t, err)
	})
}
