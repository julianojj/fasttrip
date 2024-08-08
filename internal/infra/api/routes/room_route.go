package routes

import (
	"net/http"

	"github.com/julianojj/fasttrip/internal/infra/api/controllers"
	"github.com/julianojj/fasttrip/internal/infra/api/middlewares"
)

type RoomRoute struct {
	r              *http.ServeMux
	roomController *controllers.RoomController
	authMiddleware *middlewares.AuthMiddleware
}

func NewRoomRoute(
	r *http.ServeMux,
	roomController *controllers.RoomController,
	authMiddleware *middlewares.AuthMiddleware,
) *RoomRoute {
	return &RoomRoute{
		r,
		roomController,
		authMiddleware,
	}
}

func (rr *RoomRoute) Init() {
	rr.r.HandleFunc("/register_room", rr.authMiddleware.ApplyHandler(rr.roomController.RegisterRoom))
	rr.r.HandleFunc("/get_rooms", rr.authMiddleware.ApplyHandler(rr.roomController.GetRooms))
}
