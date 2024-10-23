package domain

import (
	"context"
	"github.com/google/uuid"
)

type Measurement struct {
	ID    uuid.UUID
	Name  string
	Grams int
}

type IMeasurementRepository interface {
	Create(ctx context.Context, measurement *Measurement) error
	GetById(ctx context.Context, id uuid.UUID) (*Measurement, error)
	GetByRecipeId(ctx context.Context, ingredientId uuid.UUID, recipeId uuid.UUID) (*Measurement, int, error)
	GetAll(ctx context.Context) ([]*Measurement, error)
	UpdateLink(ctx context.Context, linkId uuid.UUID, measurementId uuid.UUID, amount int) error
	Update(ctx context.Context, measurement *Measurement) error
	DeleteById(ctx context.Context, id uuid.UUID) error
}

type IMeasurementService interface {
	Create(ctx context.Context, measurement *Measurement) error
	GetById(ctx context.Context, id uuid.UUID) (*Measurement, error)
	GetByRecipeId(ctx context.Context, ingredientId uuid.UUID, recipeId uuid.UUID) (*Measurement, int, error)
	GetAll(ctx context.Context) ([]*Measurement, error)
	UpdateLink(ctx context.Context, linkId uuid.UUID, measurementId uuid.UUID, amount int) error
	Update(ctx context.Context, measurement *Measurement) error
	DeleteById(ctx context.Context, id uuid.UUID) error
}
