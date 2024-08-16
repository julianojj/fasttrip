package routes

import (
	"net/http"

	"github.com/julianojj/fasttrip/internal/infra/api/controllers"
	"github.com/julianojj/fasttrip/internal/infra/api/middlewares"
)

type UserRoute struct {
	r              *http.ServeMux
	authController *controllers.UserController
	authMiddleware *middlewares.AuthMiddleware
}

func NewUserRoute(
	r *http.ServeMux,
	authController *controllers.UserController,
	authMiddleware *middlewares.AuthMiddleware,
) *UserRoute {
	return &UserRoute{
		r,
		authController,
		authMiddleware,
	}
}

func (ur *UserRoute) Init() {
	ur.r.HandleFunc("POST /register_user", ur.authController.RegisterUser)
	ur.r.HandleFunc("POST /auth_user", ur.authController.AuthenticateUser)
	ur.r.HandleFunc("GET /get_user/{user_id}", ur.authMiddleware.ApplyHandler(ur.authController.GetUser))
}
