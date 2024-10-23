//go:build unit_test

package unit_tests

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
	"ppo/domain"
	"ppo/mocks"
	"ppo/services"
	"ppo/tests/utils"
)

type MeasurementSuite struct {
	suite.Suite
}

func (suite *MeasurementSuite) TestIngredientTypeService_Create1(t provider.T) {
	t.Title("[Measurement create] successfully created")
	t.Tags("measurement", "create")
	t.Parallel()

	t.WithNewStep("successfully created", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		measurement := utils.NewMeasurementBuilder().
			WithId(uuid.New()).
			WithName("gram").
			WithGrams(1).
			ToDto()

		repo := mocks.NewIMeasurementRepository(t)
		repo.On("Create", ctx, measurement).Return(
			nil,
		).Once()

		logger := utils.NewMockLogger()
		service := services.NewMeasurementService(repo, logger)

		sCtx.WithNewParameters("ctx", ctx, "request", measurement)
		err := service.Create(ctx, measurement)

		sCtx.Assert().NoError(err)
	})
}

func (suite *MeasurementSuite) TestIngredientTypeService_Create2(t provider.T) {
	t.Title("[Measurement create] empty name")
	t.Tags("measurement", "create")
	t.Parallel()

	t.WithNewStep("empty name", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		measurement := utils.NewMeasurementBuilder().
			WithId(uuid.New()).
			WithName("").
			WithGrams(1).
			ToDto()

		repo := mocks.NewIMeasurementRepository(t)
		repo.On("Create", ctx, measurement).Return(
			nil,
		).Maybe()

		logger := utils.NewMockLogger()
		service := services.NewMeasurementService(repo, logger)

		sCtx.WithNewParameters("ctx", ctx, "request", measurement)
		err := service.Create(ctx, measurement)

		sCtx.Assert().Error(err)
	})
}

func (suite *MeasurementSuite) TestIngredientTypeService_DeleteById1(t provider.T) {
	t.Title("[Measurement delete by id] successfully deleted")
	t.Tags("measurement", "delete_by_id")
	t.Parallel()

	t.WithNewStep("successfully deleted", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		id := uuid.New()

		repo := mocks.NewIMeasurementRepository(t)
		repo.On("DeleteById", ctx, id).Return(
			nil,
		).Once()

		logger := utils.NewMockLogger()
		service := services.NewMeasurementService(repo, logger)

		sCtx.WithNewParameters("ctx", ctx, "request", id)
		err := service.DeleteById(ctx, id)

		sCtx.Assert().NoError(err)
	})
}

func (suite *MeasurementSuite) TestIngredientTypeService_DeleteById2(t provider.T) {
	t.Title("[Measurement delete by id] repo error")
	t.Tags("measurement", "delete_by_id")
	t.Parallel()

	t.WithNewStep("repo error", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		id := uuid.New()

		repo := mocks.NewIMeasurementRepository(t)
		repo.On("DeleteById", ctx, id).Return(
			fmt.Errorf("repo error"),
		).Once()

		logger := utils.NewMockLogger()
		service := services.NewMeasurementService(repo, logger)

		sCtx.WithNewParameters("ctx", ctx, "request", id)
		err := service.DeleteById(ctx, id)

		sCtx.Assert().Error(err)
	})
}

func (suite *MeasurementSuite) TestIngredientTypeService_GetAll1(t provider.T) {
	t.Title("[Measurement get all] successful get")
	t.Tags("measurement", "get_all")
	t.Parallel()

	t.WithNewStep("successful get", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		expMeasurements := []*domain.Measurement{
			{
				ID:    uuid.UUID{1},
				Name:  "measurement1",
				Grams: 1,
			},
			{
				ID:    uuid.UUID{2},
				Name:  "measurement2",
				Grams: 2,
			},
		}

		repo := mocks.NewIMeasurementRepository(t)
		repo.On("GetAll", ctx).Return(
			expMeasurements, nil,
		).Once()

		logger := utils.NewMockLogger()
		service := services.NewMeasurementService(repo, logger)

		sCtx.WithNewParameters("ctx", ctx)
		types, err := service.GetAll(ctx)

		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(types, expMeasurements)
	})
}

