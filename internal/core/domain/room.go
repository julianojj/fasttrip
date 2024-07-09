package domain

import (
	"github.com/google/uuid"
)

type Room struct {
	ID       string
	Category string
	Price    float64
	Capacity int
}

func NewRoom(
	category string,
	price float64,
	capacity int,
) *Room {
	return &Room{
		ID:       uuid.NewString(),
		Category: category,
		Price:    price,
		Capacity: capacity,
	}
}
