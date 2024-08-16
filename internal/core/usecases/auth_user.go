package usecases

import (
	"time"

	"github.com/julianojj/fasttrip/internal/core/domain"
	"github.com/julianojj/fasttrip/internal/core/exceptions"
	"github.com/julianojj/fasttrip/internal/infra/adapters"
)

type (
	AuthUser struct {
		userRepository domain.UserRepository
		hash           adapters.Hash
		sign           adapters.Sign
	}
	AuthUserInput struct {
		Email    string `json:"email" binding:"required" example:"juliano@test.com"`
		Password string `json:"password" binding:"required" example:"P4ssw0rd!"`
	}
	AuthUserOutput struct {
		UserID string `json:"user_id" binding:"required" example:"07410ea5-98e3-4aab-8d23-35112a67197f"`
		Token  string `json:"token" binding:"required" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjA4Mjk5MDAsInN1YiI6IjUwNmU4Mjc4LTNiMzktNDI0ZS04NGU4LWMzYTE4NzcwNzBiNyJ9.AlUCK04QFIcGwlRw0e29fRMYlzZ3V979EH3pWlFVA1g"`
	}
)

func NewAuthUser(
	userRepository domain.UserRepository,
	hash adapters.Hash,
	sign adapters.Sign,
) *AuthUser {
	return &AuthUser{
		userRepository,
		hash,
		sign,
	}
}

func (a *AuthUser) Execute(input *AuthUserInput) (*AuthUserOutput, error) {
	user, err := a.userRepository.FindByEmail(input.Email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, exceptions.ErrInvalidLogin
	}
	isMatchPassword := a.hash.DecryptPassword([]byte(user.Password), input.Password)
	if !isMatchPassword {
		return nil, exceptions.ErrInvalidLogin
	}
	expiresIn := time.Now().Add(time.Hour * 24).Unix()
	token, err := a.sign.Encode(user.ID, expiresIn)
	if err != nil {
		return nil, err
	}
	return &AuthUserOutput{
		UserID: user.ID,
		Token:  token,
	}, nil
}
