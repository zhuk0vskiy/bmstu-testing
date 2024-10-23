//go:build unit_test

package unit_tests

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
	"ppo/mocks"
	"ppo/services"
	"ppo/tests/utils"
)

type KeywordsValidatorSuite struct {
	suite.Suite
}

func (suite *KeywordsValidatorSuite) TestKeywordsValidatorService_Create1(t provider.T) {
	t.Title("[Keyword validator create] successfully created")
	t.Tags("keyword_validator", "create")
	t.Parallel()

	t.WithNewStep("successfully created", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		keyword := utils.NewKeywordBuilder().
			WithId(uuid.New()).
			WithWord("banned_word").
			ToDto()

		repo := mocks.NewIKeywordValidatorRepository(t)
		repo.On("GetAll", ctx).Return(
			map[string]uuid.UUID{
				"banned": uuid.New(),
			}, nil).Once()
		repo.On("Create", ctx, keyword).Return(
			nil,
		).Once()

		logger := utils.NewMockLogger()
		service, _ := services.NewKeywordValidatorService(ctx, repo, logger)

		sCtx.WithNewParameters("ctx", ctx, "request", keyword)
		err := service.Create(ctx, keyword)

		sCtx.Assert().NoError(err)
	})
}

func (suite *KeywordsValidatorSuite) TestKeywordsValidatorService_Create2(t provider.T) {
	t.Title("[Keyword validator create] empty word")
	t.Tags("keyword_validator", "create")
	t.Parallel()

	t.WithNewStep("successfully created", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		keyword := utils.NewKeywordBuilder().
			WithId(uuid.New()).
			WithWord("").
			ToDto()

		repo := mocks.NewIKeywordValidatorRepository(t)
		repo.On("GetAll", ctx).Return(
			map[string]uuid.UUID{
				"banned": uuid.New(),
			}, nil).Once()

		logger := utils.NewMockLogger()
		service, _ := services.NewKeywordValidatorService(ctx, repo, logger)

		sCtx.WithNewParameters("ctx", ctx, "request", keyword)
		err := service.Create(ctx, keyword)

		sCtx.Assert().Error(err)
	})
}

func (suite *KeywordsValidatorSuite) TestKeywordsValidatorService_DeleteById1(t provider.T) {
	t.Title("[keyword validator delete by id] successfully deleted")
	t.Tags("keyword_validator", "delete_by_id")
	t.Parallel()

	t.WithNewStep("successfully deleted", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		id := uuid.New()

		repo := mocks.NewIKeywordValidatorRepository(t)
		repo.On("GetAll", ctx).Return(
			map[string]uuid.UUID{
				"banned": uuid.New(),
			}, nil).Once()
		repo.On("DeleteById", ctx, id).Return(
			nil,
		).Once()

		logger := utils.NewMockLogger()
		service, _ := services.NewKeywordValidatorService(ctx, repo, logger)

		sCtx.WithNewParameters("ctx", ctx, "request", id)
		err := service.DeleteById(ctx, id)

		sCtx.Assert().NoError(err)
	})
}

func (suite *KeywordsValidatorSuite) TestKeywordsValidatorService_DeleteById2(t provider.T) {
	t.Title("[keyword validator delete by id] repo error")
	t.Tags("keyword_validator", "delete_by_id")
	t.Parallel()

	t.WithNewStep("repo_error", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		id := uuid.New()

		repo := mocks.NewIKeywordValidatorRepository(t)
		repo.On("GetAll", ctx).Return(
			map[string]uuid.UUID{
				"banned": uuid.New(),
			}, nil).Once()
		repo.On("DeleteById", ctx, id).Return(
			fmt.Errorf("repo error"),
		).Once()

		logger := utils.NewMockLogger()
		service, _ := services.NewKeywordValidatorService(ctx, repo, logger)

		sCtx.WithNewParameters("ctx", ctx, "request", id)
		err := service.DeleteById(ctx, id)

		sCtx.Assert().Error(err)
	})
}

func (suite *KeywordsValidatorSuite) TestKeywordsValidatorService_GetById1(t provider.T) {
	t.Title("[keyword validator get by id] successful get")
	t.Tags("keyword_validator", "get_by_id")
	t.Parallel()

	t.WithNewStep("successful get", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		expKeyword := utils.NewKeywordBuilder().
			WithId(uuid.New()).
			WithWord("word").
			ToDto()

		repo := mocks.NewIKeywordValidatorRepository(t)
		repo.On("GetAll", ctx).Return(
			map[string]uuid.UUID{
				"banned": uuid.New(),
			}, nil).Once()
		repo.On("GetById", ctx, expKeyword.ID).Return(
			expKeyword, nil,
		).Once()

		logger := utils.NewMockLogger()
		service, _ := services.NewKeywordValidatorService(ctx, repo, logger)

		sCtx.WithNewParameters("ctx", ctx, "request", expKeyword.ID)
		keyword, err := service.GetById(ctx, expKeyword.ID)

		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(expKeyword, keyword)
	})
}

func (suite *KeywordsValidatorSuite) TestKeywordsValidatorService_GetById2(t provider.T) {
	t.Title("[keyword validator get by id] repo error")
	t.Tags("keyword_validator", "get_by_id")
	t.Parallel()

	t.WithNewStep("repo error", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		id := uuid.New()

		repo := mocks.NewIKeywordValidatorRepository(t)
		repo.On("GetAll", ctx).Return(
			map[string]uuid.UUID{
				"banned": uuid.New(),
			}, nil).Once()
		repo.On("GetById", ctx, id).Return(
			nil, fmt.Errorf("repo error"),
		).Once()

		logger := utils.NewMockLogger()
		service, _ := services.NewKeywordValidatorService(ctx, repo, logger)

		sCtx.WithNewParameters("ctx", ctx, "request", id)
		keyword, err := service.GetById(ctx, id)

		sCtx.Assert().Error(err)
		sCtx.Assert().Nil(keyword)
	})
}

