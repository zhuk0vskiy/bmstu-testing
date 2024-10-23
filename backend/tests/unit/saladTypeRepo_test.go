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

type SaladTypeRepoSuite struct {
	suite.Suite
}

func (suite *SaladTypeRepoSuite) TestSaladTypeRepo_Create1(t provider.T) {
	t.Title("[Salad type create] successfully created")
	t.Tags("salad_type", "create")
	t.Parallel()

	t.WithNewStep("successfully created", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		saladType := utils.NewSaladTypeBuilder().
			WithId(uuid.New()).
			WithName("saladType").
			WithDescription("description").
			ToDto()

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatal(err)
		}
		defer mock.Close()
		mock.ExpectExec("insert").
			WithArgs(
				saladType.Name,
				saladType.Description,
			).
			WillReturnResult(pgxmock.NewResult("insert", 1))

		repo := postgres.NewSaladTypeRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "request", saladType)
		err = repo.Create(ctx, saladType)

		sCtx.Assert().NoError(err)
	})
}

func (suite *SaladTypeRepoSuite) TestSaladTypeRepo_Create2(t provider.T) {
	t.Title("[Salad type create] not inserted")
	t.Tags("salad_type", "create")
	t.Parallel()

	t.WithNewStep("not inserted", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		saladType := utils.NewSaladTypeBuilder().
			WithId(uuid.New()).
			WithName("saladType").
			WithDescription("description").
			ToDto()

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatal(err)
		}
		defer mock.Close()
		mock.ExpectExec("insert").
			WithArgs(
				saladType.Name,
				saladType.Description,
			).
			WillReturnError(fmt.Errorf("insert error"))

		repo := postgres.NewSaladTypeRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "request", saladType)
		err = repo.Create(ctx, saladType)

		sCtx.Assert().Error(err)
	})
}

func (suite *SaladTypeRepoSuite) TestSaladTypeRepo_GetById1(t provider.T) {
	t.Title("[Salad type get by id] successful get")
	t.Tags("salad_type", "get_by_id")
	t.Parallel()

	t.WithNewStep("successful get", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		saladType := utils.NewSaladTypeBuilder().
			WithId(uuid.New()).
			WithName("saladType").
			WithDescription("description").
			ToDto()

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatal(err)
		}
		defer mock.Close()
		mock.ExpectQuery("select").
			WithArgs(
				saladType.ID,
			).
			WillReturnRows(
				pgxmock.NewRows([]string{"ID", "Name", "Description"}).
					AddRow(saladType.ID, saladType.Name, saladType.Description),
			)

		repo := postgres.NewSaladTypeRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "request", saladType.ID)
		rType, err := repo.GetById(ctx, saladType.ID)

		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(saladType, rType)
	})
}

func (suite *SaladTypeRepoSuite) TestSaladTypeRepo_GetById2(t provider.T) {
	t.Title("[Salad type get by id] failed to get")
	t.Tags("salad_type", "get_by_id")
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

		repo := postgres.NewSaladTypeRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "request", id)
		rType, err := repo.GetById(ctx, id)

		sCtx.Assert().Error(err)
		sCtx.Assert().Nil(rType)
	})
}

func (suite *SaladTypeRepoSuite) TestSaladTypeRepo_GetAll1(t provider.T) {
	t.Title("[Salad type get all] successful get")
	t.Tags("salad_type", "get_all")
	t.Parallel()

	t.WithNewStep("successful get", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		page := 1
		expTypes := []*domain.SaladType{
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
			WithArgs(
				config.PageSize*(page-1),
				config.PageSize,
			).
			WillReturnRows(
				pgxmock.NewRows([]string{"ID", "Name", "Description"}).
					AddRow(expTypes[0].ID, expTypes[0].Name, expTypes[0].Description).
					AddRow(expTypes[1].ID, expTypes[1].Name, expTypes[1].Description),
			)

		mock.ExpectQuery("select").
			WillReturnRows(pgxmock.NewRows([]string{"Count"}).
				AddRow(2),
			)

		repo := postgres.NewSaladTypeRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "request", page)
		comments, pages, err := repo.GetAll(ctx, page)

		sCtx.Assert().NoError(err)
		sCtx.Assert().NotZero(pages)
		sCtx.Assert().Equal(comments, expTypes)
	})
}

