//go:build unit_test

package unit_tests

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
	"github.com/stretchr/testify/mock"
	"ppo/domain"
	"ppo/mocks"
	"ppo/services"
	"ppo/tests/utils"
)

type RecipeStepInteractorSuite struct {
	suite.Suite
}

func (suite *RecipeStepInteractorSuite) TestRecipeStepInteractorService_Create1(t provider.T) {
	t.Title("[Recipe step interactor create] successfully created")
	t.Tags("recipe_step_interactor", "create")
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

		recipeStepService := mocks.NewIRecipeStepService(t)
		recipeStepService.On("Create", ctx, recipeStep).Return(
			nil,
		)

		validator := mocks.NewIValidatorService(t)
		validator.On("Verify", ctx, mock.Anything).Return(
			nil,
		).Maybe()

		service := services.NewRecipeStepInteractor(recipeStepService, []domain.IValidatorService{
			validator,
		})

		sCtx.WithNewParameters("ctx", ctx, "request", recipeStep)
		err := service.Create(ctx, recipeStep)

		sCtx.Assert().NoError(err)
	})
}

func (suite *RecipeStepInteractorSuite) TestRecipeStepInteractorService_Create2(t provider.T) {
	t.Title("[Recipe step interactor create] service error")
	t.Tags("recipe_step_interactor", "create")
	t.Parallel()

	t.WithNewStep("service error", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		recipeStep := utils.NewRecipeStepBuilder().
			WithId(uuid.New()).
			WithRecipeId(uuid.New()).
			WithName("name").
			WithDescription("description").
			WithStepNum(1).
			ToDto()

		recipeStepService := mocks.NewIRecipeStepService(t)
		recipeStepService.On("Create", ctx, recipeStep).Return(
			fmt.Errorf("service error"),
		)

		validator := mocks.NewIValidatorService(t)
		validator.On("Verify", ctx, mock.Anything).Return(
			nil,
		).Maybe()

		service := services.NewRecipeStepInteractor(recipeStepService, []domain.IValidatorService{
			validator,
		})

		sCtx.WithNewParameters("ctx", ctx, "request", recipeStep)
		err := service.Create(ctx, recipeStep)

		sCtx.Assert().Error(err)
	})
}

func (suite *RecipeStepInteractorSuite) TestRecipeStepInteractorService_DeleteById1(t provider.T) {
	t.Title("[Recipe step interactor create] successfully deleted")
	t.Tags("recipe_step_interactor", "delete_by_id")
	t.Parallel()

	t.WithNewStep("successfully deleted", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		id := uuid.New()

		recipeStepService := mocks.NewIRecipeStepService(t)
		recipeStepService.On("DeleteById", ctx, id).Return(
			nil,
		)

		validator := mocks.NewIValidatorService(t)
		validator.On("Verify", ctx, mock.Anything).Return(
			nil,
		).Maybe()

		service := services.NewRecipeStepInteractor(recipeStepService, []domain.IValidatorService{
			validator,
		})

		sCtx.WithNewParameters("ctx", ctx, "request", id)
		err := service.DeleteById(ctx, id)

		sCtx.Assert().NoError(err)
	})
}

func (suite *RecipeStepInteractorSuite) TestRecipeStepInteractorService_DeleteById2(t provider.T) {
	t.Title("[Recipe step interactor create] service error")
	t.Tags("recipe_step_interactor", "delete_by_id")
	t.Parallel()

	t.WithNewStep("service error", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		id := uuid.New()

		recipeStepService := mocks.NewIRecipeStepService(t)
		recipeStepService.On("DeleteById", ctx, id).Return(
			fmt.Errorf("service error"),
		)

		validator := mocks.NewIValidatorService(t)
		validator.On("Verify", ctx, mock.Anything).Return(
			nil,
		).Maybe()

		service := services.NewRecipeStepInteractor(recipeStepService, []domain.IValidatorService{
			validator,
		})

		sCtx.WithNewParameters("ctx", ctx, "request", id)
		err := service.DeleteById(ctx, id)

		sCtx.Assert().Error(err)
	})
}

