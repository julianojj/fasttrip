package domain

import (
	"net/mail"
	"time"
	"unicode"

	"github.com/google/uuid"
	"github.com/julianojj/fastrip/internal/core/exceptions"
)

type User struct {
	ID        string
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
}

func NewUser(name, email, password string) *User {
	return &User{
		ID:        uuid.NewString(),
		Name:      name,
		Email:     email,
		Password:  password,
		CreatedAt: time.Now().UTC(),
	}
}

func (u *User) Validate() error {
	if isInvalidEmail(u.Email) {
		return exceptions.ErrInvalidEmail
	}
	if !isStrongPassword(u.Password) {
		return exceptions.ErrInvalidPassword
	}
	return nil
}

func isInvalidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err != nil
}

func (u *User) UpdatePassword(password string) {
	u.Password = password
}

func isStrongPassword(password string) bool {
	var (
		hasUpper   bool
		hasLower   bool
		hasDigit   bool
		hasSpecial bool
	)
	if len(password) < 8 {
		return false
	}
	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsDigit(char):
			hasDigit = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}
	return hasUpper && hasLower && hasDigit && hasSpecial
}