func (suite *SaladTypeRepoSuite) TestSaladTypeRepo_GetAll2(t provider.T) {
	t.Title("[Salad type get all] sql error")
	t.Tags("salad_type", "get_all")
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
			WillReturnError(fmt.Errorf("sql error"))

		repo := postgres.NewSaladTypeRepository(mock)

		sCtx.WithNewParameters("ctx", ctx)
		types, pages, err := repo.GetAll(ctx, page)

		sCtx.Assert().Error(err)
		sCtx.Assert().Zero(pages)
		sCtx.Assert().Nil(types)
	})
}

func (suite *SaladTypeRepoSuite) TestSaladTypeRepo_GetAllBySaladId1(t provider.T) {
	t.Title("[Salad type get all] successful get")
	t.Tags("salad_type", "get_all_by_salad_id")
	t.Parallel()

	t.WithNewStep("successful get", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		saladId := uuid.New()
		expTypes := []*domain.SaladType{
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
			WithArgs(
				saladId,
			).
			WillReturnRows(
				pgxmock.NewRows([]string{"ID", "Name", "Description"}).
					AddRow(expTypes[0].ID, expTypes[0].Name, expTypes[0].Description).
					AddRow(expTypes[1].ID, expTypes[1].Name, expTypes[1].Description),
			)

		repo := postgres.NewSaladTypeRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "request", saladId)
		comments, err := repo.GetAllBySaladId(ctx, saladId)

		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(comments, expTypes)
	})
}

func (suite *SaladTypeRepoSuite) TestSaladTypeRepo_GetAllBySaladId2(t provider.T) {
	t.Title("[Salad type get all] sql error")
	t.Tags("salad_type", "get_all_by_salad_id")
	t.Parallel()

	t.WithNewStep("sql error", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		saladId := uuid.New()

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatal(err)
		}
		defer mock.Close()
		mock.ExpectQuery("select").
			WillReturnError(fmt.Errorf("sql error"))

		repo := postgres.NewSaladTypeRepository(mock)

		sCtx.WithNewParameters("ctx", ctx)
		types, err := repo.GetAllBySaladId(ctx, saladId)

		sCtx.Assert().Error(err)
		sCtx.Assert().Nil(types)
	})
}

func (suite *SaladTypeRepoSuite) TestSaladTypeRepo_Update1(t provider.T) {
	t.Title("[Salad type update] successfully updated")
	t.Tags("salad_type", "update")
	t.Parallel()

	t.WithNewStep("successfully updated", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		saladType := utils.NewSaladTypeBuilder().
			WithId(uuid.New()).
			WithName("saladType").
			WithDescription("description").
			ToDto()

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatal(err)
		}
		defer mock.Close()
		mock.ExpectExec("update").
			WithArgs(
				saladType.Name,
				saladType.Description,
				saladType.ID,
			).
			WillReturnResult(pgxmock.NewResult("update", 1))

		repo := postgres.NewSaladTypeRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "request", saladType)
		err = repo.Update(ctx, saladType)

		sCtx.Assert().NoError(err)
	})
}

func (suite *SaladTypeRepoSuite) TestSaladTypeRepo_Update2(t provider.T) {
	t.Title("[Salad type update] sql error")
	t.Tags("salad_type", "update")
	t.Parallel()

	t.WithNewStep("sql error", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		saladType := utils.NewSaladTypeBuilder().
			WithId(uuid.New()).
			WithName("saladType").
			WithDescription("description").
			ToDto()

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatal(err)
		}
		defer mock.Close()
		mock.ExpectExec("update").
			WithArgs(
				saladType.Name,
				saladType.Description,
				saladType.ID,
			).
			WillReturnError(fmt.Errorf("sql error"))

		repo := postgres.NewSaladTypeRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "request", saladType)
		err = repo.Update(ctx, saladType)

		sCtx.Assert().Error(err)
	})
}

