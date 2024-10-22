package _interface

import (
	"backend/src/internal/model"
	"backend/src/internal/model/dto"
	"context"
)

type IReserveService interface {
	Add(ctx context.Context, request *dto.AddReserveRequest) error // Для создания брони
	Delete(request *dto.DeleteReserveRequest) error                // Для удаления брони
	GetAll(request *dto.GetAllReserveRequest) (equipments []*model.Reserve, err error)
	//Get(id int64) (*Reserve, error)
	//GetNotReservedRooms(startTime, endTime time.Time, studioId int64) ([]*Room, error)
	//GetNotReservedProducers(startTime time.Time, endTime time.Time, studioId int64) ([]*Producer, error)
	//GetNotReservedEquipments(startTime time.Time, endTime time.Time, studioId int64) ([]*Equipment, error)
	//GetNotReservedInstrumentalists(startTime time.Time, endTime time.Time, studioId int64) ([]*Instrumentalist, error)
}
