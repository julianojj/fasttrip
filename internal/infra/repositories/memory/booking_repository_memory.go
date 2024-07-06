package memory

import (
	"time"

	"github.com/julianojj/fastrip/internal/domain"
)

type BookingRepositoryMemory struct {
	bookings []*domain.Booking
}

var _ domain.BookingRepository = (*BookingRepositoryMemory)(nil)

func NewBookingRepositoryMemory() *BookingRepositoryMemory {
	return &BookingRepositoryMemory{
		bookings: make([]*domain.Booking, 0),
	}
}

func (brm *BookingRepositoryMemory) Save(booking *domain.Booking) error {
	brm.bookings = append(brm.bookings, booking)
	return nil
}

func (brm *BookingRepositoryMemory) CheckAvailability(checkIn time.Time, checkOut time.Time) (bool, error) {
	for _, booking := range brm.bookings {
		if booking.CheckIn.Before(checkOut) && booking.CheckOut.After(checkIn) {
			return false, nil
		}
	}
	return true, nil
}
