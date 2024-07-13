package routes

import (
	"net/http"

	"github.com/julianojj/fastrip/internal/infra/api/controllers"
)

type UserRoute struct {
	r              *http.ServeMux
	authController *controllers.UserController
}

func NewUserRoute(
	r *http.ServeMux,
	authController *controllers.UserController,
) *UserRoute {
	return &UserRoute{
		r,
		authController,
	}
}

func (ur *UserRoute) Init() {
	ur.r.HandleFunc("POST /register_user", ur.authController.RegisterUser)
	ur.r.HandleFunc("POST /auth_user", ur.authController.AuthenticateUser)
}
