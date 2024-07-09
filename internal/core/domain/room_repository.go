package domain

type RoomRepository interface {
	FindByID(id string) (*Room, error)
	FindByCategory(category string) (*Room, error)
	Save(room *Room) error
}
