package usecases

import (
	"testing"
	"time"

	"github.com/julianojj/fasttrip/internal/core/domain"
	"github.com/julianojj/fasttrip/internal/core/exceptions"
	"github.com/julianojj/fasttrip/internal/infra/repositories/memory"
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
				room := domain.NewRoom("Standard", 700, 3)
				roomRepository.Save(room)
				bookingRepository := memory.NewBookingRepositoryMemory()
				makeBooking := NewMakeBooking(roomRepository, bookingRepository)
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
				output, err := makeBooking.Execute(input)
				assert.NoError(t, err)
				assert.NotNil(t, output.BookingID)
				assert.Equal(t, 3, output.Overnight)
				assert.Equal(t, 2100.00, output.TotalAmount)
			},
		},
		{
			name: "is overlapping",
			fn: func(t *testing.T) {
				roomRepository := memory.NewRoomRepositoryMemory()
				room := domain.NewRoom("Standard", 700, 3)
				roomRepository.Save(room)
				bookingRepository := memory.NewBookingRepositoryMemory()
				makeBooking := NewMakeBooking(roomRepository, bookingRepository)
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
				booking, _ := bookingRepository.FindByID(output.BookingID)
				booking.ConfirmBooking()
				bookingRepository.Update(booking)
				_, err := makeBooking.Execute(input)
				assert.EqualError(t, err, exceptions.ErrPeriodNotAllowed.Error())
			},
		},
		{
			name: "capacity exceeded",
			fn: func(t *testing.T) {
				roomRepository := memory.NewRoomRepositoryMemory()
				room := domain.NewRoom("Standard", 700, 3)
				roomRepository.Save(room)
				bookingRepository := memory.NewBookingRepositoryMemory()
				makeBooking := NewMakeBooking(roomRepository, bookingRepository)
				checkIn := time.Now().UTC().Add(time.Hour * 24 * 1)
				checkOut := checkIn.Add(time.Hour * 24 * 3)
				input := &MakeBookingInput{
					RoomID:      room.ID,
					CheckIn:     checkIn,
					CheckOut:    checkOut,
					TotalGuests: 4,
					Email:       "juliano@test.com",
					Whatsapp:    "43999999999",
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
