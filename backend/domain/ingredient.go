package domain

import (
	"context"
	"github.com/google/uuid"
)

type Ingredient struct {
	ID       uuid.UUID
	TypeID   uuid.UUID
	Name     string
	Calories int
}

type IIngredientRepository interface {
	Create(ctx context.Context, ingredient *Ingredient) error
	GetById(ctx context.Context, id uuid.UUID) (*Ingredient, error)
	GetAll(ctx context.Context, page int) ([]*Ingredient, int, error)
	GetAllByRecipeId(ctx context.Context, id uuid.UUID) ([]*Ingredient, error)
	Link(ctx context.Context, recipeId uuid.UUID, ingredientId uuid.UUID) (uuid.UUID, error)
	Unlink(ctx context.Context, recipeId uuid.UUID, ingredientId uuid.UUID) error
	Update(ctx context.Context, ingredient *Ingredient) error
	DeleteById(ctx context.Context, id uuid.UUID) error
}

type IIngredientService interface {
	Create(ctx context.Context, salad *Ingredient) error
	GetById(ctx context.Context, id uuid.UUID) (*Ingredient, error)
	GetAll(ctx context.Context, page int) ([]*Ingredient, int, error)
	GetAllByRecipeId(ctx context.Context, id uuid.UUID) ([]*Ingredient, error)
	Link(ctx context.Context, recipeId uuid.UUID, ingredientId uuid.UUID) (uuid.UUID, error)
	Unlink(ctx context.Context, recipeId uuid.UUID, ingredientId uuid.UUID) error
	Update(ctx context.Context, salad *Ingredient) error
	DeleteById(ctx context.Context, id uuid.UUID) error
}
