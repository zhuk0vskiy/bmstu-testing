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
	"ppo/tests/utils"
)

type IngredientRepoSuite struct {
	suite.Suite
}

func (suite *IngredientRepoSuite) TestIngredientRepo_Create1(t provider.T) {
	t.Title("[Ingredient create] successfully created")
	t.Tags("ingredient", "create")
	t.Parallel()

	t.WithNewStep("successfully created", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		ingredient := utils.NewIngredientBuilder().
			WithId(uuid.New()).
			WithTypeId(uuid.New()).
			WithName("ingredient").
			WithCalories(123).
			ToDto()

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatal(err)
		}
		defer mock.Close()
		mock.ExpectExec("insert").
			WithArgs(
				ingredient.Name,
				ingredient.Calories,
				ingredient.TypeID,
			).
			WillReturnResult(pgxmock.NewResult("insert", 1))

		repo := postgres.NewIngredientRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "request", ingredient)
		err = repo.Create(ctx, ingredient)

		sCtx.Assert().NoError(err)
	})
}

func (suite *IngredientRepoSuite) TestIngredientRepo_Create2(t provider.T) {
	t.Title("[Ingredient create] not inserted")
	t.Tags("ingredient", "create")
	t.Parallel()

	t.WithNewStep("not inserted", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		ingredient := utils.NewIngredientBuilder().
			WithId(uuid.New()).
			WithTypeId(uuid.New()).
			WithName("ingredient").
			WithCalories(123).
			ToDto()

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatal(err)
		}
		defer mock.Close()
		mock.ExpectExec("insert").
			WithArgs(
				ingredient.Name,
				ingredient.Calories,
				ingredient.TypeID,
			).
			WillReturnError(fmt.Errorf("insert error"))

		repo := postgres.NewIngredientRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "request", ingredient)
		err = repo.Create(ctx, ingredient)

		sCtx.Assert().Error(err)
	})
}

func (suite *IngredientRepoSuite) TestIngredientRepo_GetById1(t provider.T) {
	t.Title("[Ingredient get by id] successful get")
	t.Tags("ingredient", "get_by_id")
	t.Parallel()

	t.WithNewStep("successful get", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		ingredient := utils.NewIngredientBuilder().
			WithId(uuid.New()).
			WithTypeId(uuid.New()).
			WithName("ingredient").
			WithCalories(123).
			ToDto()

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatal(err)
		}
		defer mock.Close()
		mock.ExpectQuery("select").
			WithArgs(
				ingredient.ID,
			).
			WillReturnRows(
				pgxmock.NewRows([]string{"ID", "Name", "Calories", "Type"}).
					AddRow(ingredient.ID, ingredient.Name, ingredient.Calories, ingredient.TypeID),
			)

		repo := postgres.NewIngredientRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "request", ingredient.ID)
		rIngredient, err := repo.GetById(ctx, ingredient.ID)

		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(ingredient, rIngredient)
	})
}

func (suite *IngredientRepoSuite) TestIngredientRepo_GetById2(t provider.T) {
	t.Title("[Ingredient get by id] failed to get")
	t.Tags("ingredient", "get_by_id")
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
				pgxmock.NewRows([]string{"ID", "Name", "Calories", "Type"}),
			)

		repo := postgres.NewIngredientRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "request", id)
		rIngredient, err := repo.GetById(ctx, id)

		sCtx.Assert().Error(err)
		sCtx.Assert().Nil(rIngredient)
	})
}

func (suite *IngredientRepoSuite) TestIngredientRepo_GetAll1(t provider.T) {
	t.Title("[Ingredient get all] successful get")
	t.Tags("ingredient", "get_all")
	t.Parallel()

	t.WithNewStep("successful get", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		page := 1
		expIngredients := []*domain.Ingredient{
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

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatal(err)
		}
		defer mock.Close()
		mock.ExpectQuery("select").
			WithArgs(
				config.PageSize*(page-1),
				config.PageSize,
			).
			WillReturnRows(
				pgxmock.NewRows([]string{"ID", "Name", "Calories", "Type"}).
					AddRow(expIngredients[0].ID, expIngredients[0].Name, expIngredients[0].Calories, expIngredients[0].TypeID).
					AddRow(expIngredients[1].ID, expIngredients[1].Name, expIngredients[1].Calories, expIngredients[1].TypeID),
			)

		mock.ExpectQuery("select").
			WillReturnRows(
				pgxmock.NewRows([]string{"count"}).
					AddRow(2),
			)

		repo := postgres.NewIngredientRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "request", page)
		comments, pages, err := repo.GetAll(ctx, page)

		sCtx.Assert().NoError(err)
		sCtx.Assert().NotZero(pages)
		sCtx.Assert().Equal(comments, expIngredients)
	})
}

