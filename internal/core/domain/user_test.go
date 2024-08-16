package domain

import (
	"testing"

	"github.com/julianojj/fasttrip/internal/core/exceptions"
	"github.com/stretchr/testify/assert"
)

func TestUser(t *testing.T) {
	tests := []struct {
		name string
		fn   func(t *testing.T)
	}{
		{
			name: "new user",
			fn: func(t *testing.T) {
				user := NewUser("John Doe", "johndoe@example.com", "P4ssw0rd!")
				assert.NotNil(t, user)
			},
		},
		{
			name: "upgrade plan pro",
			fn: func(t *testing.T) {
				user := NewUser("John Doe", "johndoe@example.com", "P4ssw0rd!")
				user.UpgradePlanToPro()
				assert.Equal(t, user.PlanType, "pro")
			},
		},
		{
			name: "update password",
			fn: func(t *testing.T) {
				user := NewUser("John Doe", "johndoe@example.com", "P4ssw0rd!")
				user.UpdatePassword("new_password")
				assert.Equal(t, user.Password, "new_password")
			},
		},
		{
			name: "return error if invalid email",
			fn: func(t *testing.T) {
				user := NewUser("John Doe", "invalid_email", "P4ssw0rd!")
				err := user.Validate()
				assert.Equal(t, exceptions.ErrInvalidEmail, err)
			},
		},
		{
			name: "return error if invalid password when not has uppercase",
			fn: func(t *testing.T) {
				user := NewUser("John Doe", "johndoe@example.com", "p4ssw0rd!")
				err := user.Validate()
				assert.Equal(t, exceptions.ErrInvalidPassword, err)
			},
		},
		{
			name: "return error if invalid password when not has lowercase",
			fn: func(t *testing.T) {
				user := NewUser("John Doe", "johndoe@example.com", "P4SSW0RD!")
				err := user.Validate()
				assert.Equal(t, exceptions.ErrInvalidPassword, err)
			},
		},
		{
			name: "return error if invalid password when not has numeric",
			fn: func(t *testing.T) {
				user := NewUser("John Doe", "johndoe@example.com", "Password!")
				err := user.Validate()
				assert.Equal(t, exceptions.ErrInvalidPassword, err)
			},
		},
		{
			name: "return error if invalid password when not has numeric",
			fn: func(t *testing.T) {
				user := NewUser("John Doe", "johndoe@example.com", "Password!")
				err := user.Validate()
				assert.Equal(t, exceptions.ErrInvalidPassword, err)
			},
		},
		{
			name: "return error if invalid password when not has special character",
			fn: func(t *testing.T) {
				user := NewUser("John Doe", "johndoe@example.com", "P4ssw0rd")
				err := user.Validate()
				assert.Equal(t, exceptions.ErrInvalidPassword, err)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.fn)
	}
}
