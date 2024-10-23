package services

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"ppo/domain"
	"ppo/pkg/logger"
)

type IngredientService struct {
	ingredientRepo domain.IIngredientRepository
	logger         logger.ILogger
}

func NewIngredientService(ingredientRepo domain.IIngredientRepository, logger logger.ILogger) domain.IIngredientService {
	return &IngredientService{
		ingredientRepo: ingredientRepo,
		logger:         logger,
	}
}

func (s *IngredientService) verify(ingredient *domain.Ingredient) error {
	if ingredient.Name == "" {
		return fmt.Errorf("empty name")
	}

	if ingredient.Calories < 0 {
		return fmt.Errorf("negative calories")
	}

	return nil
}

func (s *IngredientService) Create(ctx context.Context, ingredient *domain.Ingredient) error {
	s.logger.Infof("creating ingredient: %s", ingredient.Name)

	err := s.verify(ingredient)
	if err != nil {
		s.logger.Warnf("failed to verify ingredient: %s", err.Error())
		return fmt.Errorf("creating ingredient: %w", err)
	}

	err = s.ingredientRepo.Create(ctx, ingredient)
	if err != nil {
		s.logger.Errorf("creating ingredient error: %s", err.Error())
		return fmt.Errorf("creating ingredient: %w", err)
	}
	return nil
}

func (s *IngredientService) Update(ctx context.Context, ingredient *domain.Ingredient) error {
	s.logger.Infof("updating ingredient: %s", ingredient.Name)

	err := s.verify(ingredient)
	if err != nil {
		s.logger.Warnf("failed to verify ingredient: %s", err.Error())
		return fmt.Errorf("updating ingredient: %w", err)
	}

	err = s.ingredientRepo.Update(ctx, ingredient)
	if err != nil {
		s.logger.Errorf("updating ingredient error: %s", err.Error())
		return fmt.Errorf("updating ingredient: %w", err)
	}
	return nil
}

func (s *IngredientService) GetById(ctx context.Context, id uuid.UUID) (*domain.Ingredient, error) {
	s.logger.Infof("getting ingredient by id: %s", id.String())

	ingredient, err := s.ingredientRepo.GetById(ctx, id)
	if err != nil {
		s.logger.Errorf("getting ingredient by id error: %s", err.Error())
		return nil, fmt.Errorf("getting ingredient by id: %w", err)
	}
	return ingredient, nil
}

func (s *IngredientService) GetAll(ctx context.Context, page int) ([]*domain.Ingredient, int, error) {
	s.logger.Infof("getting all ingredients on page: %d", page)

	salads, numPages, err := s.ingredientRepo.GetAll(ctx, page)
	if err != nil {
		s.logger.Errorf("getting all ingredients error: %s", err.Error())
		return nil, 0, fmt.Errorf("getting all ingredients: %w", err)
	}
	return salads, numPages, nil
}

func (s *IngredientService) DeleteById(ctx context.Context, id uuid.UUID) error {
	s.logger.Infof("deleting ingredient by id: %s", id.String())

	err := s.ingredientRepo.DeleteById(ctx, id)
	if err != nil {
		s.logger.Errorf("deleting ingredient by id error: %s", err.Error())
		return fmt.Errorf("deleting ingredient by id: %w", err)
	}
	return nil
}

func (s *IngredientService) GetAllByRecipeId(ctx context.Context, id uuid.UUID) ([]*domain.Ingredient, error) {
	s.logger.Infof("getting all ingredients by recipe id: %s", id.String())

	typedIngredients, err := s.ingredientRepo.GetAllByRecipeId(ctx, id)
	if err != nil {
		s.logger.Errorf("getting all ingredients by recipe id error: %s", err.Error())
		return nil, fmt.Errorf("getting all ingredients by recipe id: %w", err)
	}
	return typedIngredients, nil
}

func (s *IngredientService) Link(ctx context.Context, recipeId uuid.UUID, ingredientId uuid.UUID) (uuid.UUID, error) {
	s.logger.Infof("adding ingredient %s to recipe %s", ingredientId.String(), recipeId.String())

	id, err := s.ingredientRepo.Link(ctx, recipeId, ingredientId)
	if err != nil {
		s.logger.Errorf("adding ingredient %s to recipe error: %s", ingredientId.String(), err.Error())
		return uuid.Nil, fmt.Errorf("linking ingredient to recipe: %w", err)
	}
	return id, nil
}

func (s *IngredientService) Unlink(ctx context.Context, recipeId uuid.UUID, ingredientId uuid.UUID) error {
	s.logger.Infof("removing ingredient %s from recipe %s", ingredientId.String(), recipeId.String())

	err := s.ingredientRepo.Unlink(ctx, recipeId, ingredientId)
	if err != nil {
		s.logger.Errorf("removing ingredient %s from recipe error: %s", ingredientId.String(), err.Error())
		return fmt.Errorf("unlinking ingredient from recipe: %w", err)
	}
	return nil
}
