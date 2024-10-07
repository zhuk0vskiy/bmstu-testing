package dto

import "backend/src/internal/model"

type EquipmentAndTime struct {
	Equipment    *model.Equipment
	TimeInterval *model.TimeInterval
}

type GetEquipmentRequest struct {
	Id int64
}

type GetEquipmentByStudioRequest struct {
	StudioId int64
}

type AddEquipmentRequest struct {
	Name     string
	Type     int64
	StudioId int64
}

type UpdateEquipmentRequest struct {
	Id       int64
	Name     string
	Type     int64
	StudioId int64
}

type DeleteEquipmentRequest struct {
	Id int64
}

type GetEquipmentFullTimeFreeByStudioAndTypeRequest struct {
	StudioId int64
	Type     int64
}

type GetEquipmentNotFullTimeFreeByStudioAndTypeRequest struct {
	StudioId     int64
	Type         int64
	TimeInterval *model.TimeInterval
}

type GetEquipmentNotReservedIdByStudioAndType struct {
	StudioId int64
	Type     int64
}

type GetEquipmentByReserveRequest struct {
	ReserveId int64
}
