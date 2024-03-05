package service

import (
	"github.com/chirag1807/websocket-go/api/model/request"
	"github.com/chirag1807/websocket-go/api/model/response"
	"github.com/chirag1807/websocket-go/api/repository"
)

type AuthService interface {
	UserRegistration(user request.User) error
	UserLogin(user request.User) (response.User, error)
}

type authService struct {
	authRepository repository.AuthRepository
}

func NewAuthService(a repository.AuthRepository) AuthService {
	return authService{
		authRepository: a,
	}
}

func (a authService) UserRegistration(user request.User) error {
	return a.authRepository.UserRegistration(user)
}

func (a authService) UserLogin(user request.User) (response.User, error) {
	return a.authRepository.UserLogin(user)
}