func (suite *KeywordsValidatorSuite) TestKeywordsValidatorService_Update1(t provider.T) {
	t.Title("[Keyword validator update] successfully updated")
	t.Tags("keyword_validator", "update")
	t.Parallel()

	t.WithNewStep("successfully updated", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		keyword := utils.NewKeywordBuilder().
			WithId(uuid.New()).
			WithWord("banned_word").
			ToDto()

		repo := mocks.NewIKeywordValidatorRepository(t)
		repo.On("GetAll", ctx).Return(
			map[string]uuid.UUID{
				"banned": uuid.New(),
			}, nil).Once()
		repo.On("Update", ctx, keyword).Return(
			nil,
		).Once()

		logger := utils.NewMockLogger()
		service, _ := services.NewKeywordValidatorService(ctx, repo, logger)

		sCtx.WithNewParameters("ctx", ctx, "request", keyword)
		err := service.Update(ctx, keyword)

		sCtx.Assert().NoError(err)
	})
}

func (suite *KeywordsValidatorSuite) TestKeywordsValidatorService_Update2(t provider.T) {
	t.Title("[Keyword validator update] empty word")
	t.Tags("keyword_validator", "update")
	t.Parallel()

	t.WithNewStep("empty word", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		keyword := utils.NewKeywordBuilder().
			WithId(uuid.New()).
			WithWord("").
			ToDto()

		repo := mocks.NewIKeywordValidatorRepository(t)
		repo.On("GetAll", ctx).Return(
			map[string]uuid.UUID{
				"banned": uuid.New(),
			}, nil).Once()

		logger := utils.NewMockLogger()
		service, _ := services.NewKeywordValidatorService(ctx, repo, logger)

		sCtx.WithNewParameters("ctx", ctx, "request", keyword)
		err := service.Update(ctx, keyword)

		sCtx.Assert().Error(err)
	})
}

func (suite *KeywordsValidatorSuite) TestKeywordsValidatorService_Verify1(t provider.T) {
	t.Title("[Keyword validator verify] verified")
	t.Tags("keyword_validator", "verify")
	t.Parallel()

	t.WithNewStep("verified", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		word := "to_check"

		repo := mocks.NewIKeywordValidatorRepository(t)
		repo.On("GetAll", ctx).Return(
			map[string]uuid.UUID{
				"banned": uuid.New(),
			}, nil).Once()

		logger := utils.NewMockLogger()
		service, _ := services.NewKeywordValidatorService(ctx, repo, logger)

		sCtx.WithNewParameters("ctx", ctx, "request", word)
		err := service.Verify(ctx, word)

		sCtx.Assert().NoError(err)
	})
}

func (suite *KeywordsValidatorSuite) TestKeywordsValidatorService_Verify2(t provider.T) {
	t.Title("[Keyword validator verify] verification failed")
	t.Tags("keyword_validator", "verify")
	t.Parallel()

	t.WithNewStep("verification failed", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		word := "banned"

		repo := mocks.NewIKeywordValidatorRepository(t)
		repo.On("GetAll", ctx).Return(
			map[string]uuid.UUID{
				"banned": uuid.New(),
			}, nil).Once()

		logger := utils.NewMockLogger()
		service, _ := services.NewKeywordValidatorService(ctx, repo, logger)

		sCtx.WithNewParameters("ctx", ctx, "request", word)
		err := service.Verify(ctx, word)

		sCtx.Assert().Error(err)
	})
}

func (suite *KeywordsValidatorSuite) TestKeywordsValidatorService_GetAll1(t provider.T) {
	t.Title("[Keyword validator get all] successful get")
	t.Tags("keyword_validator", "get_all")
	t.Parallel()

	t.WithNewStep("successful get", func(sCtx provider.StepCtx) {
		ctx := context.TODO()

		keywords := map[string]uuid.UUID{
			"banned1": uuid.New(),
			"banned2": uuid.New(),
			"banned3": uuid.New(),
			"banned4": uuid.New(),
			"banned5": uuid.New(),
		}

		repo := mocks.NewIKeywordValidatorRepository(t)
		repo.On("GetAll", ctx).Return(keywords, nil).Maybe()

		logger := utils.NewMockLogger()
		service, _ := services.NewKeywordValidatorService(ctx, repo, logger)

		sCtx.WithNewParameters("ctx", ctx)
		keywords, err := service.GetAll(ctx)

		sCtx.Assert().NoError(err)
	})
}

func (suite *KeywordsValidatorSuite) TestKeywordsValidatorService_GetAll2(t provider.T) {
	t.Title("[Keyword validator get all] repo error")
	t.Tags("keyword_validator", "get_all")
	t.Parallel()

	t.WithNewStep("repo error", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		repo := mocks.NewIKeywordValidatorRepository(t)
		repo.On("GetAll", ctx).
			Return(nil, fmt.Errorf("repo error")).Maybe()

		logger := utils.NewMockLogger()
		service, err := services.NewKeywordValidatorService(ctx, repo, logger)

		sCtx.WithNewParameters("ctx", ctx)
		//keywords, err := service.GetAll(ctx)

		sCtx.Assert().Error(err)
		sCtx.Assert().Nil(service)
	})
}
