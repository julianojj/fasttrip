package domain

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGuest(t *testing.T) {
	tests := []struct {
		name string
		fn   func(t *testing.T)
	}{
		{
			name: "New guest",
			fn: func(t *testing.T) {
				guest := NewGuest("1", "John Doe", "Doe", "Male", "CPF", "12345678901", "johndoe@example.com", "123456789", time.Date(1990, time.January, 1, 0, 0, 0, 0, time.UTC))
				assert.NotNil(t, guest)
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, test.fn)
	}
}
