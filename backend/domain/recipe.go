package domain

import (
	"context"
	"github.com/google/uuid"
	"ppo/services/dto"
)

type Recipe struct {
	ID               uuid.UUID
	SaladID          uuid.UUID
	Status           int
	NumberOfServings int
	TimeToCook       int
	Rating           float32
}

type IRecipeRepository interface {
	Create(ctx context.Context, recipe *Recipe) (uuid.UUID, error)
	GetById(ctx context.Context, id uuid.UUID) (*Recipe, error)
	GetBySaladId(ctx context.Context, saladId uuid.UUID) (*Recipe, error)
	GetAll(ctx context.Context, filter *dto.RecipeFilter, page int) ([]*Recipe, error)
	Update(ctx context.Context, recipe *Recipe) error
	DeleteById(ctx context.Context, id uuid.UUID) error
}

type IRecipeService interface {
	Create(ctx context.Context, user *Recipe) (uuid.UUID, error)
	GetById(ctx context.Context, id uuid.UUID) (*Recipe, error)
	GetBySaladId(ctx context.Context, saladId uuid.UUID) (*Recipe, error)
	GetAll(ctx context.Context, filter *dto.RecipeFilter, page int) ([]*Recipe, error)
	Update(ctx context.Context, recipe *Recipe) error
	DeleteById(ctx context.Context, id uuid.UUID) error
}
