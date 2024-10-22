package impl

import (
	"backend/src/internal/model"
	"backend/src/internal/model/dto"
	repositoryInterface "backend/src/internal/repository/interface"
	serviceInterface "backend/src/internal/service/interface"
	"backend/src/pkg/logger"
	"context"
	"fmt"
)

type StudioService struct {
	studioRepo repositoryInterface.IStudioRepository
	logger     logger.Interface
}

func NewStudioService(logger logger.Interface, studioRepo repositoryInterface.IStudioRepository) serviceInterface.IStudioService {
	return &StudioService{
		logger:     logger,
		studioRepo: studioRepo,
	}
}

func (s StudioService) Update(request *dto.UpdateStudioRequest) (err error) {
	if request.Id < 1 {
		s.logger.Infof("ошибка update studio: %s", fmt.Errorf("неправильный id: %w", err))
		return fmt.Errorf("неправильный id: %w", err)
	}

	if request.Name == "" {
		s.logger.Infof("ошибка update studio: %s", fmt.Errorf("пустое название: %w", err))
		return fmt.Errorf("пустое название: %w", err)
	}

	//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//defer cancel()
	ctx := context.Background()
	err = s.studioRepo.Update(ctx, &dto.UpdateStudioRequest{
		Id:   request.Id,
		Name: request.Name,
	})
	if err != nil {
		s.logger.Errorf("ошибка update studio: %s", fmt.Errorf("обновление студии: %w", err))
		return fmt.Errorf("обновление студии: %w", err)
	}

	return err
}

func (s StudioService) Get(request *dto.GetStudioRequest) (studio *model.Studio, err error) {
	if request.Id < 1 {
		s.logger.Infof("ошибка get studio: %s", fmt.Errorf("неверный id: %w", err))
		return nil, fmt.Errorf("неверный id: %w", err)
	}

	//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//defer cancel()
	ctx := context.Background()
	studio, err = s.studioRepo.Get(ctx, &dto.GetStudioRequest{
		Id: request.Id,
	})

	if err != nil {
		s.logger.Errorf("ошибка get studio: %s", fmt.Errorf("получение студии по id: %w", err))
		return nil, fmt.Errorf("получение студии по id: %w", err)
	}

	return studio, err
}

func (s StudioService) GetAll(request *dto.GetStudioAllRequest) (studios []*model.Studio, err error) {

	//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//defer cancel()
	ctx := context.Background()
	studios, err = s.studioRepo.GetAll(ctx, &dto.GetStudioAllRequest{})

	if err != nil {
		s.logger.Errorf("ошибка get all studio: %s", fmt.Errorf("получение студий по id: %w", err))
		return nil, fmt.Errorf("получение студий по id: %w", err)
	}

	return studios, err
}

func (s StudioService) Add(ctx context.Context, request *dto.AddStudioRequest) (err error) {

	if request.Name == "" {
		s.logger.Infof("ошибка add studio: %s", fmt.Errorf("пустое имя: %w", err))
		return fmt.Errorf("пустое имя: %w", err)
	}

	//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//defer cancel()

	err = s.studioRepo.Add(ctx, &dto.AddStudioRequest{
		Name: request.Name,
	})

	if err != nil {
		s.logger.Errorf("ошибка add studio: %s", fmt.Errorf("добавление студии: %w", err))
		return fmt.Errorf("добавление студии: %w", err)
	}

	return err
}

func (s StudioService) Delete(request *dto.DeleteStudioRequest) (err error) {
	if request.Id < 1 {
		s.logger.Infof("ошибка add studio: %s", fmt.Errorf("неверный id: %w", err))
		return fmt.Errorf("неверный id: %w", err)
	}

	//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//defer cancel()
	ctx := context.Background()
	err = s.studioRepo.Delete(ctx, &dto.DeleteStudioRequest{
		Id: request.Id,
	})

	if err != nil {
		s.logger.Errorf("ошибка add studio: %s", fmt.Errorf("удаление студии по id: %w", err))
		return fmt.Errorf("удаление студии по id: %w", err)
	}

	return err
}
