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

type ProducerService struct {
	producerRepo repositoryInterface.IProducerRepository
	reserveRepo  repositoryInterface.IReserveRepository
	logger       logger.Interface
}

func NewProducerService(
	logger logger.Interface,
	producerRepo repositoryInterface.IProducerRepository,
	reserveRepo repositoryInterface.IReserveRepository) serviceInterface.IProducerService {
	return &ProducerService{
		logger:       logger,
		producerRepo: producerRepo,
		reserveRepo:  reserveRepo,
	}
}

func (s ProducerService) GetByStudio(request *dto.GetProducerByStudioRequest) (producers []*model.Producer, err error) {
	if request.StudioId < 1 {
		s.logger.Infof("ошибка get producer by studio: %s", fmt.Errorf("неверный id: %w", err))
		return nil, fmt.Errorf("неверный id: %w", err)
	}

	ctx := context.Background() // , cancel := context.WithTimeout(context.Background(), cmd.TimeOut*time.Second)
	//defer cancel()
	producers, err = s.producerRepo.GetByStudio(ctx, &dto.GetProducerByStudioRequest{
		StudioId: request.StudioId,
	})

	if err != nil {
		s.logger.Errorf("ошибка get producer by studio: %s", fmt.Errorf("получение продюсеров по студии: %w", err))
		return nil, fmt.Errorf("получение продюсеров по студии: %w", err)
	}
	return producers, err
}

func (s ProducerService) Get(request *dto.GetProducerRequest) (producer *model.Producer, err error) {
	if request.Id < 1 {
		//s.logger.Infof("ошибка get producer by id: %s", fmt.Errorf("неверный id: %w", err))
		return nil, fmt.Errorf("неверный id: %w", err)
	}

	ctx := context.Background() //, cancel := context.WithTimeout(context.Background(), cmd.TimeOut*time.Second)
	//defer cancel()
	producer, err = s.producerRepo.Get(ctx, &dto.GetProducerRequest{
		Id: request.Id,
	})

	if err != nil {
		s.logger.Errorf("ошибка get producer by id: %s", fmt.Errorf("получение продюсера по id: %w", err))
		return nil, fmt.Errorf("получение продюсера по id: %w", err)
	}

	return producer, err
}

func (s ProducerService) Update(request *dto.UpdateProducerRequest) (err error) {
	if request.Id < 1 {
		s.logger.Infof("ошибка update producer: %s", fmt.Errorf("неверный id: %w", err))
		return fmt.Errorf("неверный id: %w", err)
	}

	ctx := context.Background() //, cancel := context.WithTimeout(context.Background(), cmd.TimeOut*time.Second)
	//defer cancel()
	isReserve, err := s.reserveRepo.IsProducerReserve(ctx, &dto.IsProducerReserveRequest{
		ProducerId: request.Id,
	})
	if err != nil {
		s.logger.Errorf("ошибка update producer: %s", fmt.Errorf("получение всех броней: %w", err))
		return fmt.Errorf("получение всех броней: %w", err)
	}

	if isReserve == true {
		s.logger.Infof("ошибка update producer: %s", fmt.Errorf("нельзя обновить продюсера, тк на него есть бронь: %w", err))
		return fmt.Errorf("нельзя обновить продюсера, тк на него есть бронь: %w", err)
	}

	if request.Name == "" {
		s.logger.Infof("ошибка update producer: %s", fmt.Errorf("пустое имя: %w", err))
		return fmt.Errorf("пустое имя: %w", err)
	}

	if request.StudioId < 0 {
		s.logger.Infof("ошибка update producer: %s", fmt.Errorf("id студии меньше 0"))
		return fmt.Errorf("id студии меньше 0")
	}
	//
	//if request.StartTime.s

	if request.StartHour == request.EndHour {
		s.logger.Infof("ошибка update producer: %s", fmt.Errorf("время начала работы равно времени конца"))
		return fmt.Errorf("время начала работы равно времени конца")
	}

	if request.EndHour > 23 || request.StartHour > 23 ||
		request.EndHour < 0 || request.StartHour < 0 {
		s.logger.Infof("ошибка update producer: %s", fmt.Errorf("время конца/начала работы не входит в размер суток (от 00 до 23)"))
		return fmt.Errorf("время конца/начала работы не входит в размер суток (от 00 до 23)")
	}

	err = s.producerRepo.Update(ctx, &dto.UpdateProducerRequest{
		Id:        request.Id,
		Name:      request.Name,
		StudioId:  request.StudioId,
		StartHour: request.StartHour,
		EndHour:   request.EndHour,
	})
	if err != nil {
		s.logger.Errorf("ошибка update producer: %s", fmt.Errorf("обновление продюсера: %w", err))
		return fmt.Errorf("обновление продюсера: %w", err)
	}

	return err
}

func (s ProducerService) Add(ctx context.Context, request *dto.AddProducerRequest) (err error) {

	if request.Name == "" {
		s.logger.Infof("ошибка add producer: %s", fmt.Errorf("пустое имя: %w", err))
		return fmt.Errorf("пустое имя: %w", err)
	}

	if request.StudioId < 1 {
		s.logger.Infof("ошибка add producer: %s", fmt.Errorf("id студии меньше 0"))
		return fmt.Errorf("id студии меньше 0")
	}

	if request.StartHour == request.EndHour {
		s.logger.Infof("ошибка add producer: %s", fmt.Errorf("время начала работы равно времени конца"))
		return fmt.Errorf("время начала работы равно времени конца")
	}

	if request.EndHour > 23 || request.StartHour > 23 ||
		request.StartHour < 0 || request.EndHour < 0 {
		s.logger.Infof("ошибка add producer: %s", fmt.Errorf("время конца/начала работы не входит в размер суток (от 00 до 23)"))
		return fmt.Errorf("время конца/начала работы не входит в размер суток (от 00 до 23)")
	}

	//defer cancel()
	err = s.producerRepo.Add(ctx, &dto.AddProducerRequest{
		Name:      request.Name,
		StudioId:  request.StudioId,
		StartHour: request.StartHour,
		EndHour:   request.EndHour,
	})
	if err != nil {
		s.logger.Errorf("ошибка add producer: %s", fmt.Errorf("добавление продюсера: %w", err))
		return fmt.Errorf("добавление продюсера: %w", err)
	}

	return err
}

func (s ProducerService) Delete(request *dto.DeleteProducerRequest) (err error) {
	if request.Id < 1 {
		s.logger.Infof("ошибка delete producer: %s", fmt.Errorf("неверный id: %w", err))
		return fmt.Errorf("неверный id: %w", err)
	}

	ctx := context.Background() //, cancel := context.WithTimeout(context.Background(), cmd.TimeOut*time.Second)
	//defer cancel()
	err = s.producerRepo.Delete(ctx, &dto.DeleteProducerRequest{
		Id: request.Id,
	})

	if err != nil {
		s.logger.Errorf("ошибка delete producer: %s", fmt.Errorf("удаление продюсера: %w", err))
		return fmt.Errorf("удаление продюсера: %w", err)
	}

	return err
}
