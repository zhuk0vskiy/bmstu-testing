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

type RecipeStepSuite struct {
	suite.Suite
}

func (suite *RecipeStepSuite) TestRecipeStepService_Create1(t provider.T) {
	t.Title("[RecipeStep create] successfully created")
	t.Tags("recipe_step", "create")
	t.Parallel()

	t.WithNewStep("successfully created", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		recipeStep := utils.NewRecipeStepBuilder().
			WithId(uuid.New()).
			WithRecipeId(uuid.New()).
			WithName("name").
			WithDescription("description").
			WithStepNum(1).
			ToDto()

		repo := mocks.NewIRecipeStepRepository(t)
		repo.On("Create", ctx, recipeStep).Return(
			nil,
		).Once()

		logger := utils.NewMockLogger()
		service := services.NewRecipeStepService(repo, logger)

		sCtx.WithNewParameters("ctx", ctx, "request", recipeStep)
		err := service.Create(ctx, recipeStep)

		sCtx.Assert().NoError(err)
	})
}

func (suite *RecipeStepSuite) TestRecipeStepService_Create2(t provider.T) {
	t.Title("[RecipeStep create] empty name")
	t.Tags("recipe_step", "create")
	t.Parallel()

	t.WithNewStep("empty name", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		recipeStep := utils.NewRecipeStepBuilder().
			WithId(uuid.New()).
			WithRecipeId(uuid.New()).
			WithName("").
			WithDescription("description").
			WithStepNum(1).
			ToDto()

		repo := mocks.NewIRecipeStepRepository(t)
		repo.On("Create", ctx, recipeStep).Return(
			nil,
		).Maybe()

		logger := utils.NewMockLogger()
		service := services.NewRecipeStepService(repo, logger)

		sCtx.WithNewParameters("ctx", ctx, "request", recipeStep)
		err := service.Create(ctx, recipeStep)

		sCtx.Assert().Error(err)
	})
}

func (suite *RecipeStepSuite) TestRecipeStepService_DeleteById1(t provider.T) {
	t.Title("[RecipeStep delete by id] successfully deleted")
	t.Tags("recipe_step", "delete_by_id")
	t.Parallel()

	t.WithNewStep("successfully deleted", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		id := uuid.New()

		repo := mocks.NewIRecipeStepRepository(t)
		repo.On("DeleteById", ctx, id).Return(
			nil,
		).Once()

		logger := utils.NewMockLogger()
		service := services.NewRecipeStepService(repo, logger)

		sCtx.WithNewParameters("ctx", ctx, "request", id)
		err := service.DeleteById(ctx, id)

		sCtx.Assert().NoError(err)
	})
}

func (suite *RecipeStepSuite) TestRecipeStepService_DeleteById2(t provider.T) {
	t.Title("[RecipeStep delete by id] repo error")
	t.Tags("recipe_step", "delete_by_id")
	t.Parallel()

	t.WithNewStep("repo error", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		id := uuid.New()

		repo := mocks.NewIRecipeStepRepository(t)
		repo.On("DeleteById", ctx, id).Return(
			fmt.Errorf("repo error"),
		).Once()

		logger := utils.NewMockLogger()
		service := services.NewRecipeStepService(repo, logger)

		sCtx.WithNewParameters("ctx", ctx, "request", id)
		err := service.DeleteById(ctx, id)

		sCtx.Assert().Error(err)
	})
}

func (suite *RecipeStepSuite) TestRecipeStepService_GetAllByRecipeID1(t provider.T) {
	t.Title("[RecipeStep get all by recipe] successful get")
	t.Tags("recipe_step", "get_all_by_recipe_id")
	t.Parallel()

	t.WithNewStep("successful get", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		recipeId := uuid.New()
		expSteps := []*domain.RecipeStep{
			{
				ID:          uuid.UUID{1},
				RecipeID:    recipeId,
				Name:        "ingredient1",
				Description: "description1",
				StepNum:     1,
			},
			{
				ID:          uuid.UUID{2},
				RecipeID:    recipeId,
				Name:        "ingredient2",
				Description: "description2",
				StepNum:     2,
			},
		}

		repo := mocks.NewIRecipeStepRepository(t)
		repo.On("GetAllByRecipeID", ctx, recipeId).Return(
			expSteps, nil,
		).Once()

		logger := utils.NewMockLogger()
		service := services.NewRecipeStepService(repo, logger)

		sCtx.WithNewParameters("ctx", ctx, "request", recipeId)
		steps, err := service.GetAllByRecipeID(ctx, recipeId)

		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(steps, expSteps)
	})
}

func (suite *RecipeStepSuite) TestRecipeStepService_GetAllByRecipeID2(t provider.T) {
	t.Title("[RecipeStep get all] repo error")
	t.Tags("recipe_step", "get_all_by_recipe_id")
	t.Parallel()

	t.WithNewStep("repo error", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		recipeId := uuid.New()

		repo := mocks.NewIRecipeStepRepository(t)
		repo.On("GetAllByRecipeID", ctx, recipeId).Return(
			nil, fmt.Errorf("repo error"),
		).Once()

		logger := utils.NewMockLogger()
		service := services.NewRecipeStepService(repo, logger)

		sCtx.WithNewParameters("ctx", ctx, "request", recipeId)
		steps, err := service.GetAllByRecipeID(ctx, recipeId)

		sCtx.Assert().Error(err)
		sCtx.Assert().Nil(steps)
	})
}

