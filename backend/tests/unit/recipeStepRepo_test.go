//go:build unit_test

package unit_tests

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
	"github.com/pashagolub/pgxmock/v4"
	"ppo/domain"
	"ppo/internal/storage/postgres"
	"ppo/tests/utils"
)

type RecipeStepRepoSuite struct {
	suite.Suite
}

func (suite *RecipeStepRepoSuite) TestRecipeStepRepo_Create1(t provider.T) {
	t.Title("[Recipe step create] successfully created")
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

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatal(err)
		}
		defer mock.Close()
		mock.ExpectExec("insert").
			WithArgs(
				recipeStep.Name,
				recipeStep.Description,
				recipeStep.RecipeID,
			).
			WillReturnResult(pgxmock.NewResult("insert", 1))

		repo := postgres.NewRecipeStepRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "request", recipeStep)
		err = repo.Create(ctx, recipeStep)

		sCtx.Assert().NoError(err)
	})
}

func (suite *RecipeStepRepoSuite) TestRecipeStepRepo_Create2(t provider.T) {
	t.Title("[Recipe step create] not inserted")
	t.Tags("recipe_step", "create")
	t.Parallel()

	t.WithNewStep("not inserted", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		recipeStep := utils.NewRecipeStepBuilder().
			WithId(uuid.New()).
			WithRecipeId(uuid.New()).
			WithName("name").
			WithDescription("description").
			WithStepNum(1).
			ToDto()

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatal(err)
		}
		defer mock.Close()
		mock.ExpectQuery("insert").
			WithArgs(
				recipeStep.Name,
				recipeStep.Description,
				recipeStep.RecipeID,
			).
			WillReturnError(fmt.Errorf("insert error"))

		repo := postgres.NewRecipeStepRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "request", recipeStep)
		err = repo.Create(ctx, recipeStep)

		sCtx.Assert().Error(err)
	})
}

func (suite *RecipeStepRepoSuite) TestRecipeStepRepo_GetById1(t provider.T) {
	t.Title("[Recipe step get by id] successful get")
	t.Tags("recipe_step", "get_by_id")
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

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatal(err)
		}
		defer mock.Close()
		mock.ExpectQuery("select").
			WithArgs(
				recipeStep.ID,
			).
			WillReturnRows(
				pgxmock.NewRows([]string{"ID", "Name", "Description", "RecipeId", "StepNum"}).
					AddRow(recipeStep.ID, recipeStep.Name, recipeStep.Description, recipeStep.RecipeID, recipeStep.StepNum),
			)

		repo := postgres.NewRecipeStepRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "request", recipeStep.ID)
		rRecipe, err := repo.GetById(ctx, recipeStep.ID)

		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(recipeStep, rRecipe)
	})
}

func (suite *RecipeStepRepoSuite) TestRecipeStepRepo_GetById2(t provider.T) {
	t.Title("[Recipe step get by id] failed to get")
	t.Tags("recipe_step", "get_by_id")
	t.Parallel()

	t.WithNewStep("failed to get", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		id := uuid.New()

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatal(err)
		}
		defer mock.Close()
		mock.ExpectQuery("select").
			WithArgs(
				id,
			).
			WillReturnRows(
				pgxmock.NewRows([]string{"ID", "Name", "Description", "RecipeId", "StepNum"}),
			)

		repo := postgres.NewRecipeStepRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "request", id)
		rRecipe, err := repo.GetById(ctx, id)

		sCtx.Assert().Error(err)
		sCtx.Assert().Nil(rRecipe)
	})
}

func (suite *RecipeStepRepoSuite) TestRecipeStepRepo_GetAllByRecipeId1(t provider.T) {
	t.Title("[Recipe step get all] successful get")
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

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatal(err)
		}
		defer mock.Close()
		mock.ExpectQuery("select").
			WithArgs(
				recipeId,
			).
			WillReturnRows(
				pgxmock.NewRows([]string{"ID", "Name", "Description", "RecipeId", "StepNum"}).
					AddRow(expSteps[0].ID, expSteps[0].Name, expSteps[0].Description, expSteps[0].RecipeID, expSteps[0].StepNum).
					AddRow(expSteps[1].ID, expSteps[1].Name, expSteps[1].Description, expSteps[1].RecipeID, expSteps[1].StepNum),
			)

		repo := postgres.NewRecipeStepRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "request", recipeId)
		steps, err := repo.GetAllByRecipeID(ctx, recipeId)

		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(steps, expSteps)
	})
}

func (suite *RecipeStepRepoSuite) TestRecipeStepRepo_GetAllByRecipeId2(t provider.T) {
	t.Title("[Recipe step get all] sql error")
	t.Tags("recipe_step", "get_all_by_recipe_id")
	t.Parallel()

	t.WithNewStep("sql error", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		recipeId := uuid.New()

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatal(err)
		}
		defer mock.Close()
		mock.ExpectQuery("select").
			WithArgs(
				recipeId,
			).
			WillReturnError(fmt.Errorf("sql error"))

		repo := postgres.NewRecipeStepRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "request", recipeId)
		rMeasurements, err := repo.GetAllByRecipeID(ctx, recipeId)

		sCtx.Assert().Error(err)
		sCtx.Assert().Nil(rMeasurements)
	})
}

