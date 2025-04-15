package main

import (
	"fmt"
	"net/http"

	"github.com/brunodelucasbarbosa/project-unknown/infrastructure/adapters/api"
	"github.com/brunodelucasbarbosa/project-unknown/infrastructure/config/di"
)

func main() {
	handlerUser := di.NewUserHandler()
	api := api.Router(handlerUser)

	fmt.Println("Server running on port 8080")
	http.ListenAndServe(":8080", api)
}
