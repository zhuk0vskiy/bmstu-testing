package _interface

import (
	"backend/src/internal/model"
	"backend/src/internal/model/dto"
	"context"
)

//go:generate go run github.com/vektra/mockery/v2@v2.42.1 --name=IEquipmentRepository
type IEquipmentRepository interface {
	Get(ctx context.Context, request *dto.GetEquipmentRequest) (*model.Equipment, error)                                                         // Для отдельного вывода изначальной информации на странице для отдельного оборудования при обновлении
	GetByStudio(ctx context.Context, request *dto.GetEquipmentByStudioRequest) ([]*model.Equipment, error)                                       // Для изменения оборудования по студиям
	GetFullTimeFreeByStudioAndType(ctx context.Context, request *dto.GetEquipmentFullTimeFreeByStudioAndTypeRequest) ([]*model.Equipment, error) // Для поиска всего оборудования по id студии и типу в бронировании
	GetNotFullTimeFreeByStudioAndType(ctx context.Context, request *dto.GetEquipmentNotFullTimeFreeByStudioAndTypeRequest) ([]*dto.EquipmentAndTime, error)

	Add(ctx context.Context, request *dto.AddEquipmentRequest) error       // Для добавления оборудования в таблицу
	Update(ctx context.Context, request *dto.UpdateEquipmentRequest) error // Для изменения оборудования в таблице
	Delete(ctx context.Context, request *dto.DeleteEquipmentRequest) error // Для удаления оборудования из таблицы

	GetByReserve(ctx context.Context, request *dto.GetEquipmentByReserveRequest) ([]*model.Equipment, error) // Для отдельного вывода изначальной информации на странице для отдельного оборудования при обновлении
	//dto.getby
	//GetReservedIdByStudioAndType(ctx context.Context, request *dto.GetEquipmentReservedIdByStudioAndType) ([]*int64, error)              // Для вычисления недоступного оборудования в данной студии
	//GetAll() ([]*Equipment, error) //TODO: заменить на GetByStudio
	//GetByType(id int64) ([]*Equipment, error)
}
