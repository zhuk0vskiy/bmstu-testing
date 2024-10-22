package impl

import (
	"backend/src/internal/model"
	"backend/src/internal/model/dto"
	repositoryInterface "backend/src/internal/repository/interface"
	serviceInterface "backend/src/internal/service/interface"
	"backend/src/pkg/base"
	"backend/src/pkg/logger"
	"context"
	"fmt"
)

type UserService struct {
	userRepo    repositoryInterface.IUserRepository
	reserveRepo repositoryInterface.IReserveRepository
	crypto      base.IHashCrypto
	logger      logger.Interface
}

func NewUserService(
	logger logger.Interface,
	userRepo repositoryInterface.IUserRepository,
	reserveRepo repositoryInterface.IReserveRepository,
	crypto base.IHashCrypto,
) serviceInterface.IUserService {
	return &UserService{
		logger:      logger,
		userRepo:    userRepo,
		reserveRepo: reserveRepo,
		crypto:      crypto,
	}
}

func (s UserService) GetReserves(request *dto.GetUserReservesRequest) (reserves []*model.Reserve, err error) {
	if request.Id < 1 {
		s.logger.Infof("ошибка get user %d reserves: %s", request.Id, fmt.Errorf("неверный id: %w", err))
		return nil, fmt.Errorf("неверный id: %w", err)
	}

	//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//defer cancel()
	ctx := context.Background()

	reserves, err = s.reserveRepo.GetUserReserves(ctx, &dto.GetUserReservesRequest{
		Id: request.Id,
	})

	//var reserveExt []model.ReserveExt
	//
	//for reserve := range reserves {
	//	equipmnetsId, err = s.
	//}

	if err != nil {
		s.logger.Errorf("ошибка get user %d reserves: %s", request.Id, fmt.Errorf("получение всех броней: %w", err))
		return nil, fmt.Errorf("получение всех броней: %w", err)
	}

	s.logger.Infof("пользователь %d вывел все брони", request.Id)

	return reserves, nil
}

func (s UserService) Get(request *dto.GetUserRequest) (user *model.User, err error) {
	if request.Id < 1 {
		s.logger.Infof("ошибка get user %d: %s", request.Id, fmt.Errorf("ошибка id < 1: %w", err))
		return nil, fmt.Errorf("ошибка id < 1: %w", err)
	}

	//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//defer cancel()
	ctx := context.Background()
	user, err = s.userRepo.Get(ctx, &dto.GetUserRequest{
		Id: request.Id,
	})
	if err != nil {
		s.logger.Infof("ошибка get user %d: %s", request.Id, fmt.Errorf("получение пользователя по id: %w", err))
		return nil, fmt.Errorf("получение пользователя по id: %w", err)
	}

	return user, err
}

func (s UserService) GetByLogin(request *dto.GetUserByLoginRequest) (user *model.User, err error) {
	if request.Login == "" {
		s.logger.Infof("ошибка get user by login %s: %s", request.Login, fmt.Errorf("ошибка логин пустой: %w", err))
		return nil, fmt.Errorf("ошибка логин пустой: %w", err)
	}

	//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//defer cancel()
	ctx := context.Background()
	user, err = s.userRepo.GetByLogin(ctx, &dto.GetUserByLoginRequest{
		Login: request.Login,
	})
	if err != nil {
		s.logger.Errorf("ошибка get user by login %s: %s", request.Login, fmt.Errorf("получение пользователя по логину: %w", err))
		return nil, fmt.Errorf("получение пользователя по логину: %w", err)
	}

	return user, err
}

func (s UserService) Update(request *dto.UpdateUserRequest) (err error) {

	if request.Id < 1 {
		s.logger.Infof("ошибка update user %d %s %s %s %s %s: %s",
			request.Id,
			request.Login,
			request.Password,
			request.FirstName,
			request.SecondName,
			request.ThirdName,
			fmt.Errorf("получение пользователя по логину: %w", err))
		return fmt.Errorf("id < 1: %w", err)
	}

	if request.Login == "" {
		s.logger.Infof("ошибка update user %d %s %s %s %s %s: %s",
			request.Id,
			request.Login,
			request.Password,
			request.FirstName,
			request.SecondName,
			request.ThirdName,
			fmt.Errorf("получение пользователя по логину: %w", err))
		return fmt.Errorf("пустой логин: %w", err)
	}

	if request.Password == "" {
		s.logger.Infof("ошибка update user %d %s %s %s %s %s: %s",
			request.Id,
			request.Login,
			request.Password,
			request.FirstName,
			request.SecondName,
			request.ThirdName,
			fmt.Errorf("пустой пароль: %w", err))
		return fmt.Errorf("пустой пароль: %w", err)
	}

	if request.FirstName == "" {
		s.logger.Infof("ошибка update user %d %s %s %s %s %s: %s",
			request.Id,
			request.Login,
			request.Password,
			request.FirstName,
			request.SecondName,
			request.ThirdName,
			fmt.Errorf("пустое имя %w", err))
		return fmt.Errorf("пустое имя %w", err)
	}

	if request.SecondName == "" {
		s.logger.Infof("ошибка update user %d %s %s %s %s %s: %s",
			request.Id,
			request.Login,
			request.Password,
			request.FirstName,
			request.SecondName,
			request.ThirdName,
			fmt.Errorf("пустая фамилия %w", err))
		return fmt.Errorf("пустая фамилия %w", err)
	}

	if request.ThirdName == "" {
		s.logger.Infof("ошибка update user %d %s %s %s %s %s: %s",
			request.Id,
			request.Login,
			request.Password,
			request.FirstName,
			request.SecondName,
			request.ThirdName,
			fmt.Errorf("пустое отчество %w", err))
		return fmt.Errorf("пустое отчество %w", err)
	}

	hashedPassword, err := s.crypto.GenerateHashPass(request.Password)

	if err != nil {
		s.logger.Errorf("ошибка update user %d %s %s %s %s %s: %s",
			request.Id,
			request.Login,
			request.Password,
			request.FirstName,
			request.SecondName,
			request.ThirdName,
			fmt.Errorf("генерация хэша: %w", err))
		return fmt.Errorf("генерация хэша: %w", err)
	}

	//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//defer cancel()
	ctx := context.Background()
	err = s.userRepo.Update(ctx, &dto.UpdateUserRequest{
		Id:         request.Id,
		Login:      request.Login,
		Password:   hashedPassword,
		FirstName:  request.FirstName,
		SecondName: request.SecondName,
		ThirdName:  request.ThirdName,
	})
	if err != nil {
		s.logger.Errorf("ошибка update user %d %s %s %s %s %s: %s",
			request.Id,
			request.Login,
			request.Password,
			request.FirstName,
			request.SecondName,
			request.ThirdName,
			fmt.Errorf("обновление пользователя %w", err))
		return fmt.Errorf("обновление пользователя %w", err)
	}

	return err
}

func (s UserService) Delete(request *dto.DeleteUserRequest) (err error) {
	if request.Id < 1 {
		s.logger.Infof("ошибка delete user %d: %s", request.Id, fmt.Errorf("неверный id: %w", err))
		return fmt.Errorf("неверный id: %w", err)
	}

	//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//defer cancel()
	ctx := context.Background()
	err = s.userRepo.Delete(ctx, &dto.DeleteUserRequest{
		Id: request.Id,
	})
	if err != nil {
		s.logger.Errorf("ошибка delete user %d: %s", request.Id, fmt.Errorf("удаление пользователя: %w", err))
		return fmt.Errorf("удаление пользователя: %w", err)
	}

	return err
}
