package usecases

import (
	"time"

	"github.com/julianojj/fasttrip/internal/core/domain"
	"github.com/julianojj/fasttrip/internal/core/exceptions"
)

type (
	GetBooking struct {
		roomRepository    domain.RoomRepository
		bookingRepository domain.BookingRepository
	}

	GetBookingOutput struct {
		BookingID    string                    `json:"booking_id" example:"71cf737c-228e-4973-8197-3c5cf83454a9"`
		CheckIn      time.Time                 `json:"check_in" example:"2024-06-01T11:00:00Z"`
		CheckOut     time.Time                 `json:"check_out" example:"2024-06-04T13:00:00Z"`
		TotalGuests  int                       `json:"total_guests" example:"2"`
		RoomCategory string                    `json:"room_category" example:"Standard"`
		Overnight    int                       `json:"overnight" example:"3"`
		TotalAmount  float64                   `json:"total_amount" example:"300"`
		Email        string                    `json:"email" example:"juliano@test.com"`
		Whatsapp     string                    `json:"whatsapp" example:"43999999999"`
		Guests       []*GetBookingGuestsOutput `json:"guests"`
	}

	GetBookingGuestsOutput struct {
		GuestID        string    `json:"guest_id" example:"326ebf92-1597-4850-b21e-3c1c8d7e5c6e"`
		Name           string    `json:"name" example:"John"`
		LastName       string    `json:"last_name" example:"Doe"`
		Gender         string    `json:"gender" example:"Masculino"`
		BirthDate      time.Time `json:"birth_date" example:"1990-01-01T00:00:00Z"`
		DocumentType   string    `json:"document_type" example:"CPF"`
		DocumentNumber string    `json:"document_number" example:"12345678901"`
		Email          string    `json:"email" example:"johndoe@example.com"`
		Whatsapp       string    `json:"whatsapp" example:"43888888888"`
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

func (gb *GetBooking) Execute(bookingID string) (*GetBookingOutput, error) {
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
	var guests []*GetBookingGuestsOutput
	for _, guest := range booking.Guests {
		guests = append(guests, &GetBookingGuestsOutput{
			GuestID:        guest.ID,
			Name:           guest.Name,
			LastName:       guest.LastName,
			Gender:         guest.Gender,
			BirthDate:      guest.BirthDate,
			DocumentType:   guest.DocumentType,
			DocumentNumber: guest.DocumentNumber,
			Email:          guest.Email,
			Whatsapp:       guest.WhatsApp,
		})
	}
	return &GetBookingOutput{
		BookingID:    booking.ID,
		CheckIn:      booking.Period.Start,
		CheckOut:     booking.Period.End,
		TotalGuests:  booking.TotalGuests,
		RoomCategory: room.Category,
		Overnight:    booking.CalculateOvernight(),
		TotalAmount:  booking.CalculateTotalAmount(room.Price),
		Email:        booking.Email.Value,
		Whatsapp:     booking.Whatsapp,
		Guests:       guests,
	}, nil
}
