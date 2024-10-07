package _interface

import (
	"backend/src/internal/model"
	"backend/src/internal/model/dto"
)

type IStudioService interface {
	Get(request *dto.GetStudioRequest) (*model.Studio, error)         // Для отдельного отображения при изменении студии
	GetAll(request *dto.GetStudioAllRequest) ([]*model.Studio, error) // Для вывода всех студий на начальной странице
	Update(request *dto.UpdateStudioRequest) error                    // Для изменения студии
	Add(request *dto.AddStudioRequest) error                          // Для добавления студии
	Delete(request *dto.DeleteStudioRequest) error                    // Для удаления студии
}
