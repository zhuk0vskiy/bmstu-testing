package _interface

import (
	"backend/src/internal/model"
	"backend/src/internal/model/dto"
	"context"
)

//go:generate go run github.com/vektra/mockery/v2@v2.42.1 --name=IProducerRepository
type IProducerRepository interface {
	Get(ctx context.Context, request *dto.GetProducerRequest) (*model.Producer, error)                   // Для отдельного вывода изначальной информации на странице для отдельного продюсера при обновлении
	GetByStudio(ctx context.Context, request *dto.GetProducerByStudioRequest) ([]*model.Producer, error) // Для изменения продюсеров по студии и посика незаброненных продюсеров
	Add(ctx context.Context, request *dto.AddProducerRequest) error                                      // Для вставки продюсера в таблицу
	Update(ctx context.Context, request *dto.UpdateProducerRequest) error                                // Для изменения продюсера в талблице
	Delete(ctx context.Context, request *dto.DeleteProducerRequest) error                                // Для удаления продюсера из таблицы
}
