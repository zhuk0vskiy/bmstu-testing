package _interface

import (
	"backend/src/internal/model"
	"backend/src/internal/model/dto"
	"context"
)

//go:generate go run github.com/vektra/mockery/v2@v2.42.1 --name=IReserveRepository
type IReserveRepository interface {
	//AddReserveEquipment(ctx context.Context, request *dto.AddReserveEquipmentRequest) error                                  // Добавление в таблицу многие ко многим Бронь-Оборудование
	IsRoomReserve(ctx context.Context, request *dto.IsRoomReserveRequest) (bool, error)                       // Для проверки при изменении, есть ли брони на эту комнату
	IsInstrumentalistReserve(ctx context.Context, request *dto.IsInstrumentalistReserveRequest) (bool, error) // Для проверки при изменении, есть ли брони на этого инструменталиста
	IsProducerReserve(ctx context.Context, request *dto.IsProducerReserveRequest) (bool, error)               // Для проверки при изменении, есть ли брони на этого продюсера
	IsEquipmentReserve(ctx context.Context, request *dto.IsEquipmentReserveRequest) (bool, error)             // Для проверки при изменении, есть ли брони на это оборудование

	GetByRoomId(ctx context.Context, request *dto.GetReserveByRoomIdRequest) ([]*model.Reserve, error)                       // Для бронирования
	GetByInstrumentalistId(ctx context.Context, request *dto.GetReserveByInstrumentalistIdRequest) ([]*model.Reserve, error) // Для бронирования
	GetByProducerId(ctx context.Context, request *dto.GetReserveByProducerIdRequest) ([]*model.Reserve, error)               // Для бронирования
	//GetByEquipmentsId(ctx context.Context, request *dto.GetReserveByEquipmentsIdRequest) ([]*model.Reserve, error)           // Для бронирования

	GetUserReserves(ctx context.Context, request *dto.GetUserReservesRequest) ([]*model.Reserve, error) // Для вывода пользователю его броней
	Add(ctx context.Context, request *dto.AddReserveRequest) error                                      // Для вставки в таблицу
	Delete(ctx context.Context, request *dto.DeleteReserveRequest) error                                // Для удаления из таблицы
	GetAll(ctx context.Context, request *dto.GetAllReserveRequest) ([]*model.Reserve, error)
}
