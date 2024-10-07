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

type EquipmentService struct {
	equipmentRepo repositoryInterface.IEquipmentRepository
	reserveRepo   repositoryInterface.IReserveRepository
	logger        logger.Interface
}

func NewEquipmentService(
	logger logger.Interface,
	equipmentRepo repositoryInterface.IEquipmentRepository,
	reserveRepo repositoryInterface.IReserveRepository) serviceInterface.IEquipmentService {
	return &EquipmentService{
		logger:        logger,
		equipmentRepo: equipmentRepo,
		reserveRepo:   reserveRepo,
	}
}

func (s EquipmentService) GetByReserve(request *dto.GetEquipmentByReserveRequest) (equipments []*model.Equipment, err error) {
	if request.ReserveId < 1 {
		s.logger.Infof("ошибка get equipment by reserve %d: %s", request.ReserveId, fmt.Errorf("неверный reserveId: %w", err))
		return nil, fmt.Errorf("неверный reserveId: %w", err)
	}

	//ctx, cancel := context.WithTimeout(context.Background(), cmd.TimeOut*time.Second)
	//defer cancel()
	ctx := context.Background()
	equipments, err = s.equipmentRepo.GetByReserve(ctx, &dto.GetEquipmentByReserveRequest{
		ReserveId: request.ReserveId,
	})

	if err != nil {
		s.logger.Errorf("ошибка get equipment by reserve %d: %s", request.ReserveId, err.Error())
		return nil, fmt.Errorf("получение оборудования по reserveId: %w", err)
	}

	return equipments, err
}

func (s EquipmentService) Get(request *dto.GetEquipmentRequest) (equipment *model.Equipment, err error) {
	if request.Id < 1 {
		s.logger.Infof("ошибка get equipment %d by id: %s", request.Id, fmt.Errorf("неверный id: %w", err))
		return nil, fmt.Errorf("неверный id: %w", err)
	}

	//ctx, cancel := context.WithTimeout(context.Background(), cmd.TimeOut*time.Second)
	//defer cancel()
	ctx := context.Background()
	equipment, err = s.equipmentRepo.Get(ctx, &dto.GetEquipmentRequest{
		Id: request.Id,
	})

	if err != nil {
		s.logger.Errorf("ошибка get equipment %d by id: %s", request.Id, err.Error())
		return nil, fmt.Errorf("получение оборудования по id: %w", err)
	}

	return equipment, err
}

func (s EquipmentService) Add(request *dto.AddEquipmentRequest) (err error) {

	if request.Name == "" {
		s.logger.Infof("ошибка add equipment %s %d %d: %s", request.Name, request.Type, request.StudioId, fmt.Errorf("пустое название: %w", err))
		return fmt.Errorf("пустое название: %w", err)
	}

	if request.StudioId < 1 {
		s.logger.Infof("ошибка add equipment %s %d %d: %s", request.Name, request.Type, request.StudioId, fmt.Errorf("отрицательный id студии: %w", err))
		return fmt.Errorf("отрицательный id студии: %w", err)
	}

	if request.Type <= model.OutOfFirstEquipment || request.Type >= model.OutOfLastEquipment {
		s.logger.Infof("ошибка add equipment %s %d %d: %s", request.Name, request.Type, request.StudioId, fmt.Errorf("неверный тип оборудования"))
		return fmt.Errorf("неверный тип оборудования")
	}

	//ctx, cancel := context.WithTimeout(context.Background(), cmd.TimeOut*time.Second)
	//defer cancel()
	ctx := context.Background()
	err = s.equipmentRepo.Add(ctx, &dto.AddEquipmentRequest{
		Name:     request.Name,
		StudioId: request.StudioId,
		Type:     request.Type,
	})
	if err != nil {
		s.logger.Errorf("ошибка add equipment %s %d %d: %s", request.Name, request.Type, request.StudioId, fmt.Errorf("добавление студии: %w", err))
		return fmt.Errorf("добавление студии: %w", err)
	}

	return err
}

