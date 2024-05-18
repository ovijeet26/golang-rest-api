package services

import (
	"fmt"
	"testing"

	"github.com/ovijeet26/golang-rest-api/internal/models"
	"github.com/ovijeet26/golang-rest-api/internal/repositories"
	"github.com/stretchr/testify/assert"
)

func TestGetAllUsers(t *testing.T) {

	fmt.Println("Running TestGetAllUsers")
	// Prepare mock repository
	mockRepo := &repositories.UserRepository{
		Data: map[string]models.User{
			"1": {ID: "1", Username: "user1", Password: "pass1"},
			"2": {ID: "2", Username: "user2", Password: "pass2"},
		},
	}

	// Create service
	service := NewUserService(mockRepo)

	// Call service method
	users, err := service.GetAllUsers()

	// Assertions
	assert.NoError(t, err)
	assert.Len(t, users, 2)
}

func TestCreateUser(t *testing.T) {
	// Prepare mock repository
	mockRepo := &repositories.UserRepository{
		Data: map[string]models.User{},
	}

	// Create service
	service := NewUserService(mockRepo)

	// Create user
	user := models.User{ID: "1", Username: "user1", Password: "pass1"}

	// Call service method
	err := service.CreateUser(user)

	// Assertions
	assert.NoError(t, err)
	assert.Len(t, mockRepo.Data, 1)
	assert.Contains(t, mockRepo.Data, "1")
	assert.Equal(t, user, mockRepo.Data["1"])
}

func TestAuthenticateValidCredentials(t *testing.T) {
	// Prepare mock repository
	mockRepo := &repositories.UserRepository{
		Data: map[string]models.User{
			"user1": {ID: "1", Username: "user1", Password: "pass1"},
		},
	}

	// Create service
	service := NewUserService(mockRepo)

	// Call service method
	token, err := service.Authenticate("user1", "pass1")

	// Assertions
	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}

func TestAuthenticateInvalidCredentials(t *testing.T) {
	// Prepare mock repository
	mockRepo := &repositories.UserRepository{
		Data: map[string]models.User{
			"user1": {ID: "1", Username: "user1", Password: "pass1"},
		},
	}

	// Create service
	service := NewUserService(mockRepo)

	// Call service method with invalid credentials
	token, err := service.Authenticate("user1", "wrongpassword")

	// Assertions
	assert.Error(t, err)
	assert.Empty(t, token)
}
