package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type MockUserRepository struct {
	users map[int]*User
}

// User struct for testing
// Use the User struct from repository_test.go

// UserController struct and constructor
type UserController struct {
	repo *MockUserRepository
}

func NewUserController(repo *MockUserRepository) *UserController {
	return &UserController{repo: repo}
}

// Create handler
func (uc *UserController) Create(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}
	if user.Email == "" || user.Name == "" || !isValidEmail(user.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid email"})
		return
	}
	user.ID = len(uc.repo.users) + 1
	uc.repo.users[user.ID] = &user
	c.JSON(http.StatusCreated, user)
}

// Get handler
func (uc *UserController) Get(c *gin.Context) {
	idParam := c.Param("id")
	var id int
	_, err := fmt.Sscanf(idParam, "%d", &id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	user, ok := uc.repo.users[id]
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

// Simple email validation for testing
func isValidEmail(email string) bool {
	return len(email) >= 3 && bytes.Contains([]byte(email), []byte("@"))
}

func TestUserController_Create(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockRepo := &MockUserRepository{users: make(map[int]*User)}
	controller := NewUserController(mockRepo)
	router := gin.Default()
	router.POST("/users", controller.Create)

	tests := []struct {
		name       string
		body       map[string]interface{}
		wantStatus int
	}{
		{
			name: "valid user",
			body: map[string]interface{}{
				"name":  "John Doe",
				"email": "john@example.com",
			},
			wantStatus: http.StatusCreated,
		},
		{
			name: "invalid email",
			body: map[string]interface{}{
				"name":  "John Doe",
				"email": "invalid-email",
			},
			wantStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jsonBody, _ := json.Marshal(tt.body)
			req := httptest.NewRequest("POST", "/users", bytes.NewBuffer(jsonBody))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.wantStatus, w.Code)
		})
	}
}

func TestUserController_Get(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockRepo := &MockUserRepository{
		users: map[int]*User{
			1: {ID: 1, Name: "John Doe", Email: "john@example.com"},
		},
	}
	controller := NewUserController(mockRepo)
	router := gin.Default()
	router.GET("/users/:id", controller.Get)

	t.Run("existing user", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/users/1", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		var response map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &response)
		assert.Equal(t, "John Doe", response["name"])
	})

	t.Run("non-existing user", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/users/999", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code)
	})
}
