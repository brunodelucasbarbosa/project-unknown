package service

import (
	"github.com/brunodelucasbarbosa/project-unknown/core/domain"
	"github.com/brunodelucasbarbosa/project-unknown/infrastructure/adapters/repository"
)

type ILoginService interface {
	Login(email, password string) (bool, error)
	Create(user *domain.User) (*domain.User, error)
}

type LoginService struct {
	UserRepository repository.IUserRepository
}

func NewLoginService(userRepository repository.IUserRepository) *LoginService {
	return &LoginService{
		UserRepository: userRepository,
	}
}

func (l *LoginService) Login(email, password string) (bool, error) {
	return true, nil
}

