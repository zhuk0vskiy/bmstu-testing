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

type ReserveService struct {
	reserveRepo         repositoryInterface.IReserveRepository
	producerRepo        repositoryInterface.IProducerRepository
	InstrumentalistRepo repositoryInterface.IInstrumentalistRepository
	roomRepo            repositoryInterface.IRoomRepository
	logger              logger.Interface
}

func NewReserveService(
	logger logger.Interface,
	reserveRepo repositoryInterface.IReserveRepository,
	roomRepo repositoryInterface.IRoomRepository,
	producerRepo repositoryInterface.IProducerRepository,
	instrumentalistRepo repositoryInterface.IInstrumentalistRepository) serviceInterface.IReserveService {
	return &ReserveService{
		logger:              logger,
		reserveRepo:         reserveRepo,
		producerRepo:        producerRepo,
		InstrumentalistRepo: instrumentalistRepo,
		roomRepo:            roomRepo,
	}
}

func (s ReserveService) GetAll(request *dto.GetAllReserveRequest) (equipments []*model.Reserve, err error) {

	//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//defer cancel()
	ctx := context.Background()
	equipments, err = s.reserveRepo.GetAll(ctx, &dto.GetAllReserveRequest{})
	if err != nil {
		s.logger.Errorf("ошибка get all reserve: %s", fmt.Errorf("получение всех броней: %w", err))
		return nil, fmt.Errorf("получение всех броней: %w", err)
	}

	return equipments, err
}

func (s ReserveService) Add(ctx context.Context, request *dto.AddReserveRequest) (err error) {

	if request.ProducerId < 0 {
		s.logger.Infof("ошибка add reserve: %s", fmt.Errorf("id продюсера меньше 0: %w", err))
		return fmt.Errorf("id продюсера меньше 0: %w", err)
	}

	if request.RoomId < 1 {
		s.logger.Infof("ошибка add reserve: %s", fmt.Errorf("id комнаты меньше 1: %w", err))
		return fmt.Errorf("id комнаты меньше 1: %w", err)
	}

	if request.InstrumentalistId < 0 {
		s.logger.Infof("ошибка add reserve: %s", fmt.Errorf("id инструменталиста меньше 0: %w", err))
		return fmt.Errorf("id инструменталиста меньше 0: %w", err)
	}

	if request.UserId < 1 {
		s.logger.Infof("ошибка add reserve: %s", fmt.Errorf("id пользователя меньше 1: %w", err))
		return fmt.Errorf("id пользователя меньше 1: %w", err)
	}

	for _, equipment := range request.EquipmentId {
		if equipment < 1 {
			s.logger.Infof("ошибка add reserve: %s", fmt.Errorf("id оборудования меньше 1: %w", err))
			return fmt.Errorf("id оборудования меньше 1: %w", err)
		}
	}

	//if request.TimeInterval
	if request.TimeInterval.StartTime.Unix() >= request.TimeInterval.EndTime.Unix() {
		s.logger.Infof("ошибка add reserve: %s", fmt.Errorf("время начала больше времени конца: %w", err))
		return fmt.Errorf("время начала больше времени конца: %w", err)
	}

	//room, err := s.roomRepo.Get(ctx, &dto.GetRoomRequest{
	//	Id: request.RoomId,
	//})
	//if err != nil {
	//	return fmt.Errorf("%w", err)
	//}
	//if room.StartHour > int64(request.TimeInterval.StartTime.Hour()) || room.EndHour < int64(request.TimeInterval.EndTime.Hour()) {
	//	return fmt.Errorf("выбранное время не входит во время комнаты")
	//}
	//
	//producer, err := s.producerRepo.Get(ctx, &dto.GetProducerRequest{
	//	Id: request.ProducerId,
	//})
	//if err != nil {
	//	return fmt.Errorf("%w", err)
	//}
	//if producer.StartHour > int64(request.TimeInterval.StartTime.Hour()) || producer.EndHour < int64(request.TimeInterval.EndTime.Hour()) {
	//	return fmt.Errorf("выбранное время не входит во время продюсера")
	//}
	//
	//instrumentalist, err := s.InstrumentalistRepo.Get(ctx, &dto.GetInstrumentalistRequest{
	//	Id: request.InstrumentalistId,
	//})
	//if err != nil {
	//	return fmt.Errorf("%w", err)
	//}
	//if instrumentalist.StartHour > int64(request.TimeInterval.StartTime.Hour()) || instrumentalist.EndHour < int64(request.TimeInterval.EndTime.Hour()) {
	//	return fmt.Errorf("выбранное время не входит во время инструменталиста")
	//}

	//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//defer cancel()
	err = s.reserveRepo.Add(ctx, &dto.AddReserveRequest{
		UserId:            request.UserId,
		RoomId:            request.RoomId,
		EquipmentId:       request.EquipmentId,
		ProducerId:        request.ProducerId,
		InstrumentalistId: request.InstrumentalistId,
		TimeInterval:      request.TimeInterval,
	})
	if err != nil {
		s.logger.Errorf("ошибка add reserve: %s", fmt.Errorf("добавление брони: %w", err))
		return fmt.Errorf("добавление брони: %w", err)
	}

	s.logger.Infof("пользователь %d создал бронь на комнату %d, оборудование %d, продюсера %d, инструменталиста %d",
		request.UserId,
		request.RoomId,
		request.EquipmentId,
		request.ProducerId,
		request.InstrumentalistId,
	)

	return err
}

func (s ReserveService) Delete(request *dto.DeleteReserveRequest) (err error) {
	if request.Id < 1 {
		s.logger.Infof("ошибка delete reserve: %s", fmt.Errorf("id меньше 1: %w", err))
		return fmt.Errorf("id меньше 1: %w", err)
	}

	//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//defer cancel()
	ctx := context.Background()
	err = s.reserveRepo.Delete(ctx, &dto.DeleteReserveRequest{
		Id: request.Id,
	})
	if err != nil {
		s.logger.Errorf("ошибка delete reserve: %s", fmt.Errorf("удаление брони: %w", err))
		return fmt.Errorf("удаление брони: %w", err)
	}

	s.logger.Infof("удаление брони %d", request.Id)

	return err
}
