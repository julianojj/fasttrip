package usecases

import (
	"github.com/julianojj/fastrip/internal/core/domain"
	"github.com/julianojj/fastrip/internal/core/exceptions"
)

type (
	RegisterRoom struct {
		roomRepository domain.RoomRepository
	}
	RegisterRoomInput struct {
		Category string  `json:"category" binding:"required" example:"Premium"`
		Price    float64 `json:"price" binding:"required" example:"2000"`
		Capacity int     `json:"capacity" binding:"required" example:"4"`
	}
	RegisterRoomOutput struct {
		RoomID string `json:"room_id" example:"71cf737c-228e-4973-8197-3c5cf83454a9"`
	}
)

func NewRegisterRoom(
	roomRepository domain.RoomRepository,
) *RegisterRoom {
	return &RegisterRoom{
		roomRepository,
	}
}

func (r *RegisterRoom) Execute(input *RegisterRoomInput) (*RegisterRoomOutput, error) {
	room := domain.NewRoom(input.Category, input.Price, input.Capacity)
	if err := room.Validate(); err != nil {
		return nil, err
	}
	existingRoom, err := r.roomRepository.FindByCategory(room.Category)
	if err != nil {
		return nil, err
	}
	if existingRoom != nil {
		return nil, exceptions.ErrRoomAlreadyExists
	}
	if err := r.roomRepository.Save(room); err != nil {
		return nil, err
	}
	return &RegisterRoomOutput{
		RoomID: room.ID,
	}, nil
}
