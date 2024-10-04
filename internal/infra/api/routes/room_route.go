package routes

import (
	"net/http"

	"github.com/julianojj/fasttrip/internal/infra/api/controllers"
	"github.com/julianojj/fasttrip/internal/infra/api/middlewares"
	"github.com/newrelic/go-agent/v3/newrelic"
)

type RoomRoute struct {
	r              *http.ServeMux
	app *newrelic.Application
	roomController *controllers.RoomController
	authMiddleware *middlewares.AuthMiddleware
}

func NewRoomRoute(
	r *http.ServeMux,
	app *newrelic.Application,
	roomController *controllers.RoomController,
	authMiddleware *middlewares.AuthMiddleware,
) *RoomRoute {
	return &RoomRoute{
		r,
		app,
		roomController,
		authMiddleware,
	}
}

func (rr *RoomRoute) Init() {
	rr.r.HandleFunc(newrelic.WrapHandleFunc(rr.app, "/register_room", rr.authMiddleware.ApplyHandler(rr.roomController.RegisterRoom)))
	rr.r.HandleFunc(newrelic.WrapHandleFunc(rr.app, "/get_rooms", rr.authMiddleware.ApplyHandler(rr.roomController.GetRooms)))
}
