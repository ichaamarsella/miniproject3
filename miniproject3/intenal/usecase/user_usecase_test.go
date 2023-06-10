package usecase_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"miniproject3/internal/service"
	"miniproject3/internal/usecase"
)

type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) GetUserByID(userID int) (*service.User, error) {
	args := m.Called(userID)
	return args.Get(0).(*service.User), args.Error(1)
}

func TestGetUserByID(t *testing.T) {
	// Create an instance of the mock UserService
	mockUserService := new(MockUserService)

	// Set up expectations for the mock method
	expectedUser := &service.User{
		ID:       1,
		Username: "testuser",
		Email:    "test@example.com",
	}
	mockUserService.On("GetUserByID", 1).Return(expectedUser, nil)

	// Create an instance of the use case that uses the UserService
	useCase := usecase.NewUserUseCase(mockUserService)

	// Call the method under test
	user, err := useCase.GetUserByID(1)

	// Assert the result
	assert.NoError(t, err)
	assert.Equal(t, expectedUser, user)

	// Assert that the expected method was called on the mock
	mockUserService.AssertCalled(t, "GetUserByID", 1)
}
