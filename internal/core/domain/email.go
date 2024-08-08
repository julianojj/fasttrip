package domain

import (
	"net/mail"

	"github.com/julianojj/fasttrip/internal/core/exceptions"
)

type Email struct {
	Value string
}

func NewEmail(value string) *Email {
	return &Email{
		Value: value,
	}
}

func (e *Email) Validate() error {
	if isInvalidEmail(e.Value) {
		return exceptions.ErrInvalidEmail
	}
	return nil
}

func isInvalidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err != nil
}
