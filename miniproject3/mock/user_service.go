//go:generate mockery --name UserService

package mock

import (
	"github.com/stretchr/testify/mock"
	"miniproject3/internal/service"
)

type UserService struct {
	mock.Mock
}

func (m *UserService) GetUserByID(userID int) (*service.User, error) {
	args := m.Called(userID)
	return args.Get(0).(*service.User), args.Error(1)
}