func (suite *MeasurementSuite) TestIngredientTypeService_GetAll2(t provider.T) {
	t.Title("[Measurement get all] repo error")
	t.Tags("measurements", "get_all")
	t.Parallel()

	t.WithNewStep("repo error", func(sCtx provider.StepCtx) {
		ctx := context.TODO()

		repo := mocks.NewIMeasurementRepository(t)
		repo.On("GetAll", ctx).Return(
			nil, fmt.Errorf("repo error"),
		).Once()

		logger := utils.NewMockLogger()
		service := services.NewMeasurementService(repo, logger)

		sCtx.WithNewParameters("ctx", ctx)
		comments, err := service.GetAll(ctx)

		sCtx.Assert().Error(err)
		sCtx.Assert().Nil(comments)
	})
}

func (suite *MeasurementSuite) TestIngredientTypeService_GetById1(t provider.T) {
	t.Title("[Measurement get by id] successful get")
	t.Tags("measurement", "get_by_id")
	t.Parallel()

	t.WithNewStep("successful get", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		expMeasurement := utils.NewMeasurementBuilder().
			WithId(uuid.New()).
			WithName("gram").
			WithGrams(1).
			ToDto()

		repo := mocks.NewIMeasurementRepository(t)
		repo.On("GetById", ctx, expMeasurement.ID).Return(
			expMeasurement, nil,
		).Once()

		logger := utils.NewMockLogger()
		service := services.NewMeasurementService(repo, logger)

		sCtx.WithNewParameters("ctx", ctx, "request", expMeasurement.ID)
		ingredientType, err := service.GetById(ctx, expMeasurement.ID)

		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(ingredientType, expMeasurement)
	})
}

func (suite *MeasurementSuite) TestIngredientTypeService_GetById2(t provider.T) {
	t.Title("[Measurement get by id] repo error")
	t.Tags("measurement", "get_by_id")
	t.Parallel()

	t.WithNewStep("repo error", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		id := uuid.New()

		repo := mocks.NewIMeasurementRepository(t)
		repo.On("GetById", ctx, id).Return(
			nil, fmt.Errorf("repo error"),
		).Once()

		logger := utils.NewMockLogger()
		service := services.NewMeasurementService(repo, logger)

		sCtx.WithNewParameters("ctx", ctx, "request", id)
		measurement, err := service.GetById(ctx, id)

		sCtx.Assert().Error(err)
		sCtx.Assert().Nil(measurement)
	})
}

func (suite *MeasurementSuite) TestIngredientTypeService_Update1(t provider.T) {
	t.Title("[Measurement update] successfully updated")
	t.Tags("measurement", "update")
	t.Parallel()

	t.WithNewStep("successfully updated", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		expMeasurement := utils.NewMeasurementBuilder().
			WithId(uuid.New()).
			WithName("gram").
			WithGrams(1).
			ToDto()

		repo := mocks.NewIMeasurementRepository(t)
		repo.On("Update", ctx, expMeasurement).Return(
			nil,
		).Once()

		logger := utils.NewMockLogger()
		service := services.NewMeasurementService(repo, logger)

		sCtx.WithNewParameters("ctx", ctx, "request", expMeasurement)
		err := service.Update(ctx, expMeasurement)

		sCtx.Assert().NoError(err)
	})
}

func (suite *MeasurementSuite) TestIngredientTypeService_Update2(t provider.T) {
	t.Title("[Measurement update] repo error")
	t.Tags("measurement", "update")
	t.Parallel()

	t.WithNewStep("repo error", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		expMeasurement := utils.NewMeasurementBuilder().
			WithId(uuid.New()).
			WithName("gram").
			WithGrams(1).
			ToDto()

		repo := mocks.NewIMeasurementRepository(t)
		repo.On("Update", ctx, expMeasurement).Return(
			fmt.Errorf("repo error"),
		).Once()

		logger := utils.NewMockLogger()
		service := services.NewMeasurementService(repo, logger)

		sCtx.WithNewParameters("ctx", ctx, "request", expMeasurement)
		err := service.Update(ctx, expMeasurement)

		sCtx.Assert().Error(err)
	})
}

