package domain

import "time"

type BookingRepository interface {
	Save(booking *Booking) error
	CheckAvailability(checkIn time.Time, checkOut time.Time) (bool, error)
}
