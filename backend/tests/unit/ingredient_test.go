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

type IngredientSuite struct {
	suite.Suite
}

func (suite *IngredientSuite) TestIngredientService_Create1(t provider.T) {
	t.Title("[ingredient create] successfully created")
	t.Tags("ingredient", "create")
	t.Parallel()

	t.WithNewStep("successfully created", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		Ingredient := utils.NewIngredientBuilder().
			WithId(uuid.New()).
			WithTypeId(uuid.New()).
			WithName("ingredient").
			WithCalories(123).
			ToDto()

		repo := mocks.NewIIngredientRepository(t)
		repo.On("Create", ctx, Ingredient).Return(
			nil,
		).Once()

		logger := utils.NewMockLogger()
		service := services.NewIngredientService(repo, logger)

		sCtx.WithNewParameters("ctx", ctx, "request", Ingredient)
		err := service.Create(ctx, Ingredient)

		sCtx.Assert().NoError(err)
	})
}

func (suite *IngredientSuite) TestIngredientService_Create2(t provider.T) {
	t.Title("[ingredient create] empty name")
	t.Tags("ingredient", "create")
	t.Parallel()

	t.WithNewStep("empty name", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		Ingredient := utils.NewIngredientBuilder().
			WithId(uuid.New()).
			WithTypeId(uuid.New()).
			WithName("").
			WithCalories(123).
			ToDto()

		repo := mocks.NewIIngredientRepository(t)
		repo.On("Create", ctx, Ingredient).Return(
			nil,
		).Maybe()

		logger := utils.NewMockLogger()
		service := services.NewIngredientService(repo, logger)

		sCtx.WithNewParameters("ctx", ctx, "request", Ingredient)
		err := service.Create(ctx, Ingredient)

		sCtx.Assert().Error(err)
	})
}

func (suite *IngredientSuite) TestIngredientService_DeleteById1(t provider.T) {
	t.Title("[ingredient delete by id] successfully deleted")
	t.Tags("ingredient", "delete_by_id")
	t.Parallel()

	t.WithNewStep("successfully deleted", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		id := uuid.New()

		repo := mocks.NewIIngredientRepository(t)
		repo.On("DeleteById", ctx, id).Return(
			nil,
		).Once()

		logger := utils.NewMockLogger()
		service := services.NewIngredientService(repo, logger)

		sCtx.WithNewParameters("ctx", ctx, "request", id)
		err := service.DeleteById(ctx, id)

		sCtx.Assert().NoError(err)
	})
}

func (suite *IngredientSuite) TestIngredientService_DeleteById2(t provider.T) {
	t.Title("[ingredient delete by id] repo error")
	t.Tags("ingredient", "delete_by_id")
	t.Parallel()

	t.WithNewStep("repo error", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		id := uuid.New()

		repo := mocks.NewIIngredientRepository(t)
		repo.On("DeleteById", ctx, id).Return(
			fmt.Errorf("repo error"),
		).Once()

		logger := utils.NewMockLogger()
		service := services.NewIngredientService(repo, logger)

		sCtx.WithNewParameters("ctx", ctx, "request", id)
		err := service.DeleteById(ctx, id)

		sCtx.Assert().Error(err)
	})
}

func (suite *IngredientSuite) TestIngredientService_GetAll1(t provider.T) {
	t.Title("[ingredient get all] successful get")
	t.Tags("ingredient", "get_all")
	t.Parallel()

	t.WithNewStep("successful get", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		page := 1
		expTypes := []*domain.Ingredient{
			{
				ID:       uuid.UUID{1},
				Name:     "ingredient1",
				Calories: 1,
				TypeID:   uuid.UUID{1},
			},
			{
				ID:       uuid.UUID{2},
				Name:     "ingredient2",
				Calories: 2,
				TypeID:   uuid.UUID{2},
			},
		}

		repo := mocks.NewIIngredientRepository(t)
		repo.On("GetAll", ctx, page).Return(
			expTypes, 1, nil,
		).Once()

		logger := utils.NewMockLogger()
		service := services.NewIngredientService(repo, logger)

		sCtx.WithNewParameters("ctx", ctx, "request", page)
		types, pages, err := service.GetAll(ctx, page)

		sCtx.Assert().NoError(err)
		sCtx.Assert().NotZero(pages)
		sCtx.Assert().Equal(types, expTypes)
	})
}

