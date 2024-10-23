package domain

import (
	"context"
	"github.com/google/uuid"
)

type IngredientType struct {
	ID          uuid.UUID
	Name        string
	Description string
}

type IIngredientTypeRepository interface {
	Create(ctx context.Context, ingredientType *IngredientType) error
	GetById(ctx context.Context, id uuid.UUID) (*IngredientType, error)
	GetAll(ctx context.Context) ([]*IngredientType, error)
	Update(ctx context.Context, measurement *IngredientType) error
	DeleteById(ctx context.Context, id uuid.UUID) error
}

type IIngredientTypeService interface {
	Create(ctx context.Context, measurement *IngredientType) error
	GetById(ctx context.Context, id uuid.UUID) (*IngredientType, error)
	GetAll(ctx context.Context) ([]*IngredientType, error)
	Update(ctx context.Context, measurement *IngredientType) error
	DeleteById(ctx context.Context, id uuid.UUID) error
}
