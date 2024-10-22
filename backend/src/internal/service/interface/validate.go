package _interface

import (
	"backend/src/internal/model"
	"backend/src/internal/model/dto"
	"context"
)

type IValidateTimeService interface {
	GetSuitableStuff(ctx context.Context, request *dto.GetSuitableStuffRequest) (
		[]*model.Room,
		[][]*model.Equipment,
		[]*model.Producer,
		[]*model.Instrumentalist,
		error)

	//IsStuffFree(request *dto.IsStuffFreeRequest) (bool, error)

	//GetNotReservedRooms(request *dto.GetNotReservedRoomsRequest) ([]*model.Room, error)
	//GetNotReservedProducers(request *dto.GetNotReservedProducersRequest) ([]*model.Producer, error)
	//GetNotReservedEquipments(request *dto.GetNotReservedEquipmentsRequest) ([]*model.Equipment, error)
	//GetNotReservedInstrumentalists(request *dto.GetNotReservedInstrumentalistsRequest) ([]*model.Instrumentalist, error)
}
