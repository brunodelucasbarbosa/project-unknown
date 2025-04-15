package service

import (
	"github.com/brunodelucasbarbosa/project-unknown/core/domain"
	"github.com/sirupsen/logrus"
)

func (l *LoginService) Create(user *domain.User) (*domain.User, error) {
	logrus.Info("Creating user: ", user)
	userCreated, err := l.UserRepository.Create(user)
	if err != nil {
		return nil, err
	}

	logrus.Info("User created successfully: ", userCreated)
	return userCreated, nil
}
