package domain

import (
	"context"
	"github.com/google/uuid"
	"ppo/services/dto"
)

type Salad struct {
	ID          uuid.UUID
	AuthorID    uuid.UUID
	Name        string
	Description string
}

type ISaladRepository interface {
	Create(ctx context.Context, salad *Salad) (uuid.UUID, error)
	GetById(ctx context.Context, id uuid.UUID) (*Salad, error)
	GetAll(ctx context.Context, filter *dto.RecipeFilter, page int) ([]*Salad, int, error)
	GetAllByUserId(ctx context.Context, id uuid.UUID) ([]*Salad, error)
	GetAllRatedByUser(ctx context.Context, userId uuid.UUID, page int) ([]*Salad, int, error)
	Update(ctx context.Context, salad *Salad) error
	DeleteById(ctx context.Context, id uuid.UUID) error
}

type ISaladService interface {
	Create(ctx context.Context, salad *Salad) (uuid.UUID, error)
	GetById(ctx context.Context, id uuid.UUID) (*Salad, error)
	GetAll(ctx context.Context, filter *dto.RecipeFilter, page int) ([]*Salad, int, error)
	GetAllByUserId(ctx context.Context, id uuid.UUID) ([]*Salad, error)
	GetAllRatedByUser(ctx context.Context, userId uuid.UUID, page int) ([]*Salad, int, error)
	Update(ctx context.Context, salad *Salad) error
	DeleteById(ctx context.Context, id uuid.UUID) error
}