func (suite *RecipeStepInteractorSuite) TestRecipeStepInteractorService_DeleteAllByRecipeId1(t provider.T) {
	t.Title("[Recipe step interactor create] successfully deleted")
	t.Tags("recipe_step_interactor", "delete_all_by_recipe_id")
	t.Parallel()

	t.WithNewStep("successfully deleted", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		id := uuid.New()

		recipeStepService := mocks.NewIRecipeStepService(t)
		recipeStepService.On("DeleteAllByRecipeID", ctx, id).Return(
			nil,
		)

		validator := mocks.NewIValidatorService(t)
		validator.On("Verify", ctx, mock.Anything).Return(
			nil,
		).Maybe()

		service := services.NewRecipeStepInteractor(recipeStepService, []domain.IValidatorService{
			validator,
		})

		sCtx.WithNewParameters("ctx", ctx, "request", id)
		err := service.DeleteAllByRecipeID(ctx, id)

		sCtx.Assert().NoError(err)
	})
}

func (suite *RecipeStepInteractorSuite) TestRecipeStepInteractorService_DeleteAllByRecipeId2(t provider.T) {
	t.Title("[Recipe step interactor create] service error")
	t.Tags("recipe_step_interactor", "delete_all_by_recipe_id")
	t.Parallel()

	t.WithNewStep("service error", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		id := uuid.New()

		recipeStepService := mocks.NewIRecipeStepService(t)
		recipeStepService.On("DeleteAllByRecipeID", ctx, id).Return(
			fmt.Errorf("service error"),
		)

		validator := mocks.NewIValidatorService(t)
		validator.On("Verify", ctx, mock.Anything).Return(
			nil,
		).Maybe()

		service := services.NewRecipeStepInteractor(recipeStepService, []domain.IValidatorService{
			validator,
		})

		sCtx.WithNewParameters("ctx", ctx, "request", id)
		err := service.DeleteAllByRecipeID(ctx, id)

		sCtx.Assert().Error(err)
	})
}

func (suite *RecipeStepInteractorSuite) TestRecipeStepInteractorService_GetById1(t provider.T) {
	t.Title("[Recipe step interactor create] successful get")
	t.Tags("recipe_step_interactor", "get_by_id")
	t.Parallel()

	t.WithNewStep("successful get", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		recipeStep := utils.NewRecipeStepBuilder().
			WithId(uuid.New()).
			WithRecipeId(uuid.New()).
			WithName("name").
			WithDescription("description").
			WithStepNum(1).
			ToDto()

		recipeStepService := mocks.NewIRecipeStepService(t)
		recipeStepService.On("GetById", ctx, recipeStep.ID).Return(
			recipeStep, nil,
		)

		validator := mocks.NewIValidatorService(t)
		validator.On("Verify", ctx, mock.Anything).Return(
			nil,
		).Maybe()

		service := services.NewRecipeStepInteractor(recipeStepService, []domain.IValidatorService{
			validator,
		})

		sCtx.WithNewParameters("ctx", ctx, "request", recipeStep.ID)
		step, err := service.GetById(ctx, recipeStep.ID)

		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(step, recipeStep)
	})
}

func (suite *RecipeStepInteractorSuite) TestRecipeStepInteractorService_GetById2(t provider.T) {
	t.Title("[Recipe step interactor create] service error")
	t.Tags("recipe_step_interactor", "get_by_id")
	t.Parallel()

	t.WithNewStep("service error", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		id := uuid.New()

		recipeStepService := mocks.NewIRecipeStepService(t)
		recipeStepService.On("GetById", ctx, id).Return(
			nil, fmt.Errorf("service error"),
		)

		validator := mocks.NewIValidatorService(t)
		validator.On("Verify", ctx, mock.Anything).Return(
			nil,
		).Maybe()

		service := services.NewRecipeStepInteractor(recipeStepService, []domain.IValidatorService{
			validator,
		})

		sCtx.WithNewParameters("ctx", ctx, "request", id)
		step, err := service.GetById(ctx, id)

		sCtx.Assert().Error(err)
		sCtx.Assert().Nil(step)
	})
}

