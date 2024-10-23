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

type SaladRepoSuite struct {
	suite.Suite
}

func (suite *SaladRepoSuite) TestRecipeRepo_Create1(t provider.T) {
	t.Title("[Recipe create] successfully created")
	t.Tags("recipe", "create")
	t.Parallel()

	t.WithNewStep("successfully created", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		salad := utils.NewSaladBuilder().
			WithId(uuid.New()).
			WithAuthorId(uuid.New()).
			WithName("salad").
			WithDescription("description").
			ToDto()

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatal(err)
		}
		defer mock.Close()
		mock.ExpectQuery("insert").
			WithArgs(
				salad.Name,
				salad.AuthorID,
				salad.Description,
			).
			WillReturnRows(pgxmock.NewRows([]string{"id"}).AddRow(salad.ID))

		repo := postgres.NewSaladRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "request", salad)
		id, err := repo.Create(ctx, salad)

		sCtx.Assert().NoError(err)
		sCtx.Assert().NotEqual(id, uuid.Nil)
	})
}

func (suite *SaladRepoSuite) TestRecipeRepo_Create2(t provider.T) {
	t.Title("[Recipe create] not inserted")
	t.Tags("recipe", "create")
	t.Parallel()

	t.WithNewStep("not inserted", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		salad := utils.NewSaladBuilder().
			WithId(uuid.New()).
			WithAuthorId(uuid.New()).
			WithName("salad").
			WithDescription("description").
			ToDto()

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatal(err)
		}
		defer mock.Close()
		mock.ExpectQuery("insert").
			WithArgs(
				salad.Name,
				salad.AuthorID,
				salad.Description,
			).
			WillReturnError(fmt.Errorf("insert error"))

		repo := postgres.NewSaladRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "request", salad)
		id, err := repo.Create(ctx, salad)

		sCtx.Assert().Error(err)
		sCtx.Assert().Equal(id, uuid.Nil)
	})
}

func (suite *SaladRepoSuite) TestRecipeRepo_GetById1(t provider.T) {
	t.Title("[Recipe get by id] successful get")
	t.Tags("recipe", "get_by_id")
	t.Parallel()

	t.WithNewStep("successful get", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		salad := utils.NewSaladBuilder().
			WithId(uuid.New()).
			WithAuthorId(uuid.New()).
			WithName("salad").
			WithDescription("description").
			ToDto()

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatal(err)
		}
		defer mock.Close()
		mock.ExpectQuery("select").
			WithArgs(
				salad.ID,
			).
			WillReturnRows(
				pgxmock.NewRows([]string{"ID", "Name", "AuthorId", "Description"}).
					AddRow(salad.ID, salad.Name, salad.AuthorID, salad.Description),
			)

		repo := postgres.NewSaladRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "request", salad.ID)
		rRecipe, err := repo.GetById(ctx, salad.ID)

		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(salad, rRecipe)
	})
}

func (suite *SaladRepoSuite) TestRecipeRepo_GetById2(t provider.T) {
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
				pgxmock.NewRows([]string{"ID", "Name", "AuthorId", "Description"}),
			)

		repo := postgres.NewSaladRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "request", id)
		rRecipe, err := repo.GetById(ctx, id)

		sCtx.Assert().Error(err)
		sCtx.Assert().Nil(rRecipe)
	})
}

func (suite *SaladRepoSuite) TestRecipeRepo_GetAll1(t provider.T) {
	t.Title("[Recipe get all] successful get")
	t.Tags("recipe", "get_all")
	t.Parallel()

	t.WithNewStep("successful get", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		page := 1
		expSalads := []*domain.Salad{
			{
				ID:          uuid.UUID{1},
				AuthorID:    uuid.UUID{11},
				Name:        "ingredient1",
				Description: "description1",
			},
			{
				ID:          uuid.UUID{2},
				AuthorID:    uuid.UUID{22},
				Name:        "ingredient2",
				Description: "description2",
			},
		}
		filter := &dto.RecipeFilter{
			AvailableIngredients: nil,
			MinRate:              0,
			SaladTypes:           nil,
			Status:               4,
		}

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatal(err)
		}
		defer mock.Close()
		mock.ExpectQuery("select").
			WithArgs(
				filter.AvailableIngredients,
				filter.SaladTypes,
				filter.MinRate,
				config.PageSize*(page-1),
				config.PageSize,
				filter.Status,
				1,
				1,
			).
			WillReturnRows(
				pgxmock.NewRows([]string{"ID", "Name", "Description", "AuthorId"}).
					AddRow(expSalads[0].ID, expSalads[0].Name, expSalads[0].Description, expSalads[0].AuthorID).
					AddRow(expSalads[1].ID, expSalads[1].Name, expSalads[1].Description, expSalads[1].AuthorID),
			)

		mock.ExpectQuery("select").
			WillReturnRows(pgxmock.NewRows([]string{"count"}).
				AddRow(1))

		repo := postgres.NewSaladRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "request", filter, page)
		recipes, pages, err := repo.GetAll(ctx, filter, page)

		sCtx.Assert().NoError(err)
		sCtx.Assert().NotZero(pages)
		sCtx.Assert().Equal(expSalads, recipes)
	})
}

func (suite *SaladRepoSuite) TestRecipeRepo_GetAll2(t provider.T) {
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

		repo := postgres.NewSaladRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "request", filter, page)
		rMeasurements, pages, err := repo.GetAll(ctx, filter, page)

		sCtx.Assert().Error(err)
		sCtx.Assert().Zero(pages)
		sCtx.Assert().Nil(rMeasurements)
	})
}

func (suite *SaladRepoSuite) TestRecipeRepo_Update1(t provider.T) {
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

func (suite *SaladRepoSuite) TestRecipeRepo_Update2(t provider.T) {
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

func (suite *SaladRepoSuite) TestRecipeRepo_DeleteById1(t provider.T) {
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

func (suite *SaladRepoSuite) TestRecipeRepo_DeleteById2(t provider.T) {
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
