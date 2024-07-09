package usecases

import (
	"testing"

	"github.com/julianojj/fastrip/internal/core/exceptions"
	"github.com/julianojj/fastrip/internal/infra/repositories/memory"
	"github.com/stretchr/testify/assert"
)

func TestRegisterRoom(t *testing.T) {
	tests := []struct {
		name string
		fn   func(t *testing.T)
	}{
		{
			name: "register room",
			fn: func(t *testing.T) {
				roomRepository := memory.NewRoomRepositoryMemory()
				registerRoom := NewRegisterRoom(roomRepository)
				input := &RegisterRoomInput{
					Category: "Premium",
					Price:    2000,
					Capacity: 4,
				}
				output, _ := registerRoom.Execute(input)
				room, _ := roomRepository.FindByID(output.RoomID)
				assert.Equal(t, input.Category, room.Category)
				assert.Equal(t, input.Price, room.Price)
				assert.Equal(t, input.Capacity, room.Capacity)
			},
		},
		{
			name: "return error if room already exists",
			fn: func(t *testing.T) {
				roomRepository := memory.NewRoomRepositoryMemory()
				registerRoom := NewRegisterRoom(roomRepository)
				input := &RegisterRoomInput{
					Category: "Premium",
					Price:    2000,
					Capacity: 4,
				}
				registerRoom.Execute(input)
				_, err := registerRoom.Execute(input)
				assert.EqualError(t, err, exceptions.ErrRoomAlreadyExists.Error())
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, test.fn)
	}
}
