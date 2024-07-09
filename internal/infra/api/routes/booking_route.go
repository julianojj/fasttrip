package routes

import (
	"net/http"

	"github.com/julianojj/fastrip/internal/infra/api/controllers"
)

type BookingRoute struct {
	r                 *http.ServeMux
	bookingController *controllers.BookingController
}

func NewBookingRoute(
	r *http.ServeMux,
	bookingController *controllers.BookingController,
) *BookingRoute {
	return &BookingRoute{
		r,
		bookingController,
	}
}

func (r *BookingRoute) Init() {
	r.r.HandleFunc("/make_booking", r.bookingController.MakeBooking)
}
