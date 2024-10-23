package domain

import (
	"context"
	"github.com/google/uuid"
)

const DefaultRole = "user"

type UserAuth struct {
	ID         uuid.UUID
	Username   string
	Password   string
	HashedPass string
	Role       string
}

type IAuthRepository interface {
	Register(ctx context.Context, authInfo *User) (uuid.UUID, error)
	GetByUsername(ctx context.Context, username string) (*UserAuth, error)
}

type IAuthService interface {
	Login(ctx context.Context, authInfo *UserAuth) (string, error)
	Register(ctx context.Context, authInfo *User) (string, error)
}
