package _interface

import (
	"backend/src/internal/model"
	"backend/src/internal/model/dto"
)

type IProducerService interface {
	Get(request *dto.GetProducerRequest) (*model.Producer, error)                   // Для отдельного вывода изначальной информации на странице для отдельного продюсера при обновлении
	GetByStudio(request *dto.GetProducerByStudioRequest) ([]*model.Producer, error) // Для изменения продюсеров по студии
	Add(request *dto.AddProducerRequest) error                                      // Для добавления продюсеров
	Update(request *dto.UpdateProducerRequest) error                                // Для обновления продюсеров
	Delete(request *dto.DeleteProducerRequest) error                                // Для удаления продюсеров
}
