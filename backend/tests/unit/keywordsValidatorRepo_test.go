//go:build unit_test

package unit_tests

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
	"github.com/pashagolub/pgxmock/v4"
	"ppo/internal/storage/postgres"
	"ppo/tests/utils"
)

type KeywordsRepoSuite struct {
	suite.Suite
}

func (suite *KeywordsRepoSuite) TestKeywordsRepo_Create1(t provider.T) {
	t.Title("[Keyword create] successfully created")
	t.Tags("keyword", "create")
	t.Parallel()

	t.WithNewStep("successfully created", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		keyword := utils.NewKeywordBuilder().
			WithId(uuid.New()).
			WithWord("banned_word").
			ToDto()

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatal(err)
		}
		defer mock.Close()
		mock.ExpectExec("insert").
			WithArgs(
				keyword.Word,
			).
			WillReturnResult(pgxmock.NewResult("insert", 1))

		repo := postgres.NewKeywordValidatorRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "request", keyword)
		err = repo.Create(ctx, keyword)

		sCtx.Assert().NoError(err)
	})
}

func (suite *KeywordsRepoSuite) TestKeywordsRepo_Create2(t provider.T) {
	t.Title("[Keyword create] not inserted")
	t.Tags("keyword", "create")
	t.Parallel()

	t.WithNewStep("not inserted", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		keyword := utils.NewKeywordBuilder().
			WithId(uuid.New()).
			WithWord("banned_word").
			ToDto()

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatal(err)
		}
		defer mock.Close()
		mock.ExpectExec("insert").
			WithArgs(
				keyword.Word,
			).
			WillReturnError(fmt.Errorf("insert error"))

		repo := postgres.NewKeywordValidatorRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "request", keyword)
		err = repo.Create(ctx, keyword)

		sCtx.Assert().Error(err)
	})
}

func (suite *KeywordsRepoSuite) TestKeywordsRepo_GetById1(t provider.T) {
	t.Title("[Keyword get by id] successful get")
	t.Tags("keyword", "get_by_id")
	t.Parallel()

	t.WithNewStep("successful get", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		keyword := utils.NewKeywordBuilder().
			WithId(uuid.New()).
			WithWord("banned_word").
			ToDto()

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatal(err)
		}
		defer mock.Close()
		mock.ExpectQuery("select").
			WithArgs(
				keyword.ID,
			).
			WillReturnRows(
				pgxmock.NewRows([]string{"ID", "Word"}).
					AddRow(keyword.ID, keyword.Word),
			)

		repo := postgres.NewKeywordValidatorRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "request", keyword.ID)
		rKeyword, err := repo.GetById(ctx, keyword.ID)

		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(keyword, rKeyword)
	})
}

func (suite *KeywordsRepoSuite) TestKeywordsRepo_GetById2(t provider.T) {
	t.Title("[Keyword get by id] failed to get")
	t.Tags("keyword", "get_by_id")
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
				pgxmock.NewRows([]string{"ID", "Word"}),
			)

		repo := postgres.NewKeywordValidatorRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "request", id)
		rKeyword, err := repo.GetById(ctx, id)

		sCtx.Assert().Error(err)
		sCtx.Assert().Nil(rKeyword)
	})
}

func (suite *KeywordsRepoSuite) TestKeywordsRepo_GetAll1(t provider.T) {
	t.Title("[Keyword get all] successful get")
	t.Tags("keyword", "get_all")
	t.Parallel()

	t.WithNewStep("successful get", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		keywords := map[string]uuid.UUID{
			"banned1": uuid.New(),
			"banned2": uuid.New(),
		}

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatal(err)
		}
		defer mock.Close()
		mock.ExpectQuery("select").
			WillReturnRows(
				pgxmock.NewRows([]string{"ID", "Word"}).
					AddRow(keywords["banned1"], "banned1").
					AddRow(keywords["banned2"], "banned2"),
			)

		repo := postgres.NewKeywordValidatorRepository(mock)

		sCtx.WithNewParameters("ctx", ctx)
		rKeywords, err := repo.GetAll(ctx)

		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(rKeywords, keywords)
	})
}

func (suite *KeywordsRepoSuite) TestKeywordsRepo_GetAll2(t provider.T) {
	t.Title("[Keyword get all] sql error")
	t.Tags("keyword", "get_all")
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

		repo := postgres.NewKeywordValidatorRepository(mock)

		sCtx.WithNewParameters("ctx", ctx)
		keywords, err := repo.GetAll(ctx)

		sCtx.Assert().Error(err)
		sCtx.Assert().Nil(keywords)
	})
}

func (suite *KeywordsRepoSuite) TestKeywordsRepo_Update1(t provider.T) {
	t.Title("[Keyword update] successfully updated")
	t.Tags("keyword", "update")
	t.Parallel()

	t.WithNewStep("successfully updated", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		keyword := utils.NewKeywordBuilder().
			WithId(uuid.New()).
			WithWord("banned_word").
			ToDto()

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatal(err)
		}
		defer mock.Close()
		mock.ExpectExec("update").
			WithArgs(
				keyword.Word,
				keyword.ID,
			).
			WillReturnResult(pgxmock.NewResult("update", 1))

		repo := postgres.NewKeywordValidatorRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "request", keyword)
		err = repo.Update(ctx, keyword)

		sCtx.Assert().NoError(err)
	})
}

func (suite *KeywordsRepoSuite) TestKeywordsRepo_Update2(t provider.T) {
	t.Title("[Keyword update] failed to update")
	t.Tags("keyword", "update")
	t.Parallel()

	t.WithNewStep("failed to updated", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		keyword := utils.NewKeywordBuilder().
			WithId(uuid.New()).
			WithWord("banned_word").
			ToDto()

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatal(err)
		}
		defer mock.Close()
		mock.ExpectExec("update").
			WithArgs(
				keyword.Word,
				keyword.ID,
			).
			WillReturnError(fmt.Errorf("sql error"))

		repo := postgres.NewKeywordValidatorRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "request", keyword)
		err = repo.Update(ctx, keyword)

		sCtx.Assert().Error(err)
	})
}

func (suite *KeywordsRepoSuite) TestKeywordsRepo_DeleteById1(t provider.T) {
	t.Title("[Keyword delete by id] successfully deleted")
	t.Tags("keyword", "delete_by_id")
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

		repo := postgres.NewKeywordValidatorRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "request", id)
		err = repo.DeleteById(ctx, id)

		sCtx.Assert().NoError(err)
	})
}

func (suite *KeywordsRepoSuite) TestKeywordsRepo_DeleteById2(t provider.T) {
	t.Title("[Keyword delete by id] sql error")
	t.Tags("keyword", "delete_by_id")
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

		repo := postgres.NewKeywordValidatorRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "request", id)
		err = repo.DeleteById(ctx, id)

		sCtx.Assert().Error(err)
	})
}
