package domain

import (
	"context"
	"github.com/google/uuid"
)

type RecipeStep struct {
	ID          uuid.UUID
	RecipeID    uuid.UUID
	Name        string
	Description string
	StepNum     int
}

type IRecipeStepRepository interface {
	Create(ctx context.Context, recipeStep *RecipeStep) error
	GetById(ctx context.Context, id uuid.UUID) (*RecipeStep, error)
	GetAllByRecipeID(ctx context.Context, recipeId uuid.UUID) ([]*RecipeStep, error)
	Update(ctx context.Context, recipeStep *RecipeStep) error
	DeleteById(ctx context.Context, id uuid.UUID) error
	DeleteAllByRecipeID(ctx context.Context, recipeId uuid.UUID) error
}

type IRecipeStepService interface {
	Create(ctx context.Context, recipeStep *RecipeStep) error
	GetById(ctx context.Context, id uuid.UUID) (*RecipeStep, error)
	GetAllByRecipeID(ctx context.Context, recipeId uuid.UUID) ([]*RecipeStep, error)
	Update(ctx context.Context, recipeStep *RecipeStep) error
	DeleteById(ctx context.Context, id uuid.UUID) error
	DeleteAllByRecipeID(ctx context.Context, recipeId uuid.UUID) error
}
