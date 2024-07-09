package domain

import (
	"github.com/google/uuid"
	"github.com/julianojj/fastrip/internal/core/exceptions"
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

func (r *Room) Validate() error {
	if r.Category == "" {
		return exceptions.ErrRequiredCategory
	}
	if r.Price <= 0 {
		return exceptions.ErrInvalidPrice
	}
	if r.Capacity <= 0 {
		return exceptions.ErrInvalidCapacity
	}
	return nil
}
