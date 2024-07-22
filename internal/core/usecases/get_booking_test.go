package usecases

import (
	"testing"
	"time"

	"github.com/julianojj/fastrip/internal/core/domain"
	"github.com/julianojj/fastrip/internal/core/exceptions"
	"github.com/julianojj/fastrip/internal/infra/repositories/memory"
	"github.com/stretchr/testify/assert"
)

func TestGetBooking(t *testing.T) {
	tests := []struct {
		name string
		fn   func(t *testing.T)
	}{
		{
			name: "get booking",
			fn: func(t *testing.T) {
				roomRepository := memory.NewRoomRepositoryMemory()
				room := domain.NewRoom("Standard", 700, 3)
				roomRepository.Save(room)
				bookingRepository := memory.NewBookingRepositoryMemory()
				makeBooking := NewMakeBooking(roomRepository, bookingRepository)
				getBooking := NewGetBooking(roomRepository, bookingRepository)
				checkIn := time.Now().UTC().Add(time.Hour * 24 * 1)
				checkOut := checkIn.Add(time.Hour * 24 * 3)
				input := &MakeBookingInput{
					RoomID:      room.ID,
					CheckIn:     checkIn,
					CheckOut:    checkOut,
					TotalGuests: 2,
					Email:       "juliano@test.com",
					Whatsapp:    "43999999999",
				}
				output, _ := makeBooking.Execute(input)
				outputGetBooking, _ := getBooking.Execute(output.BookingID)
				assert.Equal(t, input.CheckIn, outputGetBooking.CheckIn)
				assert.Equal(t, input.CheckOut, outputGetBooking.CheckOut)
				assert.Equal(t, input.TotalGuests, outputGetBooking.TotalGuests)
				assert.Equal(t, "Standard", outputGetBooking.RoomCategory)
				assert.Equal(t, 3, outputGetBooking.Overnight)
				assert.Equal(t, 2100.00, outputGetBooking.TotalAmount)
			},
		},
		{
			name: "return error if booking not found",
			fn: func(t *testing.T) {
				roomRepository := memory.NewRoomRepositoryMemory()
				bookingRepository := memory.NewBookingRepositoryMemory()
				getBooking := NewGetBooking(roomRepository, bookingRepository)
				bookingID := "booking_id"
				_, err := getBooking.Execute(bookingID)
				assert.EqualError(t, err, exceptions.ErrBookingNotFound.Error())
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, test.fn)
	}
}
