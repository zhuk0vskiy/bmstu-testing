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

type ProducerSuite struct {
	suite.Suite
}

func (suite *ProducerSuite) TestProducerGet01(t provider.T) {
	t.Title("[Get] success")
	t.Tags("producer", "get")
	t.Parallel()

	t.WithNewStep("successfully get", func(sCtx provider.StepCtx) {
		ctx := context.Background()
		builder := utils.ProducerBuilder{}
		producerGet := builder.
			WithId(1).
			WithName("1").
			WithStudioId(1).
			WithStartHour(1).
			WithEndHour(2).
			ToDto()

		//repo := pool.IUserRepository{t}
		producerRepo := new(mocks.IProducerRepository)
		reserveRepo := new(mocks.IReserveRepository)

		producer := &model.Producer{
			Id:        producerGet.Id,
			Name:      producerGet.Name,
			StudioId:  producerGet.StudioId,
			StartHour: producerGet.StartHour,
			EndHour:   producerGet.EndHour,
		}

		producerRepo.On("Get", ctx, &dto.GetProducerRequest{
			Id: producerGet.Id,
		}).Return(
			producer, nil,
		).Once()

		logger := utils.NewMockLogger()

		service := impl.NewProducerService(logger, producerRepo, reserveRepo)
		sCtx.WithNewParameters("ctx", ctx, "request", producerGet)

		producerRes, err := service.Get(&dto.GetProducerRequest{
			Id: producerGet.Id,
		})

		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(producer, producerRes)
	})
}

func (suite *ProducerSuite) TestProducerGet02(t provider.T) {
	t.Title("[Get] incorrect id")
	t.Tags("producer", "get")
	t.Parallel()

	t.WithNewStep("fail get", func(sCtx provider.StepCtx) {
		ctx := context.Background()
		builder := utils.ProducerBuilder{}
		producerGet := builder.
			WithId(1).
			WithName("1").
			WithStudioId(1).
			WithStartHour(1).
			WithEndHour(2).
			ToDto()

		//repo := pool.IUserRepository{t}
		producerRepo := new(mocks.IProducerRepository)
		reserveRepo := new(mocks.IReserveRepository)

		producer := &model.Producer{
			Id:        producerGet.Id,
			Name:      producerGet.Name,
			StudioId:  producerGet.StudioId,
			StartHour: producerGet.StartHour,
			EndHour:   producerGet.EndHour,
		}

		producerRepo.On("Get", ctx, &dto.GetProducerRequest{
			Id: producerGet.Id,
		}).Return(
			nil, fmt.Errorf("incorrect id"),
		).Once()

		logger := utils.NewMockLogger()

		service := impl.NewProducerService(logger, producerRepo, reserveRepo)
		sCtx.WithNewParameters("ctx", ctx, "request", producerGet)

		producer, err := service.Get(&dto.GetProducerRequest{
			Id: producerGet.Id,
		})

		sCtx.Assert().Error(err)
		sCtx.Assert().Nil(producer)
	})
}
