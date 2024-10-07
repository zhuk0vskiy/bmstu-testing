package impl

import (
	"backend/src/internal/model"
	"backend/src/internal/model/dto"
	repositoryInterface "backend/src/internal/repository/interface"
	serviceInterface "backend/src/internal/service/interface"
	"backend/src/pkg/logger"
	"context"
	"fmt"
	"slices"
)

type ValidateTimeService struct {
	logger              logger.Interface
	roomRepo            repositoryInterface.IRoomRepository
	equipmentRepo       repositoryInterface.IEquipmentRepository
	producerRepo        repositoryInterface.IProducerRepository
	instrumentalistRepo repositoryInterface.IInstrumentalistRepository
	reserveRepo         repositoryInterface.IReserveRepository
	//reservedEquipmentRepo repositoryInterface.IReservedEquipmentRepository
}

func NewValidateTimeService(
	logger logger.Interface,
	roomRepo repositoryInterface.IRoomRepository,
	equipmentRepo repositoryInterface.IEquipmentRepository,
	producerRepo repositoryInterface.IProducerRepository,
	instrumentalistRepo repositoryInterface.IInstrumentalistRepository,
	reserveRepo repositoryInterface.IReserveRepository,
	// reservedEquipmentRepo repositoryInterface.IReservedEquipmentRepository,
) serviceInterface.IValidateTimeService {
	return &ValidateTimeService{
		logger:              logger,
		roomRepo:            roomRepo,
		equipmentRepo:       equipmentRepo,
		producerRepo:        producerRepo,
		instrumentalistRepo: instrumentalistRepo,
		reserveRepo:         reserveRepo,
		//reservedEquipmentRepo: reservedEquipmentRepo,
	}
}

func (s ValidateTimeService) GetSuitableStuff(request *dto.GetSuitableStuffRequest) (
	notReservedRooms []*model.Room,
	notReservedEquipments [][]*model.Equipment,
	notReservedProducers []*model.Producer,
	notReservedInstrumentalists []*model.Instrumentalist,
	err error) {

	if request.StudioId < 1 {
		s.logger.Infof("ошибка get suitable stuff by studio id %d: %s", request.StudioId, fmt.Errorf("id студии меньше 1"))
		return nil, nil, nil, nil, fmt.Errorf("id студии меньше 1")
	}

	ctx := context.Background() //, cancel := context.WithTimeout(context.Background(), cmd.TimeOut*time.Second)
	//defer cancel()
	notReservedRooms, err = s.getNotReservedRooms(ctx, &dto.GetNotReservedRoomsRequest{
		ChoosenInterval: request.TimeInterval,
		StudioId:        request.StudioId,
	})
	if err != nil {
		s.logger.Errorf("ошибка get suitable stuff by studio id %d: %s", request.StudioId, fmt.Errorf("поиск доступных комнат"))
		return nil,
			nil,
			nil,
			nil,
			fmt.Errorf("поиск доступных комнат")
	}

	notReservedEquipments, err = s.getNotReservedEquipments(ctx, &dto.GetNotReservedEquipmentsRequest{
		ChoosenInterval: request.TimeInterval,
		StudioId:        request.StudioId,
	})
	if err != nil {
		s.logger.Errorf("ошибка get suitable stuff by studio id %d: %s", request.StudioId, fmt.Errorf("поиск доступного оборудования"))
		return nil, nil, nil, nil, fmt.Errorf("поиск доступного оборудования")
	}

	notReservedProducers, err = s.getNotReservedProducers(ctx, &dto.GetNotReservedProducersRequest{
		ChoosenInterval: request.TimeInterval,
		StudioId:        request.StudioId,
	})
	if err != nil {
		s.logger.Errorf("ошибка get suitable stuff by studio id %d: %s", request.StudioId, fmt.Errorf("поиск доступного продюсера"))
		return nil, nil, nil, nil, fmt.Errorf("поиск доступного продюсера")
	}

	notReservedInstrumentalists, err = s.getNotReservedInstrumentalists(ctx, &dto.GetNotReservedInstrumentalistsRequest{
		ChoosenInterval: request.TimeInterval,
		StudioId:        request.StudioId,
	})
	if err != nil {
		s.logger.Errorf("ошибка get suitable stuff by studio id %d: %s", request.StudioId, fmt.Errorf("поиск доступного инструменталиста"))
		return nil, nil, nil, nil, fmt.Errorf("поиск доступного инструменталиста")
	}

	return notReservedRooms, notReservedEquipments, notReservedProducers, notReservedInstrumentalists, err
}

