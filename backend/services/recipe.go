package services

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"ppo/domain"
	"ppo/pkg/logger"
	"ppo/services/dto"
)

type RecipeService struct {
	recipeRepo domain.IRecipeRepository
	logger     logger.ILogger
}

func NewRecipeService(recipeRepo domain.IRecipeRepository, logger logger.ILogger) domain.IRecipeService {
	return &RecipeService{
		recipeRepo: recipeRepo,
		logger:     logger,
	}
}

func (s *RecipeService) verify(recipe *domain.Recipe) error {
	if recipe.NumberOfServings <= 0 {
		return fmt.Errorf("negative or zero number of servings")
	}

	if recipe.TimeToCook <= 0 {
		return fmt.Errorf("negative or zero time to cook")
	}

	return nil
}

func (s *RecipeService) Create(ctx context.Context, recipe *domain.Recipe) (uuid.UUID, error) {
	s.logger.Infof("create recipe ща salad: %s", recipe.SaladID.String())

	id := uuid.Nil
	err := s.verify(recipe)
	if err != nil {
		s.logger.Warnf("failed to verify recipe: %s", err.Error())
		return id, fmt.Errorf("creating recipe: %w", err)
	}

	id, err = s.recipeRepo.Create(ctx, recipe)
	if err != nil {
		s.logger.Errorf("creating recipe error: %s", err.Error())
		return id, fmt.Errorf("creating recipe: %w", err)
	}

	return id, nil
}

func (s *RecipeService) Update(ctx context.Context, recipe *domain.Recipe) error {
	s.logger.Infof("updating recipe with id: %s", recipe.ID.String())

	err := s.verify(recipe)
	if err != nil {
		s.logger.Warnf("failed to verify recipe: %s", err.Error())
		return fmt.Errorf("updating recipe: %w", err)
	}

	err = s.recipeRepo.Update(ctx, recipe)
	if err != nil {
		s.logger.Errorf("updating recipe error: %s", err.Error())
		return fmt.Errorf("updating recipe: %w", err)
	}
	return nil
}

func (s *RecipeService) GetById(ctx context.Context, id uuid.UUID) (*domain.Recipe, error) {
	s.logger.Infof("getting recipe by id: %s", id.String())

	recipe, err := s.recipeRepo.GetById(ctx, id)
	if err != nil {
		s.logger.Errorf("getting recipe by id error: %s", err.Error())
		return nil, fmt.Errorf("getting recipe by id: %w", err)
	}
	return recipe, nil
}

func (s *RecipeService) GetBySaladId(ctx context.Context, saladId uuid.UUID) (*domain.Recipe, error) {
	s.logger.Infof("getting recipe by salad id: %s", saladId.String())

	recipe, err := s.recipeRepo.GetBySaladId(ctx, saladId)
	if err != nil {
		s.logger.Errorf("getting recipe by salad id error: %s", err.Error())
		return nil, fmt.Errorf("getting recipe by salad id: %w", err)
	}
	return recipe, nil
}

func (s *RecipeService) GetAll(ctx context.Context, filter *dto.RecipeFilter, page int) ([]*domain.Recipe, error) {
	s.logger.Infof("getting all recipes with filter: %+v", filter)

	recipes, err := s.recipeRepo.GetAll(ctx, filter, page)
	if err != nil {
		s.logger.Errorf("getting all recipes error: %s", err.Error())
		return nil, fmt.Errorf("getting all recipes: %w", err)
	}
	return recipes, nil
}

func (s *RecipeService) DeleteById(ctx context.Context, id uuid.UUID) error {
	s.logger.Infof("deleting recipe by id: %s", id.String())

	err := s.recipeRepo.DeleteById(ctx, id)
	if err != nil {
		s.logger.Errorf("deleting recipe by id error: %s", err.Error())
		return fmt.Errorf("deleting recipe by id: %w", err)
	}
	return nil
}
