package main

import (
	"log/slog"
	"net/http"
	"os"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/julianojj/fasttrip/docs"
	"github.com/julianojj/fasttrip/internal/core/usecases"
	"github.com/julianojj/fasttrip/internal/infra/adapters"
	"github.com/julianojj/fasttrip/internal/infra/api/controllers"
	"github.com/julianojj/fasttrip/internal/infra/api/middlewares"
	"github.com/julianojj/fasttrip/internal/infra/api/routes"
	"github.com/julianojj/fasttrip/internal/infra/repositories/memory"
	"github.com/newrelic/go-agent/v3/integrations/logcontext-v2/nrslog"
	"github.com/newrelic/go-agent/v3/newrelic"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title			Fasttrip API
// @version		v1.8.0
// @description	Fasttrip API permite executar operações para cadastrar quartos, fazer reservas, checkin e checkout.
// @host			api.fasttrip.com.br
func main() {
	r := http.NewServeMux()

	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName("fasttrip"),
		newrelic.ConfigLicense(os.Getenv("NR_KEY")),
		newrelic.ConfigAppLogForwardingEnabled(true),
	)
	if err != nil {
		panic(err)
	}

	jsonHandler := nrslog.JSONHandler(app, os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelInfo,
	})
	logger := slog.New(jsonHandler)
	slog.SetDefault(logger)

	roomRepository := memory.NewRoomRepositoryMemory()
	bookingRepository := memory.NewBookingRepositoryMemory()
	userRepository := memory.NewUserRepositoryMemory()
	hash := adapters.NewBcrypt()
	sign := adapters.NewJWT()

	makeBooking := usecases.NewMakeBooking(roomRepository, bookingRepository)
	getBookings := usecases.NewGetBookings(roomRepository, bookingRepository)
	getBooking := usecases.NewGetBooking(roomRepository, bookingRepository)
	registerRoom := usecases.NewRegisterRoom(roomRepository)
	getRooms := usecases.NewGetRooms(roomRepository)
	registerUser := usecases.NewRegisterUser(userRepository, hash)
	authUser := usecases.NewAuthUser(userRepository, hash, sign)
	getUser := usecases.NewGetUser(userRepository)

	bookingController := controllers.NewBookingController(makeBooking, getBookings, getBooking)
	roomController := controllers.NewRoomController(registerRoom, getRooms)
	userController := controllers.NewUserController(registerUser, authUser, getUser)

	authMiddleware := middlewares.NewAuthMiddleware(sign)

	routes.NewBookingRoute(r, app, bookingController, authMiddleware).Init()
	routes.NewRoomRoute(r, app, roomController, authMiddleware).Init()
	routes.NewUserRoute(r, app, userController, authMiddleware).Init()

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
