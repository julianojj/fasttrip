package usecases

import (
	"time"

	"github.com/julianojj/fastrip/internal/core/domain"
	"github.com/julianojj/fastrip/internal/core/exceptions"
)

type (
	GetBooking struct {
		roomRepository    domain.RoomRepository
		bookingRepository domain.BookingRepository
	}

	GetBookingOutput struct {
		BookingID    string    `json:"booking_id" example:"71cf737c-228e-4973-8197-3c5cf83454a9"`
		CheckIn      time.Time `json:"check_in" example:"2024-06-01T11:00:00Z"`
		CheckOut     time.Time `json:"check_out" example:"2024-06-04T13:00:00Z"`
		TotalGuests  int       `json:"total_guests" example:"2"`
		RoomCategory string    `json:"room_category" example:"Standard"`
		Overnight    int       `json:"overnight" example:"3"`
		TotalAmount  float64   `json:"total_amount" example:"300"`
	}
)

func NewGetBooking(
	roomRepository domain.RoomRepository,
	bookingRepository domain.BookingRepository,
) *GetBooking {
	return &GetBooking{
		roomRepository,
		bookingRepository,
	}
}

func (gb *GetBooking) Execute(bookingID string) (*GetBookingsOutput, error) {
	booking, err := gb.bookingRepository.FindByID(bookingID)
	if err != nil {
		return nil, err
	}
	if booking == nil {
		return nil, exceptions.ErrBookingNotFound
	}
	room, err := gb.roomRepository.FindByID(booking.RoomID)
	if err != nil {
		return nil, err
	}
	if room == nil {
		return nil, exceptions.ErrRoomNotFound
	}
	return &GetBookingsOutput{
		BookingID:    booking.ID,
		CheckIn:      booking.CheckIn,
		CheckOut:     booking.CheckOut,
		TotalGuests:  booking.TotalGuests,
		RoomCategory: room.Category,
		Overnight:    booking.CalculateOvernight(),
		TotalAmount:  booking.CalculateTotalAmount(room.Price),
	}, nil
}