func (suite *RecipeStepRepoSuite) TestRecipeStepRepo_Update1(t provider.T) {
	t.Title("[Recipe step update] successfully updated")
	t.Tags("recipe_step", "update")
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

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatal(err)
		}
		defer mock.Close()
		mock.ExpectBegin()
		mock.ExpectQuery("select").
			WithArgs(
				recipeStep.RecipeID,
				recipeStep.ID,
			).
			WillReturnRows(pgxmock.NewRows([]string{"maxStepNum", "prevStepNum"}).
				AddRow(1, 1),
			)

		mock.ExpectExec("update").
			WithArgs(
				recipeStep.RecipeID,
				1,
				recipeStep.StepNum,
				-1,
			).
			WillReturnResult(pgxmock.NewResult("update", 1))

		mock.ExpectExec("update").
			WithArgs(
				recipeStep.Name,
				recipeStep.Description,
				recipeStep.StepNum,
				recipeStep.ID,
			).
			WillReturnResult(pgxmock.NewResult("update", 1))
		mock.ExpectCommit()

		repo := postgres.NewRecipeStepRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "request", recipeStep)
		err = repo.Update(ctx, recipeStep)

		sCtx.Assert().NoError(err)
	})
}

func (suite *RecipeStepRepoSuite) TestRecipeStepRepo_Update2(t provider.T) {
	t.Title("[Recipe step update] failed to update")
	t.Tags("recipe_step", "update")
	t.Parallel()

	t.WithNewStep("failed to updated", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		recipeStep := utils.NewRecipeStepBuilder().
			WithId(uuid.New()).
			WithRecipeId(uuid.New()).
			WithName("name").
			WithDescription("description").
			WithStepNum(1).
			ToDto()

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatal(err)
		}
		defer mock.Close()
		mock.ExpectQuery("select").
			WithArgs(
				recipeStep.RecipeID,
				recipeStep.ID,
			).
			WillReturnError(fmt.Errorf("sql error"))

		repo := postgres.NewRecipeStepRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "request", recipeStep)
		err = repo.Update(ctx, recipeStep)

		sCtx.Assert().Error(err)
	})
}

func (suite *RecipeStepRepoSuite) TestRecipeStepRepo_DeleteById1(t provider.T) {
	t.Title("[Recipe step delete by id] successfully deleted")
	t.Tags("recipe_step", "delete_by_id")
	t.Parallel()

	t.WithNewStep("successfully deleted", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		id := uuid.New()
		stepNum := 1
		recipeId := uuid.New()

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatal(err)
		}
		defer mock.Close()
		mock.ExpectBegin()
		mock.ExpectQuery("select").
			WithArgs(
				id,
			).
			WillReturnRows(pgxmock.NewRows([]string{"recipeId", "stepNum"}).
				AddRow(recipeId, stepNum),
			)
		mock.ExpectExec("update").
			WithArgs(
				recipeId,
				stepNum,
			).
			WillReturnResult(pgxmock.NewResult("update", 1))

		mock.ExpectExec("delete").
			WithArgs(
				id,
			).
			WillReturnResult(pgxmock.NewResult("delete", 1))
		mock.ExpectCommit()

		repo := postgres.NewRecipeStepRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "request", id)
		err = repo.DeleteById(ctx, id)

		sCtx.Assert().NoError(err)
	})
}

func (suite *RecipeStepRepoSuite) TestRecipeStepRepo_DeleteById2(t provider.T) {
	t.Title("[Recipe step delete by id] sql error")
	t.Tags("recipe_step", "delete_by_id")
	t.Parallel()

	t.WithNewStep("sql error", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		id := uuid.New()

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatal(err)
		}
		defer mock.Close()
		mock.ExpectExec("delete").
			WithArgs(
				id,
			).
			WillReturnError(fmt.Errorf("sql error"))

		repo := postgres.NewRecipeStepRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "request", id)
		err = repo.DeleteById(ctx, id)

		sCtx.Assert().Error(err)
	})
}

func (suite *RecipeStepRepoSuite) TestRecipeStepRepo_DeleteAllByRecipeId1(t provider.T) {
	t.Title("[Recipe step delete by id] successfully deleted")
	t.Tags("recipe_step", "delete_all_by_recipe_id")
	t.Parallel()

	t.WithNewStep("successfully deleted", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		id := uuid.New()

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatal(err)
		}
		defer mock.Close()
		mock.ExpectExec("delete").
			WithArgs(
				id,
			).
			WillReturnResult(pgxmock.NewResult("delete", 1))

		repo := postgres.NewRecipeStepRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "request", id)
		err = repo.DeleteAllByRecipeID(ctx, id)

		sCtx.Assert().NoError(err)
	})
}

func (suite *RecipeStepRepoSuite) TestRecipeStepRepo_DeleteAllByRecipeId2(t provider.T) {
	t.Title("[Recipe step delete by id] sql error")
	t.Tags("recipe_step", "delete_all_by_recipe_id")
	t.Parallel()

	t.WithNewStep("sql error", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		id := uuid.New()

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatal(err)
		}
		defer mock.Close()
		mock.ExpectExec("delete").
			WithArgs(
				id,
			).
			WillReturnError(fmt.Errorf("sql error"))

		repo := postgres.NewRecipeStepRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "request", id)
		err = repo.DeleteAllByRecipeID(ctx, id)

		sCtx.Assert().Error(err)
	})
}