func (suite *IngredientSuite) TestIngredientService_GetAll2(t provider.T) {
	t.Title("[ingredient get all] repo error")
	t.Tags("ingredient", "get_all")
	t.Parallel()

	t.WithNewStep("repo error", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		page := 1

		repo := mocks.NewIIngredientRepository(t)
		repo.On("GetAll", ctx, page).Return(
			nil, 1, fmt.Errorf("repo error"),
		).Once()

		logger := utils.NewMockLogger()
		service := services.NewIngredientService(repo, logger)

		sCtx.WithNewParameters("ctx", ctx, "request", page)
		comments, pages, err := service.GetAll(ctx, page)

		sCtx.Assert().Error(err)
		sCtx.Assert().Zero(pages)
		sCtx.Assert().Nil(comments)
	})
}

func (suite *IngredientSuite) TestIngredientService_GetById1(t provider.T) {
	t.Title("[ingredient get by id] successful get")
	t.Tags("ingredient", "get_by_id")
	t.Parallel()

	t.WithNewStep("successful get", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		expIngredient := utils.NewIngredientBuilder().
			WithId(uuid.New()).
			WithTypeId(uuid.New()).
			WithName("ingredient").
			WithCalories(123).
			ToDto()

		repo := mocks.NewIIngredientRepository(t)
		repo.On("GetById", ctx, expIngredient.ID).Return(
			expIngredient, nil,
		).Once()

		logger := utils.NewMockLogger()
		service := services.NewIngredientService(repo, logger)

		sCtx.WithNewParameters("ctx", ctx, "request", expIngredient.ID)
		Ingredient, err := service.GetById(ctx, expIngredient.ID)

		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(Ingredient, expIngredient)
	})
}

func (suite *IngredientSuite) TestIngredientService_GetById2(t provider.T) {
	t.Title("[ingredient get by id] repo error")
	t.Tags("ingredient", "get_by_id")
	t.Parallel()

	t.WithNewStep("repo error", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		id := uuid.New()

		repo := mocks.NewIIngredientRepository(t)
		repo.On("GetById", ctx, id).Return(
			nil, fmt.Errorf("repo error"),
		).Once()

		logger := utils.NewMockLogger()
		service := services.NewIngredientService(repo, logger)

		sCtx.WithNewParameters("ctx", ctx, "request", id)
		comment, err := service.GetById(ctx, id)

		sCtx.Assert().Error(err)
		sCtx.Assert().Nil(comment)
	})
}

func (suite *IngredientSuite) TestIngredientService_Update1(t provider.T) {
	t.Title("[ingredient update] successfully updated")
	t.Tags("ingredient", "update")
	t.Parallel()

	t.WithNewStep("successfully updated", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		expIngredient := utils.NewIngredientBuilder().
			WithId(uuid.New()).
			WithTypeId(uuid.New()).
			WithName("ingredient").
			WithCalories(123).
			ToDto()

		repo := mocks.NewIIngredientRepository(t)
		repo.On("Update", ctx, expIngredient).Return(
			nil,
		).Once()

		logger := utils.NewMockLogger()
		service := services.NewIngredientService(repo, logger)

		sCtx.WithNewParameters("ctx", ctx, "request", expIngredient)
		err := service.Update(ctx, expIngredient)

		sCtx.Assert().NoError(err)
	})
}

func (suite *IngredientSuite) TestIngredientService_Update2(t provider.T) {
	t.Title("[ingredient update] repo error")
	t.Tags("ingredient", "update")
	t.Parallel()

	t.WithNewStep("repo error", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		expIngredient := utils.NewIngredientBuilder().
			WithId(uuid.New()).
			WithTypeId(uuid.New()).
			WithName("ingredient").
			WithCalories(123).
			ToDto()

		repo := mocks.NewIIngredientRepository(t)
		repo.On("Update", ctx, expIngredient).Return(
			fmt.Errorf("repo error"),
		).Once()

		logger := utils.NewMockLogger()
		service := services.NewIngredientService(repo, logger)

		sCtx.WithNewParameters("ctx", ctx, "request", expIngredient)
		err := service.Update(ctx, expIngredient)

		sCtx.Assert().Error(err)
	})
}

func (suite *IngredientSuite) TestIngredientService_Link1(t provider.T) {
	t.Title("[ingredient link] successfully linked")
	t.Tags("ingredient", "link")
	t.Parallel()

	t.WithNewStep("successfully linked", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		recipeId := uuid.New()
		ingredientId := uuid.New()

		repo := mocks.NewIIngredientRepository(t)
		repo.On("Link", ctx, recipeId, ingredientId).Return(
			uuid.New(), nil,
		).Once()

		logger := utils.NewMockLogger()
		service := services.NewIngredientService(repo, logger)

		sCtx.WithNewParameters("ctx", ctx, "request", recipeId, ingredientId)
		linkId, err := service.Link(ctx, recipeId, ingredientId)

		sCtx.Assert().NoError(err)
		sCtx.Assert().NotEqual(linkId, uuid.New())
	})
}

