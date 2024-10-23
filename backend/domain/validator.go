package domain

import (
	"context"
)

type IValidatorService interface {
	Verify(ctx context.Context, word string) error
}
