package _interface

import (
	"backend/src/internal/model"
	"backend/src/internal/model/dto"
)

type IEquipmentService interface {
	Get(request *dto.GetEquipmentRequest) (*model.Equipment, error)                   // Для отдельного вывода изначальной информации на странице для отдельного оборудования при изменении
	GetByStudio(request *dto.GetEquipmentByStudioRequest) ([]*model.Equipment, error) // Для изменения оборудования по студиям
	Add(request *dto.AddEquipmentRequest) error                                       // Для добавления оборудования
	Update(request *dto.UpdateEquipmentRequest) error                                 // Для изменения оборудования
	Delete(request *dto.DeleteEquipmentRequest) error                                 // Для удаления оборудования
	GetByReserve(request *dto.GetEquipmentByReserveRequest) ([]*model.Equipment, error)
	// GetAll() ([]*Equipment, error)
	//GetByType(id int64) ([]*Equipment, error)
}
