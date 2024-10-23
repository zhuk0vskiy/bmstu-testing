package services

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"ppo/domain"
	"ppo/pkg/logger"
)

type MeasurementService struct {
	measurementRepo domain.IMeasurementRepository
	logger          logger.ILogger
}

func NewMeasurementService(measurementRepo domain.IMeasurementRepository, logger logger.ILogger) domain.IMeasurementService {
	return &MeasurementService{
		measurementRepo: measurementRepo,
		logger:          logger,
	}
}

func (s *MeasurementService) verify(measurement *domain.Measurement) error {
	if measurement.Name == "" {
		return fmt.Errorf("empty name")
	}

	if measurement.Grams <= 0 {
		return fmt.Errorf("negative or zero grams count")
	}

	return nil
}

func (s *MeasurementService) Create(ctx context.Context, measurement *domain.Measurement) error {
	s.logger.Infof("create measurement: %s", measurement.Name)

	err := s.verify(measurement)
	if err != nil {
		s.logger.Warnf("fail to verify measurement: %s", err.Error())
		return fmt.Errorf("creating measurement unit: %w", err)
	}

	err = s.measurementRepo.Create(ctx, measurement)
	if err != nil {
		s.logger.Errorf("creating measurement unit: %s", err.Error())
		return fmt.Errorf("creating measurement unit: %w", err)
	}

	return nil
}

func (s *MeasurementService) Update(ctx context.Context, measurement *domain.Measurement) error {
	s.logger.Infof("updating measurement: %s", measurement.ID.String())

	err := s.verify(measurement)
	if err != nil {
		s.logger.Warnf("fail to verify measurement: %s", err.Error())
		return fmt.Errorf("updating measurement unit: %w", err)
	}

	err = s.measurementRepo.Update(ctx, measurement)
	if err != nil {
		s.logger.Errorf("updating measurement unit error: %s", err.Error())
		return fmt.Errorf("updating measurement unit: %w", err)
	}
	return nil
}

func (s *MeasurementService) GetById(ctx context.Context, id uuid.UUID) (*domain.Measurement, error) {
	s.logger.Infof("getting measurement by id: %s", id.String())

	measurement, err := s.measurementRepo.GetById(ctx, id)
	if err != nil {
		s.logger.Errorf("getting measurement by id error: %s", err.Error())
		return nil, fmt.Errorf("getting measurement unit by id: %w", err)
	}
	return measurement, nil
}

func (s *MeasurementService) GetByRecipeId(ctx context.Context,
	ingredientId uuid.UUID,
	recipeId uuid.UUID) (*domain.Measurement, int, error) {
	s.logger.Infof("getting measurement with recipe id: %s, ingredient id: %s",
		recipeId.String(), ingredientId.String())

	measurement, count, err := s.measurementRepo.GetByRecipeId(ctx, ingredientId, recipeId)
	if err != nil {
		s.logger.Errorf("getting measurement unit by recipe id error: %s", err.Error())
		return nil, 0, fmt.Errorf("getting measurement unit by recipe id: %w", err)
	}
	return measurement, count, nil
}

func (s *MeasurementService) GetAll(ctx context.Context) ([]*domain.Measurement, error) {
	s.logger.Infof("getting all measurements")

	measurements, err := s.measurementRepo.GetAll(ctx)
	if err != nil {
		s.logger.Errorf("getting all measurements error: %s", err.Error())
		return nil, fmt.Errorf("getting all measurement units: %w", err)
	}
	return measurements, nil
}

func (s *MeasurementService) DeleteById(ctx context.Context, id uuid.UUID) error {
	s.logger.Infof("deleting measurement by id: %s", id.String())
	err := s.measurementRepo.DeleteById(ctx, id)
	if err != nil {
		s.logger.Errorf("deleting measurement by id error: %s", err.Error())
		return fmt.Errorf("deleting measurement unit by id: %w", err)
	}
	return nil
}

func (s *MeasurementService) UpdateLink(ctx context.Context, linkId uuid.UUID, measurementId uuid.UUID, amount int) error {
	s.logger.Infof("updating measurement link: %s", linkId.String())

	if amount <= 0 {
		s.logger.Warnf("updating measurement link amount must be greater than zero")
		return fmt.Errorf("negative or zero amount")
	}

	err := s.measurementRepo.UpdateLink(ctx, linkId, measurementId, amount)
	if err != nil {
		s.logger.Errorf("updating measurement link error: %s", err.Error())
		return fmt.Errorf("updating measurement unit by link id: %w", err)
	}
	return nil
}