func (suite *RecipeStepSuite) TestRecipeStepService_GetById1(t provider.T) {
	t.Title("[RecipeStep get by id] successful get")
	t.Tags("recipe_step", "get_by_id")
	t.Parallel()

	t.WithNewStep("successful get", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		expStep := utils.NewRecipeStepBuilder().
			WithId(uuid.New()).
			WithRecipeId(uuid.New()).
			WithName("name").
			WithDescription("description").
			WithStepNum(1).
			ToDto()

		repo := mocks.NewIRecipeStepRepository(t)
		repo.On("GetById", ctx, expStep.ID).Return(
			expStep, nil,
		).Once()

		logger := utils.NewMockLogger()
		service := services.NewRecipeStepService(repo, logger)

		sCtx.WithNewParameters("ctx", ctx, "request", expStep.ID)
		Ingredient, err := service.GetById(ctx, expStep.ID)

		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(Ingredient, expStep)
	})
}

func (suite *RecipeStepSuite) TestRecipeStepService_GetById2(t provider.T) {
	t.Title("[RecipeStep get by id] repo error")
	t.Tags("recipe_step", "get_by_id")
	t.Parallel()

	t.WithNewStep("repo error", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		id := uuid.New()

		repo := mocks.NewIRecipeStepRepository(t)
		repo.On("GetById", ctx, id).Return(
			nil, fmt.Errorf("repo error"),
		).Once()

		logger := utils.NewMockLogger()
		service := services.NewRecipeStepService(repo, logger)

		sCtx.WithNewParameters("ctx", ctx, "request", id)
		comment, err := service.GetById(ctx, id)

		sCtx.Assert().Error(err)
		sCtx.Assert().Nil(comment)
	})
}

func (suite *RecipeStepSuite) TestRecipeStepService_Update1(t provider.T) {
	t.Title("[RecipeStep update] successfully updated")
	t.Tags("recipe_step", "update")
	t.Parallel()

	t.WithNewStep("successfully updated", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		expStep := utils.NewRecipeStepBuilder().
			WithId(uuid.New()).
			WithRecipeId(uuid.New()).
			WithName("name").
			WithDescription("description").
			WithStepNum(1).
			ToDto()

		repo := mocks.NewIRecipeStepRepository(t)
		repo.On("Update", ctx, expStep).Return(
			nil,
		).Once()

		logger := utils.NewMockLogger()
		service := services.NewRecipeStepService(repo, logger)

		sCtx.WithNewParameters("ctx", ctx, "request", expStep)
		err := service.Update(ctx, expStep)

		sCtx.Assert().NoError(err)
	})
}

func (suite *RecipeStepSuite) TestRecipeStepService_Update2(t provider.T) {
	t.Title("[RecipeStep update] repo error")
	t.Tags("recipe_step", "update")
	t.Parallel()

	t.WithNewStep("repo error", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		expStep := utils.NewRecipeStepBuilder().
			WithId(uuid.New()).
			WithRecipeId(uuid.New()).
			WithName("name").
			WithDescription("description").
			WithStepNum(1).
			ToDto()

		repo := mocks.NewIRecipeStepRepository(t)
		repo.On("Update", ctx, expStep).Return(
			fmt.Errorf("repo error"),
		).Once()

		logger := utils.NewMockLogger()
		service := services.NewRecipeStepService(repo, logger)

		sCtx.WithNewParameters("ctx", ctx, "request", expStep)
		err := service.Update(ctx, expStep)

		sCtx.Assert().Error(err)
	})
}

func (suite *RecipeStepSuite) TestRecipeStepService_DeleteAllByRecipeID1(t provider.T) {
	t.Title("[RecipeStep delete all by recipe id] successfully deleted")
	t.Tags("recipe_step", "delete_all_by_recipe_id")
	t.Parallel()

	t.WithNewStep("successfully deleted", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		recipeId := uuid.New()

		repo := mocks.NewIRecipeStepRepository(t)
		repo.On("DeleteAllByRecipeID", ctx, recipeId).Return(
			nil,
		).Once()

		logger := utils.NewMockLogger()
		service := services.NewRecipeStepService(repo, logger)

		sCtx.WithNewParameters("ctx", ctx, "request", recipeId)
		err := service.DeleteAllByRecipeID(ctx, recipeId)

		sCtx.Assert().NoError(err)
	})
}

func (suite *RecipeStepSuite) TestRecipeStepService_DeleteAllByRecipeID2(t provider.T) {
	t.Title("[RecipeStep delete all by recipe id] repo error")
	t.Tags("recipe_step", "delete_all_by_recipe_id")
	t.Parallel()

	t.WithNewStep("repo error", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		recipeId := uuid.New()

		repo := mocks.NewIRecipeStepRepository(t)
		repo.On("DeleteAllByRecipeID", ctx, recipeId).Return(
			fmt.Errorf("repo error"),
		).Once()

		logger := utils.NewMockLogger()
		service := services.NewRecipeStepService(repo, logger)

		sCtx.WithNewParameters("ctx", ctx, "request", recipeId)
		err := service.DeleteAllByRecipeID(ctx, recipeId)

		sCtx.Assert().Error(err)
	})
}
