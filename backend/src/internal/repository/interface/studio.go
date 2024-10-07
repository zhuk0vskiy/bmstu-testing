package _interface

import (
	"backend/src/internal/model"
	"backend/src/internal/model/dto"
	"context"
)

//go:generate go run github.com/vektra/mockery/v2@v2.42.1 --name=IStudioRepository
type IStudioRepository interface {
	Get(ctx context.Context, request *dto.GetStudioRequest) (*model.Studio, error)         // Для отдельного отображения при изменении студии
	GetAll(ctx context.Context, request *dto.GetStudioAllRequest) ([]*model.Studio, error) // Для вывода всех студий на начальной странице
	Update(ctx context.Context, request *dto.UpdateStudioRequest) error                    // Для изменения студии
	Add(ctx context.Context, request *dto.AddStudioRequest) error                          // Для добавления студии
	Delete(ctx context.Context, request *dto.DeleteStudioRequest) error                    // Для удаления студии
}
