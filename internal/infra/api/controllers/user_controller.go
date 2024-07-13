package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/julianojj/fastrip/internal/core/exceptions"
	"github.com/julianojj/fastrip/internal/core/usecases"
)

type UserController struct {
	registerUser *usecases.RegisterUser
}

func NewUserController(
	registerUser *usecases.RegisterUser,
) *UserController {
	return &UserController{
		registerUser,
	}
}

// RegisterUser godoc
//
//	@Summary		Registra novo usuário
//	@Description	Registra novo usuário
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			input	body		usecases.RegisterUserInput	true	"RegisterUserInput"
//	@Success		201		{object}	usecases.RegisterUserOutput
//	@Failure		400		{object}	ErrorResponse
//	@Failure		422		{object}	ErrorResponse
//	@Failure		500		{object}	ErrorResponse
//	@Router			/register_user [post]
func (c *UserController) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var input *usecases.RegisterUserInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		handleError(w, http.StatusBadRequest, "Invalid request body", err)
		return
	}

	output, err := c.registerUser.Execute(input)
	if err != nil {
		statusCode := http.StatusInternalServerError
		switch err.(type) {
		case *exceptions.DomainException:
			statusCode = http.StatusUnprocessableEntity
		}
		handleError(w, statusCode, err.Error(), err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(&output); err != nil {
		handleError(w, http.StatusInternalServerError, "Failed to encode response", err)
		return
	}
}
