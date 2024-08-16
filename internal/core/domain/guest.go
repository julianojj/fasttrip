package domain

import (
	"time"

	"github.com/google/uuid"
)

type Guest struct {
	ID             string
	BookingID      string
	Name           string
	LastName       string
	Gender         string
	BirthDate      time.Time
	DocumentType   string
	DocumentNumber string
	Email          string
	WhatsApp       string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func NewGuest(
	bookingID,
	name,
	lastName,
	gender,
	documentType,
	documentNumber,
	email,
	whatsapp string,
	birthDate time.Time,
) *Guest {
	return &Guest{
		ID:             uuid.NewString(),
		BookingID:      bookingID,
		Name:           name,
		LastName:       lastName,
		Gender:         gender,
		DocumentType:   documentType,
		DocumentNumber: documentNumber,
		Email:          email,
		WhatsApp:       whatsapp,
		BirthDate:      birthDate,
		CreatedAt:      time.Now().UTC(),
		UpdatedAt:      time.Now().UTC(),
	}
}
