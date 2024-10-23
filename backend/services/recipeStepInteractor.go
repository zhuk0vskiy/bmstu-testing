package services

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"ppo/domain"
	"strings"
)

type RecipeStepInteractor struct {
	recipeStepService domain.IRecipeStepService
	validators        []domain.IValidatorService
}

func NewRecipeStepInteractor(
	recipeStepService domain.IRecipeStepService,
	validators []domain.IValidatorService) *RecipeStepInteractor {
	return &RecipeStepInteractor{
		recipeStepService: recipeStepService,
		validators:        validators,
	}
}

func (i *RecipeStepInteractor) verifyText(ctx context.Context, str string) error {
	for _, word := range strings.Fields(str) {
		for _, validator := range i.validators {
			err := validator.Verify(ctx, word)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (i *RecipeStepInteractor) Create(ctx context.Context, recipeStep *domain.RecipeStep) error {
	err := i.verifyText(ctx, recipeStep.Name)
	if err != nil {
		return fmt.Errorf("recipe step interactor (name): %w", err)
	}

	err = i.verifyText(ctx, recipeStep.Description)
	if err != nil {
		return fmt.Errorf("recipe step interactor (description): %w", err)
	}

	err = i.recipeStepService.Create(ctx, recipeStep)
	if err != nil {
		return fmt.Errorf("recipe step interactor: %w", err)
	}
	return nil
}

func (i *RecipeStepInteractor) Update(ctx context.Context, recipeStep *domain.RecipeStep) error {
	err := i.verifyText(ctx, recipeStep.Name)
	if err != nil {
		return fmt.Errorf("recipe step interactor (name): %w", err)
	}

	err = i.verifyText(ctx, recipeStep.Description)
	if err != nil {
		return fmt.Errorf("recipe step interactor (description): %w", err)
	}

	err = i.recipeStepService.Update(ctx, recipeStep)
	if err != nil {
		return fmt.Errorf("recipe step interactor: %w", err)
	}
	return nil
}

func (i *RecipeStepInteractor) GetById(ctx context.Context, id uuid.UUID) (*domain.RecipeStep, error) {
	return i.recipeStepService.GetById(ctx, id)
}

func (i *RecipeStepInteractor) GetAllByRecipeID(ctx context.Context, recipeId uuid.UUID) ([]*domain.RecipeStep, error) {
	return i.recipeStepService.GetAllByRecipeID(ctx, recipeId)
}

func (i *RecipeStepInteractor) DeleteById(ctx context.Context, id uuid.UUID) error {
	return i.recipeStepService.DeleteById(ctx, id)
}

func (i *RecipeStepInteractor) DeleteAllByRecipeID(ctx context.Context, recipeId uuid.UUID) error {
	return i.recipeStepService.DeleteAllByRecipeID(ctx, recipeId)
}
