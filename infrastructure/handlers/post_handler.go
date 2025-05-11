package handlers

import (
	"net/http"

	"github.com/brunodelucasbarbosa/project-unknown/infrastructure/handlers/dto"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

type HandlerPost struct {
	Validator *validator.Validate
}

func (h *HandlerPost) HandleCreatePost(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "id")

	var payload dto.CreatePostDto

	if err := parseAndValidateOrSendError(w, r, payload, h.Validator); err != nil {
		http.Error(w, `{"error": "Invalid payload"}`, http.StatusBadRequest)
		return
	}

	d := payload.ToDomain()
	d.UserID = &userId

}

func (h *HandlerPost) HandleGetPost(w http.ResponseWriter, r *http.Request) {

}
