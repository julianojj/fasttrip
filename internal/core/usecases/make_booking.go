package usecases

import (
	"log/slog"
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
		Email       string    `json:"email" binding:"required" example:"juliano@test.com"`
		Whatsapp    string    `json:"whatsapp" binding:"required" example:"43999999999"`
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
		slog.Error(
			"failed to find room by ID",
			"error", err,
		)
		return nil, err
	}
	if existingRoom == nil {
		slog.Error(
			"room not found",
			"room_id", input.RoomID,
		)
		return nil, exceptions.ErrRoomNotFound
	}
	isExceededCapacity := input.TotalGuests > existingRoom.Capacity
	if isExceededCapacity {
		slog.Error(
			"capacity exceeded",
			"room_id", input.RoomID,
			"capacity", existingRoom.Capacity,
			"total_guests", input.TotalGuests,
		)
		return nil, exceptions.ErrCapacityExceeded
	}
	isAvailable, err := mb.bookingRepository.CheckAvailability(input.CheckIn, input.CheckOut)
	if err != nil {
		slog.Error(
			"failed to check availability",
			"error", err,
		)
		return nil, err
	}
	if !isAvailable {
		slog.Error(
			"period not allowed",
			"room_id", input.RoomID,
			"check_in", input.CheckIn,
			"check_out", input.CheckOut,
		)
		return nil, exceptions.ErrPeriodNotAllowed
	}
	booking := domain.NewBooking(input.RoomID, input.CheckIn, input.CheckOut, input.TotalGuests, input.Email, input.Whatsapp)
	if err := booking.Validate(); err != nil {
		slog.Error(
			"invalid booking data",
			"error", err,
		)
		return nil, err
	}
	if err := mb.bookingRepository.Save(booking); err != nil {
		slog.Error(
			"failed to save booking",
			"error", err,
		)
		return nil, err
	}
	slog.Info(
		"success maked booking",
		"booking_id", booking.ID,
	)
	return &MakeBookingOutput{
		BookingID:   booking.ID,
		Overnight:   booking.CalculateOvernight(),
		TotalAmount: booking.CalculateTotalAmount(existingRoom.Price),
	}, nil
}