func (suite *IngredientRepoSuite) TestIngredientRepo_GetAll2(t provider.T) {
	t.Title("[Ingredient get all] sql error")
	t.Tags("ingredient", "get_all")
	t.Parallel()

	t.WithNewStep("sql error", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		page := 1

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatal(err)
		}
		defer mock.Close()
		mock.ExpectQuery("select").
			WithArgs(
				config.PageSize*(page-1),
				config.PageSize,
			).
			WillReturnError(fmt.Errorf("sql error"))

		repo := postgres.NewIngredientRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "request", page)
		ingredients, pages, err := repo.GetAll(ctx, page)

		sCtx.Assert().Error(err)
		sCtx.Assert().Zero(pages)
		sCtx.Assert().Nil(ingredients)
	})
}

func (suite *IngredientRepoSuite) TestIngredientRepo_GetAllByRecipeId1(t provider.T) {
	t.Title("[Ingredient get all by recipe id] successful get")
	t.Tags("ingredient", "get_all_by_recipe_id")
	t.Parallel()

	t.WithNewStep("successful get", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		id := uuid.New()
		expIngredients := []*domain.Ingredient{
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
				pgxmock.NewRows([]string{"ID", "Name", "Calories", "Type"}).
					AddRow(expIngredients[0].ID, expIngredients[0].Name, expIngredients[0].Calories, expIngredients[0].TypeID).
					AddRow(expIngredients[1].ID, expIngredients[1].Name, expIngredients[1].Calories, expIngredients[1].TypeID),
			)

		repo := postgres.NewIngredientRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "request", id)
		ingredients, err := repo.GetAllByRecipeId(ctx, id)

		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(ingredients, expIngredients)
	})
}

func (suite *IngredientRepoSuite) TestIngredientRepo_GetAllByRecipeId2(t provider.T) {
	t.Title("[Ingredient get all by recipe id] successful get")
	t.Tags("ingredient", "get_all_by_recipe_id")
	t.Parallel()

	t.WithNewStep("successful get", func(sCtx provider.StepCtx) {
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
			WillReturnError(fmt.Errorf("sql error"))

		repo := postgres.NewIngredientRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "request", id)
		ingredients, err := repo.GetAllByRecipeId(ctx, id)

		sCtx.Assert().NoError(err)
		sCtx.Assert().Nil(ingredients)
	})
}

func (suite *IngredientRepoSuite) TestIngredientRepo_Update1(t provider.T) {
	t.Title("[Ingredient update] successfully updated")
	t.Tags("ingredient", "update")
	t.Parallel()

	t.WithNewStep("successfully updated", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		ingredient := utils.NewIngredientBuilder().
			WithId(uuid.New()).
			WithTypeId(uuid.New()).
			WithName("ingredient").
			WithCalories(123).
			ToDto()

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatal(err)
		}
		defer mock.Close()
		mock.ExpectExec("update").
			WithArgs(
				ingredient.Name,
				ingredient.Calories,
				ingredient.TypeID,
				ingredient.ID,
			).
			WillReturnResult(pgxmock.NewResult("update", 1))

		repo := postgres.NewIngredientRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "request", ingredient)
		err = repo.Update(ctx, ingredient)

		sCtx.Assert().NoError(err)
	})
}

func (suite *IngredientRepoSuite) TestIngredientRepo_Update2(t provider.T) {
	t.Title("[Ingredient update] failed to update")
	t.Tags("ingredient", "update")
	t.Parallel()

	t.WithNewStep("failed to updated", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		ingredient := utils.NewIngredientBuilder().
			WithId(uuid.New()).
			WithTypeId(uuid.New()).
			WithName("ingredient").
			WithCalories(123).
			ToDto()

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatal(err)
		}
		defer mock.Close()
		mock.ExpectExec("update").
			WithArgs(
				ingredient.Name,
				ingredient.Calories,
				ingredient.TypeID,
				ingredient.ID,
			).
			WillReturnError(fmt.Errorf("sql error"))

		repo := postgres.NewIngredientRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "request", ingredient)
		err = repo.Update(ctx, ingredient)

		sCtx.Assert().Error(err)
	})
}

