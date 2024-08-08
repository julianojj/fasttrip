package usecases

import (
	"time"

	"github.com/julianojj/fasttrip/internal/core/domain"
)

type (
	GetBookings struct {
		roomRepository    domain.RoomRepository
		bookingRepository domain.BookingRepository
	}

	GetBookingsOutput struct {
		BookingID    string    `json:"booking_id" example:"71cf737c-228e-4973-8197-3c5cf83454a9"`
		CheckIn      time.Time `json:"check_in" example:"2024-06-01T11:00:00Z"`
		CheckOut     time.Time `json:"check_out" example:"2024-06-04T13:00:00Z"`
		TotalGuests  int       `json:"total_guests" example:"2"`
		RoomCategory string    `json:"room_category" example:"Standard"`
		Overnight    int       `json:"overnight" example:"3"`
		TotalAmount  float64   `json:"total_amount" example:"300"`
	}
)

func NewGetBookings(
	roomRepository domain.RoomRepository,
	bookingRepository domain.BookingRepository,
) *GetBookings {
	return &GetBookings{
		roomRepository,
		bookingRepository,
	}
}

func (gb *GetBookings) Execute() ([]*GetBookingsOutput, error) {
	bookings, err := gb.bookingRepository.FindAll()
	if err != nil {
		return nil, err
	}
	outputs := make([]*GetBookingsOutput, len(bookings))
	for i, booking := range bookings {
		room, err := gb.roomRepository.FindByID(booking.RoomID)
		if err != nil {
			return nil, err
		}
		outputs[i] = &GetBookingsOutput{
			BookingID:    booking.ID,
			CheckIn:      booking.Period.Start,
			CheckOut:     booking.Period.End,
			TotalGuests:  booking.TotalGuests,
			RoomCategory: room.Category,
			Overnight:    booking.CalculateOvernight(),
			TotalAmount:  booking.CalculateTotalAmount(room.Price),
		}
	}
	return outputs, nil
}
