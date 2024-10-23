//go:build unit_test

package unit_tests

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
	"ppo/domain"
	"ppo/mocks"
	"ppo/services"
	"ppo/tests/utils"
)

type CommentSuite struct {
	suite.Suite
}

func (suite *CommentSuite) TestCommentService_Create1(t provider.T) {
	t.Title("[Comment Create] successfully created")
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

		repo := mocks.NewICommentRepository(t)
		repo.On("Create", ctx, comment).Return(
			nil,
		).Once()

		logger := utils.NewMockLogger()
		service := services.NewCommentService(repo, logger)

		sCtx.WithNewParameters("ctx", ctx, "request", comment)
		err := service.Create(ctx, comment)

		sCtx.Assert().NoError(err)
	})
}

func (suite *CommentSuite) TestCommentService_Create2(t provider.T) {
	t.Title("[Comment create] rating below minimum border")
	t.Tags("comment", "create")
	t.Parallel()

	t.WithNewStep("rating below minimum border", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		comment := utils.NewCommentBuilder().
			WithId(uuid.New()).
			WithAuthorId(uuid.New()).
			WithSaladId(uuid.New()).
			WithText("text").
			WithRating(-1).
			ToDto()

		repo := mocks.NewICommentRepository(t)
		repo.On("Create", ctx, comment).Return(
			nil,
		).Maybe()

		logger := utils.NewMockLogger()
		service := services.NewCommentService(repo, logger)

		sCtx.WithNewParameters("ctx", ctx, "request", comment)
		err := service.Create(ctx, comment)

		sCtx.Assert().Error(err)
	})
}

func (suite *CommentSuite) TestCommentService_DeleteById1(t provider.T) {
	t.Title("[Comment delete by id] successfully deleted")
	t.Tags("comment", "delete_by_id")
	t.Parallel()

	t.WithNewStep("successfully deleted", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		id := uuid.New()

		repo := mocks.NewICommentRepository(t)
		repo.On("DeleteById", ctx, id).Return(
			nil,
		).Once()

		logger := utils.NewMockLogger()
		service := services.NewCommentService(repo, logger)

		sCtx.WithNewParameters("ctx", ctx, "request", id)
		err := service.DeleteById(ctx, id)

		sCtx.Assert().NoError(err)
	})
}

func (suite *CommentSuite) TestCommentService_DeleteById2(t provider.T) {
	t.Title("[Comment delete by id] repo error")
	t.Tags("comment", "delete_by_id")
	t.Parallel()

	t.WithNewStep("repo error", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		id := uuid.New()

		repo := mocks.NewICommentRepository(t)
		repo.On("DeleteById", ctx, id).Return(
			fmt.Errorf("repo error"),
		).Once()

		logger := utils.NewMockLogger()
		service := services.NewCommentService(repo, logger)

		sCtx.WithNewParameters("ctx", ctx, "request", id)
		err := service.DeleteById(ctx, id)

		sCtx.Assert().Error(err)
	})
}

func (suite *CommentSuite) TestCommentService_GetAllBySaladId1(t provider.T) {
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

		repo := mocks.NewICommentRepository(t)
		repo.On("GetAllBySaladID", ctx, id, page).Return(
			expComments, 1, nil,
		).Once()

		logger := utils.NewMockLogger()
		service := services.NewCommentService(repo, logger)

		sCtx.WithNewParameters("ctx", ctx, "request", id)
		comments, pages, err := service.GetAllBySaladID(ctx, id, page)

		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(comments, expComments)
		sCtx.Assert().NotEmpty(pages)
	})
}

func (suite *CommentSuite) TestCommentService_GetAllBySaladId2(t provider.T) {
	t.Title("[Comment get all by salad id] repo error")
	t.Tags("comment", "get_all_by_salad_id")
	t.Parallel()

	t.WithNewStep("repo error", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		id := uuid.New()
		page := 1

		repo := mocks.NewICommentRepository(t)
		repo.On("GetAllBySaladID", ctx, id, page).Return(
			nil, 0, fmt.Errorf("repo error"),
		).Once()

		logger := utils.NewMockLogger()
		service := services.NewCommentService(repo, logger)

		sCtx.WithNewParameters("ctx", ctx, "request", id)
		comments, pages, err := service.GetAllBySaladID(ctx, id, page)

		sCtx.Assert().Error(err)
		sCtx.Assert().Nil(comments)
		sCtx.Assert().Zero(pages)
	})
}