func (suite *IngredientRepoSuite) TestIngredientRepo_DeleteById1(t provider.T) {
	t.Title("[Ingredient delete by id] successfully deleted")
	t.Tags("ingredient", "delete_by_id")
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

		repo := postgres.NewIngredientRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "request", id)
		err = repo.DeleteById(ctx, id)

		sCtx.Assert().NoError(err)
	})
}

func (suite *IngredientRepoSuite) TestIngredientRepo_DeleteById2(t provider.T) {
	t.Title("[Ingredient delete by id] sql error")
	t.Tags("ingredient", "delete_by_id")
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

		repo := postgres.NewIngredientRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "request", id)
		err = repo.DeleteById(ctx, id)

		sCtx.Assert().Error(err)
	})
}

func (suite *IngredientRepoSuite) TestIngredientRepo_Link1(t provider.T) {
	t.Title("[Ingredient link] successfully linked")
	t.Tags("ingredient", "link")
	t.Parallel()

	t.WithNewStep("successfully linked", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		recipeId := uuid.New()
		ingredientId := uuid.New()

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatal(err)
		}
		defer mock.Close()
		mock.ExpectQuery("insert").
			WithArgs(
				recipeId,
				ingredientId,
			).
			WillReturnRows(pgxmock.NewRows([]string{"id"}).AddRow(uuid.New()))

		repo := postgres.NewIngredientRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "request", recipeId, ingredientId)
		id, err := repo.Link(ctx, recipeId, ingredientId)

		sCtx.Assert().NoError(err)
		sCtx.Assert().NotEqual(uuid.Nil, id)
	})
}

func (suite *IngredientRepoSuite) TestIngredientRepo_Link2(t provider.T) {
	t.Title("[Ingredient link] failed to link")
	t.Tags("ingredient", "link")
	t.Parallel()

	t.WithNewStep("failed to link", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		recipeId := uuid.New()
		ingredientId := uuid.New()

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatal(err)
		}
		defer mock.Close()
		mock.ExpectQuery("insert").
			WithArgs(
				recipeId,
				ingredientId,
			).
			WillReturnError(fmt.Errorf("sql error"))

		repo := postgres.NewIngredientRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "request", recipeId, ingredientId)
		id, err := repo.Link(ctx, recipeId, ingredientId)

		sCtx.Assert().Error(err)
		sCtx.Assert().Equal(uuid.Nil, id)
	})
}

func (suite *IngredientRepoSuite) TestIngredientRepo_Unlink1(t provider.T) {
	t.Title("[Ingredient unlink] successfully unlinked")
	t.Tags("ingredient", "unlink")
	t.Parallel()

	t.WithNewStep("successfully unlinked", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		recipeId := uuid.New()
		ingredientId := uuid.New()

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatal(err)
		}
		defer mock.Close()
		mock.ExpectExec("delete").
			WithArgs(
				recipeId,
				ingredientId,
			).
			WillReturnResult(pgxmock.NewResult("delete", 1))

		repo := postgres.NewIngredientRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "request", recipeId, ingredientId)
		err = repo.Unlink(ctx, recipeId, ingredientId)

		sCtx.Assert().NoError(err)
	})
}

func (suite *IngredientRepoSuite) TestIngredientRepo_Unlink2(t provider.T) {
	t.Title("[Ingredient unlink] failed to unlink")
	t.Tags("ingredient", "unlink")
	t.Parallel()

	t.WithNewStep("failed to unlink", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		recipeId := uuid.New()
		ingredientId := uuid.New()

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatal(err)
		}
		defer mock.Close()
		mock.ExpectExec("delete").
			WithArgs(
				recipeId,
				ingredientId,
			).
			WillReturnError(fmt.Errorf("sql error"))

		repo := postgres.NewIngredientRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "request", recipeId, ingredientId)
		err = repo.Unlink(ctx, recipeId, ingredientId)

		sCtx.Assert().NoError(err)
	})
}
