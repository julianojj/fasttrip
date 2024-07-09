package memory

import (
	"time"

	"github.com/julianojj/fastrip/internal/core/domain"
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

func (brm *BookingRepositoryMemory) FindByID(id string) (*domain.Booking, error) {
	for _, booking := range brm.bookings {
		if booking.ID == id {
			return booking, nil
		}
	}
	return nil, nil
}

func (brm *BookingRepositoryMemory) CheckAvailability(checkIn time.Time, checkOut time.Time) (bool, error) {
	for _, booking := range brm.bookings {
		if booking.CheckIn.Before(checkOut) && booking.CheckOut.After(checkIn) && booking.Status != "PENDING_PAYMENT" {
			return false, nil
		}
	}
	return true, nil
}

func (brm *BookingRepositoryMemory) Update(booking *domain.Booking) error {
	for i, b := range brm.bookings {
		if b.ID == booking.ID {
			brm.bookings[i] = booking
			return nil
		}
	}
	return nil
}
