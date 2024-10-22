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

type StudioSuite struct {
	suite.Suite
}

func (suite *StudioSuite) TestStudioGet01(t provider.T) {
	t.Title("[Get] success")
	t.Tags("studio", "get")
	t.Parallel()

	t.WithNewStep("successfully get", func(sCtx provider.StepCtx) {
		ctx := context.Background()
		builder := utils.StudioBuilder{}
		studio := builder.
			WithId(1).
			WithName("1").
			ToDto()

		//repo := pool.IUserRepository{t}
		studioRepo := new(mocks.IStudioRepository)

		studioRepo.On("Get", ctx, &dto.GetStudioRequest{
			Id: studio.Id,
		}).Return(
			studio, nil,
		).Once()

		logger := utils.NewMockLogger()

		service := impl.NewStudioService(logger, studioRepo)
		sCtx.WithNewParameters("ctx", ctx, "request", studio)

		studioRes, err := service.Get(&dto.GetStudioRequest{
			Id: studio.Id,
		})

		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(studio, studioRes)
	})
}

func (suite *StudioSuite) TestStudioGet02(t provider.T) {
	t.Title("[Get] failed")
	t.Tags("studio", "get")
	t.Parallel()

	t.WithNewStep("failed to get", func(sCtx provider.StepCtx) {
		ctx := context.Background()
		builder := utils.StudioBuilder{}
		studio := builder.
			WithId(1).
			WithName("").
			ToDto()

		//repo := pool.IUserRepository{t}
		studioRepo := new(mocks.IStudioRepository)

		studioRepo.On("Get", ctx, &dto.GetStudioRequest{
			Id: studio.Id,
		}).Return(
			nil, fmt.Errorf("failed to get"),
		).Once()

		logger := utils.NewMockLogger()

		service := impl.NewStudioService(logger, studioRepo)
		sCtx.WithNewParameters("ctx", ctx, "request", studio)

		studio, err := service.Get(&dto.GetStudioRequest{
			Id: studio.Id,
		})

		sCtx.Assert().Error(err)
		sCtx.Assert().Nil(studio)
	})
}

func (suite *StudioSuite) TestStudioDelete01(t provider.T) {
	t.Title("[Delete] success")
	t.Tags("studio", "delete")
	t.Parallel()

	t.WithNewStep("successfully delete", func(sCtx provider.StepCtx) {
		ctx := context.Background()

		request := utils.StudioFabric{Id: 1}.CorrectStudioDelete()

		//repo := pool.IUserRepository{t}
		studioRepo := new(mocks.IStudioRepository)

		studioRepo.On("Delete", ctx, request).Return(
			nil,
		).Once()

		logger := utils.NewMockLogger()

		service := impl.NewStudioService(logger, studioRepo)
		sCtx.WithNewParameters("ctx", ctx, "request", request)

		err := service.Delete(request)

		sCtx.Assert().NoError(err)
	})
}

func (suite *StudioSuite) TestStudioDelete02(t provider.T) {
	t.Title("[Delete] failed")
	t.Tags("studio", "delete")
	t.Parallel()

	t.WithNewStep("failed to delete", func(sCtx provider.StepCtx) {
		ctx := context.Background()

		request := utils.StudioFabric{Id: 0}.IncorrectStudioDelete()

		//repo := pool.IUserRepository{t}
		studioRepo := new(mocks.IStudioRepository)

		studioRepo.On("Delete", ctx, request).Return(
			fmt.Errorf("no studio with this id"),
		).Once()

		logger := utils.NewMockLogger()

		service := impl.NewStudioService(logger, studioRepo)
		sCtx.WithNewParameters("ctx", ctx, "request", request)

		err := service.Delete(request)

		sCtx.Assert().Error(err)
	})
}

func (suite *StudioSuite) TestStudioAdd01(t provider.T) {
	t.Title("[Delete] failed")
	t.Tags("studio", "delete")
	t.Parallel()

	t.WithNewStep("failed to delete", func(sCtx provider.StepCtx) {
		ctx := context.Background()

		request := utils.StudioFabric{Id: 1}.CorrectStudioAdd()

		//repo := pool.IUserRepository{t}
		studioRepo := new(mocks.IStudioRepository)

		studioRepo.On("Add", ctx, request).Return(
			nil,
		).Once()

		logger := utils.NewMockLogger()

		service := impl.NewStudioService(logger, studioRepo)
		sCtx.WithNewParameters("ctx", ctx, "request", request)

		err := service.Add(ctx, request)

		sCtx.Assert().NoError(err)
	})
}

func (suite *StudioSuite) TestStudioAdd02(t provider.T) {
	t.Title("[Delete] failed")
	t.Tags("studio", "delete")
	t.Parallel()

	t.WithNewStep("failed to delete", func(sCtx provider.StepCtx) {
		ctx := context.Background()

		request := utils.StudioFabric{Id: 1}.IncorrectStudioAdd()

		//repo := pool.IUserRepository{t}
		studioRepo := new(mocks.IStudioRepository)

		studioRepo.On("Add", ctx, request).Return(
			nil,
		).Once()

		logger := utils.NewMockLogger()

		service := impl.NewStudioService(logger, studioRepo)
		sCtx.WithNewParameters("ctx", ctx, "request", request)

		err := service.Add(ctx, request)

		sCtx.Assert().Error(err)
	})
}
