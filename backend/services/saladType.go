package services

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"ppo/domain"
	"ppo/pkg/logger"
)

type SaladTypeService struct {
	saladTypeRepo domain.ISaladTypeRepository
	logger        logger.ILogger
}

func NewSaladTypeService(saladTypeRepo domain.ISaladTypeRepository, logger logger.ILogger) domain.ISaladTypeService {
	return &SaladTypeService{
		saladTypeRepo: saladTypeRepo,
		logger:        logger,
	}
}

func (s *SaladTypeService) verify(saladType *domain.SaladType) error {
	if saladType.Name == "" {
		return fmt.Errorf("empty name")
	}
	return nil
}

func (s *SaladTypeService) Create(ctx context.Context, saladType *domain.SaladType) error {
	s.logger.Infof("creating salad type: %+v", saladType)

	err := s.verify(saladType)
	if err != nil {
		s.logger.Warnf("failed to verify salad type: %s", err.Error())
		return fmt.Errorf("creating salad type: %w", err)
	}

	err = s.saladTypeRepo.Create(ctx, saladType)
	if err != nil {
		s.logger.Errorf("creating salad type error: %s", err.Error())
		return fmt.Errorf("creating salad type: %w", err)
	}
	return nil
}

func (s *SaladTypeService) Update(ctx context.Context, saladType *domain.SaladType) error {
	s.logger.Infof("updating salad type: %+v", saladType)

	err := s.verify(saladType)
	if err != nil {
		s.logger.Warnf("failed to verify salad type: %s", err.Error())
		return fmt.Errorf("updating salad type: %w", err)
	}

	err = s.saladTypeRepo.Update(ctx, saladType)
	if err != nil {
		s.logger.Errorf("updating salad type error: %s", err.Error())
		return fmt.Errorf("updating salad type: %w", err)
	}
	return nil
}

func (s *SaladTypeService) GetById(ctx context.Context, id uuid.UUID) (*domain.SaladType, error) {
	s.logger.Infof("getting salad type by id %s", id.String())

	saladType, err := s.saladTypeRepo.GetById(ctx, id)
	if err != nil {
		s.logger.Errorf("getting salad type by id %s error: %s", id.String(), err.Error())
		return nil, fmt.Errorf("getting salad type by id: %w", err)
	}
	return saladType, nil
}

func (s *SaladTypeService) GetAll(ctx context.Context, page int) ([]*domain.SaladType, int, error) {
	s.logger.Infof("gettign all salad types on page %d", page)

	saladTypes, numPages, err := s.saladTypeRepo.GetAll(ctx, page)
	if err != nil {
		s.logger.Errorf("getting salad types error: %s", err.Error())
		return nil, 0, fmt.Errorf("getting all salad types: %w", err)
	}
	return saladTypes, numPages, nil
}

func (s *SaladTypeService) GetAllBySaladId(ctx context.Context, saladId uuid.UUID) ([]*domain.SaladType, error) {
	s.logger.Infof("getting salad types by salad id %s", saladId.String())

	saladTypes, err := s.saladTypeRepo.GetAllBySaladId(ctx, saladId)
	if err != nil {
		s.logger.Errorf("getting salad types error: %s", err.Error())
		return nil, fmt.Errorf("getting all types of salad: %w", err)
	}
	return saladTypes, err
}

func (s *SaladTypeService) DeleteById(ctx context.Context, id uuid.UUID) error {
	s.logger.Infof("deleting salad type by id %s", id.String())

	err := s.saladTypeRepo.DeleteById(ctx, id)
	if err != nil {
		s.logger.Errorf("deleting salad type by id %s error: %s", id.String(), err.Error())
		return fmt.Errorf("deleting salad type by id: %w", err)
	}
	return nil
}

func (s *SaladTypeService) Link(ctx context.Context, saladId uuid.UUID, saladTypeId uuid.UUID) error {
	s.logger.Infof("linking salad type %s to salad %s", saladTypeId.String(), saladId.String())

	err := s.saladTypeRepo.Link(ctx, saladId, saladTypeId)
	if err != nil {
		s.logger.Errorf("linking salad type error: %s", err.Error())
		return fmt.Errorf("linking salad type: %w", err)
	}
	return nil
}

func (s *SaladTypeService) Unlink(ctx context.Context, saladId uuid.UUID, saladTypeId uuid.UUID) error {
	s.logger.Infof("unlinking salad type %s from salad %s", saladTypeId.String(), saladId.String())

	err := s.saladTypeRepo.Unlink(ctx, saladId, saladTypeId)
	if err != nil {
		s.logger.Errorf("unlinking salad type error: %s", err.Error())
		return fmt.Errorf("unlinking salad type: %w", err)
	}
	return nil
}
