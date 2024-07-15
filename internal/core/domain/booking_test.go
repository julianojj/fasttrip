package domain

import (
	"testing"
	"time"

	"github.com/julianojj/fastrip/internal/core/exceptions"
	"github.com/stretchr/testify/assert"
)

func TestBooking(t *testing.T) {
	tests := []struct {
		name string
		fn   func(t *testing.T)
	}{
		{
			name: "new booking",
			fn: func(t *testing.T) {
				checkIn := time.Now().UTC()
				checkOut := checkIn.Add(time.Hour * 24 * 3)
				booking := NewBooking("1", checkIn, checkOut, 2, "juliano@test.com", "43999999999")
				assert.NotNil(t, booking)
			},
		},
		{
			name: "calculate overnight",
			fn: func(t *testing.T) {
				checkIn := time.Now().UTC()
				checkOut := checkIn.Add(time.Hour * 24 * 3)
				booking := NewBooking("1", checkIn, checkOut, 2, "juliano@test.com", "43999999999")
				assert.Equal(t, 3, booking.CalculateOvernight())
			},
		},
		{
			name: "calculate total amount",
			fn: func(t *testing.T) {
				checkIn := time.Now().UTC()
				checkOut := checkIn.Add(time.Hour * 24 * 3)
				booking := NewBooking("1", checkIn, checkOut, 2, "juliano@test.com", "43999999999")
				assert.Equal(t, 3000.00, booking.CalculateTotalAmount(1000))
			},
		},
		{
			name: "return erro if invalid checkin",
			fn: func(t *testing.T) {
				checkIn := time.Now().UTC().Add(-time.Hour * 24 * 3)
				checkOut := checkIn
				booking := NewBooking("1", checkIn, checkOut, 2, "juliano@test.com", "43999999999")
				err := booking.Validate()
				assert.EqualError(t, err, exceptions.ErrInvalidPeriod.Error())
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, test.fn)
	}
}
