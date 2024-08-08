package usecases

import "github.com/julianojj/fasttrip/internal/core/domain"

type (
	GetRooms struct {
		roomRepository domain.RoomRepository
	}
	GetRoomsOutput struct {
		RoomID   string  `json:"room_id" example:"71cf737c-228e-4973-8197-3c5cf83454a9"`
		Category string  `json:"category" example:"Premium"`
		Price    float64 `json:"price" example:"2000"`
		Capacity int     `json:"capacity" example:"4"`
	}
)

func NewGetRooms(
	roomRepository domain.RoomRepository,
) *GetRooms {
	return &GetRooms{
		roomRepository,
	}
}

func (g *GetRooms) Execute() ([]*GetRoomsOutput, error) {
	rooms, err := g.roomRepository.FindAll()
	if err != nil {
		return nil, err
	}
	outputs := make([]*GetRoomsOutput, len(rooms))
	for i, room := range rooms {
		outputs[i] = &GetRoomsOutput{
			RoomID:   room.ID,
			Category: room.Category,
			Price:    room.Price,
			Capacity: room.Capacity,
		}
	}
	return outputs, nil
}
