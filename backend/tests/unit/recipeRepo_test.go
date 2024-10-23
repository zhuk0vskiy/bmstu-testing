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
	"ppo/internal/config"
	"ppo/internal/storage/postgres"
	"ppo/services/dto"
	"ppo/tests/utils"
)

type RecipeRepoSuite struct {
	suite.Suite
}

func (suite *RecipeRepoSuite) TestRecipeRepo_Create1(t provider.T) {
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

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatal(err)
		}
		defer mock.Close()
		mock.ExpectQuery("insert").
			WithArgs(
				recipe.SaladID,
				recipe.Status,
				recipe.NumberOfServings,
				recipe.TimeToCook,
			).
			WillReturnRows(pgxmock.NewRows([]string{"id"}).AddRow(recipe.ID))

		repo := postgres.NewRecipeRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "request", recipe)
		id, err := repo.Create(ctx, recipe)

		sCtx.Assert().NoError(err)
		sCtx.Assert().NotEqual(id, uuid.Nil)
	})
}

func (suite *RecipeRepoSuite) TestRecipeRepo_Create2(t provider.T) {
	t.Title("[Recipe create] not inserted")
	t.Tags("recipe", "create")
	t.Parallel()

	t.WithNewStep("not inserted", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		recipe := utils.NewRecipeBuilder().
			WithId(uuid.New()).
			WithSaladId(uuid.New()).
			WithStatus(1).
			WithNumberOfServings(1).
			WithTimeToCook(1).
			WithRating(5.0).
			ToDto()

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatal(err)
		}
		defer mock.Close()
		mock.ExpectQuery("insert").
			WithArgs(
				recipe.SaladID,
				recipe.Status,
				recipe.NumberOfServings,
				recipe.TimeToCook,
			).
			WillReturnError(fmt.Errorf("insert error"))

		repo := postgres.NewRecipeRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "request", recipe)
		id, err := repo.Create(ctx, recipe)

		sCtx.Assert().Error(err)
		sCtx.Assert().Equal(id, uuid.Nil)
	})
}

func (suite *RecipeRepoSuite) TestRecipeRepo_GetById1(t provider.T) {
	t.Title("[Recipe get by id] successful get")
	t.Tags("recipe", "get_by_id")
	t.Parallel()

	t.WithNewStep("successful get", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		recipe := utils.NewRecipeBuilder().
			WithId(uuid.New()).
			WithSaladId(uuid.New()).
			WithStatus(1).
			WithNumberOfServings(1).
			WithTimeToCook(1).
			WithRating(5.0).
			ToDto()

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatal(err)
		}
		defer mock.Close()
		mock.ExpectQuery("select").
			WithArgs(
				recipe.ID,
			).
			WillReturnRows(
				pgxmock.NewRows([]string{"ID", "SaladId", "Status", "NumberOfServings", "TimeToCook", "Rating"}).
					AddRow(recipe.ID, recipe.SaladID, recipe.Status, recipe.NumberOfServings, recipe.TimeToCook, recipe.Rating),
			)

		repo := postgres.NewRecipeRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "request", recipe.ID)
		rRecipe, err := repo.GetById(ctx, recipe.ID)

		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(recipe, rRecipe)
	})
}

func (suite *RecipeRepoSuite) TestRecipeRepo_GetById2(t provider.T) {
	t.Title("[Recipe get by id] failed to get")
	t.Tags("recipe", "get_by_id")
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
				pgxmock.NewRows([]string{"ID", "SaladId", "Status", "NumberOfServings", "TimeToCook", "Rating"}),
			)

		repo := postgres.NewRecipeRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "request", id)
		rRecipe, err := repo.GetById(ctx, id)

		sCtx.Assert().Error(err)
		sCtx.Assert().Nil(rRecipe)
	})
}

func (suite *RecipeRepoSuite) TestRecipeRepo_GetBySaladId1(t provider.T) {
	t.Title("[Recipe get by salad id] successful get")
	t.Tags("recipe", "get_by_id")
	t.Parallel()

	t.WithNewStep("successful get", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		recipe := utils.NewRecipeBuilder().
			WithId(uuid.New()).
			WithSaladId(uuid.New()).
			WithStatus(1).
			WithNumberOfServings(1).
			WithTimeToCook(1).
			WithRating(5.0).
			ToDto()

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatal(err)
		}
		defer mock.Close()
		mock.ExpectQuery("select").
			WithArgs(
				recipe.SaladID,
			).
			WillReturnRows(
				pgxmock.NewRows([]string{"ID", "SaladId", "Status", "NumberOfServings", "TimeToCook", "Rating"}).
					AddRow(recipe.ID, recipe.SaladID, recipe.Status, recipe.NumberOfServings, recipe.TimeToCook, recipe.Rating),
			)

		repo := postgres.NewRecipeRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "request", recipe.SaladID)
		rRecipe, err := repo.GetBySaladId(ctx, recipe.SaladID)

		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(recipe, rRecipe)
	})
}

