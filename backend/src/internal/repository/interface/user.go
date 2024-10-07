package _interface

import (
	"backend/src/internal/model"
	"backend/src/internal/model/dto"
	"context"
)

//go:generate go run github.com/vektra/mockery/v2@v2.42.1 --name=IUserRepository
type IUserRepository interface {
	Get(ctx context.Context, request *dto.GetUserRequest) (*model.User, error)               // Для получения данных пользователя
	GetByLogin(ctx context.Context, request *dto.GetUserByLoginRequest) (*model.User, error) // Для поиска при авторизации пользователя
	Add(ctx context.Context, request *dto.AddUserRequest) error                              // Для добаления нового пользователя в таблицу
	Update(ctx context.Context, request *dto.UpdateUserRequest) error                        // Для обновления пользователя в таблице
	Delete(ctx context.Context, request *dto.DeleteUserRequest) error                        // Для удаления пользователя из таблицы
	//GetReserves(ctx context.Context, request *dto.GetUserReservesRequest) ([]*Reserve, error) // Для вывода пользователей его броней
}
