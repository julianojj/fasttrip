package main

import (
	"log/slog"
	"net/http"
	"os"

	_ "github.com/julianojj/fastrip/docs"
	"github.com/julianojj/fastrip/internal/core/usecases"
	"github.com/julianojj/fastrip/internal/infra/adapters"
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

	jsonHandler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelInfo,
	})
	logger := slog.New(jsonHandler)
	slog.SetDefault(logger)

	roomRepository := memory.NewRoomRepositoryMemory()
	bookingRepository := memory.NewBookingRepositoryMemory()
	userRepository := memory.NewUserRepositoryMemory()
	hash := adapters.NewBcrypt()

	makeBooking := usecases.NewMakeBooking(roomRepository, bookingRepository)
	getBookings := usecases.NewGetBookings(roomRepository, bookingRepository)
	registerRoom := usecases.NewRegisterRoom(roomRepository)
	getRooms := usecases.NewGetRooms(roomRepository)
	registerUser := usecases.NewRegisterUser(userRepository, hash)

	bookingController := controllers.NewBookingController(makeBooking, getBookings)
	roomController := controllers.NewRoomController(registerRoom, getRooms)
	userController := controllers.NewUserController(registerUser)

	routes.NewBookingRoute(r, bookingController).Init()
	routes.NewRoomRoute(r, roomController).Init()
	routes.NewUserRoute(r, userController).Init()

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
