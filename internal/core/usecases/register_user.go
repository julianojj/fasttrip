package usecases

import (
	"log/slog"

	"github.com/julianojj/fasttrip/internal/core/domain"
	"github.com/julianojj/fasttrip/internal/core/exceptions"
	"github.com/julianojj/fasttrip/internal/infra/adapters"
)

type (
	RegisterUser struct {
		userRepository domain.UserRepository
		hash           adapters.Hash
	}
	RegisterUserInput struct {
		Name     string `json:"name" binding:"required" example:"Juliano"`
		Email    string `json:"email" binding:"required" example:"juliano@test.com"`
		Password string `json:"password" binding:"required" example:"P4ssw0rd!"`
	}
	RegisterUserOutput struct {
		ID string `json:"id" binding:"required" example:"506e8278-3b39-424e-84e8-c3a1877070b7"`
	}
)

func NewRegisterUser(
	userRepository domain.UserRepository,
	hash adapters.Hash,
) *RegisterUser {
	return &RegisterUser{
		userRepository,
		hash,
	}
}

func (u *RegisterUser) Execute(input *RegisterUserInput) (*RegisterUserOutput, error) {
	existingUser, err := u.userRepository.FindByEmail(input.Email)
	if err != nil {
		slog.Error(
			"failed to find user by email",
			"error", err,
		)
		return nil, err
	}
	if existingUser != nil {
		slog.Error(
			"email already registered",
		)
		return nil, exceptions.ErrEmailAlreadyExists
	}
	user := domain.NewUser(input.Name, input.Email, input.Password)
	if err := user.Validate(); err != nil {
		slog.Error(
			"invalid user data",
			"error", err,
		)
		return nil, err
	}
	encryptedPassword, err := u.hash.EncryptPassword(input.Password)
	if err != nil {
		slog.Error(
			"failed to encrypt password",
			"error", err,
		)
		return nil, err
	}
	user.UpdatePassword(string(encryptedPassword))
	if err := u.userRepository.Save(user); err != nil {
		slog.Error(
			"failed to save user",
			"error", err,
		)
		return nil, err
	}
	slog.Info(
		"user registered successfully",
		"user_id", user.ID,
	)
	return &RegisterUserOutput{
		ID: user.ID,
	}, nil
}
