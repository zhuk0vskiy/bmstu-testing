package domain

import (
	"context"
	"github.com/google/uuid"
)

const (
	MinRate = 1
	MaxRate = 5
)

type Comment struct {
	ID       uuid.UUID
	AuthorID uuid.UUID
	SaladID  uuid.UUID
	Text     string
	Rating   int
}

type ICommentRepository interface {
	Create(ctx context.Context, comment *Comment) error
	GetById(ctx context.Context, id uuid.UUID) (*Comment, error)
	GetBySaladAndUser(ctx context.Context, saladId uuid.UUID, userId uuid.UUID) (*Comment, error)
	GetAllBySaladID(ctx context.Context, saladId uuid.UUID, page int) ([]*Comment, int, error)
	Update(ctx context.Context, comment *Comment) error
	DeleteById(ctx context.Context, id uuid.UUID) error
}

type ICommentService interface {
	Create(ctx context.Context, comment *Comment) error
	GetById(ctx context.Context, id uuid.UUID) (*Comment, error)
	GetBySaladAndUser(ctx context.Context, saladId uuid.UUID, userId uuid.UUID) (*Comment, error)
	GetAllBySaladID(ctx context.Context, saladId uuid.UUID, page int) ([]*Comment, int, error)
	Update(ctx context.Context, user *Comment) error
	DeleteById(ctx context.Context, id uuid.UUID) error
}
