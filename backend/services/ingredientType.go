package services

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"ppo/domain"
	"ppo/pkg/logger"
)

type IngredientTypeService struct {
	ingredientTypeRepo domain.IIngredientTypeRepository
	logger             logger.ILogger
}

func NewIngredientTypeService(ingredientTypeRepo domain.IIngredientTypeRepository, logger logger.ILogger) domain.IIngredientTypeService {
	return &IngredientTypeService{
		ingredientTypeRepo: ingredientTypeRepo,
		logger:             logger,
	}
}

func (s *IngredientTypeService) verify(ingredientType *domain.IngredientType) error {
	if ingredientType.Name == "" {
		return fmt.Errorf("empty name")
	}
	return nil
}

func (s *IngredientTypeService) Create(ctx context.Context, ingredientType *domain.IngredientType) error {
	s.logger.Infof("create ingredient type: %s", ingredientType.Name)

	err := s.verify(ingredientType)
	if err != nil {
		s.logger.Warnf("failed to verify ingredient type: %s", err.Error())
		return fmt.Errorf("creating ingredient type: %w", err)
	}

	err = s.ingredientTypeRepo.Create(ctx, ingredientType)
	if err != nil {
		s.logger.Errorf("creating ingredient type error: %s", err.Error())
		return fmt.Errorf("creating ingredient type: %w", err)
	}

	return nil
}

func (s *IngredientTypeService) Update(ctx context.Context, ingredientType *domain.IngredientType) error {
	s.logger.Infof("update ingredient type: %s", ingredientType.Name)

	err := s.verify(ingredientType)
	if err != nil {
		s.logger.Warnf("failed to verify ingredient type: %s", err.Error())
		return fmt.Errorf("updating ingredient type: %w", err)
	}

	err = s.ingredientTypeRepo.Update(ctx, ingredientType)
	if err != nil {
		s.logger.Errorf("updating ingredient type error: %s", err.Error())
		return fmt.Errorf("updating ingredient type: %w", err)
	}
	return nil
}

func (s *IngredientTypeService) GetById(ctx context.Context, id uuid.UUID) (*domain.IngredientType, error) {
	s.logger.Infof("getting ingredient type by id: %s", id.String())

	ingredientType, err := s.ingredientTypeRepo.GetById(ctx, id)
	if err != nil {
		s.logger.Errorf("getting ingredient type by id error: %s", err.Error())
		return nil, fmt.Errorf("getting ingredient type by id: %w", err)
	}
	return ingredientType, nil
}

func (s *IngredientTypeService) GetAll(ctx context.Context) ([]*domain.IngredientType, error) {
	s.logger.Infof("getting all ingredient types")

	measurements, err := s.ingredientTypeRepo.GetAll(ctx)
	if err != nil {
		s.logger.Errorf("getting all ingredient types error: %s", err.Error())
		return nil, fmt.Errorf("getting all ingredient types: %w", err)
	}
	return measurements, nil
}

func (s *IngredientTypeService) DeleteById(ctx context.Context, id uuid.UUID) error {
	s.logger.Infof("deleting ingredient type by id: %s", id.String())

	err := s.ingredientTypeRepo.DeleteById(ctx, id)
	if err != nil {
		s.logger.Errorf("deleting ingredient type error: %s", err.Error())
		return fmt.Errorf("deleting ingredient type by id: %w", err)
	}
	return nil
}
