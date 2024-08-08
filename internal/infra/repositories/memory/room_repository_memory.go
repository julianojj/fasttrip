package memory

import (
	"github.com/julianojj/fasttrip/internal/core/domain"
)

type RoomRepositoryMemory struct {
	rooms []*domain.Room
}

var _ domain.RoomRepository = (*RoomRepositoryMemory)(nil)

func NewRoomRepositoryMemory() *RoomRepositoryMemory {
	return &RoomRepositoryMemory{
		rooms: make([]*domain.Room, 0),
	}
}

func (rrm *RoomRepositoryMemory) Save(room *domain.Room) error {
	rrm.rooms = append(rrm.rooms, room)
	return nil
}

func (rrm *RoomRepositoryMemory) FindByID(id string) (*domain.Room, error) {
	for _, room := range rrm.rooms {
		if room.ID == id {
			return room, nil
		}
	}
	return nil, nil
}

func (rrm *RoomRepositoryMemory) FindByCategory(category string) (*domain.Room, error) {
	for _, room := range rrm.rooms {
		if room.Category == category {
			return room, nil
		}
	}
	return nil, nil
}

func (rrm *RoomRepositoryMemory) FindAll() ([]*domain.Room, error) {
	return rrm.rooms, nil
}
