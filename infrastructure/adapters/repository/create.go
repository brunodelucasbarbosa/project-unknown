package repository

import (
	"github.com/brunodelucasbarbosa/project-unknown/core/domain"
	"github.com/sirupsen/logrus"
)

func (r *UserRepository) Create(user *domain.User) (*domain.User, error) {
	stm := `INSERT INTO users (name, password, age, email, phone, address, document, username, isActive, createdAt, updatedAt) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING id, email, name, createdAt`
	logrus.Info("SQL Statement: ", stm)
	var userCreated domain.User

	err := r.db.QueryRow(stm, user.Name, user.Password, user.Age, user.Email, user.Phone, user.Address, user.Document, user.Username, user.IsActive, user.CreatedAt, user.UpdatedAt).
		Scan(&userCreated.ID, &userCreated.Email, &userCreated.Name, &userCreated.CreatedAt)
	if err != nil {
		logrus.Error("Error inserting user into database: ", err)
		return nil, err
	}

	return &userCreated, nil
}
