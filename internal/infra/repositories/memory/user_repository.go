package memory

import "github.com/julianojj/fastrip/internal/core/domain"

type UserRepositoryMemory struct {
	users []*domain.User
}

var _ domain.UserRepository = (*UserRepositoryMemory)(nil)

func NewUserRepositoryMemory() *UserRepositoryMemory {
	return &UserRepositoryMemory{
		users: make([]*domain.User, 0),
	}
}

func (r *UserRepositoryMemory) Save(user *domain.User) error {
	r.users = append(r.users, user)
	return nil
}

func (r *UserRepositoryMemory) FindByEmail(email string) (*domain.User, error) {
	for _, user := range r.users {
		if user.Email.Value == email {
			return user, nil
		}
	}
	return nil, nil
}

func (r *UserRepositoryMemory) FindByID(id string) (*domain.User, error) {
	for _, user := range r.users {
		if user.ID == id {
			return user, nil
		}
	}
	return nil, nil
}
