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
	"ppo/services/dto"
	"ppo/tests/utils"
)

type RecipeSuite struct {
	suite.Suite
}

func (suite *RecipeSuite) TestRecipeService_Create1(t provider.T) {
	t.Title("[Recipe create] successfully created")
	t.Tags("recipe", "create")
	t.Parallel()

	t.WithNewStep("successfully created", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		recipe := utils.NewRecipeBuilder().
			WithId(uuid.New()).
			WithSaladId(uuid.New()).
			WithStatus(1).
			WithNumberOfServings(1).
			WithTimeToCook(1).
			WithRating(5.0).
			ToDto()

		repo := mocks.NewIRecipeRepository(t)
		repo.On("Create", ctx, recipe).Return(
			uuid.New(), nil,
		).Once()

		logger := utils.NewMockLogger()
		service := services.NewRecipeService(repo, logger)

		sCtx.WithNewParameters("ctx", ctx, "request", recipe)
		id, err := service.Create(ctx, recipe)

		sCtx.Assert().NoError(err)
		sCtx.Assert().NotEqual(id, uuid.Nil)
	})
}

func (suite *RecipeSuite) TestRecipeService_Create2(t provider.T) {
	t.Title("[Recipe create] repo error")
	t.Tags("recipe", "create")
	t.Parallel()

	t.WithNewStep("successfully created", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		recipe := utils.NewRecipeBuilder().
			WithId(uuid.New()).
			WithSaladId(uuid.New()).
			WithStatus(1).
			WithNumberOfServings(1).
			WithTimeToCook(1).
			WithRating(5.0).
			ToDto()

		repo := mocks.NewIRecipeRepository(t)
		repo.On("Create", ctx, recipe).Return(
			uuid.Nil, fmt.Errorf("repo error"),
		).Once()

		logger := utils.NewMockLogger()
		service := services.NewRecipeService(repo, logger)

		sCtx.WithNewParameters("ctx", ctx, "request", recipe)
		id, err := service.Create(ctx, recipe)

		sCtx.Assert().Error(err)
		sCtx.Assert().Equal(id, uuid.Nil)
	})
}

func (suite *RecipeSuite) TestRecipeService_DeleteById1(t provider.T) {
	t.Title("[Recipe delete by id] successfully deleted")
	t.Tags("recipe", "delete_by_id")
	t.Parallel()

	t.WithNewStep("successfully deleted", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		id := uuid.New()

		repo := mocks.NewIRecipeRepository(t)
		repo.On("DeleteById", ctx, id).Return(
			nil,
		).Once()

		logger := utils.NewMockLogger()
		service := services.NewRecipeService(repo, logger)

		sCtx.WithNewParameters("ctx", ctx, "request", id)
		err := service.DeleteById(ctx, id)

		sCtx.Assert().NoError(err)
	})
}

func (suite *RecipeSuite) TestRecipeService_DeleteById2(t provider.T) {
	t.Title("[Recipe delete by id] repo error")
	t.Tags("recipe", "delete_by_id")
	t.Parallel()

	t.WithNewStep("repo error", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		id := uuid.New()

		repo := mocks.NewIRecipeRepository(t)
		repo.On("DeleteById", ctx, id).Return(
			fmt.Errorf("repo error"),
		).Once()

		logger := utils.NewMockLogger()
		service := services.NewRecipeService(repo, logger)

		sCtx.WithNewParameters("ctx", ctx, "request", id)
		err := service.DeleteById(ctx, id)

		sCtx.Assert().Error(err)
	})
}

func (suite *RecipeSuite) TestRecipeService_GetAll1(t provider.T) {
	t.Title("[Recipe get all] successful get")
	t.Tags("recipe", "get_all")
	t.Parallel()

	t.WithNewStep("successful get", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		page := 1
		expRecipes := []*domain.Recipe{
			{
				ID:               uuid.UUID{1},
				SaladID:          uuid.UUID{11},
				Status:           1,
				NumberOfServings: 1,
				TimeToCook:       1,
				Rating:           5.0,
			},
			{
				ID:               uuid.UUID{2},
				SaladID:          uuid.UUID{22},
				Status:           2,
				NumberOfServings: 2,
				TimeToCook:       2,
				Rating:           5.0,
			},
		}
		filter := &dto.RecipeFilter{
			AvailableIngredients: nil,
			MinRate:              0,
			SaladTypes:           nil,
			Status:               1,
		}

		repo := mocks.NewIRecipeRepository(t)
		repo.On("GetAll", ctx, filter, page).Return(
			expRecipes, nil,
		).Once()

		logger := utils.NewMockLogger()
		service := services.NewRecipeService(repo, logger)

		sCtx.WithNewParameters("ctx", ctx, "request", filter, page)
		recipes, err := service.GetAll(ctx, filter, page)

		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(recipes, expRecipes)
	})
}

func (suite *RecipeSuite) TestRecipeService_GetAll2(t provider.T) {
	t.Title("[Recipe get all] repo error")
	t.Tags("recipe", "get_all")
	t.Parallel()

	t.WithNewStep("repo error", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		page := 1
		filter := &dto.RecipeFilter{
			AvailableIngredients: nil,
			MinRate:              0,
			SaladTypes:           nil,
			Status:               1,
		}

		repo := mocks.NewIRecipeRepository(t)
		repo.On("GetAll", ctx, filter, page).Return(
			nil, fmt.Errorf("repo error"),
		).Once()

		logger := utils.NewMockLogger()
		service := services.NewRecipeService(repo, logger)

		sCtx.WithNewParameters("ctx", ctx, "request", filter, page)
		recipes, err := service.GetAll(ctx, filter, page)

		sCtx.Assert().Error(err)
		sCtx.Assert().Nil(recipes)
	})
}