func (s ValidateTimeService) getNotReservedRooms(ctx context.Context, request *dto.GetNotReservedRoomsRequest) (notReservedRooms []*model.Room, err error) {

	if request.StudioId <= 0 {
		s.logger.Infof("ошибка get not reserved rooms by studio id %d: %s", request.StudioId, fmt.Errorf("id студии меньше 1"))
		return nil, fmt.Errorf("id студии меньше 1")
	}

	//if request.StartTime.Unix() >= request.EndTime.Unix() {
	//	return nil, fmt.Errorf("")
	//}
	//request.ChoosenInterval

	rooms, err := s.roomRepo.GetByStudio(ctx, &dto.GetRoomByStudioRequest{
		StudioId: request.StudioId,
	})
	if err != nil {
		s.logger.Errorf("ошибка get not reserved rooms by studio id %d: %s", request.StudioId, fmt.Errorf("get all rooms error"))
		return nil, fmt.Errorf("get all rooms error")
	}

	for _, room := range rooms {
		reserveFlag := false
		//fmt.Println("searching for producer:", producer.Id)
		//fmt.Println("id -- ", room.Id)

		if int64(request.ChoosenInterval.StartTime.Hour()) >= room.StartHour &&
			int64(request.ChoosenInterval.EndTime.Hour()) <= room.EndHour {
			reserves, err := s.reserveRepo.GetByRoomId(ctx, &dto.GetReserveByRoomIdRequest{RoomId: room.Id})
			if err != nil {
				s.logger.Errorf("ошибка get not reserved rooms by studio id %d: %s", request.StudioId, fmt.Errorf("get all reserves error"))
				return nil, fmt.Errorf("get all reserves error")
			}
			for _, reserve := range reserves {
				if isIntervalsIntersect(*request.ChoosenInterval, *reserve.TimeInterval) == true {
					reserveFlag = true
					break
				}
			}
			if reserveFlag == false {
				notReservedRooms = append(notReservedRooms, room)
			}
		}
	}

	//fmt.Println(len(notReservedRooms))
	return notReservedRooms, err
}

func (s ValidateTimeService) getNotReservedProducers(ctx context.Context, request *dto.GetNotReservedProducersRequest) (notReservedProducers []*model.Producer, err error) {

	if request.StudioId <= 0 {
		s.logger.Infof("ошибка get not reserved producers by studio id %d: %s", request.StudioId, fmt.Errorf("id студии меньше 1"))
		return nil, fmt.Errorf("id студии меньше 1")
	}

	producers, err := s.producerRepo.GetByStudio(ctx, &dto.GetProducerByStudioRequest{
		StudioId: request.StudioId,
	})
	if err != nil {
		s.logger.Errorf("ошибка get not reserved producers by studio id %d: %s", request.StudioId, fmt.Errorf("get all rooms error"))
		return nil, fmt.Errorf("get all rooms error")
	}

	for _, producer := range producers {
		reserveFlag := false

		if int64(request.ChoosenInterval.StartTime.Hour()) >= producer.StartHour &&
			int64(request.ChoosenInterval.EndTime.Hour()) <= producer.EndHour {
			reserves, err := s.reserveRepo.GetByProducerId(ctx, &dto.GetReserveByProducerIdRequest{ProducerId: producer.Id}) //TODO: спросить, нормально ли бегать в бд при каждой итерации
			if err != nil {
				s.logger.Errorf("ошибка get not reserved producers by studio id %d: %s", request.StudioId, fmt.Errorf("get all reserves error"))
				return nil, fmt.Errorf("get all reserves error")
			}
			for _, reserve := range reserves {
				if isIntervalsIntersect(*request.ChoosenInterval, *reserve.TimeInterval) == true {
					reserveFlag = true
					break
				}
			}
			if reserveFlag == false {
				notReservedProducers = append(notReservedProducers, producer)
			}
		}
	}

	//fmt.Println(len(notReservedRooms))
	//for i, _ := range notReservedProducers {
	//	fmt.Println(notReservedProducers[i])
	//}
	return notReservedProducers, err
}

