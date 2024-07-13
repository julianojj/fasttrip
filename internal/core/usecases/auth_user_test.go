package usecases

import (
	"testing"

	"github.com/julianojj/fastrip/internal/core/exceptions"
	"github.com/julianojj/fastrip/internal/infra/adapters"
	"github.com/julianojj/fastrip/internal/infra/repositories/memory"
	"github.com/stretchr/testify/assert"
)

func TestAuthUser(t *testing.T) {
	tests := []struct {
		name string
		fn   func(t *testing.T)
	}{
		{
			name: "Auth user",
			fn: func(t *testing.T) {
				userRepository := memory.NewUserRepositoryMemory()
				hash := adapters.NewBcrypt()
				sign := adapters.NewJWT()
				registerUser := NewRegisterUser(userRepository, hash)
				input := &RegisterUserInput{
					Name:     "John Doe",
					Email:    "johndoe@example.com",
					Password: "P4ssw0rd!",
				}
				registerUser.Execute(input)
				authUser := NewAuthUser(userRepository, hash, sign)
				inputAuthUser := &AuthUserInput{
					Email:    input.Email,
					Password: input.Password,
				}
				output, _ := authUser.Execute(inputAuthUser)
				assert.NotEmpty(t, output)
			},
		},
		{
			name: "Return exception if invalid email",
			fn: func(t *testing.T) {
				userRepository := memory.NewUserRepositoryMemory()
				hash := adapters.NewBcrypt()
				sign := adapters.NewJWT()
				authUser := NewAuthUser(userRepository, hash, sign)
				inputAuthUser := &AuthUserInput{
					Email:    "invalid_email",
					Password: "P4ssw0rd!",
				}
				_, err := authUser.Execute(inputAuthUser)
				assert.EqualError(t, err, exceptions.ErrInvalidLogin.Error())
			},
		},
		{
			name: "Return exception if invalid password",
			fn: func(t *testing.T) {
				userRepository := memory.NewUserRepositoryMemory()
				hash := adapters.NewBcrypt()
				sign := adapters.NewJWT()
				registerUser := NewRegisterUser(userRepository, hash)
				input := &RegisterUserInput{
					Name:     "John Doe",
					Email:    "johndoe@example.com",
					Password: "P4ssw0rd!",
				}
				registerUser.Execute(input)
				authUser := NewAuthUser(userRepository, hash, sign)
				inputAuthUser := &AuthUserInput{
					Email:    input.Email,
					Password: "invalid_password",
				}
				_, err := authUser.Execute(inputAuthUser)
				assert.EqualError(t, err, exceptions.ErrInvalidLogin.Error())
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.fn)
	}
}
