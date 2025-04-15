package domain

import (
	"time"
)

type User struct {
	ID        *string
	Name      *string
	Password  *string
	Age       *int
	Email     *string
	Phone     *string
	Address   *string
	Document  *string
	Username  *string
	IsActive  *bool
	Contents  *[]Content
	Comments  *[]Comment
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

func NewUser(name, password, email, phone, address, document, username string, age int) *User {
	return &User{
		ID:        nil,
		Name:      &name,
		Password:  &password,
		Email:     &email,
		Phone:     &phone,
		Address:   &address,
		Document:  &document,
		Username:  &username,
		Age:       &age,
		IsActive:  nil,
		CreatedAt: nil,
		UpdatedAt: nil,
	}
}