func (s EquipmentService) Update(request *dto.UpdateEquipmentRequest) (err error) {
	if request.Id < 1 {
		s.logger.Infof("ошибка update equipment %d %d %d: %s", request.Id, request.Type, request.Type, fmt.Errorf("неверный id: %w", err))
		return fmt.Errorf("неверный id: %w", err)
	}

	//ctx, cancel := context.WithTimeout(context.Background(), cmd.TimeOut*time.Second)
	//defer cancel()
	ctx := context.Background()
	isReserve, err := s.reserveRepo.IsEquipmentReserve(ctx, &dto.IsEquipmentReserveRequest{
		EquipmentId: request.Id,
	})
	if err != nil {
		s.logger.Errorf("ошибка update equipment %d %d %d: %s", request.Id, request.Type, request.Type, fmt.Errorf("получение броней c оборудованием: %w", err))
		return fmt.Errorf("получение броней c оборудованием: %w", err)
	}

	if isReserve == true {
		s.logger.Infof("ошибка update equipment %d %d %d: %s", request.Id, request.Type, request.Type, fmt.Errorf("получение броней c оборудованием: %w", err))
		return fmt.Errorf("нельзя обновить оборудование, тк на него есть бронь")
	}

	if request.Name == "" {
		s.logger.Infof("ошибка update equipment %d %d %d: %s", request.Id, request.Type, request.Type, fmt.Errorf("ошибка пустого название: %w", err))
		return fmt.Errorf("ошибка пустого название: %w", err)
	}

	if request.StudioId < 1 {
		s.logger.Infof("ошибка update equipment %d %d %d: %s", request.Id, request.Type, request.Type, fmt.Errorf("отрицательный id студии: %w", err))
		return fmt.Errorf("отрицательный id студии: %w", err)
	}

	if request.Type <= model.OutOfFirstEquipment || request.Type >= model.OutOfLastEquipment {
		s.logger.Infof("ошибка update equipment %d %d %d: %s", request.Id, request.Type, request.Type, fmt.Errorf("неверный тип оборудования"))
		return fmt.Errorf("неверный тип оборудования")
	}

	//ctx, cancel = context.WithTimeout(context.Background(), cmd.TimeOut*time.Second)
	//defer cancel()
	err = s.equipmentRepo.Update(ctx, &dto.UpdateEquipmentRequest{
		Id:       request.Id,
		Name:     request.Name,
		StudioId: request.StudioId,
		Type:     request.Type,
	})
	if err != nil {
		s.logger.Errorf("ошибка update equipment %d %d %d: %s", request.Id, request.Type, request.Type, fmt.Errorf("обновление студии: %w", err))
		return fmt.Errorf("обновление студии: %w", err)
	}

	return err
}

func (s EquipmentService) Delete(request *dto.DeleteEquipmentRequest) (err error) {
	if request.Id < 1 {
		s.logger.Infof("ошибка delete equipment %d: %s", request.Id, fmt.Errorf("неправильный id: %w", err))
		return fmt.Errorf("неправильный id: %w", err)
	}

	//ctx, cancel := context.WithTimeout(context.Background(), cmd.TimeOut*time.Second)
	//defer cancel()
	ctx := context.Background()
	err = s.equipmentRepo.Delete(ctx, &dto.DeleteEquipmentRequest{
		Id: request.Id,
	})

	if err != nil {
		s.logger.Errorf("ошибка delete equipment %d: %s", request.Id, fmt.Errorf("ошибка получения студий по id: %w", err))
		return fmt.Errorf("ошибка получения студий по id: %w", err)
	}

	return err
}

func (s EquipmentService) GetByStudio(request *dto.GetEquipmentByStudioRequest) (equipments []*model.Equipment, err error) {
	if request.StudioId < 1 {
		s.logger.Infof("ошибка get equipment by studio %d: %s", request.StudioId, fmt.Errorf("неверный id: %w", err))
		return nil, fmt.Errorf("неверный id: %w", err)
	}

	//ctx, cancel := context.WithTimeout(context.Background(), cmd.TimeOut*time.Second)
	//defer cancel()
	ctx := context.Background()
	equipments, err = s.equipmentRepo.GetByStudio(ctx, &dto.GetEquipmentByStudioRequest{
		StudioId: request.StudioId,
	})

	if err != nil {
		s.logger.Errorf("ошибка get equipment by studio %d: %s", request.StudioId, fmt.Errorf("получение оборудования по типу: %w", err))
		return nil, fmt.Errorf("получение оборудования по типу: %w", err)
	}

	return equipments, err
}
