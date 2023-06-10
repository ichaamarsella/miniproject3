package usecase

import (
	"miniproject3/internal/service"
)

type UserUseCase struct {
	userService service.UserService
}

func NewUserUseCase(userService service.UserService) *UserUseCase {
	return &UserUseCase{
		userService: userService,
	}
}

func (uc *UserUseCase) GetUserByID(userID int) (*service.User, error) {
	user, err := uc.userService.GetUserByID(userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}
