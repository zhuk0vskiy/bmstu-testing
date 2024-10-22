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

type InstrumentalistService struct {
	instrumentalistRepo repositoryInterface.IInstrumentalistRepository
	reserveRepo         repositoryInterface.IReserveRepository
	logger              logger.Interface
}

func NewInstrumentalistService(
	logger logger.Interface,
	instrumentalistRepo repositoryInterface.IInstrumentalistRepository,
	reserveRepo repositoryInterface.IReserveRepository) serviceInterface.IInstrumentalistService {
	return &InstrumentalistService{
		logger:              logger,
		instrumentalistRepo: instrumentalistRepo,
		reserveRepo:         reserveRepo,
	}
}

func (s InstrumentalistService) GetByStudio(request *dto.GetInstrumentalistByStudioRequest) (instrumentalists []*model.Instrumentalist, err error) {
	if request.StudioId < 1 {
		s.logger.Infof("ошибка get instrumentalist by studio: %s", fmt.Errorf("получение оборудования по типу: %w", err))
		return nil, fmt.Errorf("неверный id: %w", err)
	}

	//ctx, cancel := context.WithTimeout(context.Background(), cmd.TimeOut*time.Second)
	//defer cancel()
	ctx := context.Background()
	instrumentalists, err = s.instrumentalistRepo.GetByStudio(ctx, &dto.GetInstrumentalistByStudioRequest{
		StudioId: request.StudioId,
	})

	if err != nil {
		s.logger.Errorf("ошибка get instrumentalist by studio: %s", fmt.Errorf("получение инструменталистов по студии: %w", err))
		return nil, fmt.Errorf("получение инструменталистов по студии: %w", err)
	}
	return instrumentalists, err
}

func (s InstrumentalistService) Get(request *dto.GetInstrumentalistRequest) (instrumentalist *model.Instrumentalist, err error) {
	if request.Id < 1 {
		//s.logger.Infof("ошибка get instrumentalist by id: %s", fmt.Errorf("неверный id: %w", err))
		return nil, fmt.Errorf("неверный id: %w", err)
	}

	//ctx, cancel := context.WithTimeout(context.Background(), cmd.TimeOut*time.Second)
	//defer cancel()
	ctx := context.Background()
	instrumentalist, err = s.instrumentalistRepo.Get(ctx, &dto.GetInstrumentalistRequest{
		Id: request.Id,
	})

	if err != nil {
		s.logger.Errorf("ошибка get instrumentalist by id: %s", fmt.Errorf("получение инструменталиста по id: %w", err))
		return nil, fmt.Errorf("получение инструменталиста по id: %w", err)
	}

	return instrumentalist, err
}

//func (s InstrumentalistService) GetAll() (instrumentalists []*domain.Instrumentalist, err error) {
//	instrumentalists, err = s.instrumentalistRepo.GetAll()
//
//	if err != nil {
//		return nil, fmt.Errorf("получение всех инструменталистов: %w", err)
//	}
//
//	return instrumentalists, nil
//}

