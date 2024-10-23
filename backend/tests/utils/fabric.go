package utils

import (
	"github.com/google/uuid"
	"ppo/domain"
)

type AuthInfoObjectMother struct {
}

func (a *AuthInfoObjectMother) Default() *domain.UserAuth {
	return &domain.UserAuth{
		ID:         uuid.New(),
		Username:   "username",
		Password:   "user",
		HashedPass: "hashedPass",
		Role:       "user",
	}
}

func (a *AuthInfoObjectMother) UsernameNotFound() *domain.UserAuth {
	return &domain.UserAuth{
		ID:         uuid.New(),
		Username:   "notFound",
		Password:   "name",
		HashedPass: "hashedPass",
		Role:       "user",
	}
}
