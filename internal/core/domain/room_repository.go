package domain

type RoomRepository interface {
	FindByID(id string) (*Room, error)
}
