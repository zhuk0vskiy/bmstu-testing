package repository_test

import (
	"backend/src/internal/model"
	"backend/src/internal/repository/interface/mocks"
	"backend/src/tests/utils"
	"context"
	"fmt"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
	"time"
)

type ReserveRepositorySuite struct {
	suite.Suite

	mockRepo mocks.IReserveRepository
}

func (suite *ReserveRepositorySuite) TestReserveAdd01(t provider.T) {
	t.Title("[Add] success")
	t.Tags("reserve", "add")
	t.Parallel()

	t.WithNewStep("successfully add", func(sCtx provider.StepCtx) {
		ctx := context.Background()

		request := utils.ReserveFabric{Id: 1}.CorrectReserveAdd()

		suite.mockRepo.On("Add", ctx, request).Return(
			nil,
		).Once()

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		err := suite.mockRepo.Add(ctx, request)

		sCtx.Assert().NoError(err)
	})
}

func (suite *ReserveRepositorySuite) TestReserveAdd02(t provider.T) {
	t.Title("[Add] failed")
	t.Tags("reserve", "add")
	t.Parallel()

	t.WithNewStep("failed to add", func(sCtx provider.StepCtx) {
		ctx := context.Background()

		request := utils.ReserveFabric{Id: 1}.IncorrectReserveAdd()

		suite.mockRepo.On("Add", ctx, request).Return(
			fmt.Errorf("failed to add"),
		).Once()

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		err := suite.mockRepo.Add(ctx, request)

		sCtx.Assert().Error(err)
	})
}

func (suite *ReserveRepositorySuite) TestReserveGetAll01(t provider.T) {
	t.Title("[GetAll] success")
	t.Tags("reserve", "getAll")
	t.Parallel()

	t.WithNewStep("success to getAll", func(sCtx provider.StepCtx) {
		ctx := context.Background()

		request := utils.ReserveFabric{Id: 1}.CorrectReserveGetAll()

		reserves := []*model.Reserve{
			{
				UserId:            1,
				RoomId:            1,
				ProducerId:        1,
				InstrumentalistId: 1,
				TimeInterval: &model.TimeInterval{
					StartTime: time.Date(2024, 4, 14, 12, 0, 0, 0, time.UTC),
					EndTime:   time.Date(2024, 4, 14, 13, 0, 0, 0, time.UTC),
				},
			},
		}

		suite.mockRepo.On("GetAll", ctx, request).Return(
			reserves, nil,
		).Once()

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		reserves, err := suite.mockRepo.GetAll(ctx, request)

		sCtx.Assert().NoError(err)
		sCtx.Assert().NotNil(reserves)

	})
}