func (suite *SaladTypeRepoSuite) TestSaladTypeRepo_DeleteById1(t provider.T) {
	t.Title("[Salad type delete by id] successfully deleted")
	t.Tags("salad_type", "delete_by_id")
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

		repo := postgres.NewSaladTypeRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "request", id)
		err = repo.DeleteById(ctx, id)

		sCtx.Assert().NoError(err)
	})
}

func (suite *SaladTypeRepoSuite) TestSaladTypeRepo_DeleteById2(t provider.T) {
	t.Title("[Salad type delete by id] sql error")
	t.Tags("salad_type", "delete_by_id")
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

		repo := postgres.NewSaladTypeRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "request", id)
		err = repo.DeleteById(ctx, id)

		sCtx.Assert().Error(err)
	})
}

func (suite *SaladTypeRepoSuite) TestSaladTypeRepo_Link1(t provider.T) {
	t.Title("[Salad type link] successfully linked")
	t.Tags("salad_type", "link")
	t.Parallel()

	t.WithNewStep("successfully deleted", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		saladId := uuid.New()
		saladTypeId := uuid.New()

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatal(err)
		}
		defer mock.Close()
		mock.ExpectExec("insert").
			WithArgs(
				saladId,
				saladTypeId,
			).
			WillReturnResult(pgxmock.NewResult("insert", 1))

		repo := postgres.NewSaladTypeRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "request", saladId, saladTypeId)
		err = repo.Link(ctx, saladId, saladTypeId)

		sCtx.Assert().NoError(err)
	})
}

func (suite *SaladTypeRepoSuite) TestSaladTypeRepo_Link2(t provider.T) {
	t.Title("[Salad type link] sql error")
	t.Tags("salad_type", "link")
	t.Parallel()

	t.WithNewStep("sql error", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		saladId := uuid.New()
		saladTypeId := uuid.New()

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatal(err)
		}
		defer mock.Close()
		mock.ExpectExec("insert").
			WithArgs(
				saladId,
				saladTypeId,
			).
			WillReturnError(fmt.Errorf("sql error"))

		repo := postgres.NewSaladTypeRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "request", saladId, saladTypeId)
		err = repo.Link(ctx, saladId, saladTypeId)

		sCtx.Assert().Error(err)
	})
}

func (suite *SaladTypeRepoSuite) TestSaladTypeRepo_Unlink1(t provider.T) {
	t.Title("[Salad type link] successfully unlinked")
	t.Tags("salad_type", "unlink")
	t.Parallel()

	t.WithNewStep("successfully unlinked", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		saladId := uuid.New()
		saladTypeId := uuid.New()

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatal(err)
		}
		defer mock.Close()
		mock.ExpectExec("delete").
			WithArgs(
				saladId,
				saladTypeId,
			).
			WillReturnResult(pgxmock.NewResult("delete", 1))

		repo := postgres.NewSaladTypeRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "request", saladId, saladTypeId)
		err = repo.Unlink(ctx, saladId, saladTypeId)

		sCtx.Assert().NoError(err)
	})
}

func (suite *SaladTypeRepoSuite) TestSaladTypeRepo_Unlink2(t provider.T) {
	t.Title("[Salad type unlink] sql error")
	t.Tags("salad_type", "unlink")
	t.Parallel()

	t.WithNewStep("sql error", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		saladId := uuid.New()
		saladTypeId := uuid.New()

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatal(err)
		}
		defer mock.Close()
		mock.ExpectExec("delete").
			WithArgs(
				saladId,
				saladTypeId,
			).
			WillReturnError(fmt.Errorf("sql error"))

		repo := postgres.NewSaladTypeRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "request", saladId, saladTypeId)
		err = repo.Unlink(ctx, saladId, saladTypeId)

		sCtx.Assert().Error(err)
	})
}
