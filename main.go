package main

import (
	"net/http"

	_ "github.com/julianojj/fastrip/docs"
	"github.com/julianojj/fastrip/internal/core/usecases"
	"github.com/julianojj/fastrip/internal/infra/api/controllers"
	"github.com/julianojj/fastrip/internal/infra/api/routes"
	"github.com/julianojj/fastrip/internal/infra/repositories/memory"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title			Fastrip API
// @version		v1.3.0
// @description	Fastrip API permite executar operações para cadastrar quartos, fazer reservas, checkin e checkout.
// @host			localhost:8080
func main() {
	r := http.NewServeMux()

	roomRepository := memory.NewRoomRepositoryMemory()
	bookingRepository := memory.NewBookingRepositoryMemory()

	makeBooking := usecases.NewMakeBooking(roomRepository, bookingRepository)
	getBookings := usecases.NewGetBookings(roomRepository, bookingRepository)
	registerRoom := usecases.NewRegisterRoom(roomRepository)
	getRooms := usecases.NewGetRooms(roomRepository)

	bookingController := controllers.NewBookingController(makeBooking, getBookings)
	roomController := controllers.NewRoomController(registerRoom, getRooms)

	routes.NewBookingRoute(r, bookingController).Init()
	routes.NewRoomRoute(r, roomController).Init()

	r.HandleFunc("/swagger/*", httpSwagger.WrapHandler)

	http.ListenAndServe(":8080", corsHandler(r))
}

func corsHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}