func (s ValidateTimeService) getNotReservedInstrumentalists(ctx context.Context, request *dto.GetNotReservedInstrumentalistsRequest) (notReservedInstrumentalists []*model.Instrumentalist, err error) {

	if request.StudioId <= 0 {
		s.logger.Infof("ошибка get not reserved instrumentaists by studio id %d: %s", request.StudioId, fmt.Errorf("id студии меньше 1"))
		return nil, fmt.Errorf("id студии меньше 1")
	}

	instrumentalists, err := s.instrumentalistRepo.GetByStudio(ctx, &dto.GetInstrumentalistByStudioRequest{
		StudioId: request.StudioId,
	})

	if err != nil {
		s.logger.Errorf("ошибка get not reserved instrumentaists by studio id %d: %s", request.StudioId, fmt.Errorf("get all rooms error"))
		return nil, fmt.Errorf("get all rooms error")
	}

	for _, instrumentalist := range instrumentalists {
		reserveFlag := false
		//fmt.Println("searching for producer:", producer.Id)

		if int64(request.ChoosenInterval.StartTime.Hour()) >= instrumentalist.StartHour &&
			int64(request.ChoosenInterval.EndTime.Hour()) <= instrumentalist.EndHour {
			reserves, err := s.reserveRepo.GetByInstrumentalistId(ctx, &dto.GetReserveByInstrumentalistIdRequest{InstrumentalistId: instrumentalist.Id}) //TODO: спросить, нормально ли бегать в бд при каждой итерации
			if err != nil {
				s.logger.Errorf("ошибка get not reserved instrumentaists by studio id %d: %s", request.StudioId, fmt.Errorf("get all reserves error"))
				return nil, fmt.Errorf("get all reserves error")
			}
			for _, reserve := range reserves {
				if isIntervalsIntersect(*request.ChoosenInterval, *reserve.TimeInterval) == true {
					reserveFlag = true
					break
				}
				//reserveStartHour := int64(reserve.StartTime.Hour())
				//reserveEndHour := int64(reserve.EndTime.Hour())
				//
				////fmt.Println("	looking reserve:", reserve.Id)
				//if (choosenStartHour >= reserveStartHour && choosenStartHour < reserveEndHour) ||
				//	(choosenEndHour <= reserveEndHour && choosenEndHour > reserveStartHour) ||
				//	(choosenStartHour <= reserveStartHour && choosenEndHour >= reserveEndHour) {
				//	reserveFlag = true
				//	break
				//}
			}
			if reserveFlag == false {
				notReservedInstrumentalists = append(notReservedInstrumentalists, instrumentalist)
			}
		}
	}

	//fmt.Println(len(notReservedRooms))
	return notReservedInstrumentalists, err
}

func (s ValidateTimeService) getNotReservedEquipments(ctx context.Context, request *dto.GetNotReservedEquipmentsRequest) (notReservedEquipments [][]*model.Equipment, err error) {
	//ctx = context.Background()
	notReservedEquipments = make([][]*model.Equipment, 0)

	if request.StudioId <= 0 {
		s.logger.Infof("ошибка get not reserved equipments by studio id %d: %s", request.StudioId, fmt.Errorf("id студии меньше 1"))
		return nil, fmt.Errorf("id студии меньше 1")
	}

	if request.ChoosenInterval.StartTime.Unix() >= request.ChoosenInterval.EndTime.Unix() {
		s.logger.Infof("ошибка get not reserved equipments by studio id %d: %s", request.StudioId, fmt.Errorf("время начала больше или равно времени конца"))
		return nil, fmt.Errorf("время начала больше или равно времени конца")
	}

	for equipmentType := int64(model.OutOfFirstEquipment + 1); equipmentType < int64(model.OutOfLastEquipment); equipmentType++ {
		//fmt.Println(equipmentType)
		notFullTimeFreeEquipments := make([]*model.Equipment, 0)
		equipmentsAndTimes, err := s.equipmentRepo.GetNotFullTimeFreeByStudioAndType(
			ctx,
			&dto.GetEquipmentNotFullTimeFreeByStudioAndTypeRequest{
				StudioId: request.StudioId,
				Type:     equipmentType,
				TimeInterval: &model.TimeInterval{
					StartTime: request.ChoosenInterval.StartTime,
					EndTime:   request.ChoosenInterval.EndTime,
				},
			})
		for _, equipmentAndTime := range equipmentsAndTimes {
			if isIntervalsIntersect(*equipmentAndTime.TimeInterval, *request.ChoosenInterval) == false {
				notFullTimeFreeEquipments = append(notFullTimeFreeEquipments, equipmentAndTime.Equipment)
			}
		}

		if err != nil {
			s.logger.Errorf("ошибка get not reserved equipments by studio id %d: %s", request.StudioId, fmt.Errorf("ошибка получения не полностью свободного оборудования по студии и типу из репозитория: %w", err))
			return nil, fmt.Errorf("ошибка получения не полностью забронированного оборудования по студии и типу из репозитория: %w", err)
		}

		fullTimeFreeEquipments, err := s.equipmentRepo.GetFullTimeFreeByStudioAndType(
			ctx,
			&dto.GetEquipmentFullTimeFreeByStudioAndTypeRequest{
				StudioId: request.StudioId,
				Type:     equipmentType,
			})

		if err != nil {
			s.logger.Errorf("ошибка get not reserved equipments by studio id %d: %s", request.StudioId, fmt.Errorf("ошибка получения полностью свободного оборудования по студии и типу из репозитория: %w", err))
			return nil, fmt.Errorf("ошибка получения полностью свободного оборудования по студии и типу из репозитория: %w", err)
		}

		notFullTimeFreeEquipments = slices.Concat(notFullTimeFreeEquipments, fullTimeFreeEquipments)

		notReservedEquipments = append(notReservedEquipments, notFullTimeFreeEquipments)

	}

	return notReservedEquipments, err
}
