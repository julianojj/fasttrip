package routes

import (
	"net/http"

	"github.com/julianojj/fastrip/internal/infra/api/controllers"
)

type RoomRoute struct {
	r              *http.ServeMux
	roomController *controllers.RoomController
}

func NewRoomRoute(
	r *http.ServeMux,
	roomController *controllers.RoomController,
) *RoomRoute {
	return &RoomRoute{
		r,
		roomController,
	}
}

func (r *RoomRoute) Init() {
	r.r.HandleFunc("/register_room", r.roomController.RegisterRoom)
	r.r.HandleFunc("/get_rooms", r.roomController.GetRooms)
}
