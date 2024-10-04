package routes

import (
	"net/http"

	"github.com/julianojj/fasttrip/internal/infra/api/controllers"
	"github.com/julianojj/fasttrip/internal/infra/api/middlewares"
	"github.com/newrelic/go-agent/v3/newrelic"
)

type UserRoute struct {
	r              *http.ServeMux
	app            *newrelic.Application
	authController *controllers.UserController
	authMiddleware *middlewares.AuthMiddleware
}

func NewUserRoute(
	r *http.ServeMux,
	app *newrelic.Application,
	authController *controllers.UserController,
	authMiddleware *middlewares.AuthMiddleware,
) *UserRoute {
	return &UserRoute{
		r,
		app,
		authController,
		authMiddleware,
	}
}

func (ur *UserRoute) Init() {
	ur.r.HandleFunc(newrelic.WrapHandleFunc(ur.app, "POST /register_user", ur.authController.RegisterUser))
	ur.r.HandleFunc(newrelic.WrapHandleFunc(ur.app, "POST /auth_user", ur.authController.AuthenticateUser))
	ur.r.HandleFunc(newrelic.WrapHandleFunc(ur.app, "GET /get_user/{user_id}", ur.authMiddleware.ApplyHandler(ur.authController.GetUser)))
}
