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

type IngredientTypeSuite struct {
	suite.Suite
}

func (suite *IngredientTypeSuite) TestIngredientTypeService_Create1(t provider.T) {
	t.Title("[Ingredient type create] successfully created")
	t.Tags("ingredient_type", "create")
	t.Parallel()

	t.WithNewStep("successfully created", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		ingredientType := utils.NewIngredientTypeBuilder().
			WithId(uuid.New()).
			WithName("type").
			WithDescription("description").
			ToDto()

		repo := mocks.NewIIngredientTypeRepository(t)
		repo.On("Create", ctx, ingredientType).Return(
			nil,
		).Once()

		logger := utils.NewMockLogger()
		service := services.NewIngredientTypeService(repo, logger)

		sCtx.WithNewParameters("ctx", ctx, "request", ingredientType)
		err := service.Create(ctx, ingredientType)

		sCtx.Assert().NoError(err)
	})
}

func (suite *IngredientTypeSuite) TestIngredientTypeService_Create2(t provider.T) {
	t.Title("[Ingredient type create] empty name")
	t.Tags("ingredient_type", "create")
	t.Parallel()

	t.WithNewStep("empty name", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		ingredientType := utils.NewIngredientTypeBuilder().
			WithId(uuid.New()).
			WithName("").
			WithDescription("description").
			ToDto()

		repo := mocks.NewIIngredientTypeRepository(t)
		repo.On("Create", ctx, ingredientType).Return(
			nil,
		).Maybe()

		logger := utils.NewMockLogger()
		service := services.NewIngredientTypeService(repo, logger)

		sCtx.WithNewParameters("ctx", ctx, "request", ingredientType)
		err := service.Create(ctx, ingredientType)

		sCtx.Assert().Error(err)
	})
}

func (suite *IngredientTypeSuite) TestIngredientTypeService_DeleteById1(t provider.T) {
	t.Title("[Ingredient type delete by id] successfully deleted")
	t.Tags("ingredient_type", "delete_by_id")
	t.Parallel()

	t.WithNewStep("successfully deleted", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		id := uuid.New()

		repo := mocks.NewIIngredientTypeRepository(t)
		repo.On("DeleteById", ctx, id).Return(
			nil,
		).Once()

		logger := utils.NewMockLogger()
		service := services.NewIngredientTypeService(repo, logger)

		sCtx.WithNewParameters("ctx", ctx, "request", id)
		err := service.DeleteById(ctx, id)

		sCtx.Assert().NoError(err)
	})
}

func (suite *IngredientTypeSuite) TestIngredientTypeService_DeleteById2(t provider.T) {
	t.Title("[Ingredient type delete by id] repo error")
	t.Tags("ingredient_type", "delete_by_id")
	t.Parallel()

	t.WithNewStep("repo error", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		id := uuid.New()

		repo := mocks.NewIIngredientTypeRepository(t)
		repo.On("DeleteById", ctx, id).Return(
			fmt.Errorf("repo error"),
		).Once()

		logger := utils.NewMockLogger()
		service := services.NewIngredientTypeService(repo, logger)

		sCtx.WithNewParameters("ctx", ctx, "request", id)
		err := service.DeleteById(ctx, id)

		sCtx.Assert().Error(err)
	})
}

func (suite *IngredientTypeSuite) TestIngredientTypeService_GetAll1(t provider.T) {
	t.Title("[Ingredient type get all] successful get")
	t.Tags("ingredient_type", "get_all")
	t.Parallel()

	t.WithNewStep("successful get", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		id := uuid.New()
		expTypes := []*domain.IngredientType{
			{
				ID:          uuid.UUID{1},
				Name:        "type1",
				Description: "description1",
			},
			{
				ID:          uuid.UUID{2},
				Name:        "type2",
				Description: "description2",
			},
		}

		repo := mocks.NewIIngredientTypeRepository(t)
		repo.On("GetAll", ctx).Return(
			expTypes, nil,
		).Once()

		logger := utils.NewMockLogger()
		service := services.NewIngredientTypeService(repo, logger)

		sCtx.WithNewParameters("ctx", ctx, "request", id)
		types, err := service.GetAll(ctx)

		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(types, expTypes)
	})
}

