package service_test

import (
	"backend/src/internal/model"
	"backend/src/internal/model/dto"
	"backend/src/internal/repository/interface/mocks"
	"backend/src/internal/service/impl"
	"backend/src/tests/utils"
	"context"
	"fmt"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
)

type InstrumentalistSuite struct {
	suite.Suite
}

func (suite *InstrumentalistSuite) TestInstrumentalistGet01(t provider.T) {
	t.Title("[Get] success")
	t.Tags("instrumentalist", "get")
	t.Parallel()

	t.WithNewStep("successfully get", func(sCtx provider.StepCtx) {
		ctx := context.Background()

		builder := utils.InstrumentalistBuilder{}
		instrumentalistAdd := builder.
			WithId(1).
			WithName("1").
			WithStudioId(1).
			WithStartHour(1).
			WithEndHour(2).
			ToDto()

		//repo := pool.IUserRepository{t}
		instrumentalistRepo := new(mocks.IInstrumentalistRepository)
		reserveRepo := new(mocks.IReserveRepository)

		instrumentalist := &model.Instrumentalist{
			Id:        instrumentalistAdd.Id,
			Name:      instrumentalistAdd.Name,
			StudioId:  instrumentalistAdd.StudioId,
			StartHour: instrumentalistAdd.StartHour,
			EndHour:   instrumentalistAdd.EndHour,
		}

		instrumentalistRepo.On("Get", ctx, &dto.GetInstrumentalistRequest{
			Id: instrumentalistAdd.Id,
		}).Return(
			instrumentalist, nil,
		).Once()

		logger := utils.NewMockLogger()

		service := impl.NewInstrumentalistService(logger, instrumentalistRepo, reserveRepo)
		sCtx.WithNewParameters("ctx", ctx, "request", instrumentalistAdd)

		instrumentalistRes, err := service.Get(&dto.GetInstrumentalistRequest{
			Id: instrumentalistAdd.Id,
		})

		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(instrumentalist, instrumentalistRes)
	})
}

func (suite *InstrumentalistSuite) TestInstrumentalistGet02(t provider.T) {
	t.Title("[Get] failed")
	t.Tags("instrumentalist", "get")
	t.Parallel()

	t.WithNewStep("fail to get", func(sCtx provider.StepCtx) {
		ctx := context.Background()

		builder := utils.InstrumentalistBuilder{}
		instrumentalist := builder.
			WithId(1).
			WithName("").
			WithStudioId(1).
			WithStartHour(1).
			WithEndHour(2).
			ToDto()

		//repo := pool.IUserRepository{t}
		instrumentalistRepo := new(mocks.IInstrumentalistRepository)
		reserveRepo := new(mocks.IReserveRepository)

		instrumentalistRepo.On("Get", ctx, &dto.GetInstrumentalistRequest{
			Id: instrumentalist.Id,
		}).Return(
			nil, fmt.Errorf("no instrumentalist with this id"),
		).Once()

		logger := utils.NewMockLogger()

		service := impl.NewInstrumentalistService(logger, instrumentalistRepo, reserveRepo)
		sCtx.WithNewParameters("ctx", ctx, "request", instrumentalist)

		instrumentalist, err := service.Get(&dto.GetInstrumentalistRequest{
			Id: instrumentalist.Id,
		})

		sCtx.Assert().Error(err)
		sCtx.Assert().Nil(instrumentalist)
	})
}

func (suite *InstrumentalistSuite) TestInstrumentalistDelete01(t provider.T) {
	t.Title("[Delete] success")
	t.Tags("instrumentalist", "delete")
	t.Parallel()

	t.WithNewStep("successfully delete", func(sCtx provider.StepCtx) {
		ctx := context.Background()

		request := utils.InstrumentalistFabric{Id: 1}.CorrectInstrumentalistDelete()

		//repo := pool.IUserRepository{t}
		instrumentalistRepo := new(mocks.IInstrumentalistRepository)
		reserveRepo := new(mocks.IReserveRepository)

		instrumentalistRepo.On("Delete", ctx, request).Return(
			nil,
		).Once()

		logger := utils.NewMockLogger()

		service := impl.NewInstrumentalistService(logger, instrumentalistRepo, reserveRepo)
		sCtx.WithNewParameters("ctx", ctx, "request", request)

		err := service.Delete(request)

		sCtx.Assert().NoError(err)
	})
}

func (suite *InstrumentalistSuite) TestInstrumentalistDelete02(t provider.T) {
	t.Title("[Delete] failed")
	t.Tags("instrumentalist", "delete")
	t.Parallel()

	t.WithNewStep("failed to delete", func(sCtx provider.StepCtx) {
		ctx := context.Background()

		request := utils.InstrumentalistFabric{Id: 0}.IncorrectInstrumentalistDelete()

		//repo := pool.IUserRepository{t}
		instrumentalistRepo := new(mocks.IInstrumentalistRepository)
		reserveRepo := new(mocks.IReserveRepository)

		instrumentalistRepo.On("Delete", ctx, request).Return(
			fmt.Errorf("no instrumentalist with this id"),
		).Once()

		logger := utils.NewMockLogger()

		service := impl.NewInstrumentalistService(logger, instrumentalistRepo, reserveRepo)
		sCtx.WithNewParameters("ctx", ctx, "request", request)

		err := service.Delete(request)

		sCtx.Assert().Error(err)
	})
}

func (suite *InstrumentalistSuite) TestInstrumentalistAdd01(t provider.T) {
	t.Title("[Delete] failed")
	t.Tags("instrumentalist", "delete")
	t.Parallel()

	t.WithNewStep("failed to delete", func(sCtx provider.StepCtx) {
		ctx := context.Background()

		request := utils.InstrumentalistFabric{Id: 1}.CorrectInstrumentalistAdd()

		//repo := pool.IUserRepository{t}
		instrumentalistRepo := new(mocks.IInstrumentalistRepository)
		reserveRepo := new(mocks.IReserveRepository)

		instrumentalistRepo.On("Add", ctx, request).Return(
			nil,
		).Once()

		logger := utils.NewMockLogger()

		service := impl.NewInstrumentalistService(logger, instrumentalistRepo, reserveRepo)
		sCtx.WithNewParameters("ctx", ctx, "request", request)

		err := service.Add(request)

		sCtx.Assert().NoError(err)
	})
}

func (suite *InstrumentalistSuite) TestInstrumentalistAdd02(t provider.T) {
	t.Title("[Delete] failed")
	t.Tags("instrumentalist", "delete")
	t.Parallel()

	t.WithNewStep("failed to delete", func(sCtx provider.StepCtx) {
		ctx := context.Background()

		request := utils.InstrumentalistFabric{Id: 1}.IncorrectInstrumentalistAdd()

		//repo := pool.IUserRepository{t}
		instrumentalistRepo := new(mocks.IInstrumentalistRepository)
		reserveRepo := new(mocks.IReserveRepository)

		instrumentalistRepo.On("Add", ctx, request).Return(
			nil,
		).Once()

		logger := utils.NewMockLogger()

		service := impl.NewInstrumentalistService(logger, instrumentalistRepo, reserveRepo)
		sCtx.WithNewParameters("ctx", ctx, "request", request)

		err := service.Add(request)

		sCtx.Assert().Error(err)
	})
}
