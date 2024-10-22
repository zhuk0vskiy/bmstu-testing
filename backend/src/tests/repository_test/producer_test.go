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

type ProducerRepositorySuite struct {
	suite.Suite

	mockRepo mocks.IProducerRepository
}

func (suite *ProducerRepositorySuite) TestProducerGet01(t provider.T) {
	t.Title("[Get] success")
	t.Tags("repository", "producer", "get")
	t.Parallel()

	t.WithNewStep("successfully get", func(sCtx provider.StepCtx) {
		ctx := context.Background()

		//request := utils.ProducerFabric{Id: 1}.CorrectProducerGet()

		builder := utils.ProducerBuilder{}
		producer := builder.
			WithId(1).
			WithName("1").
			WithStudioId(1).
			WithStartHour(1).
			WithEndHour(2).
			ToDto()

		//repo := pool.IUserRepository{t}
		suite.mockRepo.On("Get", ctx, &dto.GetProducerRequest{
			Id: producer.Id,
		}).Return(
			producer, nil,
		).Once()

		sCtx.WithNewParameters("ctx", ctx, "request", producer)
		producer, err := suite.mockRepo.Get(ctx, &dto.GetProducerRequest{
			Id: producer.Id,
		})

		sCtx.Assert().NoError(err)
		sCtx.Assert().NotNil(producer)
	})
}

func (suite *ProducerRepositorySuite) TestProducerGet02(t provider.T) {
	t.Title("[Get] incorrect id")
	t.Tags("producer", "get")
	t.Parallel()

	t.WithNewStep("fail get", func(sCtx provider.StepCtx) {
		ctx := context.Background()

		//request := utils.ProducerFabric{Id: 0}.IncorrectProducerGet()

		builder := utils.ProducerBuilder{}
		producer := builder.
			WithId(1).
			WithName("1").
			WithStudioId(1).
			WithStartHour(1).
			WithEndHour(2).
			ToDto()

		//repo := pool.IUserRepository{t}
		suite.mockRepo.On("Get", ctx, &dto.GetProducerRequest{
			Id: producer.Id,
		}).Return(
			nil, fmt.Errorf("invalid id"),
		).Once()

		sCtx.WithNewParameters("ctx", ctx, "request", producer)
		producer, err := suite.mockRepo.Get(ctx, &dto.GetProducerRequest{
			Id: producer.Id,
		})

		sCtx.Assert().Error(err)
		sCtx.Assert().Nil(producer)
	})
}

func (suite *ProducerRepositorySuite) TestProducerDelete01(t provider.T) {
	t.Title("[Delete] success")
	t.Tags("producer", "delete")
	t.Parallel()

	t.WithNewStep("successfully delete", func(sCtx provider.StepCtx) {
		ctx := context.Background()

		request := utils.ProducerFabric{Id: 1}.CorrectProducerDelete()

		//repo := pool.IUserRepository{t}
		producerRepo := new(mocks.IProducerRepository)
		reserveRepo := new(mocks.IReserveRepository)

		producerRepo.On("Delete", ctx, request).Return(
			nil,
		).Once()

		logger := utils.NewMockLogger()

		service := impl.NewProducerService(logger, producerRepo, reserveRepo)
		sCtx.WithNewParameters("ctx", ctx, "request", request)

		err := service.Delete(request)

		sCtx.Assert().NoError(err)
	})
}

func (suite *ProducerRepositorySuite) TestProducerDelete02(t provider.T) {
	t.Title("[Delete] failed")
	t.Tags("producer", "delete")
	t.Parallel()

	t.WithNewStep("failed to delete", func(sCtx provider.StepCtx) {
		ctx := context.Background()

		request := utils.ProducerFabric{Id: 0}.IncorrectProducerDelete()

		//repo := pool.IUserRepository{t}
		producerRepo := new(mocks.IProducerRepository)
		reserveRepo := new(mocks.IReserveRepository)

		producerRepo.On("Delete", ctx, request).Return(
			fmt.Errorf("no producer with this id"),
		).Once()

		logger := utils.NewMockLogger()

		service := impl.NewProducerService(logger, producerRepo, reserveRepo)
		sCtx.WithNewParameters("ctx", ctx, "request", request)

		err := service.Delete(request)

		sCtx.Assert().Error(err)
	})
}

func (suite *ProducerRepositorySuite) TestProducerAdd01(t provider.T) {
	t.Title("[Delete] failed")
	t.Tags("producer", "delete")
	t.Parallel()

	t.WithNewStep("failed to delete", func(sCtx provider.StepCtx) {
		ctx := context.Background()

		request := utils.ProducerFabric{Id: 1}.CorrectProducerAdd()

		//repo := pool.IUserRepository{t}
		producerRepo := new(mocks.IProducerRepository)
		reserveRepo := new(mocks.IReserveRepository)

		producerRepo.On("Add", ctx, request).Return(
			nil,
		).Once()

		logger := utils.NewMockLogger()

		service := impl.NewProducerService(logger, producerRepo, reserveRepo)
		sCtx.WithNewParameters("ctx", ctx, "request", request)

		err := service.Add(ctx, request)

		sCtx.Assert().NoError(err)
	})
}

func (suite *ProducerRepositorySuite) TestProducerAdd02(t provider.T) {
	t.Title("[Delete] failed")
	t.Tags("producer", "delete")
	t.Parallel()

	t.WithNewStep("failed to delete", func(sCtx provider.StepCtx) {
		ctx := context.Background()

		request := utils.ProducerFabric{Id: 1}.IncorrectProducerAdd()

		//repo := pool.IUserRepository{t}
		producerRepo := new(mocks.IProducerRepository)
		reserveRepo := new(mocks.IReserveRepository)

		producerRepo.On("Add", ctx, request).Return(
			nil,
		).Once()

		logger := utils.NewMockLogger()

		service := impl.NewProducerService(logger, producerRepo, reserveRepo)
		sCtx.WithNewParameters("ctx", ctx, "request", request)

		err := service.Add(ctx, request)

		sCtx.Assert().Error(err)
	})
}
