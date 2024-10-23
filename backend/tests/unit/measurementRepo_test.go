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

type MeasurementRepoSuite struct {
	suite.Suite
}

func (suite *MeasurementRepoSuite) TestMeasurementRepo_Create1(t provider.T) {
	t.Title("[Measurement create] successfully created")
	t.Tags("measurement", "create")
	t.Parallel()

	t.WithNewStep("successfully created", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		measurement := utils.NewMeasurementBuilder().
			WithId(uuid.New()).
			WithName("gram").
			WithGrams(1).
			ToDto()

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatal(err)
		}
		defer mock.Close()
		mock.ExpectExec("insert").
			WithArgs(
				measurement.Name,
				measurement.Grams,
			).
			WillReturnResult(pgxmock.NewResult("insert", 1))

		repo := postgres.NewMeasrementRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "request", measurement)
		err = repo.Create(ctx, measurement)

		sCtx.Assert().NoError(err)
	})
}

func (suite *MeasurementRepoSuite) TestMeasurementRepo_Create2(t provider.T) {
	t.Title("[Measurement create] not inserted")
	t.Tags("measurement", "create")
	t.Parallel()

	t.WithNewStep("not inserted", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		measurement := utils.NewMeasurementBuilder().
			WithId(uuid.New()).
			WithName("gram").
			WithGrams(1).
			ToDto()

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatal(err)
		}
		defer mock.Close()
		mock.ExpectExec("insert").
			WithArgs(
				measurement.Name,
				measurement.Grams,
			).
			WillReturnError(fmt.Errorf("insert error"))

		repo := postgres.NewMeasrementRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "request", measurement)
		err = repo.Create(ctx, measurement)

		sCtx.Assert().Error(err)
	})
}

func (suite *MeasurementRepoSuite) TestMeasurementRepo_GetById1(t provider.T) {
	t.Title("[Measurement get by id] successful get")
	t.Tags("measurement", "get_by_id")
	t.Parallel()

	t.WithNewStep("successful get", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		measurement := utils.NewMeasurementBuilder().
			WithId(uuid.New()).
			WithName("gram").
			WithGrams(1).
			ToDto()

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatal(err)
		}
		defer mock.Close()
		mock.ExpectQuery("select").
			WithArgs(
				measurement.ID,
			).
			WillReturnRows(
				pgxmock.NewRows([]string{"ID", "Name", "Grams"}).
					AddRow(measurement.ID, measurement.Name, measurement.Grams),
			)

		repo := postgres.NewMeasrementRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "request", measurement.ID)
		rMeasurement, err := repo.GetById(ctx, measurement.ID)

		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(measurement, rMeasurement)
	})
}

func (suite *MeasurementRepoSuite) TestMeasurementRepo_GetById2(t provider.T) {
	t.Title("[Measurement get by id] failed to get")
	t.Tags("measurement", "get_by_id")
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
				pgxmock.NewRows([]string{"ID", "Name", "Grams"}),
			)

		repo := postgres.NewMeasrementRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "request", id)
		rMeasurement, err := repo.GetById(ctx, id)

		sCtx.Assert().Error(err)
		sCtx.Assert().Nil(rMeasurement)
	})
}

func (suite *MeasurementRepoSuite) TestMeasurementRepo_GetByRecipeId1(t provider.T) {
	t.Title("[Measurement get by recipe id] successful get")
	t.Tags("measurement", "get_by_recipe_id")
	t.Parallel()

	t.WithNewStep("successful get", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		ingredientId := uuid.New()
		recipeId := uuid.New()
		amount := 1
		measurement := utils.NewMeasurementBuilder().
			WithId(uuid.New()).
			WithName("gram").
			WithGrams(1).
			ToDto()

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatal(err)
		}
		defer mock.Close()
		mock.ExpectQuery("select").
			WithArgs(
				recipeId,
				ingredientId,
			).
			WillReturnRows(
				pgxmock.NewRows([]string{"ID", "Name", "Grams", "Amount"}).
					AddRow(measurement.ID, measurement.Name, measurement.Grams, amount),
			)

		repo := postgres.NewMeasrementRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "request", ingredientId, recipeId)
		rMeasurement, amount, err := repo.GetByRecipeId(ctx, ingredientId, recipeId)

		sCtx.Assert().NoError(err)
		sCtx.Assert().NotZero(amount)
		sCtx.Assert().Equal(measurement, rMeasurement)
	})
}

func (suite *MeasurementRepoSuite) TestMeasurementRepo_GetByRecipeId2(t provider.T) {
	t.Title("[Measurement get by recipe id] sql error")
	t.Tags("measurement", "get_by_recipe_id")
	t.Parallel()

	t.WithNewStep("sql error", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		ingredientId := uuid.New()
		recipeId := uuid.New()

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatal(err)
		}
		defer mock.Close()
		mock.ExpectQuery("select").
			WithArgs(
				recipeId,
				ingredientId,
			).
			WillReturnError(fmt.Errorf("sql error"))

		repo := postgres.NewMeasrementRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "request", ingredientId, recipeId)
		rMeasurement, _, err := repo.GetByRecipeId(ctx, ingredientId, recipeId)

		sCtx.Assert().Error(err)
		sCtx.Assert().Nil(rMeasurement)
	})
}

