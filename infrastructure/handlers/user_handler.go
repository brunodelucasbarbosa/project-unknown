package handlers

import (
	"net/http"

	"github.com/brunodelucasbarbosa/project-unknown/core/domain"
	"github.com/brunodelucasbarbosa/project-unknown/core/service"
	"github.com/brunodelucasbarbosa/project-unknown/infrastructure/handlers/dto"
	"github.com/brunodelucasbarbosa/project-unknown/utils/responses"
	"github.com/go-playground/validator/v10"
)

type HandlerUser struct {
	Validator    *validator.Validate
	LoginService service.ILoginService
}

type IHandlerUser interface {
	HandleLogin(w http.ResponseWriter, r *http.Request)
	HandleCreateUser(w http.ResponseWriter, r *http.Request)
}

func (h *HandlerUser) HandleLogin(w http.ResponseWriter, r *http.Request) {
	var payload dto.UserLoginPayload

	if err := responses.ParseJsonRequestBodyOrSendErrorResponse(w, r, &payload); err != nil {
		return
	}

	if err := h.Validator.Struct(payload); err != nil {
		responses.SendErrorResponse(w, http.StatusBadRequest, "VALIDATION_ERROR", err)
		return
	}

	_, _ = h.LoginService.Login(payload.Email, payload.Password)
}

func (h *HandlerUser) HandleCreateUser(w http.ResponseWriter, r *http.Request) {
	var payload dto.UserCreatePayload

	if err := responses.ParseJsonRequestBodyOrSendErrorResponse(w, r, &payload); err != nil {
		return
	}

	if err := h.Validator.Struct(payload); err != nil {
		responses.SendErrorResponse(w, http.StatusBadRequest, "VALIDATION_ERROR", err)
		return
	}
	u := domain.NewUser(payload.Name, payload.Password, payload.Email, payload.Phone, payload.Address, payload.Document, payload.Username, payload.Age)

	user, err := h.LoginService.Create(u)

	if err != nil {
		responses.SendErrorResponse(w, http.StatusInternalServerError, "INTERNAL_SERVER_ERROR", err)
		return
	}

	responses.SendResponse(w, http.StatusCreated, user)

}
