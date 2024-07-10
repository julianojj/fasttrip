package domain

type RoomRepository interface {
	FindByID(id string) (*Room, error)
	FindByCategory(category string) (*Room, error)
	FindAll() ([]*Room, error)
	Save(room *Room) error
}
