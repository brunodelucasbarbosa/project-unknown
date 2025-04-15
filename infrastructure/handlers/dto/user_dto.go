package dto

type UserLoginPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type UserCreatePayload struct {
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required"`
	Age      int    `json:"age" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Phone    string `json:"phone" validate:"required"`
	Address  string `json:"address" validate:"required"`
	Document string `json:"document" validate:"required"`
	Username string `json:"username" validate:"required"`
}
