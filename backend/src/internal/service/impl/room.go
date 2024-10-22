package impl

import (
	"backend/src/internal/model"
	"backend/src/internal/model/dto"
	repositoryInterface "backend/src/internal/repository/interface"
	serviceInterface "backend/src/internal/service/interface"
	"context"
	"fmt"
)

type RoomService struct {
	roomRepo    repositoryInterface.IRoomRepository
	reserveRepo repositoryInterface.IReserveRepository
}

func NewRoomService(
	roomRepo repositoryInterface.IRoomRepository,
	reserveRepo repositoryInterface.IReserveRepository) serviceInterface.IRoomService {
	return &RoomService{
		roomRepo:    roomRepo,
		reserveRepo: reserveRepo,
	}
}

func (s RoomService) Get(request *dto.GetRoomRequest) (room *model.Room, err error) {
	if request.Id < 1 {
		return nil, fmt.Errorf("неверный id: %w", err)
	}

	//ctx, cancel := context.WithTimeout(context.Background(), cmd.TimeOut*time.Second)
	//defer cancel()
	ctx := context.Background()
	room, err = s.roomRepo.Get(ctx, &dto.GetRoomRequest{
		Id: request.Id,
	})

	if err != nil {
		return nil, fmt.Errorf("получение комнаты по id: %w", err)
	}

	return room, err
}

func (s RoomService) GetByStudio(request *dto.GetRoomByStudioRequest) (rooms []*model.Room, err error) {
	if request.StudioId < 1 {
		return nil, fmt.Errorf("неверный id: %w", err)
	}

	//ctx, cancel := context.WithTimeout(context.Background(), cmd.TimeOut*time.Second)
	//defer cancel()
	ctx := context.Background()
	rooms, err = s.roomRepo.GetByStudio(ctx, &dto.GetRoomByStudioRequest{
		StudioId: request.StudioId,
	})

	if err != nil {
		return nil, fmt.Errorf("получение комнат по студии: %w", err)
	}

	return rooms, err
}

func (s RoomService) Update(request *dto.UpdateRoomRequest) (err error) {
	//TODO: сделать транзакцию
	if request.Id < 1 {
		return fmt.Errorf("неверный id: %w", err)
	}

	//ctx, cancel := context.WithTimeout(context.Background(), cmd.TimeOut*time.Second)
	//defer cancel()
	ctx := context.Background()
	isReserve, err := s.reserveRepo.IsRoomReserve(ctx, &dto.IsRoomReserveRequest{
		RoomId: request.Id,
	})

	if err != nil {
		return fmt.Errorf("получение всех броней: %w", err)
	}

	if isReserve == true {
		return fmt.Errorf("нельзя обновить комнату, тк на нее есть бронь: %w", err)
	}

	if request.Name == "" {
		return fmt.Errorf("пустое название: %w", err)
	}

	if request.StudioId < 1 {
		return fmt.Errorf("id студии меньше 0")
	}
	//
	//if room.StartTime.s

	if request.StartHour == request.EndHour {
		return fmt.Errorf("время начала работы равно времени конца")
	}

	if request.EndHour > 23 || request.StartHour > 23 ||
		request.EndHour < 0 || request.StartHour < 0 {
		return fmt.Errorf("время конца/начала работы не входит в размер суток (от 00 до 23)")
	}

	err = s.roomRepo.Update(ctx, &dto.UpdateRoomRequest{
		Id:        request.Id,
		Name:      request.Name,
		StudioId:  request.StudioId,
		StartHour: request.StartHour,
		EndHour:   request.EndHour,
	})
	if err != nil {
		return fmt.Errorf("обновление комнаты: %w", err)
	}

	return err
}

func (s RoomService) Add(ctx context.Context, request *dto.AddRoomRequest) (err error) {

	if request.Name == "" {
		return fmt.Errorf("пустое название: %w", err)
	}

	if request.StudioId < 1 {
		return fmt.Errorf("id комнаты меньше 0")
	}

	if request.StartHour >= request.EndHour {
		return fmt.Errorf("время начала работы равно времени конца")
	}

	if request.EndHour > 23 || request.StartHour > 23 ||
		request.EndHour < 0 || request.StartHour < 0 {
		return fmt.Errorf("время конца/начала работы не входит в размер суток (от 00 до 23)")
	}

	//ctx, cancel := context.WithTimeout(context.Background(), cmd.TimeOut*time.Second)
	//defer cancel()
	err = s.roomRepo.Add(ctx, &dto.AddRoomRequest{
		Name:      request.Name,
		StudioId:  request.StudioId,
		StartHour: request.StartHour,
		EndHour:   request.EndHour,
	})
	if err != nil {
		return fmt.Errorf("добавление комнаты: %w", err)
	}

	return err
}

func (s RoomService) Delete(request *dto.DeleteRoomRequest) (err error) {
	if request.Id < 1 {
		return fmt.Errorf("неверный id: %w", err)
	}

	//ctx, cancel := context.WithTimeout(context.Background(), cmd.TimeOut*time.Second)
	//defer cancel()
	ctx := context.Background()
	err = s.roomRepo.Delete(ctx, &dto.DeleteRoomRequest{
		Id: request.Id,
	})

	if err != nil {
		return fmt.Errorf("удаление комнаты: %w", err)
	}

	return err
}
