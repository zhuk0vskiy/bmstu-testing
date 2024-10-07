package _interface

import (
	"backend/src/internal/model"
	"backend/src/internal/model/dto"
	"context"
)

//go:generate go run github.com/vektra/mockery/v2@v2.42.1 --name=IInstrumentalistRepository
type IInstrumentalistRepository interface {
	Get(ctx context.Context, request *dto.GetInstrumentalistRequest) (*model.Instrumentalist, error)                   // Для отдельного вывода изначальной информации на странице для отдельного инструменталистаы при обновлении
	GetByStudio(ctx context.Context, request *dto.GetInstrumentalistByStudioRequest) ([]*model.Instrumentalist, error) // Для изменения инструменталистов по студии и поиска незаброненных
	Add(ctx context.Context, request *dto.AddInstrumentalistRequest) error                                             // Для вставки инструменталиста в таблицу
	Update(ctx context.Context, request *dto.UpdateInstrumentalistRequest) error                                       // Для изменения инструменталиста в талблице
	Delete(ctx context.Context, request *dto.DeleteInstrumentalistRequest) error                                       // Для удаления инструменталистов из таблицы
}
