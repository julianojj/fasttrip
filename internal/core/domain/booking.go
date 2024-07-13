package domain

import (
	"time"

	"github.com/google/uuid"
)

type Booking struct {
	ID          string
	RoomID      string
	Period      *Period
	Status      string
	TotalGuests int
	Email       string
	Whatsapp    string
}

func NewBooking(
	roomID string,
	checkIn time.Time,
	checkOut time.Time,
	totalGuests int,
) *Booking {
	return &Booking{
		ID:     uuid.NewString(),
		RoomID: roomID,
		Period: &Period{
			Start: checkIn,
			End:   checkOut,
		},
		TotalGuests: totalGuests,
		Status:      "PENDING_PAYMENT",
	}
}

func (b *Booking) Validate() error {
	if err := b.Period.Validate(); err != nil {
		return err
	}
	return nil
}

func (b *Booking) CalculateOvernight() int {
	return b.Period.DurationInDays()
}

func (b *Booking) CalculateTotalAmount(amount float64) float64 {
	return float64(b.CalculateOvernight()) * amount
}

func (b *Booking) ConfirmBooking() {
	b.Status = "CONFIRMED"
}