func (suite *IngredientTypeSuite) TestIngredientTypeService_GetAll2(t provider.T) {
	t.Title("[Ingredient type get all] repo error")
	t.Tags("ingredient_type", "get_all")
	t.Parallel()

	t.WithNewStep("repo error", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		id := uuid.New()

		repo := mocks.NewIIngredientTypeRepository(t)
		repo.On("GetAll", ctx).Return(
			nil, fmt.Errorf("repo error"),
		).Once()

		logger := utils.NewMockLogger()
		service := services.NewIngredientTypeService(repo, logger)

		sCtx.WithNewParameters("ctx", ctx, "request", id)
		comments, err := service.GetAll(ctx)

		sCtx.Assert().Error(err)
		sCtx.Assert().Nil(comments)
	})
}

func (suite *IngredientTypeSuite) TestIngredientTypeService_GetById1(t provider.T) {
	t.Title("[Ingredient type get by id] successful get")
	t.Tags("ingredient_type", "get_by_id")
	t.Parallel()

	t.WithNewStep("successful get", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		expType := utils.NewIngredientTypeBuilder().
			WithId(uuid.New()).
			WithName("type").
			WithDescription("description").
			ToDto()

		repo := mocks.NewIIngredientTypeRepository(t)
		repo.On("GetById", ctx, expType.ID).Return(
			expType, nil,
		).Once()

		logger := utils.NewMockLogger()
		service := services.NewIngredientTypeService(repo, logger)

		sCtx.WithNewParameters("ctx", ctx, "request", expType.ID)
		ingredientType, err := service.GetById(ctx, expType.ID)

		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(ingredientType, expType)
	})
}

func (suite *IngredientTypeSuite) TestIngredientTypeService_GetById2(t provider.T) {
	t.Title("[Ingredient type get by id] repo error")
	t.Tags("ingredient_type", "get_by_id")
	t.Parallel()

	t.WithNewStep("repo error", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		id := uuid.New()

		repo := mocks.NewIIngredientTypeRepository(t)
		repo.On("GetById", ctx, id).Return(
			nil, fmt.Errorf("repo error"),
		).Once()

		logger := utils.NewMockLogger()
		service := services.NewIngredientTypeService(repo, logger)

		sCtx.WithNewParameters("ctx", ctx, "request", id)
		comment, err := service.GetById(ctx, id)

		sCtx.Assert().Error(err)
		sCtx.Assert().Nil(comment)
	})
}

func (suite *IngredientTypeSuite) TestIngredientTypeService_Update1(t provider.T) {
	t.Title("[Ingredient type update] successfully updated")
	t.Tags("ingredient_type", "update")
	t.Parallel()

	t.WithNewStep("successfully updated", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		expType := utils.NewIngredientTypeBuilder().
			WithId(uuid.New()).
			WithName("type").
			WithDescription("description").
			ToDto()

		repo := mocks.NewIIngredientTypeRepository(t)
		repo.On("Update", ctx, expType).Return(
			nil,
		).Once()

		logger := utils.NewMockLogger()
		service := services.NewIngredientTypeService(repo, logger)

		sCtx.WithNewParameters("ctx", ctx, "request", expType)
		err := service.Update(ctx, expType)

		sCtx.Assert().NoError(err)
	})
}

func (suite *IngredientTypeSuite) TestIngredientTypeService_Update2(t provider.T) {
	t.Title("[Ingredient type update] repo error")
	t.Tags("ingredient_type", "update")
	t.Parallel()

	t.WithNewStep("repo error", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		expType := utils.NewIngredientTypeBuilder().
			WithId(uuid.New()).
			WithName("type").
			WithDescription("description").
			ToDto()

		repo := mocks.NewIIngredientTypeRepository(t)
		repo.On("Update", ctx, expType).Return(
			fmt.Errorf("repo error"),
		).Once()

		logger := utils.NewMockLogger()
		service := services.NewIngredientTypeService(repo, logger)

		sCtx.WithNewParameters("ctx", ctx, "request", expType)
		err := service.Update(ctx, expType)

		sCtx.Assert().Error(err)
	})
}
