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

type InstrumentalistRepositorySuite struct {
	suite.Suite

	mockRepo mocks.IInstrumentalistRepository
}

func (suite *InstrumentalistRepositorySuite) TestInstrumentalistGet01(t provider.T) {
	t.Title("[Get] success")
	t.Tags("repository", "instrumentalist", "get")
	t.Parallel()

	t.WithNewStep("successfully get", func(sCtx provider.StepCtx) {
		ctx := context.Background()

		//request := utils.InstrumentalistFabric{Id: 1}.CorrectInstrumentalistGet()

		builder := utils.InstrumentalistBuilder{}
		instrumentalist := builder.
			WithId(1).
			WithName("1").
			WithStudioId(1).
			WithStartHour(1).
			WithEndHour(2).
			ToDto()

		//repo := pool.IUserRepository{t}
		suite.mockRepo.On("Get", ctx, &dto.GetInstrumentalistRequest{
			Id: instrumentalist.Id,
		}).Return(
			instrumentalist, nil,
		).Once()

		sCtx.WithNewParameters("ctx", ctx, "request", instrumentalist)
		instrumentalist, err := suite.mockRepo.Get(ctx, &dto.GetInstrumentalistRequest{
			Id: instrumentalist.Id,
		})

		sCtx.Assert().NoError(err)
		sCtx.Assert().NotNil(instrumentalist)
	})
}

func (suite *InstrumentalistRepositorySuite) TestInstrumentalistGet02(t provider.T) {
	t.Title("[Get] incorrect id")
	t.Tags("instrumentalist", "get")
	t.Parallel()

	t.WithNewStep("fail get", func(sCtx provider.StepCtx) {
		ctx := context.Background()

		//request := utils.InstrumentalistFabric{Id: 0}.IncorrectInstrumentalistGet()

		builder := utils.InstrumentalistBuilder{}
		instrumentalist := builder.
			WithId(1).
			WithName("1").
			WithStudioId(1).
			WithStartHour(1).
			WithEndHour(2).
			ToDto()

		//repo := pool.IUserRepository{t}
		suite.mockRepo.On("Get", ctx, &dto.GetInstrumentalistRequest{
			Id: instrumentalist.Id,
		}).Return(
			nil, fmt.Errorf("invalid id"),
		).Once()

		sCtx.WithNewParameters("ctx", ctx, "request", instrumentalist)
		instrumentalist, err := suite.mockRepo.Get(ctx, &dto.GetInstrumentalistRequest{
			Id: instrumentalist.Id,
		})

		sCtx.Assert().Error(err)
		sCtx.Assert().Nil(instrumentalist)
	})
}

func (suite *InstrumentalistRepositorySuite) TestInstrumentalistDelete01(t provider.T) {
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

func (suite *InstrumentalistRepositorySuite) TestInstrumentalistDelete02(t provider.T) {
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

func (suite *InstrumentalistRepositorySuite) TestInstrumentalistAdd01(t provider.T) {
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

		err := service.Add(ctx, request)

		sCtx.Assert().NoError(err)
	})
}

func (suite *InstrumentalistRepositorySuite) TestInstrumentalistAdd02(t provider.T) {
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

		err := service.Add(ctx, request)

		sCtx.Assert().Error(err)
	})
}
