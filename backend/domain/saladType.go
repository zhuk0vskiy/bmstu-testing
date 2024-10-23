package domain

import (
	"context"
	"github.com/google/uuid"
)

type SaladType struct {
	ID          uuid.UUID
	Name        string
	Description string
}

type ISaladTypeRepository interface {
	Create(ctx context.Context, saladType *SaladType) error
	GetById(ctx context.Context, id uuid.UUID) (*SaladType, error)
	GetAll(ctx context.Context, page int) ([]*SaladType, int, error)
	GetAllBySaladId(ctx context.Context, saladId uuid.UUID) ([]*SaladType, error)
	Update(ctx context.Context, saladType *SaladType) error
	Link(ctx context.Context, saladId uuid.UUID, saladTypeId uuid.UUID) error
	Unlink(ctx context.Context, saladId uuid.UUID, saladTypeId uuid.UUID) error
	DeleteById(ctx context.Context, id uuid.UUID) error
}

type ISaladTypeService interface {
	Create(ctx context.Context, saladType *SaladType) error
	GetById(ctx context.Context, id uuid.UUID) (*SaladType, error)
	GetAll(ctx context.Context, page int) ([]*SaladType, int, error)
	GetAllBySaladId(ctx context.Context, saladId uuid.UUID) ([]*SaladType, error)
	Update(ctx context.Context, measurement *SaladType) error
	Link(ctx context.Context, saladId uuid.UUID, saladTypeId uuid.UUID) error
	Unlink(ctx context.Context, saladId uuid.UUID, saladTypeId uuid.UUID) error
	DeleteById(ctx context.Context, id uuid.UUID) error
}
