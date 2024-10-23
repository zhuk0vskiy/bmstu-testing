package domain

import (
	"context"
	"github.com/google/uuid"
)

type IRecipeStepInteractor interface {
	Create(ctx context.Context, recipeStep *RecipeStep) error
	GetById(ctx context.Context, id uuid.UUID) (*RecipeStep, error)
	GetAllByRecipeID(ctx context.Context, recipeId uuid.UUID) ([]*RecipeStep, error)
	Update(ctx context.Context, recipeStep *RecipeStep) error
	DeleteById(ctx context.Context, id uuid.UUID) error
	DeleteAllByRecipeID(ctx context.Context, recipeId uuid.UUID) error
}