func (suite *CommentSuite) TestCommentService_GetById1(t provider.T) {
	t.Title("[Comment get by id] successful get")
	t.Tags("comment", "get_by_id")
	t.Parallel()

	t.WithNewStep("successful get", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		id := uuid.New()
		expComment := utils.NewCommentBuilder().
			WithId(id).
			WithAuthorId(uuid.New()).
			WithSaladId(uuid.New()).
			WithText("text").
			WithRating(5).
			ToDto()

		repo := mocks.NewICommentRepository(t)
		repo.On("GetById", ctx, expComment.ID).Return(
			expComment, nil,
		).Once()

		logger := utils.NewMockLogger()
		service := services.NewCommentService(repo, logger)

		sCtx.WithNewParameters("ctx", ctx, "request", expComment.ID)
		comment, err := service.GetById(ctx, expComment.ID)

		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(comment, expComment)
	})
}

func (suite *CommentSuite) TestCommentService_GetById2(t provider.T) {
	t.Title("[Comment get by id] repo error")
	t.Tags("comment", "get_by_id")
	t.Parallel()

	t.WithNewStep("repo error", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		id := uuid.New()

		repo := mocks.NewICommentRepository(t)
		repo.On("GetById", ctx, id).Return(
			nil, fmt.Errorf("repo error"),
		).Once()

		logger := utils.NewMockLogger()
		service := services.NewCommentService(repo, logger)

		sCtx.WithNewParameters("ctx", ctx, "request", id)
		comment, err := service.GetById(ctx, id)

		sCtx.Assert().Error(err)
		sCtx.Assert().Nil(comment)
	})
}

func (suite *CommentSuite) TestCommentService_GetBySaladAndUser1(t provider.T) {
	t.Title("[Comment get by salad and user] successful get")
	t.Tags("comment", "get_by_salad_and_user")
	t.Parallel()

	t.WithNewStep("successful get", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		expComment := utils.NewCommentBuilder().
			WithId(uuid.New()).
			WithAuthorId(uuid.New()).
			WithSaladId(uuid.New()).
			WithText("text").
			WithRating(5).
			ToDto()

		repo := mocks.NewICommentRepository(t)
		repo.On("GetBySaladAndUser", ctx, expComment.SaladID, expComment.AuthorID).Return(
			expComment, nil,
		).Once()

		logger := utils.NewMockLogger()
		service := services.NewCommentService(repo, logger)

		sCtx.WithNewParameters("ctx", ctx, "request", expComment.SaladID, expComment.AuthorID)
		comment, err := service.GetBySaladAndUser(ctx, expComment.SaladID, expComment.AuthorID)

		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(comment, expComment)
	})
}

func (suite *CommentSuite) TestCommentService_GetBySaladAndUser2(t provider.T) {
	t.Title("[Comment get by salad and user] repo error")
	t.Tags("comment", "get_by_salad_and_user")
	t.Parallel()

	t.WithNewStep("repo error", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		saladId := uuid.New()
		userId := uuid.New()

		repo := mocks.NewICommentRepository(t)
		repo.On("GetBySaladAndUser", ctx, saladId, userId).Return(
			nil, fmt.Errorf("repo error"),
		).Once()

		logger := utils.NewMockLogger()
		service := services.NewCommentService(repo, logger)

		sCtx.WithNewParameters("ctx", ctx, "request", saladId, userId)
		comment, err := service.GetBySaladAndUser(ctx, saladId, userId)

		sCtx.Assert().Error(err)
		sCtx.Assert().Nil(comment)
	})
}

func (suite *CommentSuite) TestCommentService_Update1(t provider.T) {
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

		repo := mocks.NewICommentRepository(t)
		repo.On("Update", ctx, comment).Return(
			nil,
		).Once()

		logger := utils.NewMockLogger()
		service := services.NewCommentService(repo, logger)

		sCtx.WithNewParameters("ctx", ctx, "request", comment)
		err := service.Update(ctx, comment)

		sCtx.Assert().NoError(err)
	})
}

func (suite *CommentSuite) TestCommentService_Update2(t provider.T) {
	t.Title("[Comment update] repo error")
	t.Tags("comment", "update")
	t.Parallel()

	t.WithNewStep("repo error", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		comment := utils.NewCommentBuilder().
			WithId(uuid.New()).
			WithAuthorId(uuid.New()).
			WithSaladId(uuid.New()).
			WithText("text").
			WithRating(5).
			ToDto()

		repo := mocks.NewICommentRepository(t)
		repo.On("Update", ctx, comment).Return(
			fmt.Errorf("repo error"),
		).Once()

		logger := utils.NewMockLogger()
		service := services.NewCommentService(repo, logger)

		sCtx.WithNewParameters("ctx", ctx, "request", comment)
		err := service.Update(ctx, comment)

		sCtx.Assert().Error(err)
	})
}
