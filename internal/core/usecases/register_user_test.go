package usecases

import (
	"testing"

	"github.com/julianojj/fasttrip/internal/core/exceptions"
	"github.com/julianojj/fasttrip/internal/infra/adapters"
	"github.com/julianojj/fasttrip/internal/infra/repositories/memory"
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
				getUser := NewGetUser(userRepository)
				input := &RegisterUserInput{
					Name:     "John Doe",
					Email:    "johndoe@example.com",
					Password: "P4ssw0rd!",
				}
				output, _ := registerUser.Execute(input)
				user, _ := getUser.Execute(output.ID)
				assert.NotEmpty(t, user.ID)
				assert.Equal(t, input.Name, user.Name)
				assert.Equal(t, "free", user.PlanType)
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