func (suite *RecipeRepoSuite) TestRecipeRepo_GetBySaladId2(t provider.T) {
	t.Title("[Recipe get by salad id] failed to get")
	t.Tags("recipe", "get_by_id")
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
				pgxmock.NewRows([]string{"ID", "SaladId", "Status", "NumberOfServings", "TimeToCook", "Rating"}),
			)

		repo := postgres.NewRecipeRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "request", id)
		rRecipe, err := repo.GetBySaladId(ctx, id)

		sCtx.Assert().Error(err)
		sCtx.Assert().Nil(rRecipe)
	})
}

func (suite *RecipeRepoSuite) TestRecipeRepo_GetAll1(t provider.T) {
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

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatal(err)
		}
		defer mock.Close()
		mock.ExpectQuery("select").
			WithArgs(
				"ingredientId",
				"typeId",
				filter.MinRate,
				config.PageSize*(page-1),
				config.PageSize,
			).
			WillReturnRows(
				pgxmock.NewRows([]string{"ID", "SaladId", "Status", "NumberOfServings", "TimeToCook", "Rating"}).
					AddRow(expRecipes[0].ID, expRecipes[0].SaladID, expRecipes[0].Status, expRecipes[0].NumberOfServings, expRecipes[0].TimeToCook, expRecipes[0].Rating).
					AddRow(expRecipes[1].ID, expRecipes[1].SaladID, expRecipes[1].Status, expRecipes[1].NumberOfServings, expRecipes[1].TimeToCook, expRecipes[1].Rating),
			)

		repo := postgres.NewRecipeRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "request", filter, page)
		recipes, err := repo.GetAll(ctx, filter, page)

		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(recipes, expRecipes)
	})
}

func (suite *RecipeRepoSuite) TestRecipeRepo_GetAll2(t provider.T) {
	t.Title("[Recipe get all] sql error")
	t.Tags("recipe", "get_all")
	t.Parallel()

	t.WithNewStep("sql error", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		page := 1
		filter := &dto.RecipeFilter{
			AvailableIngredients: nil,
			MinRate:              0,
			SaladTypes:           nil,
			Status:               1,
		}

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatal(err)
		}
		defer mock.Close()
		mock.ExpectQuery("select").
			WithArgs(
				"ingredientId",
				"typeId",
				filter.MinRate,
				config.PageSize*(page-1),
				config.PageSize,
			).
			WillReturnError(fmt.Errorf("sql error"))

		repo := postgres.NewRecipeRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "request", filter, page)
		rMeasurements, err := repo.GetAll(ctx, filter, page)

		sCtx.Assert().Error(err)
		sCtx.Assert().Nil(rMeasurements)
	})
}

func (suite *RecipeRepoSuite) TestRecipeRepo_Update1(t provider.T) {
	t.Title("[Recipe update] successfully updated")
	t.Tags("recipe", "update")
	t.Parallel()

	t.WithNewStep("successfully updated", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		recipe := utils.NewRecipeBuilder().
			WithId(uuid.New()).
			WithSaladId(uuid.New()).
			WithStatus(1).
			WithNumberOfServings(1).
			WithTimeToCook(1).
			WithRating(5.0).
			ToDto()

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatal(err)
		}
		defer mock.Close()
		mock.ExpectExec("update").
			WithArgs(
				recipe.Status,
				recipe.NumberOfServings,
				recipe.TimeToCook,
				recipe.Rating,
				recipe.ID,
			).
			WillReturnResult(pgxmock.NewResult("update", 1))

		repo := postgres.NewRecipeRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "request", recipe)
		err = repo.Update(ctx, recipe)

		sCtx.Assert().NoError(err)
	})
}

func (suite *RecipeRepoSuite) TestRecipeRepo_Update2(t provider.T) {
	t.Title("[Recipe update] failed to update")
	t.Tags("recipe", "update")
	t.Parallel()

	t.WithNewStep("failed to updated", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		recipe := utils.NewRecipeBuilder().
			WithId(uuid.New()).
			WithSaladId(uuid.New()).
			WithStatus(1).
			WithNumberOfServings(1).
			WithTimeToCook(1).
			WithRating(5.0).
			ToDto()

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatal(err)
		}
		defer mock.Close()
		mock.ExpectExec("update").
			WithArgs(
				recipe.Status,
				recipe.NumberOfServings,
				recipe.TimeToCook,
				recipe.Rating,
				recipe.ID,
			).
			WillReturnError(fmt.Errorf("sql error"))

		repo := postgres.NewRecipeRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "request", recipe)
		err = repo.Update(ctx, recipe)

		sCtx.Assert().Error(err)
	})
}

func (suite *RecipeRepoSuite) TestRecipeRepo_DeleteById1(t provider.T) {
	t.Title("[Recipe delete by id] successfully deleted")
	t.Tags("recipe", "delete_by_id")
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

		repo := postgres.NewRecipeRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "request", id)
		err = repo.DeleteById(ctx, id)

		sCtx.Assert().NoError(err)
	})
}

func (suite *RecipeRepoSuite) TestRecipeRepo_DeleteById2(t provider.T) {
	t.Title("[Recipe delete by id] sql error")
	t.Tags("recipe", "delete_by_id")
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

		repo := postgres.NewRecipeRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "request", id)
		err = repo.DeleteById(ctx, id)

		sCtx.Assert().Error(err)
	})
}