func (suite *MeasurementSuite) TestIngredientTypeService_GetByRecipeId1(t provider.T) {
	t.Title("[Measurement get by recipe] successfully updated")
	t.Tags("measurement", "get_by_recipe")
	t.Parallel()

	t.WithNewStep("successful get", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		recipeId := uuid.New()
		ingredientId := uuid.New()
		expMeasurement := utils.NewMeasurementBuilder().
			WithId(uuid.New()).
			WithName("gram").
			WithGrams(1).
			ToDto()

		repo := mocks.NewIMeasurementRepository(t)
		repo.On("GetByRecipeId", ctx, ingredientId, recipeId).Return(
			expMeasurement, 1, nil,
		).Once()

		logger := utils.NewMockLogger()
		service := services.NewMeasurementService(repo, logger)

		sCtx.WithNewParameters("ctx", ctx, "request", ingredientId, recipeId)
		measurement, p, err := service.GetByRecipeId(ctx, ingredientId, recipeId)

		sCtx.Assert().NoError(err)
		sCtx.Assert().NotZero(p)
		sCtx.Assert().Equal(measurement, expMeasurement)
	})
}

func (suite *MeasurementSuite) TestIngredientTypeService_GetByRecipeId2(t provider.T) {
	t.Title("[Measurement get by recipe] repo error")
	t.Tags("measurement", "get_by_recipe")
	t.Parallel()

	t.WithNewStep("successful get", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		recipeId := uuid.New()
		ingredientId := uuid.New()

		repo := mocks.NewIMeasurementRepository(t)
		repo.On("GetByRecipeId", ctx, ingredientId, recipeId).Return(
			nil, 0, fmt.Errorf("repo error"),
		).Once()

		logger := utils.NewMockLogger()
		service := services.NewMeasurementService(repo, logger)

		sCtx.WithNewParameters("ctx", ctx, "request", ingredientId, recipeId)
		measurement, p, err := service.GetByRecipeId(ctx, ingredientId, recipeId)

		sCtx.Assert().Error(err)
		sCtx.Assert().Zero(p)
		sCtx.Assert().Nil(measurement)
	})
}

func (suite *MeasurementSuite) TestIngredientTypeService_UpdateLink1(t provider.T) {
	t.Title("[Measurement update link] successfully updated")
	t.Tags("measurement", "update_link")
	t.Parallel()

	t.WithNewStep("successfully updated", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		linkId := uuid.New()
		measurementId := uuid.New()
		amount := 1

		repo := mocks.NewIMeasurementRepository(t)
		repo.On("UpdateLink", ctx, linkId, measurementId, amount).Return(
			nil,
		).Once()

		logger := utils.NewMockLogger()
		service := services.NewMeasurementService(repo, logger)

		sCtx.WithNewParameters("ctx", ctx, "request", linkId, measurementId, amount)
		err := service.UpdateLink(ctx, linkId, measurementId, amount)

		sCtx.Assert().NoError(err)
	})
}

func (suite *MeasurementSuite) TestIngredientTypeService_UpdateLink2(t provider.T) {
	t.Title("[Measurement update link] repo error")
	t.Tags("measurement", "update_link")
	t.Parallel()

	t.WithNewStep("successfully updated", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		linkId := uuid.New()
		measurementId := uuid.New()
		amount := 1

		repo := mocks.NewIMeasurementRepository(t)
		repo.On("UpdateLink", ctx, linkId, measurementId, amount).Return(
			fmt.Errorf("repo error"),
		).Once()

		logger := utils.NewMockLogger()
		service := services.NewMeasurementService(repo, logger)

		sCtx.WithNewParameters("ctx", ctx, "request", linkId, measurementId, amount)
		err := service.UpdateLink(ctx, linkId, measurementId, amount)

		sCtx.Assert().Error(err)
	})
}
