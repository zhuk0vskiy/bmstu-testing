//go:build integration_test

package integration_tests

import (
	"context"
	"github.com/google/uuid"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
	"ppo/domain"
	"ppo/internal/storage/postgres"
	"ppo/tests/utils"
)

type ITRecipeStepSuite struct {
	suite.Suite
	repo domain.IRecipeStepRepository
}

func (s *ITRecipeStepSuite) BeforeAll(t provider.T) {
	t.Title("init test repository")
	s.repo = postgres.NewRecipeStepRepository(testDbInstance)
	t.Tags("recipe_step")
}

func (s *ITRecipeStepSuite) Test_RecipeStepRepo_Create1(t provider.T) {
	t.Title("[Create] successfully created")
	t.Tags("integration_test", "postgres", "create")
	t.Parallel()

	t.WithNewStep("Successfully created", func(sCtx provider.StepCtx) {
		ctx := context.TODO()

		recipeStep := utils.NewRecipeStepBuilder().
			WithId(uuid.New()).
			WithRecipeId(uuid.UUID{3}).
			WithName("first").
			WithDescription("description").
			WithStepNum(1).
			ToDto()

		sCtx.WithNewParameters("ctx", ctx, "request", recipeStep)
		err := s.repo.Create(ctx, recipeStep)

		sCtx.Assert().NoError(err)
	})
}

func (s *ITRecipeStepSuite) Test_RecipeStepRepo_Create2(t provider.T) {
	t.Title("[Create] recipe not found")
	t.Tags("integration_test", "postgres", "create")
	t.Parallel()

	t.WithNewStep("Recipe not found", func(sCtx provider.StepCtx) {
		ctx := context.TODO()

		recipeStep := utils.NewRecipeStepBuilder().
			WithId(uuid.New()).
			WithRecipeId(uuid.UUID{111}).
			WithName("first").
			WithDescription("description").
			WithStepNum(1).
			ToDto()

		sCtx.WithNewParameters("ctx", ctx, "request", recipeStep)
		err := s.repo.Create(ctx, recipeStep)

		sCtx.Assert().Error(err)
	})
}

func (s *ITRecipeStepSuite) Test_RecipeStepRepo_Update1(t provider.T) {
	t.Title("[Update] successfully updated")
	t.Tags("integration_test", "postgres", "update")
	t.Parallel()

	t.WithNewStep("Successfully updated", func(sCtx provider.StepCtx) {
		ctx := context.TODO()

		recipeStep := utils.NewRecipeStepBuilder().
			WithId(uuid.UUID{1}).
			WithRecipeId(uuid.UUID{2}).
			WithName("first").
			WithDescription("first").
			WithStepNum(1).
			ToDto()

		sCtx.WithNewParameters("ctx", ctx, "request", recipeStep)
		err := s.repo.Update(ctx, recipeStep)

		sCtx.Assert().NoError(err)
	})
}

func (s *ITRecipeStepSuite) Test_RecipeStepRepo_Update2(t provider.T) {
	t.Title("[Update] step num greater then max")
	t.Tags("integration_test", "postgres", "update")
	t.Parallel()

	t.WithNewStep("Step num greater then max", func(sCtx provider.StepCtx) {
		ctx := context.TODO()

		recipeStep := utils.NewRecipeStepBuilder().
			WithId(uuid.UUID{1}).
			WithRecipeId(uuid.UUID{2}).
			WithName("first").
			WithDescription("first").
			WithStepNum(2).
			ToDto()

		sCtx.WithNewParameters("ctx", ctx, "request", recipeStep)
		err := s.repo.Update(ctx, recipeStep)

		sCtx.Assert().Error(err)
	})
}

