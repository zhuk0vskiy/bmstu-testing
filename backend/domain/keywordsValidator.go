package domain

import (
	"context"
	"github.com/google/uuid"
)

type KeyWord struct {
	ID   uuid.UUID
	Word string
}

type IKeywordValidatorRepository interface {
	Create(ctx context.Context, word *KeyWord) error
	GetById(ctx context.Context, id uuid.UUID) (*KeyWord, error)
	GetAll(ctx context.Context) (map[string]uuid.UUID, error)
	Update(ctx context.Context, word *KeyWord) error
	DeleteById(ctx context.Context, id uuid.UUID) error
}

type IKeywordValidatorService interface {
	IValidatorService
	Create(ctx context.Context, word *KeyWord) error
	GetById(ctx context.Context, id uuid.UUID) (*KeyWord, error)
	GetAll(ctx context.Context) (map[string]uuid.UUID, error)
	Update(ctx context.Context, word *KeyWord) error
	DeleteById(ctx context.Context, id uuid.UUID) error
}
