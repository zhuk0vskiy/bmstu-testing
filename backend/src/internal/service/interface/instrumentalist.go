package _interface

import (
	"backend/src/internal/model"
	"backend/src/internal/model/dto"
	"context"
)

type IInstrumentalistService interface {
	Get(request *dto.GetInstrumentalistRequest) (*model.Instrumentalist, error)                   // Для отдельного вывода изначальной информации на странице для отдельного инструменталиста при обновлении
	GetByStudio(request *dto.GetInstrumentalistByStudioRequest) ([]*model.Instrumentalist, error) // Для изменения инструменталистов по студии
	Add(ctx context.Context, request *dto.AddInstrumentalistRequest) error                        // Для добавления инструменталистов
	Update(request *dto.UpdateInstrumentalistRequest) error                                       // Для обновления инструменталистов
	Delete(request *dto.DeleteInstrumentalistRequest) error                                       // Для удаления инструменталистов
}
