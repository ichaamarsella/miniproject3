package service

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type UserService struct {
	apiURL string
}

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func NewUserService(apiURL string) *UserService {
	return &UserService{
		apiURL: apiURL,
	}
}

func (s *UserService) GetUserByID(userID int) (*User, error) {
	url := fmt.Sprintf("%s/users/%d", s.apiURL, userID)
	resp, err := http.Get(url)
	if err != nil {
