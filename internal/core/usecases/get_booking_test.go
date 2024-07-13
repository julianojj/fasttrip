package usecases

import (
	"testing"
	"time"

	"github.com/julianojj/fastrip/internal/core/exceptions"
	"github.com/julianojj/fastrip/internal/infra/repositories/memory"
	"github.com/stretchr/testify/assert"
)

func TestGetBookin(t *testing.T) {
	tests := []struct {
		name string
		fn   func(t *testing.T)
	}{
		{
			name: "get booking",
			fn: func(t *testing.T) {
				roomRepository := memory.NewRoomRepositoryMemory()
				bookingRepository := memory.NewBookingRepositoryMemory()
				makeBooking := NewMakeBooking(roomRepository, bookingRepository)
				getBooking := NewGetBooking(roomRepository, bookingRepository)
				checkIn := time.Now().UTC()
				checkOut := checkIn.Add(time.Hour * 24 * 3)
				input := &MakeBookingInput{
					RoomID:      "1",
					CheckIn:     checkIn,
					CheckOut:    checkOut,
					TotalGuests: 2,
				}
				output, _ := makeBooking.Execute(input)
				outputGetBooking, _ := getBooking.Execute(output.BookingID)
				assert.Equal(t, input.CheckIn, outputGetBooking.CheckIn)
				assert.Equal(t, input.CheckOut, outputGetBooking.CheckOut)
				assert.Equal(t, input.TotalGuests, outputGetBooking.TotalGuests)
				assert.Equal(t, "Standard", outputGetBooking.RoomCategory)
				assert.Equal(t, 3, outputGetBooking.Overnight)
				assert.Equal(t, 300.00, outputGetBooking.TotalAmount)
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
		{
			name: "return error if room not found",
			fn: func(t *testing.T) {
				// TODO
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, test.fn)
	}
}