func (suite *RecipeStepInteractorSuite) TestRecipeStepInteractorService_Update1(t provider.T) {
	t.Title("[Recipe step interactor create] successfully updated")
	t.Tags("recipe_step_interactor", "update")
	t.Parallel()

	t.WithNewStep("successfully updated", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		recipeStep := utils.NewRecipeStepBuilder().
			WithId(uuid.New()).
			WithRecipeId(uuid.New()).
			WithName("name").
			WithDescription("description").
			WithStepNum(1).
			ToDto()

		recipeStepService := mocks.NewIRecipeStepService(t)
		recipeStepService.On("Update", ctx, recipeStep).Return(
			nil,
		)

		validator := mocks.NewIValidatorService(t)
		validator.On("Verify", ctx, mock.Anything).Return(
			nil,
		).Maybe()

		service := services.NewRecipeStepInteractor(recipeStepService, []domain.IValidatorService{
			validator,
		})

		sCtx.WithNewParameters("ctx", ctx, "request", recipeStep.ID)
		err := service.Update(ctx, recipeStep)

		sCtx.Assert().NoError(err)
	})
}

func (suite *RecipeStepInteractorSuite) TestRecipeStepInteractorService_Update2(t provider.T) {
	t.Title("[Recipe step interactor create] service error")
	t.Tags("recipe_step_interactor", "update")
	t.Parallel()

	t.WithNewStep("service error", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		recipeStep := utils.NewRecipeStepBuilder().
			WithId(uuid.New()).
			WithRecipeId(uuid.New()).
			WithName("name").
			WithDescription("description").
			WithStepNum(1).
			ToDto()

		recipeStepService := mocks.NewIRecipeStepService(t)
		recipeStepService.On("Update", ctx, recipeStep).Return(
			fmt.Errorf("service error"),
		)

		validator := mocks.NewIValidatorService(t)
		validator.On("Verify", ctx, mock.Anything).Return(
			nil,
		).Maybe()

		service := services.NewRecipeStepInteractor(recipeStepService, []domain.IValidatorService{
			validator,
		})

		sCtx.WithNewParameters("ctx", ctx, "request", recipeStep.ID)
		err := service.Update(ctx, recipeStep)

		sCtx.Assert().Error(err)
	})
}

func (suite *RecipeStepInteractorSuite) TestRecipeStepInteractorService_GetAllByRecipeId1(t provider.T) {
	t.Title("[Recipe step interactor create] successful get")
	t.Tags("recipe_step_interactor", "get_all_by_recipe_id")
	t.Parallel()

	t.WithNewStep("successful get", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		recipeId := uuid.New()
		expSteps := []*domain.RecipeStep{}

		recipeStepService := mocks.NewIRecipeStepService(t)
		recipeStepService.On("GetAllByRecipeID", ctx, recipeId).Return(
			expSteps, nil,
		)

		validator := mocks.NewIValidatorService(t)
		validator.On("Verify", ctx, mock.Anything).Return(
			nil,
		).Maybe()

		service := services.NewRecipeStepInteractor(recipeStepService, []domain.IValidatorService{
			validator,
		})

		sCtx.WithNewParameters("ctx", ctx, "request", recipeId)
		steps, err := service.GetAllByRecipeID(ctx, recipeId)

		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(steps, expSteps)
	})
}

func (suite *RecipeStepInteractorSuite) TestRecipeStepInteractorService_GetAllByRecipeId2(t provider.T) {
	t.Title("[Recipe step interactor create] service error")
	t.Tags("recipe_step_interactor", "get_all_by_recipe_id")
	t.Parallel()

	t.WithNewStep("service error", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		recipeId := uuid.New()

		recipeStepService := mocks.NewIRecipeStepService(t)
		recipeStepService.On("GetAllByRecipeID", ctx, recipeId).Return(
			nil, fmt.Errorf("service error"),
		)

		validator := mocks.NewIValidatorService(t)
		validator.On("Verify", ctx, mock.Anything).Return(
			nil,
		).Maybe()

		service := services.NewRecipeStepInteractor(recipeStepService, []domain.IValidatorService{
			validator,
		})

		sCtx.WithNewParameters("ctx", ctx, "request", recipeId)
		steps, err := service.GetAllByRecipeID(ctx, recipeId)

		sCtx.Assert().Error(err)
		sCtx.Assert().Nil(steps)
	})
}
