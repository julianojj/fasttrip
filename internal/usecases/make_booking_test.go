package usecases

import (
	"testing"
	"time"

	"github.com/julianojj/fastrip/internal/exceptions"
	"github.com/julianojj/fastrip/internal/infra/repositories/memory"
	"github.com/stretchr/testify/assert"
)

func TestMakeBooking(t *testing.T) {
	tests := []struct {
		name string
		fn   func(t *testing.T)
	}{
		{
			name: "make booking",
			fn: func(t *testing.T) {
				roomRepository := memory.NewRoomRepositoryMemory()
				bookingRepository := memory.NewBookingRepositoryMemory()
				makeBooking := NewMakeBooking(roomRepository, bookingRepository)
				checkIn := time.Now().UTC()
				checkOut := checkIn.Add(time.Hour * 24 * 3)
				input := &MakeBookingInput{
					RoomID:      "1",
					CheckIn:     checkIn,
					CheckOut:    checkOut,
					TotalGuests: 2,
				}
				output, err := makeBooking.Execute(input)
				assert.NoError(t, err)
				assert.NotNil(t, output.BookingID)
				assert.Equal(t, 3, output.Overnight)
			},
		},
		{
			name: "is overlapping",
			fn: func(t *testing.T) {
				roomRepository := memory.NewRoomRepositoryMemory()
				bookingRepository := memory.NewBookingRepositoryMemory()
				makeBooking := NewMakeBooking(roomRepository, bookingRepository)
				checkIn := time.Now().UTC()
				checkOut := checkIn.Add(time.Hour * 24 * 3)
				input := &MakeBookingInput{
					RoomID:      "1",
					CheckIn:     checkIn,
					CheckOut:    checkOut,
					TotalGuests: 2,
				}
				makeBooking.Execute(input)
				output, err := makeBooking.Execute(input)
				assert.EqualError(t, err, exceptions.ErrPeriodNotAllowed.Error())
				assert.Nil(t, output)

			},
		},
		{
			name: "capacity exceeded",
			fn: func(t *testing.T) {
				roomRepository := memory.NewRoomRepositoryMemory()
				bookingRepository := memory.NewBookingRepositoryMemory()
				makeBooking := NewMakeBooking(roomRepository, bookingRepository)
				checkIn := time.Now().UTC()
				checkOut := checkIn.Add(time.Hour * 24 * 3)
				input := &MakeBookingInput{
					RoomID:      "1",
					CheckIn:     checkIn,
					CheckOut:    checkOut,
					TotalGuests: 4,
				}
				output, err := makeBooking.Execute(input)
				assert.EqualError(t, err, exceptions.ErrCapacityExceeded.Error())
				assert.Nil(t, output)
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, test.fn)
	}
}
