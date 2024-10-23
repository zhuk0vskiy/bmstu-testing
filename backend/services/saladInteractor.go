package services

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"ppo/domain"
	"ppo/services/dto"
	"strings"
)

type SaladInteractor struct {
	saladService domain.ISaladService
	validators   []domain.IValidatorService
}

func NewSaladInteractor(
	saladService domain.ISaladService,
	validators []domain.IValidatorService) domain.ISaladInteractor {
	return &SaladInteractor{
		saladService: saladService,
		validators:   validators,
	}
}

func (i *SaladInteractor) Create(ctx context.Context, salad *domain.Salad) (uuid.UUID, error) {
	id := uuid.Nil
	for _, word := range strings.Fields(salad.Name) {
		for _, validator := range i.validators {
			if err := validator.Verify(ctx, word); err != nil {
				return id, fmt.Errorf("salad interactor (name): %w", err)
			}
		}
	}

	for _, word := range strings.Fields(salad.Description) {
		for _, validator := range i.validators {
			if err := validator.Verify(ctx, word); err != nil {
				return id, fmt.Errorf("salad interactor (description): %w", err)
			}
		}
	}

	id, err := i.saladService.Create(ctx, salad)
	if err != nil {
		return id, fmt.Errorf("salad interactor: %w", err)
	}
	return id, nil
}

func (i *SaladInteractor) Update(ctx context.Context, salad *domain.Salad) error {
	for _, word := range strings.Fields(salad.Name) {
		for _, validator := range i.validators {
			if err := validator.Verify(ctx, word); err != nil {
				return fmt.Errorf("salad interactor (name): %w", err)
			}
		}
	}

	for _, word := range strings.Fields(salad.Description) {
		for _, validator := range i.validators {
			if err := validator.Verify(ctx, word); err != nil {
				return fmt.Errorf("salad interactor (description): %w", err)
			}
		}
	}

	err := i.saladService.Update(ctx, salad)
	if err != nil {
		return fmt.Errorf("salad interactor: %w", err)
	}
	return nil
}

func (i *SaladInteractor) GetById(ctx context.Context, id uuid.UUID) (*domain.Salad, error) {
	return i.saladService.GetById(ctx, id)
}

func (i *SaladInteractor) GetAll(ctx context.Context, filter *dto.RecipeFilter, page int) ([]*domain.Salad, int, error) {
	return i.saladService.GetAll(ctx, filter, page)
}

func (i *SaladInteractor) GetAllByUserId(ctx context.Context, id uuid.UUID) ([]*domain.Salad, error) {
	return i.saladService.GetAllByUserId(ctx, id)
}

func (i *SaladInteractor) DeleteById(ctx context.Context, id uuid.UUID) error {
	return i.saladService.DeleteById(ctx, id)
}

func (i *SaladInteractor) GetAllRatedByUser(ctx context.Context, userId uuid.UUID, page int) ([]*domain.Salad, int, error) {
	return i.saladService.GetAllRatedByUser(ctx, userId, page)
}
