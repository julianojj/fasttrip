package usecases

import (
	"testing"

	"github.com/julianojj/fastrip/internal/core/exceptions"
	"github.com/julianojj/fastrip/internal/infra/adapters"
	"github.com/julianojj/fastrip/internal/infra/repositories/memory"
	"github.com/stretchr/testify/assert"
)

func TestRegisterUser(t *testing.T) {
	tests := []struct {
		name string
		fn   func(t *testing.T)
	}{
		{
			name: "Register new user",
			fn: func(t *testing.T) {
				userRepository := memory.NewUserRepositoryMemory()
				hash := adapters.NewBcrypt()
				registerUser := NewRegisterUser(userRepository, hash)
				input := &RegisterUserInput{
					Name:     "John Doe",
					Email:    "johndoe@example.com",
					Password: "P4ssw0rd!",
				}
				registerUser.Execute(input)
				user, _ := userRepository.FindByEmail(input.Email)
				assert.Equal(t, user.Name, input.Name)
				assert.Equal(t, user.Email.Value, input.Email)
				assert.NotEmpty(t, user.ID)
				assert.NotEqual(t, user.Password, input.Password)
			},
		},
		{
			name: "Returtn exception if user already registered",
			fn: func(t *testing.T) {
				userRepository := memory.NewUserRepositoryMemory()
				hash := adapters.NewBcrypt()
				registerUser := NewRegisterUser(userRepository, hash)
				input := &RegisterUserInput{
					Name:     "John Doe",
					Email:    "johndoe@example.com",
					Password: "P4ssw0rd!",
				}
				registerUser.Execute(input)
				_, err := registerUser.Execute(input)
				assert.Error(t, err)
				assert.Equal(t, exceptions.ErrEmailAlreadyExists, err)
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, test.fn)
	}
}