func (s InstrumentalistService) Update(request *dto.UpdateInstrumentalistRequest) (err error) {
	if request.Id < 1 {
		s.logger.Infof("ошибка update instrumentalist: %s", fmt.Errorf("неверный id: %w", err))
		return fmt.Errorf("неверный id: %w", err)
	}

	//ctx, cancel := context.WithTimeout(context.Background(), cmd.TimeOut*time.Second)
	//defer cancel()
	ctx := context.Background()
	isReserve, err := s.reserveRepo.IsInstrumentalistReserve(ctx, &dto.IsInstrumentalistReserveRequest{
		InstrumentalistId: request.Id,
	})
	if err != nil {
		s.logger.Errorf("ошибка update instrumentalist: %s", fmt.Errorf("инструменталист забронирован: %w", err))
		return fmt.Errorf("ошибка при проверке: %w", err)
	}

	if isReserve == true {
		s.logger.Infof("ошибка update instrumentalist: %s", fmt.Errorf("нельзя обновить инструменталиста, тк на него есть бронь: %w", err))
		return fmt.Errorf("нельзя обновить инструменталиста, тк на него есть бронь: %w", err)
	}

	if request.Name == "" {
		s.logger.Infof("ошибка update instrumentalist: %s", fmt.Errorf("пустое имя: %w", err))
		return fmt.Errorf("пустое имя: %w", err)
	}

	if request.StudioId < 0 {
		s.logger.Infof("ошибка update instrumentalist: %s", fmt.Errorf("id студии меньше 0"))
		return fmt.Errorf("id студии меньше 0")
	}
	//
	//if request.StartTime.s

	if request.StartHour == request.EndHour {
		s.logger.Infof("ошибка update instrumentalist: %s", fmt.Errorf("время начала работы равно времени конца"))
		return fmt.Errorf("время начала работы равно времени конца")
	}

	if request.EndHour > 23 || request.StartHour > 23 ||
		request.EndHour < 0 || request.StartHour < 0 {
		s.logger.Infof("ошибка update instrumentalist: %s", fmt.Errorf("время конца/начала работы не входит в размер суток (от 00 до 23)"))
		return fmt.Errorf("время конца/начала работы не входит в размер суток (от 00 до 23)")
	}

	err = s.instrumentalistRepo.Update(ctx, &dto.UpdateInstrumentalistRequest{
		Id:        request.Id,
		Name:      request.Name,
		StudioId:  request.StudioId,
		StartHour: request.StartHour,
		EndHour:   request.EndHour,
	})
	if err != nil {
		s.logger.Errorf("ошибка update instrumentalist: %s", fmt.Errorf("обновление инструменталиста: %w", err))
		return fmt.Errorf("обновление инструменталиста: %w", err)
	}

	return err
}

func (s InstrumentalistService) Add(ctx context.Context, request *dto.AddInstrumentalistRequest) (err error) {

	if request.Name == "" {
		s.logger.Infof("ошибка add instrumentalist: %s", fmt.Errorf("пустое имя: %w", err))
		return fmt.Errorf("пустое имя: %w", err)
	}

	if request.StudioId < 1 {
		s.logger.Infof("ошибка add instrumentalist: %s", fmt.Errorf("id студии меньше 0"))
		return fmt.Errorf("id студии меньше 0")
	}

	if request.StartHour == request.EndHour {
		s.logger.Infof("ошибка add instrumentalist: %s", fmt.Errorf("время начала работы равно времени конца"))
		return fmt.Errorf("время начала работы равно времени конца")
	}

	if request.EndHour > 23 || request.StartHour > 23 ||
		request.StartHour < 0 || request.EndHour < 0 {
		s.logger.Infof("ошибка add instrumentalist: %s", fmt.Errorf("время конца/начала работы не входит в размер суток (от 00 до 23)"))
		return fmt.Errorf("время конца/начала работы не входит в размер суток (от 00 до 23)")
	}

	//ctx, cancel := context.WithTimeout(context.Background(), cmd.TimeOut*time.Second)
	//defer cancel()
	ctx = context.Background()
	err = s.instrumentalistRepo.Add(ctx, &dto.AddInstrumentalistRequest{
		Name:      request.Name,
		StudioId:  request.StudioId,
		StartHour: request.StartHour,
		EndHour:   request.EndHour,
	})
	if err != nil {
		s.logger.Errorf("ошибка add instrumentalist: %s", fmt.Errorf("добавление инструменталиста: %w", err))
		return fmt.Errorf("добавление инструменталиста: %w", err)
	}

	return err
}

func (s InstrumentalistService) Delete(request *dto.DeleteInstrumentalistRequest) (err error) {
	if request.Id < 1 {
		s.logger.Infof("ошибка delete instrumentalist: %s", fmt.Errorf("delet instrumentalist: request.Id < 1"))
		return fmt.Errorf("delet instrumentalist: request.Id < 1")
	}

	//ctx, cancel := context.WithTimeout(context.Background(), cmd.TimeOut*time.Second)
	//defer cancel()
	ctx := context.Background()
	err = s.instrumentalistRepo.Delete(ctx, &dto.DeleteInstrumentalistRequest{
		Id: request.Id,
	})

	if err != nil {
		s.logger.Errorf("ошибка delete instrumentalist: %s", fmt.Errorf("удаление инструменталиста: %w", err))
		return fmt.Errorf("удаление инструменталиста: %w", err)
	}

	return err
}
