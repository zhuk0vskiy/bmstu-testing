package services

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"ppo/domain"
	"ppo/pkg/logger"
)

type RecipeStepService struct {
	recipeStepRepo domain.IRecipeStepRepository
	logger         logger.ILogger
}

func NewRecipeStepService(recipeStepRepo domain.IRecipeStepRepository, logger logger.ILogger) domain.IRecipeStepService {
	return &RecipeStepService{
		recipeStepRepo: recipeStepRepo,
		logger:         logger,
	}
}

func (s *RecipeStepService) verify(step *domain.RecipeStep) error {
	if step.Name == "" {
		return fmt.Errorf("empty name")
	}

	if step.Description == "" {
		return fmt.Errorf("empty description")
	}

	if step.StepNum <= 0 {
		return fmt.Errorf("negative or zero step num")
	}

	return nil
}

func (s *RecipeStepService) Create(ctx context.Context, recipeStep *domain.RecipeStep) error {
	s.logger.Infof("create recipeStep %+v", recipeStep)

	err := s.verify(recipeStep)
	if err != nil {
		s.logger.Warnf("failed to verify recipeStep %+v", recipeStep)
		return fmt.Errorf("creating recipe step: %w", err)
	}

	err = s.recipeStepRepo.Create(ctx, recipeStep)
	if err != nil {
		s.logger.Errorf("creating recipe step error: %s", err.Error())
		return fmt.Errorf("creating recipe step: %w", err)
	}
	return nil
}

//func (s *RecipeStepService) Update(ctx context.Context, recipeStep *domain.RecipeStep) error {
//	err := s.verify(recipeStep)
//	if err != nil {
//		return fmt.Errorf("updating recipe step: %w", err)
//	}
//
//	beforeUpdateStep, err := s.recipeStepRepo.GetById(ctx, recipeStep.ID)
//	if err != nil {
//		return fmt.Errorf("updating recipe step: %w", err)
//	}
//
//	steps, err := s.recipeStepRepo.GetAllByRecipeID(ctx, beforeUpdateStep.RecipeID)
//	if err != nil {
//		return fmt.Errorf("updating recipe step: %w", err)
//	}
//
//	if recipeStep.StepNum > len(steps) {
//		return fmt.Errorf("updating recipe step: invalid step num")
//	} else {
//		if beforeUpdateStep.StepNum != recipeStep.StepNum {
//			for _, step := range steps[beforeUpdateStep.StepNum:] {
//				step.StepNum--
//				err := s.recipeStepRepo.Update(ctx, step)
//				if err != nil {
//					return fmt.Errorf("updating recipe step: %w", err)
//				}
//			}
//		}
//	}
//
//	err = s.recipeStepRepo.Update(ctx, recipeStep)
//	if err != nil {
//		return fmt.Errorf("updating recipe step: %w", err)
//	}
//
//	return nil
//}

// FIXME: should use update below this comment
func (s *RecipeStepService) Update(ctx context.Context, recipeStep *domain.RecipeStep) error {
	s.logger.Infof("updating recipeStep %+v", recipeStep)

	err := s.verify(recipeStep)
	if err != nil {
		s.logger.Warnf("failed to verify recipeStep %+v", recipeStep)
		return fmt.Errorf("updating recipe step: %w", err)
	}

	err = s.recipeStepRepo.Update(ctx, recipeStep)
	if err != nil {
		s.logger.Errorf("updating recipe step error: %s", err.Error())
		return fmt.Errorf("updating recipe step: %w", err)
	}

	return nil
}

func (s *RecipeStepService) GetById(ctx context.Context, id uuid.UUID) (*domain.RecipeStep, error) {
	s.logger.Infof("getting recipe step by id: %s", id.String())

	recipeStep, err := s.recipeStepRepo.GetById(ctx, id)
	if err != nil {
		s.logger.Errorf("getting recipe step by id error: %s", err.Error())
		return nil, fmt.Errorf("getting recipe step by id: %w", err)
	}
	return recipeStep, nil
}

func (s *RecipeStepService) GetAllByRecipeID(ctx context.Context, recipeId uuid.UUID) ([]*domain.RecipeStep, error) {
	s.logger.Infof("getting all recipe steps by recipeId: %s", recipeId.String())

	recipeSteps, err := s.recipeStepRepo.GetAllByRecipeID(ctx, recipeId)
	if err != nil {
		s.logger.Errorf("getting all recipe steps of recipe error: %s", err.Error())
		return nil, fmt.Errorf("getting all steps of recipe: %w", err)
	}
	return recipeSteps, nil
}

//func (s *RecipeStepService) DeleteById(ctx context.Context, id uuid.UUID) error {
//	deletedStep, err := s.recipeStepRepo.GetById(ctx, id)
//	if err != nil {
//		return fmt.Errorf("deleting step by id: %w", err)
//	}
//
//	err = s.recipeStepRepo.DeleteById(ctx, id)
//	if err != nil {
//		return fmt.Errorf("deleting step by id: %w", err)
//	}
//
//	steps, err := s.recipeStepRepo.GetAllByRecipeID(ctx, deletedStep.RecipeID)
//	if err != nil {
//		return fmt.Errorf("deleting step by id: %w", err)
//	}
//
//	deletedStepNum := deletedStep.StepNum
//	if deletedStepNum != len(steps)+1 {
//		for _, step := range steps[deletedStepNum-1:] {
//			step.StepNum--
//			err := s.recipeStepRepo.Update(ctx, step)
//			if err != nil {
//				return fmt.Errorf("deleting step by id: %w", err)
//			}
//		}
//	}
//
//	return nil
//}

// FIXME: should user delete by id below this comment
func (s *RecipeStepService) DeleteById(ctx context.Context, id uuid.UUID) error {
	s.logger.Infof("deleting recipeStep by id: %s", id.String())

	err := s.recipeStepRepo.DeleteById(ctx, id)
	if err != nil {
		s.logger.Errorf("deleting recipeStep by id error: %s", err.Error())
		return fmt.Errorf("deleting step by id: %w", err)
	}

	return nil
}

func (s *RecipeStepService) DeleteAllByRecipeID(ctx context.Context, recipeId uuid.UUID) error {
	s.logger.Infof("deleting all recipeSteps by recipeId: %s", recipeId.String())

	err := s.recipeStepRepo.DeleteAllByRecipeID(ctx, recipeId)
	if err != nil {
		s.logger.Errorf("deleting all steps of recipe error: %s", err.Error())
		return fmt.Errorf("deleting all step of recipe: %w", err)
	}
	return nil
}