func (s *ITRecipeStepSuite) Test_RecipeStepRepo_GetAllByRecipeId1(t provider.T) {
	t.Title("[Get all by recipe id] success")
	t.Tags("integration_test", "postgres", "get_all_by_recipe_id")
	t.Parallel()

	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		ctx := context.TODO()

		recipeId := uuid.UUID{4}
		expSteps := []*domain.RecipeStep{
			{
				ID:          uuid.UUID{8},
				Name:        "first",
				Description: "first",
				RecipeID:    uuid.UUID{4},
				StepNum:     1,
			},
			{
				ID:          uuid.UUID{9},
				Name:        "second",
				Description: "second",
				RecipeID:    uuid.UUID{4},
				StepNum:     2,
			},
			{
				ID:          uuid.UUID{10},
				Name:        "third",
				Description: "third",
				RecipeID:    uuid.UUID{4},
				StepNum:     3,
			},
		}

		sCtx.WithNewParameters("ctx", ctx, "request", recipeId)
		steps, err := s.repo.GetAllByRecipeID(ctx, recipeId)

		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(expSteps, steps)
	})
}

func (s *ITRecipeStepSuite) Test_RecipeStepRepo_GetAllByRecipeId2(t provider.T) {
	t.Title("[Get all by recipe id] recipe not found")
	t.Tags("integration_test", "postgres", "get_all_by_recipe_id")
	t.Parallel()

	t.WithNewStep("Recipe not found", func(sCtx provider.StepCtx) {
		ctx := context.TODO()

		recipeId := uuid.UUID{111}
		expSteps := []*domain.RecipeStep{}

		sCtx.WithNewParameters("ctx", ctx, "request", recipeId)
		steps, err := s.repo.GetAllByRecipeID(ctx, recipeId)

		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(steps, expSteps)
	})
}

func (s *ITRecipeStepSuite) Test_RecipeStepRepo_GetById1(t provider.T) {
	t.Title("[Get by id] success")
	t.Tags("integration_test", "postgres", "get_by_id")
	t.Parallel()

	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		ctx := context.TODO()

		recipeStep := utils.NewRecipeStepBuilder().
			WithId(uuid.UUID{8}).
			WithRecipeId(uuid.UUID{4}).
			WithName("first").
			WithDescription("first").
			WithStepNum(1).
			ToDto()

		sCtx.WithNewParameters("ctx", ctx, "request", recipeStep.ID)
		step, err := s.repo.GetById(ctx, recipeStep.ID)

		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(step, recipeStep)
	})
}

func (s *ITRecipeStepSuite) Test_RecipeStepRepo_GetById2(t provider.T) {
	t.Title("[Get by id] step not found")
	t.Tags("integration_test", "postgres", "get_by_id")
	t.Parallel()

	t.WithNewStep("Step not found", func(sCtx provider.StepCtx) {
		ctx := context.TODO()

		id := uuid.Nil

		sCtx.WithNewParameters("ctx", ctx, "request", id)
		step, err := s.repo.GetById(ctx, id)

		sCtx.Assert().Error(err)
		sCtx.Assert().Nil(step)
	})
}

func (s *ITRecipeStepSuite) Test_RecipeStepRepo_DeleteALlByRecipeId1(t provider.T) {
	t.Title("[Delete all by recipe id] success")
	t.Tags("integration_test", "postgres", "delete_all_by_recipe_id")
	t.Parallel()

	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		ctx := context.TODO()

		id := uuid.UUID{3}

		sCtx.WithNewParameters("ctx", ctx, "request", id)
		err := s.repo.DeleteAllByRecipeID(ctx, id)

		sCtx.Assert().NoError(err)
	})
}

func (s *ITRecipeStepSuite) Test_RecipeStepRepo_DeleteById1(t provider.T) {
	t.Title("[Delete by id] success")
	t.Tags("integration_test", "postgres", "delete_by_id")
	t.Parallel()

	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		ctx := context.TODO()

		id := uuid.UUID{6}

		sCtx.WithNewParameters("ctx", ctx, "request", id)
		err := s.repo.DeleteById(ctx, id)

		sCtx.Assert().NoError(err)
	})
}
