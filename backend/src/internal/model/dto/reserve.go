package dto

import "backend/src/internal/model"

type AddReserveRequest struct {
	UserId            int64
	RoomId            int64
	ProducerId        int64
	InstrumentalistId int64
	TimeInterval      *model.TimeInterval
	EquipmentId       []int64
}

type DeleteReserveRequest struct {
	Id int64
}

type AddReserveEquipmentRequest struct {
	ReserveId   int64
	EquipmentId int64
}

type IsRoomReserveRequest struct {
	RoomId int64
}

type IsInstrumentalistReserveRequest struct {
	InstrumentalistId int64
}

type IsProducerReserveRequest struct {
	ProducerId int64
}

type IsEquipmentReserveRequest struct {
	EquipmentId int64
}

type GetReserveByRoomIdRequest struct {
	RoomId int64
}

type GetReserveByInstrumentalistIdRequest struct {
	InstrumentalistId int64
}

type GetReserveByProducerIdRequest struct {
	ProducerId int64
}

type GetReserveByEquipmentsIdRequest struct {
	EquipmentId int64
}

type GetAllReserveRequest struct {
}