func (suite *MeasurementRepoSuite) TestMeasurementRepo_GetAll1(t provider.T) {
	t.Title("[Measurement get all] successful get")
	t.Tags("measurement", "get_all")
	t.Parallel()

	t.WithNewStep("successful get", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		expMeasurements := []*domain.Measurement{
			{
				ID:    uuid.UUID{1},
				Name:  "measurement1",
				Grams: 1,
			},
			{
				ID:    uuid.UUID{2},
				Name:  "measurement2",
				Grams: 2,
			},
		}

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatal(err)
		}
		defer mock.Close()
		mock.ExpectQuery("select").
			WillReturnRows(
				pgxmock.NewRows([]string{"ID", "Name", "Grams"}).
					AddRow(expMeasurements[0].ID, expMeasurements[0].Name, expMeasurements[0].Grams).
					AddRow(expMeasurements[1].ID, expMeasurements[1].Name, expMeasurements[1].Grams),
			)

		repo := postgres.NewMeasrementRepository(mock)

		sCtx.WithNewParameters("ctx", ctx)
		rMeasurements, err := repo.GetAll(ctx)

		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(expMeasurements, rMeasurements)
	})
}

func (suite *MeasurementRepoSuite) TestMeasurementRepo_GetAll2(t provider.T) {
	t.Title("[Measurement get all] sql error")
	t.Tags("measurement", "get_all")
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

		repo := postgres.NewMeasrementRepository(mock)

		sCtx.WithNewParameters("ctx", ctx)
		rMeasurements, err := repo.GetAll(ctx)

		sCtx.Assert().Error(err)
		sCtx.Assert().Nil(rMeasurements)
	})
}

func (suite *MeasurementRepoSuite) TestMeasurementRepo_Update1(t provider.T) {
	t.Title("[Measurement update] successfully updated")
	t.Tags("measurement", "update")
	t.Parallel()

	t.WithNewStep("successfully updated", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		measurement := utils.NewMeasurementBuilder().
			WithId(uuid.New()).
			WithName("gram").
			WithGrams(1).
			ToDto()

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatal(err)
		}
		defer mock.Close()
		mock.ExpectExec("update").
			WithArgs(
				measurement.Name,
				measurement.Grams,
				measurement.ID,
			).
			WillReturnResult(pgxmock.NewResult("update", 1))

		repo := postgres.NewMeasrementRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "request", measurement)
		err = repo.Update(ctx, measurement)

		sCtx.Assert().NoError(err)
	})
}

func (suite *MeasurementRepoSuite) TestMeasurementRepo_Update2(t provider.T) {
	t.Title("[Measurement update] failed to update")
	t.Tags("measurement", "update")
	t.Parallel()

	t.WithNewStep("failed to updated", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		measurement := utils.NewMeasurementBuilder().
			WithId(uuid.New()).
			WithName("gram").
			WithGrams(1).
			ToDto()

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatal(err)
		}
		defer mock.Close()
		mock.ExpectExec("update").
			WithArgs(
				measurement.Name,
				measurement.Grams,
				measurement.ID,
			).
			WillReturnError(fmt.Errorf("sql error"))

		repo := postgres.NewMeasrementRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "request", measurement)
		err = repo.Update(ctx, measurement)

		sCtx.Assert().Error(err)
	})
}

func (suite *MeasurementRepoSuite) TestMeasurementRepo_DeleteById1(t provider.T) {
	t.Title("[Measurement delete by id] successfully deleted")
	t.Tags("measurement", "delete_by_id")
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

		repo := postgres.NewMeasrementRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "request", id)
		err = repo.DeleteById(ctx, id)

		sCtx.Assert().NoError(err)
	})
}

func (suite *MeasurementRepoSuite) TestMeasurementRepo_DeleteById2(t provider.T) {
	t.Title("[Measurement delete by id] sql error")
	t.Tags("measurement", "delete_by_id")
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

		repo := postgres.NewMeasrementRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "request", id)
		err = repo.DeleteById(ctx, id)

		sCtx.Assert().Error(err)
	})
}

func (suite *MeasurementRepoSuite) TestMeasurementRepo_UpdateLink1(t provider.T) {
	t.Title("[Measurement update link] successfully updated")
	t.Tags("measurement", "update_link")
	t.Parallel()

	t.WithNewStep("successfully updated", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		linkId := uuid.New()
		measurementId := uuid.New()
		amount := 1

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatal(err)
		}
		defer mock.Close()
		mock.ExpectExec("update").
			WithArgs(
				measurementId,
				amount,
				linkId,
			).
			WillReturnResult(pgxmock.NewResult("update", 1))

		repo := postgres.NewMeasrementRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "request", linkId, measurementId, amount)
		err = repo.UpdateLink(ctx, linkId, measurementId, amount)

		sCtx.Assert().NoError(err)
	})
}

func (suite *MeasurementRepoSuite) TestMeasurementRepo_UpdateLink2(t provider.T) {
	t.Title("[Measurement update link] sql error")
	t.Tags("measurement", "update_link")
	t.Parallel()

	t.WithNewStep("sql error", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		linkId := uuid.New()
		measurementId := uuid.New()
		amount := 1

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatal(err)
		}
		defer mock.Close()
		mock.ExpectExec("update").
			WithArgs(
				measurementId,
				amount,
				linkId,
			).
			WillReturnError(fmt.Errorf("sql error"))

		repo := postgres.NewMeasrementRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "request", linkId, measurementId, amount)
		err = repo.UpdateLink(ctx, linkId, measurementId, amount)

		sCtx.Assert().Error(err)
	})
}
