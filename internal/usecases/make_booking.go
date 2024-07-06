package usecases

import (
	"time"

	"github.com/julianojj/fastrip/internal/domain"
	"github.com/julianojj/fastrip/internal/exceptions"
)

type (
	MakeBooking struct {
		roomRepository    domain.RoomRepository
		bookingRepository domain.BookingRepository
	}
	MakeBookingInput struct {
		RoomID      string    `json:"room_id"`
		CheckIn     time.Time `json:"check_in"`
		CheckOut    time.Time `json:"check_out"`
		TotalGuests int       `json:"total_guests"`
	}
	MakeBookingOutput struct {
		BookingID   string  `json:"booking_id"`
		Overnight   int     `json:"overnight"`
		TotalAmount float64 `json:"total_amount"`
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
