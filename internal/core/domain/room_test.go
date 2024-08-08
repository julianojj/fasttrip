package domain

import (
	"testing"

	"github.com/julianojj/fasttrip/internal/core/exceptions"
	"github.com/stretchr/testify/assert"
)

func TestRoom(t *testing.T) {
	tests := []struct {
		name string
		fn   func(t *testing.T)
	}{
		{
			name: "new room",
			fn: func(t *testing.T) {
				room := NewRoom("Premium", 2000, 4)
				assert.NotNil(t, room)
			},
		},
		{
			name: "return error if empty category",
			fn: func(t *testing.T) {
				room := NewRoom("", 2000, 4)
				err := room.Validate()
				assert.ErrorIs(t, err, exceptions.ErrRequiredCategory)
			},
		},
		{
			name: "return error if price is negative",
			fn: func(t *testing.T) {
				room := NewRoom("Premium", -2000, 4)
				err := room.Validate()
				assert.ErrorIs(t, err, exceptions.ErrInvalidPrice)
			},
		},
		{
			name: "return error if category is negative",
			fn: func(t *testing.T) {
				room := NewRoom("Premium", 2000, -4)
				err := room.Validate()
				assert.ErrorIs(t, err, exceptions.ErrInvalidCapacity)
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, test.fn)
	}
}
