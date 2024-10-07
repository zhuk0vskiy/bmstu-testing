package dto

import (
	"backend/src/internal/model"
)

type GetSuitableStuffRequest struct {
	TimeInterval *model.TimeInterval
	StudioId     int64
}

type GetNotReservedRoomsRequest struct {
	ChoosenInterval *model.TimeInterval
	StudioId        int64
}

type GetNotReservedProducersRequest struct {
	ChoosenInterval *model.TimeInterval
	StudioId        int64
}

type GetNotReservedEquipmentsRequest struct {
	ChoosenInterval *model.TimeInterval
	StudioId        int64
}

type GetNotReservedInstrumentalistsRequest struct {
	ChoosenInterval *model.TimeInterval
	StudioId        int64
}

type IsStuffFreeRequest struct {
	RoomId            int64
	ProducerId        int64
	InstrumentalistId int64
	EquipmentId       []int64
}
