package usecases

import (
	"time"

	"github.com/julianojj/fastrip/internal/core/domain"
	"github.com/julianojj/fastrip/internal/core/exceptions"
)

type (
	MakeBooking struct {
		roomRepository    domain.RoomRepository
		bookingRepository domain.BookingRepository
	}
	MakeBookingInput struct {
		RoomID      string    `json:"room_id" binding:"required" example:"1"`
		CheckIn     time.Time `json:"check_in" binding:"required" example:"2024-06-01T11:00:00Z"`
		CheckOut    time.Time `json:"check_out" binding:"required" example:"2024-06-04T13:00:00Z"`
		TotalGuests int       `json:"total_guests" binding:"required" example:"2"`
	}
	MakeBookingOutput struct {
		BookingID   string  `json:"booking_id" example:"71cf737c-228e-4973-8197-3c5cf83454a9"`
		Overnight   int     `json:"overnight" example:"3"`
		TotalAmount float64 `json:"total_amount" example:"300"`
	}
)

func NewMakeBooking(
	roomRepository domain.RoomRepository,
	bookingRepository domain.BookingRepository,
) *MakeBooking {
	return &MakeBooking{
		roomRepository,
		bookingRepository,
	}
}

func (mb *MakeBooking) Execute(input *MakeBookingInput) (*MakeBookingOutput, error) {
	existingRoom, err := mb.roomRepository.FindByID(input.RoomID)
	if err != nil {
		return nil, err
	}
	if existingRoom == nil {
		return nil, exceptions.ErrRoomNotFound
	}
	isExceededCapacity := input.TotalGuests > existingRoom.Capacity
	if isExceededCapacity {
		return nil, exceptions.ErrCapacityExceeded
	}
	isAvailable, err := mb.bookingRepository.CheckAvailability(input.CheckIn, input.CheckOut)
	if err != nil {
		return nil, err
	}
	if !isAvailable {
		return nil, exceptions.ErrPeriodNotAllowed
	}
	booking := domain.NewBooking(input.RoomID, input.CheckIn, input.CheckOut, input.TotalGuests)
	if err := mb.bookingRepository.Save(booking); err != nil {
		return nil, err
	}
	return &MakeBookingOutput{
		BookingID:   booking.ID,
		Overnight:   booking.CalculateOvernight(),
		TotalAmount: booking.CalculateTotalAmount(existingRoom.Price),
	}, nil
}
