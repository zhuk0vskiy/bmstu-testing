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

type CommentRepoSuite struct {
	suite.Suite
}

func (suite *CommentRepoSuite) TestCommentRepo_Create1(t provider.T) {
	t.Title("[Comment create] successfully created")
	t.Tags("comment", "create")
	t.Parallel()

	t.WithNewStep("successfully created", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		comment := utils.NewCommentBuilder().
			WithId(uuid.New()).
			WithAuthorId(uuid.New()).
			WithSaladId(uuid.New()).
			WithText("text").
			WithRating(5).
			ToDto()

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatal(err)
		}
		defer mock.Close()
		mock.ExpectExec("insert").
			WithArgs(
				comment.AuthorID,
				comment.SaladID,
				comment.Text,
				comment.Rating,
			).
			WillReturnResult(pgxmock.NewResult("insert", 1))

		repo := postgres.NewCommentRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "request", comment)
		err = repo.Create(ctx, comment)

		sCtx.Assert().NoError(err)
	})
}

func (suite *CommentRepoSuite) TestCommentRepo_Create2(t provider.T) {
	t.Title("[Comment create] not inserted")
	t.Tags("comment", "create")
	t.Parallel()

	t.WithNewStep("not inserted", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		comment := utils.NewCommentBuilder().
			WithId(uuid.New()).
			WithAuthorId(uuid.New()).
			WithSaladId(uuid.New()).
			WithText("text").
			WithRating(5).
			ToDto()

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatal(err)
		}
		defer mock.Close()
		mock.ExpectExec("insert").
			WithArgs(
				comment.AuthorID,
				comment.SaladID,
				comment.Text,
				comment.Rating,
			).
			WillReturnError(fmt.Errorf("insert error"))

		repo := postgres.NewCommentRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "request", comment)
		err = repo.Create(ctx, comment)

		sCtx.Assert().Error(err)
	})
}

func (suite *CommentRepoSuite) TestCommentRepo_GetById1(t provider.T) {
	t.Title("[Comment get by id] successful get")
	t.Tags("comment", "get_by_id")
	t.Parallel()

	t.WithNewStep("successful get", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		comment := utils.NewCommentBuilder().
			WithId(uuid.New()).
			WithAuthorId(uuid.New()).
			WithSaladId(uuid.New()).
			WithText("text").
			WithRating(5).
			ToDto()

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatal(err)
		}
		defer mock.Close()
		mock.ExpectQuery("select").
			WithArgs(
				comment.ID,
			).
			WillReturnRows(
				pgxmock.NewRows([]string{"ID", "AuthorID", "SaladId", "Text", "Rating"}).
					AddRow(comment.ID, comment.AuthorID, comment.SaladID, comment.Text, comment.Rating),
			)

		repo := postgres.NewCommentRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "request", comment.ID)
		rComment, err := repo.GetById(ctx, comment.ID)

		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(comment, rComment)
	})
}

func (suite *CommentRepoSuite) TestCommentRepo_GetById2(t provider.T) {
	t.Title("[Comment get by id] failed to get")
	t.Tags("comment", "get_by_id")
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
				pgxmock.NewRows([]string{"ID", "AuthorID", "SaladId", "Text", "Rating"}),
			)

		repo := postgres.NewCommentRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "request", id)
		rComment, err := repo.GetById(ctx, id)

		sCtx.Assert().Error(err)
		sCtx.Assert().Nil(rComment)
	})
}

func (suite *CommentRepoSuite) TestCommentRepo_GetBySaladAndUser1(t provider.T) {
	t.Title("[Comment get by salad and user] successful get")
	t.Tags("comment", "get_by_salad_and_user")
	t.Parallel()

	t.WithNewStep("successful get", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		comment := utils.NewCommentBuilder().
			WithId(uuid.New()).
			WithAuthorId(uuid.New()).
			WithSaladId(uuid.New()).
			WithText("text").
			WithRating(5).
			ToDto()

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatal(err)
		}
		defer mock.Close()
		mock.ExpectQuery("select").
			WithArgs(
				comment.SaladID,
				comment.AuthorID,
			).
			WillReturnRows(
				pgxmock.NewRows([]string{"ID", "AuthorID", "SaladId", "Text", "Rating"}).
					AddRow(comment.ID, comment.AuthorID, comment.SaladID, comment.Text, comment.Rating),
			)

		repo := postgres.NewCommentRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "request", comment.SaladID, comment.AuthorID)
		rComment, err := repo.GetBySaladAndUser(ctx, comment.SaladID, comment.AuthorID)

		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(comment, rComment)
	})
}

func (suite *CommentRepoSuite) TestCommentRepo_GetBySaladAndUser2(t provider.T) {
	t.Title("[Comment get by salad and user] failed to get")
	t.Tags("comment", "get_by_salad_and_user")
	t.Parallel()

	t.WithNewStep("failed to get", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		saladId := uuid.New()
		authorId := uuid.New()

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatal(err)
		}
		defer mock.Close()
		mock.ExpectQuery("select").
			WithArgs(
				saladId,
				authorId,
			).
			WillReturnRows(
				pgxmock.NewRows([]string{"ID", "AuthorID", "SaladId", "Text", "Rating"}),
			)

		repo := postgres.NewCommentRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "request", saladId, authorId)
		rComment, err := repo.GetBySaladAndUser(ctx, saladId, authorId)

		sCtx.Assert().Error(err)
		sCtx.Assert().Nil(rComment)
	})
}

