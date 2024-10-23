package services

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"ppo/domain"
	"ppo/pkg/logger"
	"ppo/services/dto"
)

type SaladService struct {
	saladRepo domain.ISaladRepository
	logger    logger.ILogger
}

func NewSaladService(saladRepo domain.ISaladRepository, logger logger.ILogger) domain.ISaladService {
	return &SaladService{
		saladRepo: saladRepo,
		logger:    logger,
	}
}

func (s SaladService) Create(ctx context.Context, salad *domain.Salad) (uuid.UUID, error) {
	s.logger.Infof("Salad created: %+v", salad)

	id := uuid.Nil
	if salad.Name == "" {
		s.logger.Warnf("empty salad name")
		return id, fmt.Errorf("empty salad name")
	}

	id, err := s.saladRepo.Create(ctx, salad)
	if err != nil {
		s.logger.Errorf("creating salad error: %s", err.Error())
		return id, fmt.Errorf("creating salad: %w", err)
	}
	return id, nil
}

func (s SaladService) Update(ctx context.Context, salad *domain.Salad) error {
	s.logger.Infof("updating salad: %+v", salad)

	if salad.Name == "" {
		s.logger.Warnf("empty salad name")
		return fmt.Errorf("updating salad: empty salad name")
	}

	err := s.saladRepo.Update(ctx, salad)
	if err != nil {
		s.logger.Errorf("updating salad error: %s", err.Error())
		return fmt.Errorf("updating salad: %w", err)
	}
	return nil
}

func (s SaladService) GetById(ctx context.Context, id uuid.UUID) (*domain.Salad, error) {
	s.logger.Infof("getting salad by id: %s", id.String())

	salad, err := s.saladRepo.GetById(ctx, id)
	if err != nil {
		s.logger.Errorf("getting salad by id error: %s", err.Error())
		return nil, fmt.Errorf("getting salad by id: %w", err)
	}
	return salad, nil
}

func (s SaladService) GetAll(ctx context.Context, filter *dto.RecipeFilter, page int) ([]*domain.Salad, int, error) {
	s.logger.Infof("getting all salads by filter: %+v", filter)

	salads, numPages, err := s.saladRepo.GetAll(ctx, filter, page)
	if err != nil {
		s.logger.Errorf("getting salads by filter error: %s", err.Error())
		return nil, 0, fmt.Errorf("getting all salads: %w", err)
	}
	return salads, numPages, nil
}

func (s SaladService) DeleteById(ctx context.Context, id uuid.UUID) error {
	s.logger.Infof("deleting salad by id: %s", id.String())

	err := s.saladRepo.DeleteById(ctx, id)
	if err != nil {
		s.logger.Errorf("deleting salad by id error: %s", err.Error())
		return fmt.Errorf("deleting salad by id: %w", err)
	}
	return nil
}

func (s SaladService) GetAllByUserId(ctx context.Context, id uuid.UUID) ([]*domain.Salad, error) {
	s.logger.Infof("getting all salads by user id: %s", id.String())

	authorSalads, err := s.saladRepo.GetAllByUserId(ctx, id)
	if err != nil {
		s.logger.Errorf("getting salads by user id error: %s", err.Error())
		return nil, fmt.Errorf("getting all salads by author id: %w", err)
	}
	return authorSalads, nil
}

func (s SaladService) GetAllRatedByUser(ctx context.Context, userId uuid.UUID, page int) ([]*domain.Salad, int, error) {
	s.logger.Infof("getting all salads rated by user with id: %s", userId.String())

	salads, numPages, err := s.saladRepo.GetAllRatedByUser(ctx, userId, page)
	if err != nil {
		s.logger.Errorf("getting salads rated by user error: %s", err.Error())
		return nil, 0, fmt.Errorf("getting all salads rated by user: %w", err)
	}
	return salads, numPages, nil
}
