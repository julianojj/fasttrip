package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/julianojj/fasttrip/internal/core/exceptions"
	"github.com/julianojj/fasttrip/internal/core/usecases"
)

type UserController struct {
	registerUser *usecases.RegisterUser
	authUser     *usecases.AuthUser
	getUser      *usecases.GetUser
}

func NewUserController(
	registerUser *usecases.RegisterUser,
	authUser *usecases.AuthUser,
	getUser *usecases.GetUser,
) *UserController {
	return &UserController{
		registerUser,
		authUser,
		getUser,
	}
}

// Authenticate user godoc
//
//	@Summary		Autentica usuário
//	@Description	Autentica usuário
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			input	body		usecases.AuthUserInput	true	"AuthUserInput"
//	@Success		200		{object}	usecases.AuthUserOutput
//	@Failure		401		{object}	ErrorResponse
//	@Failure		500		{object}	ErrorResponse
//	@Router			/auth_user [post]
func (c *UserController) AuthenticateUser(w http.ResponseWriter, r *http.Request) {
	var input *usecases.AuthUserInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		handleError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	output, err := c.authUser.Execute(input)
	if err != nil {
		statusCode := http.StatusInternalServerError
		switch err.(type) {
		case *exceptions.DomainException:
			statusCode = http.StatusUnprocessableEntity
		case *exceptions.UnauthorizedException:
			statusCode = http.StatusUnauthorized
		}

		handleError(w, statusCode, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(&output); err != nil {
		handleError(w, http.StatusInternalServerError, "Failed to encode response")
		return
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
		handleError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	output, err := c.registerUser.Execute(input)
	if err != nil {
		statusCode := http.StatusInternalServerError
		switch err.(type) {
		case *exceptions.DomainException:
			statusCode = http.StatusUnprocessableEntity
		}
		handleError(w, statusCode, err.Error())
		return
	}

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(&output); err != nil {
		handleError(w, http.StatusInternalServerError, "Failed to encode response")
		return
	}
}

// GetUser godoc
//
//	@Summary		Pegar usuário
//	@Description	Pegar usuário por ID
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string	true	"Insert you bearer token"
//	@Param			user_id			path		string	true	"User ID"
//	@Success		200				{array}		usecases.GetUserOutput
//	@Failure		404				{object}	ErrorResponse
//	@Failure		500				{object}	ErrorResponse
//	@Router			/get_user/{user_id} [get]
func (c *UserController) GetUser(w http.ResponseWriter, r *http.Request) {
	params := r.PathValue("user_id")
	output, err := c.getUser.Execute(params)
	if err != nil {
		statusCode := http.StatusInternalServerError
		switch err.(type) {
		case *exceptions.NotFoundException:
			statusCode = http.StatusNotFound
		}
		handleError(w, statusCode, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(&output); err != nil {
		handleError(w, http.StatusInternalServerError, "Failed to encode response")
		return
	}
}
