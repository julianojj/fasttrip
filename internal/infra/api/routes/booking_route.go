package routes

import (
	"net/http"

	"github.com/julianojj/fasttrip/internal/infra/api/controllers"
	"github.com/julianojj/fasttrip/internal/infra/api/middlewares"
)

type BookingRoute struct {
	r                 *http.ServeMux
	bookingController *controllers.BookingController
	authMiddleware    *middlewares.AuthMiddleware
}

func NewBookingRoute(
	r *http.ServeMux,
	bookingController *controllers.BookingController,
	authMiddleware *middlewares.AuthMiddleware,
) *BookingRoute {
	return &BookingRoute{
		r,
		bookingController,
		authMiddleware,
	}
}

func (br *BookingRoute) Init() {
	br.r.HandleFunc("/make_booking", br.authMiddleware.ApplyHandler(br.bookingController.MakeBooking))
	br.r.HandleFunc("/get_bookings", br.authMiddleware.ApplyHandler(br.bookingController.GetBookings))
	br.r.HandleFunc("/get_booking/{id}", br.authMiddleware.ApplyHandler(br.bookingController.GetBooking))
}
