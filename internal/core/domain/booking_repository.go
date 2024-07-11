package domain

import "time"

type BookingRepository interface {
	Save(booking *Booking) error
	FindByID(id string) (*Booking, error)
	FindAll() ([]*Booking, error)
	CheckAvailability(checkIn time.Time, checkOut time.Time) (bool, error)
	Update(booking *Booking) error
}
