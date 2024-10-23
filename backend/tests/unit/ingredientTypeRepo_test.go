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

type IngredientTypeRepoSuite struct {
	suite.Suite
}

func (suite *IngredientTypeRepoSuite) TestIngredientTypeRepo_Create1(t provider.T) {
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

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatal(err)
		}
		defer mock.Close()
		mock.ExpectExec("insert").
			WithArgs(
				ingredientType.Name,
				ingredientType.Description,
			).
			WillReturnResult(pgxmock.NewResult("insert", 1))

		repo := postgres.NewIngredientTypeRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "request", ingredientType)
		err = repo.Create(ctx, ingredientType)

		sCtx.Assert().NoError(err)
	})
}

func (suite *IngredientTypeRepoSuite) TestIngredientTypeRepo_Create2(t provider.T) {
	t.Title("[Ingredient type create] not inserted")
	t.Tags("ingredient_type", "create")
	t.Parallel()

	t.WithNewStep("not inserted", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		ingredientType := utils.NewIngredientTypeBuilder().
			WithId(uuid.New()).
			WithName("type").
			WithDescription("description").
			ToDto()

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatal(err)
		}
		defer mock.Close()
		mock.ExpectExec("insert").
			WithArgs(
				ingredientType.Name,
				ingredientType.Description,
			).
			WillReturnError(fmt.Errorf("insert error"))

		repo := postgres.NewIngredientTypeRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "request", ingredientType)
		err = repo.Create(ctx, ingredientType)

		sCtx.Assert().Error(err)
	})
}

func (suite *IngredientTypeRepoSuite) TestIngredientTypeRepo_GetById1(t provider.T) {
	t.Title("[Ingredient type get by id] successful get")
	t.Tags("ingredient_type", "get_by_id")
	t.Parallel()

	t.WithNewStep("successful get", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		ingredientType := utils.NewIngredientTypeBuilder().
			WithId(uuid.New()).
			WithName("type").
			WithDescription("description").
			ToDto()

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatal(err)
		}
		defer mock.Close()
		mock.ExpectQuery("select").
			WithArgs(
				ingredientType.ID,
			).
			WillReturnRows(
				pgxmock.NewRows([]string{"ID", "Name", "Description"}).
					AddRow(ingredientType.ID, ingredientType.Name, ingredientType.Description),
			)

		repo := postgres.NewIngredientTypeRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "request", ingredientType.ID)
		rType, err := repo.GetById(ctx, ingredientType.ID)

		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(ingredientType, rType)
	})
}

func (suite *IngredientTypeRepoSuite) TestIngredientTypeRepo_GetById2(t provider.T) {
	t.Title("[Ingredient type get by id] failed to get")
	t.Tags("ingredient_type", "get_by_id")
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
				pgxmock.NewRows([]string{"ID", "Name", "Description"}),
			)

		repo := postgres.NewIngredientTypeRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "request", id)
		rType, err := repo.GetById(ctx, id)

		sCtx.Assert().Error(err)
		sCtx.Assert().Nil(rType)
	})
}

func (suite *IngredientTypeRepoSuite) TestIngredientTypeRepo_GetAll1(t provider.T) {
	t.Title("[Ingredient type get all] successful get")
	t.Tags("ingredient_type", "get_all")
	t.Parallel()

	t.WithNewStep("successful get", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
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

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatal(err)
		}
		defer mock.Close()
		mock.ExpectQuery("select").
			WillReturnRows(
				pgxmock.NewRows([]string{"ID", "Name", "Description"}).
					AddRow(expTypes[0].ID, expTypes[0].Name, expTypes[0].Description).
					AddRow(expTypes[1].ID, expTypes[1].Name, expTypes[1].Description),
			)

		repo := postgres.NewIngredientTypeRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "request")
		comments, err := repo.GetAll(ctx)

		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(comments, expTypes)
	})
}

func (suite *IngredientTypeRepoSuite) TestIngredientTypeRepo_GetAll2(t provider.T) {
	t.Title("[Ingredient type get all] sql error")
	t.Tags("ingredient_type", "get_all")
	t.Parallel()

	t.WithNewStep("sql error", func(sCtx provider.StepCtx) {
		ctx := context.TODO()

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatal(err)
		}
		defer mock.Close()
		mock.ExpectQuery("select").
			WillReturnError(fmt.Errorf("sql error"))

		repo := postgres.NewIngredientTypeRepository(mock)

		sCtx.WithNewParameters("ctx", ctx)
		types, err := repo.GetAll(ctx)

		sCtx.Assert().Error(err)
		sCtx.Assert().Nil(types)
	})
}

func (suite *IngredientTypeRepoSuite) TestIngredientTypeRepo_Update1(t provider.T) {
	t.Title("[Ingredient type update] successfully updated")
	t.Tags("ingredient_type", "update")
	t.Parallel()

	t.WithNewStep("successfully updated", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		ingredientType := utils.NewIngredientTypeBuilder().
			WithId(uuid.New()).
			WithName("type").
			WithDescription("description").
			ToDto()

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatal(err)
		}
		defer mock.Close()
		mock.ExpectExec("update").
			WithArgs(
				ingredientType.Name,
				ingredientType.Description,
				ingredientType.ID,
			).
			WillReturnResult(pgxmock.NewResult("update", 1))

		repo := postgres.NewIngredientTypeRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "request", ingredientType)
		err = repo.Update(ctx, ingredientType)

		sCtx.Assert().NoError(err)
	})
}

func (suite *IngredientTypeRepoSuite) TestIngredientTypeRepo_Update2(t provider.T) {
	t.Title("[Ingredient type update] sql error")
	t.Tags("ingredient_type", "update")
	t.Parallel()

	t.WithNewStep("sql error", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		ingredientType := utils.NewIngredientTypeBuilder().
			WithId(uuid.New()).
			WithName("type").
			WithDescription("description").
			ToDto()

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatal(err)
		}
		defer mock.Close()
		mock.ExpectExec("update").
			WithArgs(
				ingredientType.Name,
				ingredientType.Description,
				ingredientType.ID,
			).
			WillReturnError(fmt.Errorf("sql error"))

		repo := postgres.NewIngredientTypeRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "request", ingredientType)
		err = repo.Update(ctx, ingredientType)

		sCtx.Assert().Error(err)
	})
}

func (suite *IngredientTypeRepoSuite) TestIngredientTypeRepo_DeleteById1(t provider.T) {
	t.Title("[Ingredient type delete by id] successfully deleted")
	t.Tags("ingredient_type", "delete_by_id")
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

		repo := postgres.NewIngredientTypeRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "request", id)
		err = repo.DeleteById(ctx, id)

		sCtx.Assert().NoError(err)
	})
}

func (suite *IngredientTypeRepoSuite) TestIngredientTypeRepo_DeleteById2(t provider.T) {
	t.Title("[Ingredient type delete by id] sql error")
	t.Tags("ingredient_type", "delete_by_id")
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

		repo := postgres.NewIngredientTypeRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "request", id)
		err = repo.DeleteById(ctx, id)

		sCtx.Assert().Error(err)
	})
}
