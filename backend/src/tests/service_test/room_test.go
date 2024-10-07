package service_test

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

type RoomSuite struct {
	suite.Suite
}

func (suite *RoomSuite) TestRoomGet01(t provider.T) {
	t.Title("[Get] success")
	t.Tags("room", "get")
	t.Parallel()

	t.WithNewStep("successfully get", func(sCtx provider.StepCtx) {
		ctx := context.Background()

		builder := utils.RoomBuilder{}
		roomAdd := builder.
			WithId(1).
			WithName("1").
			WithStudioId(1).
			ToDto()

		//repo := pool.IUserRepository{t}
		roomRepo := new(mocks.IRoomRepository)
		reserveRepo := new(mocks.IReserveRepository)

		roomRepo.On("Get", ctx, &dto.GetRoomRequest{
			Id: roomAdd.Id,
		}).Return(
			roomAdd, nil,
		)

		service := impl.NewRoomService(roomRepo, reserveRepo)
		sCtx.WithNewParameters("ctx", ctx, "request", roomAdd)

		room, err := service.Get(&dto.GetRoomRequest{
			Id: roomAdd.Id,
		})

		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(room, roomAdd)
	})
}

func (suite *RoomSuite) TestRoomGet02(t provider.T) {
	t.Title("[Get] failed")
	t.Tags("room", "get")
	t.Parallel()

	t.WithNewStep("fail to get", func(sCtx provider.StepCtx) {
		ctx := context.Background()

		builder := utils.RoomBuilder{}
		roomAdd := builder.
			WithId(1).
			WithName("2").
			WithStudioId(1).
			WithStartHour(1).
			WithEndHour(2).
			ToDto()

		//repo := pool.IUserRepository{t}
		roomRepo := new(mocks.IRoomRepository)
		reserveRepo := new(mocks.IReserveRepository)

		roomRepo.On("Get", ctx, &dto.GetRoomRequest{
			Id: roomAdd.Id,
		}).Return(
			nil, fmt.Errorf("no room with this id"),
		).Once()

		service := impl.NewRoomService(roomRepo, reserveRepo)
		sCtx.WithNewParameters("ctx", ctx, "request", roomAdd)

		room, err := service.Get(&dto.GetRoomRequest{
			Id: roomAdd.Id,
		})

		sCtx.Assert().Error(err)
		sCtx.Assert().Nil(room)
	})
}

func (suite *RoomSuite) TestRoomDelete01(t provider.T) {
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

func (suite *RoomSuite) TestRoomDelete02(t provider.T) {
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

func (suite *RoomSuite) TestRoomAdd01(t provider.T) {
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

func (suite *RoomSuite) TestRoomAdd02(t provider.T) {
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