func (suite *RecipeSuite) TestRecipeService_GetById1(t provider.T) {
	t.Title("[Recipe get by id] successful get")
	t.Tags("recipe", "get_by_id")
	t.Parallel()

	t.WithNewStep("successful get", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		expRecipe := utils.NewRecipeBuilder().
			WithId(uuid.New()).
			WithSaladId(uuid.New()).
			WithStatus(1).
			WithNumberOfServings(1).
			WithTimeToCook(1).
			WithRating(5.0).
			ToDto()

		repo := mocks.NewIRecipeRepository(t)
		repo.On("GetById", ctx, expRecipe.ID).Return(
			expRecipe, nil,
		).Once()

		logger := utils.NewMockLogger()
		service := services.NewRecipeService(repo, logger)

		sCtx.WithNewParameters("ctx", ctx, "request", expRecipe.ID)
		recipe, err := service.GetById(ctx, expRecipe.ID)

		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(recipe, expRecipe)
	})
}

func (suite *RecipeSuite) TestRecipeService_GetById2(t provider.T) {
	t.Title("[Recipe get by id] repo error")
	t.Tags("recipe", "get_by_id")
	t.Parallel()

	t.WithNewStep("repo error", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		id := uuid.New()

		repo := mocks.NewIRecipeRepository(t)
		repo.On("GetById", ctx, id).Return(
			nil, fmt.Errorf("repo error"),
		).Once()

		logger := utils.NewMockLogger()
		service := services.NewRecipeService(repo, logger)

		sCtx.WithNewParameters("ctx", ctx, "request", id)
		recipe, err := service.GetById(ctx, id)

		sCtx.Assert().Error(err)
		sCtx.Assert().Nil(recipe)
	})
}

func (suite *RecipeSuite) TestRecipeService_Update1(t provider.T) {
	t.Title("[Recipe update] successfully updated")
	t.Tags("recipe", "update")
	t.Parallel()

	t.WithNewStep("successfully updated", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		expRecipe := utils.NewRecipeBuilder().
			WithId(uuid.New()).
			WithSaladId(uuid.New()).
			WithStatus(1).
			WithNumberOfServings(1).
			WithTimeToCook(1).
			WithRating(5.0).
			ToDto()

		repo := mocks.NewIRecipeRepository(t)
		repo.On("Update", ctx, expRecipe).Return(
			nil,
		).Once()

		logger := utils.NewMockLogger()
		service := services.NewRecipeService(repo, logger)

		sCtx.WithNewParameters("ctx", ctx, "request", expRecipe)
		err := service.Update(ctx, expRecipe)

		sCtx.Assert().NoError(err)
	})
}

func (suite *RecipeSuite) TestRecipeService_Update2(t provider.T) {
	t.Title("[Recipe update] repo error")
	t.Tags("recipe", "update")
	t.Parallel()

	t.WithNewStep("repo error", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		expRecipe := utils.NewRecipeBuilder().
			WithId(uuid.New()).
			WithSaladId(uuid.New()).
			WithStatus(1).
			WithNumberOfServings(1).
			WithTimeToCook(1).
			WithRating(5.0).
			ToDto()

		repo := mocks.NewIRecipeRepository(t)
		repo.On("Update", ctx, expRecipe).Return(
			fmt.Errorf("repo error"),
		).Once()

		logger := utils.NewMockLogger()
		service := services.NewRecipeService(repo, logger)

		sCtx.WithNewParameters("ctx", ctx, "request", expRecipe)
		err := service.Update(ctx, expRecipe)

		sCtx.Assert().Error(err)
	})
}

func (suite *RecipeSuite) TestRecipeService_GetBySaladId1(t provider.T) {
	t.Title("[Recipe get by salad id] successful get")
	t.Tags("recipe", "get_by_salad_id")
	t.Parallel()

	t.WithNewStep("successful get", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		expRecipe := utils.NewRecipeBuilder().
			WithId(uuid.New()).
			WithSaladId(uuid.New()).
			WithStatus(1).
			WithNumberOfServings(1).
			WithTimeToCook(1).
			WithRating(5.0).
			ToDto()

		repo := mocks.NewIRecipeRepository(t)
		repo.On("GetBySaladId", ctx, expRecipe.SaladID).Return(
			expRecipe, nil,
		).Once()

		logger := utils.NewMockLogger()
		service := services.NewRecipeService(repo, logger)

		sCtx.WithNewParameters("ctx", ctx, "request", expRecipe.SaladID)
		recipe, err := service.GetBySaladId(ctx, expRecipe.SaladID)

		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(recipe, expRecipe)
	})
}

func (suite *RecipeSuite) TestRecipeService_GetBySaladId2(t provider.T) {
	t.Title("[Recipe get by salad id] repo error")
	t.Tags("recipe", "get_by_salad_id")
	t.Parallel()

	t.WithNewStep("successful get", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		saladId := uuid.New()

		repo := mocks.NewIRecipeRepository(t)
		repo.On("GetBySaladId", ctx, saladId).Return(
			nil, fmt.Errorf("repo error"),
		).Once()

		logger := utils.NewMockLogger()
		service := services.NewRecipeService(repo, logger)

		sCtx.WithNewParameters("ctx", ctx, "request", saladId)
		recipe, err := service.GetBySaladId(ctx, saladId)

		sCtx.Assert().Error(err)
		sCtx.Assert().Nil(recipe)
	})
}
