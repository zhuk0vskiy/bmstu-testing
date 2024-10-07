package repository_test

import (
	"backend/src/internal/model/dto"
	"backend/src/internal/repository/interface/mocks"
	"backend/src/internal/service/impl"
	"backend/src/tests/utils"
	"context"
	"fmt"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
)

type RoomRepositorySuite struct {
	suite.Suite

	mockRepo mocks.IRoomRepository
}

func (suite *RoomRepositorySuite) TestRoomGet01(t provider.T) {
	t.Title("[Get] success")
	t.Tags("repository", "room", "get")
	t.Parallel()

	t.WithNewStep("successfully get", func(sCtx provider.StepCtx) {
		ctx := context.Background()

		//request := utils.RoomFabric{Id: 1}.CorrectRoomGet()

		builder := utils.RoomBuilder{}
		room := builder.
			WithId(1).
			WithName("1").
			WithStudioId(1).
			WithStartHour(1).
			WithEndHour(2).
			ToDto()

		//repo := pool.IUserRepository{t}
		suite.mockRepo.On("Get", ctx, &dto.GetRoomRequest{
			Id: room.Id,
		}).Return(
			room, nil,
		).Once()

		sCtx.WithNewParameters("ctx", ctx, "request", room)
		room, err := suite.mockRepo.Get(ctx, &dto.GetRoomRequest{
			Id: room.Id,
		})

		sCtx.Assert().NoError(err)
		sCtx.Assert().NotNil(room)
	})
}

func (suite *RoomRepositorySuite) TestRoomGet02(t provider.T) {
	t.Title("[Get] incorrect id")
	t.Tags("room", "get")
	t.Parallel()

	t.WithNewStep("fail get", func(sCtx provider.StepCtx) {
		ctx := context.Background()

		//request := utils.RoomFabric{Id: 0}.IncorrectRoomGet()

		builder := utils.RoomBuilder{}
		room := builder.
			WithId(1).
			WithName("1").
			WithStudioId(1).
			WithStartHour(1).
			WithEndHour(2).
			ToDto()

		//repo := pool.IUserRepository{t}
		suite.mockRepo.On("Get", ctx, &dto.GetRoomRequest{
			Id: room.Id,
		}).Return(
			nil, fmt.Errorf("invalid id"),
		).Once()

		sCtx.WithNewParameters("ctx", ctx, "request", room)
		room, err := suite.mockRepo.Get(ctx, &dto.GetRoomRequest{
			Id: room.Id,
		})

		sCtx.Assert().Error(err)
		sCtx.Assert().Nil(room)
	})
}

func (suite *RoomRepositorySuite) TestRoomDelete01(t provider.T) {
	t.Title("[Delete] success")
	t.Tags("room", "delete")
	t.Parallel()

	t.WithNewStep("successfully delete", func(sCtx provider.StepCtx) {
		ctx := context.Background()

		request := utils.RoomFabric{Id: 1}.CorrectRoomDelete()

		//repo := pool.IUserRepository{t}
		roomRepo := new(mocks.IRoomRepository)
		reserveRepo := new(mocks.IReserveRepository)

		roomRepo.On("Delete", ctx, request).Return(
			nil,
		).Once()

		service := impl.NewRoomService(roomRepo, reserveRepo)
		sCtx.WithNewParameters("ctx", ctx, "request", request)

		err := service.Delete(request)

		sCtx.Assert().NoError(err)
	})
}

func (suite *RoomRepositorySuite) TestRoomDelete02(t provider.T) {
	t.Title("[Delete] failed")
	t.Tags("room", "delete")
	t.Parallel()

	t.WithNewStep("failed to delete", func(sCtx provider.StepCtx) {
		ctx := context.Background()

		request := utils.RoomFabric{Id: 0}.IncorrectRoomDelete()

		//repo := pool.IUserRepository{t}
		roomRepo := new(mocks.IRoomRepository)
		reserveRepo := new(mocks.IReserveRepository)

		roomRepo.On("Delete", ctx, request).Return(
			fmt.Errorf("no room with this id"),
		).Once()

		service := impl.NewRoomService(roomRepo, reserveRepo)
		sCtx.WithNewParameters("ctx", ctx, "request", request)

		err := service.Delete(request)

		sCtx.Assert().Error(err)
	})
}

func (suite *RoomRepositorySuite) TestRoomAdd01(t provider.T) {
	t.Title("[Delete] failed")
	t.Tags("room", "delete")
	t.Parallel()

	t.WithNewStep("failed to delete", func(sCtx provider.StepCtx) {
		ctx := context.Background()

		request := utils.RoomFabric{Id: 1}.CorrectRoomAdd()

		//repo := pool.IUserRepository{t}
		roomRepo := new(mocks.IRoomRepository)
		reserveRepo := new(mocks.IReserveRepository)

		roomRepo.On("Add", ctx, request).Return(
			nil,
		).Once()

		service := impl.NewRoomService(roomRepo, reserveRepo)
		sCtx.WithNewParameters("ctx", ctx, "request", request)

		err := service.Add(request)

		sCtx.Assert().NoError(err)
	})
}

func (suite *RoomRepositorySuite) TestRoomAdd02(t provider.T) {
	t.Title("[Delete] failed")
	t.Tags("room", "delete")
	t.Parallel()

	t.WithNewStep("failed to delete", func(sCtx provider.StepCtx) {
		ctx := context.Background()

		request := utils.RoomFabric{Id: 1}.IncorrectRoomAdd()

		//repo := pool.IUserRepository{t}
		roomRepo := new(mocks.IRoomRepository)
		reserveRepo := new(mocks.IReserveRepository)

		roomRepo.On("Add", ctx, request).Return(
			nil,
		).Once()

		service := impl.NewRoomService(roomRepo, reserveRepo)
		sCtx.WithNewParameters("ctx", ctx, "request", request)

		err := service.Add(request)

		sCtx.Assert().Error(err)
	})
}