func (suite *CommentRepoSuite) TestCommentRepo_GetAllBySaladId1(t provider.T) {
	t.Title("[Comment get all by salad id] successful get")
	t.Tags("comment", "get_all_by_salad_id")
	t.Parallel()

	t.WithNewStep("successful get", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		id := uuid.New()
		page := 1
		expComments := []*domain.Comment{
			{
				ID:       uuid.UUID{1},
				AuthorID: uuid.UUID{11},
				SaladID:  id,
				Text:     "some text",
				Rating:   5,
			},
			{
				ID:       uuid.UUID{2},
				AuthorID: uuid.UUID{22},
				SaladID:  id,
				Text:     "some text",
				Rating:   5,
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
				config.PageSize*(page-1),
				config.PageSize,
			).
			WillReturnRows(
				pgxmock.NewRows([]string{"ID", "AuthorID", "SaladId", "Text", "Rating"}).
					AddRow(expComments[0].ID, expComments[0].AuthorID, expComments[0].SaladID, expComments[0].Text, expComments[0].Rating).
					AddRow(expComments[1].ID, expComments[1].AuthorID, expComments[1].SaladID, expComments[1].Text, expComments[1].Rating),
			)

		mock.ExpectQuery("select").
			WithArgs(
				id,
			).
			WillReturnRows(
				pgxmock.NewRows([]string{"count"}).
					AddRow(2),
			)

		repo := postgres.NewCommentRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "request", id, page)
		comments, pages, err := repo.GetAllBySaladID(ctx, id, page)

		sCtx.Assert().NoError(err)
		sCtx.Assert().NotZero(pages)
		sCtx.Assert().Equal(comments, expComments)
	})
}

func (suite *CommentRepoSuite) TestCommentRepo_GetAllBySaladId2(t provider.T) {
	t.Title("[Comment get all by salad id] sql error")
	t.Tags("comment", "get_all_by_salad_id")
	t.Parallel()

	t.WithNewStep("sql error", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		id := uuid.New()
		page := 1

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatal(err)
		}
		defer mock.Close()
		mock.ExpectQuery("select").
			WithArgs(
				id,
				config.PageSize*(page-1),
				config.PageSize,
			).
			WillReturnError(fmt.Errorf("sql error"))

		repo := postgres.NewCommentRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "request", id, page)
		comments, pages, err := repo.GetAllBySaladID(ctx, id, page)

		sCtx.Assert().Error(err)
		sCtx.Assert().Zero(pages)
		sCtx.Assert().Nil(comments)
	})
}

func (suite *CommentRepoSuite) TestCommentRepo_Update1(t provider.T) {
	t.Title("[Comment update] successfully updated")
	t.Tags("comment", "update")
	t.Parallel()

	t.WithNewStep("successfully updated", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		comment := utils.NewCommentBuilder().
			WithId(uuid.New()).
			WithAuthorId(uuid.New()).
			WithSaladId(uuid.New()).
			WithText("text").
			WithRating(5).
			ToDto()

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatal(err)
		}
		defer mock.Close()
		mock.ExpectExec("update").
			WithArgs(
				comment.Rating,
				comment.Text,
				comment.ID,
			).
			WillReturnResult(pgxmock.NewResult("update", 1))

		repo := postgres.NewCommentRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "request", comment)
		err = repo.Update(ctx, comment)

		sCtx.Assert().NoError(err)
	})
}

func (suite *CommentRepoSuite) TestCommentRepo_Update2(t provider.T) {
	t.Title("[Comment update] sql error")
	t.Tags("comment", "update")
	t.Parallel()

	t.WithNewStep("sql error", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		comment := utils.NewCommentBuilder().
			WithId(uuid.New()).
			WithAuthorId(uuid.New()).
			WithSaladId(uuid.New()).
			WithText("text").
			WithRating(5).
			ToDto()

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatal(err)
		}
		defer mock.Close()
		mock.ExpectExec("update").
			WithArgs(
				comment.Rating,
				comment.Text,
				comment.ID,
			).
			WillReturnError(fmt.Errorf("sql error"))

		repo := postgres.NewCommentRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "request", comment)
		err = repo.Update(ctx, comment)

		sCtx.Assert().Error(err)
	})
}

func (suite *CommentRepoSuite) TestCommentRepo_DeleteById1(t provider.T) {
	t.Title("[Comment delete by id] successfully deleted")
	t.Tags("comment", "delete_by_id")
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

		repo := postgres.NewCommentRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "request", id)
		err = repo.DeleteById(ctx, id)

		sCtx.Assert().NoError(err)
	})
}

func (suite *CommentRepoSuite) TestCommentRepo_DeleteById2(t provider.T) {
	t.Title("[Comment delete by id] sql error")
	t.Tags("comment", "delete_by_id")
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

		repo := postgres.NewCommentRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "request", id)
		err = repo.DeleteById(ctx, id)

		sCtx.Assert().Error(err)
	})
}
