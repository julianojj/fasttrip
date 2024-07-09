package domain

import (
	"time"

	"github.com/google/uuid"
)

type Booking struct {
	ID          string
	RoomID      string
	CheckIn     time.Time
	CheckOut    time.Time
	Status      string
	TotalGuests int
}

func NewBooking(
	roomID string,
	checkIn time.Time,
	checkOut time.Time,
	totalGuests int,
) *Booking {
	return &Booking{
		ID:          uuid.NewString(),
		RoomID:      roomID,
		CheckIn:     checkIn,
		CheckOut:    checkOut,
		TotalGuests: totalGuests,
		Status:      "PENDING_PAYMENT",
	}
}

func (b *Booking) CalculateOvernight() int {
	return int(b.CheckOut.Sub(b.CheckIn).Hours() / 24)
}

func (b *Booking) CalculateTotalAmount(amount float64) float64 {
	return float64(b.CalculateOvernight()) * amount
}

func (b *Booking) ConfirmBooking() {
	b.Status = "CONFIRMED"
}
