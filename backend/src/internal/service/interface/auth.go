package _interface

import (
	"backend/src/internal/model/dto"
	"context"
)

type IAuthService interface {
	SignIn(request *dto.SignInRequest) error
	LogIn(ctx context.Context, request *dto.LogInRequest) (string, error)
}
