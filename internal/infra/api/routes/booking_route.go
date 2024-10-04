package routes

import (
	"net/http"

	"github.com/julianojj/fasttrip/internal/infra/api/controllers"
	"github.com/julianojj/fasttrip/internal/infra/api/middlewares"
	"github.com/newrelic/go-agent/v3/newrelic"
)

type BookingRoute struct {
	r                 *http.ServeMux
	app               *newrelic.Application
	bookingController *controllers.BookingController
	authMiddleware    *middlewares.AuthMiddleware
}

func NewBookingRoute(
	r *http.ServeMux,
	app *newrelic.Application,
	bookingController *controllers.BookingController,
	authMiddleware *middlewares.AuthMiddleware,
) *BookingRoute {
	return &BookingRoute{
		r,
		app,
		bookingController,
		authMiddleware,
	}
}

func (br *BookingRoute) Init() {
	br.r.HandleFunc(newrelic.WrapHandleFunc(br.app, "/make_booking", br.authMiddleware.ApplyHandler(br.bookingController.MakeBooking)))
	br.r.HandleFunc(newrelic.WrapHandleFunc(br.app, "/get_bookings", br.authMiddleware.ApplyHandler(br.bookingController.GetBookings)))
	br.r.HandleFunc(newrelic.WrapHandleFunc(br.app, "/get_booking/{id}", br.authMiddleware.ApplyHandler(br.bookingController.GetBooking)))
}
