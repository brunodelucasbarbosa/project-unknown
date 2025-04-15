package repository

import (
	"database/sql"

	"github.com/brunodelucasbarbosa/project-unknown/core/domain"
)

type IUserRepository interface {
	Create(user *domain.User) (*domain.User, error)
}

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) IUserRepository {
	return &UserRepository{
		db: db,
	}
}
