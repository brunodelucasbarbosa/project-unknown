package di

import (
	"database/sql"

	"github.com/brunodelucasbarbosa/project-unknown/core/service"
	"github.com/brunodelucasbarbosa/project-unknown/infrastructure/adapters/repository"
	handlers "github.com/brunodelucasbarbosa/project-unknown/infrastructure/handlers"
	"github.com/go-playground/validator/v10"
	_ "github.com/lib/pq"
)

func NewUserHandler() handlers.IHandlerUser {
	conn := `host=localhost port=5432 user=user password=password dbname=mydatabase sslmode=disable`
	userDb, _ := sql.Open("postgres", conn)
	userRepository := repository.NewUserRepository(userDb)
	loginService := service.NewLoginService(userRepository)

	v := validator.New(validator.WithRequiredStructEnabled())
	return &handlers.HandlerUser{
		Validator:    v,
		LoginService: loginService,
	}
}
