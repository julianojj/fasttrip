package memory

import (
	"github.com/julianojj/fastrip/internal/domain"
)

type RoomRepositoryMemory struct {
	rooms []*domain.Room
}

var _ domain.RoomRepository = (*RoomRepositoryMemory)(nil)

func NewRoomRepositoryMemory() *RoomRepositoryMemory {
	return &RoomRepositoryMemory{
		rooms: []*domain.Room{
			{
				ID:       "1",
				Category: "Standard",
				Price:    100,
				Capacity: 3,
			},
		},
	}
}

func (rrm *RoomRepositoryMemory) FindByID(id string) (*domain.Room, error) {
	for _, room := range rrm.rooms {
		if room.ID == id {
			return room, nil
		}
	}
	return nil, nil
}