func (suite *IngredientSuite) TestIngredientService_Link2(t provider.T) {
	t.Title("[ingredient link] repo error")
	t.Tags("ingredient", "link")
	t.Parallel()

	t.WithNewStep("repo error", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		recipeId := uuid.New()
		ingredientId := uuid.New()

		repo := mocks.NewIIngredientRepository(t)
		repo.On("Link", ctx, recipeId, ingredientId).Return(
			uuid.Nil, fmt.Errorf("repo error"),
		).Once()

		logger := utils.NewMockLogger()
		service := services.NewIngredientService(repo, logger)

		sCtx.WithNewParameters("ctx", ctx, "request", recipeId, ingredientId)
		linkId, err := service.Link(ctx, recipeId, ingredientId)

		sCtx.Assert().Error(err)
		sCtx.Assert().Equal(linkId, uuid.Nil)
	})
}

func (suite *IngredientSuite) TestIngredientService_Unlink1(t provider.T) {
	t.Title("[ingredient unlink] successfully unlinked")
	t.Tags("ingredient", "unlink")
	t.Parallel()

	t.WithNewStep("successfully unlinked", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		recipeId := uuid.New()
		ingredientId := uuid.New()

		repo := mocks.NewIIngredientRepository(t)
		repo.On("Unlink", ctx, recipeId, ingredientId).Return(
			nil,
		).Once()

		logger := utils.NewMockLogger()
		service := services.NewIngredientService(repo, logger)

		sCtx.WithNewParameters("ctx", ctx, "request", recipeId, ingredientId)
		err := service.Unlink(ctx, recipeId, ingredientId)

		sCtx.Assert().NoError(err)
	})
}

func (suite *IngredientSuite) TestIngredientService_Unlink2(t provider.T) {
	t.Title("[ingredient unlink] repo error")
	t.Tags("ingredient", "unlink")
	t.Parallel()

	t.WithNewStep("repo error", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		recipeId := uuid.New()
		ingredientId := uuid.New()

		repo := mocks.NewIIngredientRepository(t)
		repo.On("Unlink", ctx, recipeId, ingredientId).Return(
			fmt.Errorf("repo error"),
		).Once()

		logger := utils.NewMockLogger()
		service := services.NewIngredientService(repo, logger)

		sCtx.WithNewParameters("ctx", ctx, "request", recipeId, ingredientId)
		err := service.Unlink(ctx, recipeId, ingredientId)

		sCtx.Assert().Error(err)
	})
}

func (suite *IngredientSuite) TestIngredientService_GetAllIngredientsByRecipeId1(t provider.T) {
	t.Title("[ingredient get all by recipe id] successful get")
	t.Tags("ingredient", "get_all_by_recipe")
	t.Parallel()

	t.WithNewStep("successful get", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		recipeId := uuid.New()
		expTypes := []*domain.Ingredient{
			{
				ID:       uuid.UUID{1},
				Name:     "ingredient1",
				Calories: 1,
				TypeID:   uuid.UUID{1},
			},
			{
				ID:       uuid.UUID{2},
				Name:     "ingredient2",
				Calories: 2,
				TypeID:   uuid.UUID{2},
			},
		}

		repo := mocks.NewIIngredientRepository(t)
		repo.On("GetAllByRecipeId", ctx, recipeId).Return(
			expTypes, nil,
		).Once()

		logger := utils.NewMockLogger()
		service := services.NewIngredientService(repo, logger)

		sCtx.WithNewParameters("ctx", ctx, "request", recipeId)
		ingredients, err := service.GetAllByRecipeId(ctx, recipeId)

		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(ingredients, expTypes)
	})
}

func (suite *IngredientSuite) TestIngredientService_GetAllIngredientsByRecipeId2(t provider.T) {
	t.Title("[ingredient get all by recipe id] repo error")
	t.Tags("ingredient", "get_all_by_recipe")
	t.Parallel()

	t.WithNewStep("repo error", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		recipeId := uuid.New()

		repo := mocks.NewIIngredientRepository(t)
		repo.On("GetAllByRecipeId", ctx, recipeId).Return(
			nil, fmt.Errorf("repo error"),
		).Once()

		logger := utils.NewMockLogger()
		service := services.NewIngredientService(repo, logger)

		sCtx.WithNewParameters("ctx", ctx, "request", recipeId)
		ingredients, err := service.GetAllByRecipeId(ctx, recipeId)

		sCtx.Assert().Error(err)
		sCtx.Assert().Nil(ingredients)
	})
}
