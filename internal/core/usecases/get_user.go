package usecases

import (
	"github.com/julianojj/fasttrip/internal/core/domain"
	"github.com/julianojj/fasttrip/internal/core/exceptions"
)

type (
	GetUser struct {
		userRepository domain.UserRepository
	}
	GetUserOutput struct {
		ID       string `json:"id" example:"3b51a987-029c-42cb-9450-9a066f025b9f"`
		Name     string `json:"name" example:"John Doe"`
		PlanType string `json:"plan_type" example:"free"`
	}
)

func NewGetUser(userRepository domain.UserRepository) *GetUser {
	return &GetUser{
		userRepository,
	}
}

func (g *GetUser) Execute(userID string) (*GetUserOutput, error) {
	user, err := g.userRepository.FindByID(userID)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, exceptions.ErrUserNotFound
	}
	return &GetUserOutput{
		ID:       user.ID,
		Name:     user.Name,
		PlanType: user.PlanType,
	}, nil
}
