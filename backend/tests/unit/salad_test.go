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
	"ppo/services/dto"
	"ppo/tests/utils"
)

type SaladSuite struct {
	suite.Suite
}

func (suite *SaladSuite) TestSaladService_Create1(t provider.T) {
	t.Title("[Salad create] successfully created")
	t.Tags("salad", "create")
	t.Parallel()

	t.WithNewStep("successfully created", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		salad := utils.NewSaladBuilder().
			WithId(uuid.New()).
			WithAuthorId(uuid.New()).
			WithName("salad").
			WithDescription("description").
			ToDto()

		repo := mocks.NewISaladRepository(t)
		repo.On("Create", ctx, salad).Return(
			uuid.New(), nil,
		).Once()

		logger := utils.NewMockLogger()
		service := services.NewSaladService(repo, logger)

		sCtx.WithNewParameters("ctx", ctx, "request", salad)
		id, err := service.Create(ctx, salad)

		sCtx.Assert().NoError(err)
		sCtx.Assert().NotEqual(id, uuid.Nil)
	})
}

func (suite *SaladSuite) TestSaladService_Create2(t provider.T) {
	t.Title("[Salad create] empty name")
	t.Tags("salad", "create")
	t.Parallel()

	t.WithNewStep("empty name", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		salad := utils.NewSaladBuilder().
			WithId(uuid.New()).
			WithAuthorId(uuid.New()).
			WithName("").
			WithDescription("description").
			ToDto()

		repo := mocks.NewISaladRepository(t)
		repo.On("Create", ctx, salad).Return(
			uuid.Nil, nil,
		).Maybe()

		logger := utils.NewMockLogger()
		service := services.NewSaladService(repo, logger)

		sCtx.WithNewParameters("ctx", ctx, "request", salad)
		id, err := service.Create(ctx, salad)

		sCtx.Assert().Error(err)
		sCtx.Assert().Equal(id, uuid.Nil)
	})
}

func (suite *SaladSuite) TestSaladService_DeleteById1(t provider.T) {
	t.Title("[Salad delete by id] successfully deleted")
	t.Tags("salad", "delete_by_id")
	t.Parallel()

	t.WithNewStep("successfully deleted", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		id := uuid.New()

		repo := mocks.NewISaladRepository(t)
		repo.On("DeleteById", ctx, id).Return(
			nil,
		).Once()

		logger := utils.NewMockLogger()
		service := services.NewSaladService(repo, logger)

		sCtx.WithNewParameters("ctx", ctx, "request", id)
		err := service.DeleteById(ctx, id)

		sCtx.Assert().NoError(err)
	})
}

func (suite *SaladSuite) TestSaladService_DeleteById2(t provider.T) {
	t.Title("[Salad delete by id] repo error")
	t.Tags("salad", "delete_by_id")
	t.Parallel()

	t.WithNewStep("repo error", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		id := uuid.New()

		repo := mocks.NewISaladRepository(t)
		repo.On("DeleteById", ctx, id).Return(
			fmt.Errorf("repo error"),
		).Once()

		logger := utils.NewMockLogger()
		service := services.NewSaladService(repo, logger)

		sCtx.WithNewParameters("ctx", ctx, "request", id)
		err := service.DeleteById(ctx, id)

		sCtx.Assert().Error(err)
	})
}

func (suite *SaladSuite) TestSaladService_GetAll1(t provider.T) {
	t.Title("[Salad get all] successful get")
	t.Tags("salad", "get_all")
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

		repo := mocks.NewISaladRepository(t)
		repo.On("GetAll", ctx, filter, page).Return(
			expSalads, 1, nil,
		).Once()

		logger := utils.NewMockLogger()
		service := services.NewSaladService(repo, logger)

		sCtx.WithNewParameters("ctx", ctx, "request", page, filter)
		salads, pages, err := service.GetAll(ctx, filter, page)

		sCtx.Assert().NoError(err)
		sCtx.Assert().NotZero(pages)
		sCtx.Assert().Equal(salads, expSalads)
	})
}

func (suite *SaladSuite) TestSaladService_GetAll2(t provider.T) {
	t.Title("[Salad get all] repo error")
	t.Tags("salad", "get_all")
	t.Parallel()

	t.WithNewStep("repo error", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		page := 1
		filter := &dto.RecipeFilter{
			AvailableIngredients: nil,
			MinRate:              0,
			SaladTypes:           nil,
			Status:               4,
		}

		repo := mocks.NewISaladRepository(t)
		repo.On("GetAll", ctx, filter, page).Return(
			nil, 1, fmt.Errorf("repo error"),
		).Once()

		logger := utils.NewMockLogger()
		service := services.NewSaladService(repo, logger)

		sCtx.WithNewParameters("ctx", ctx, "request", filter, page)
		salads, pages, err := service.GetAll(ctx, filter, page)

		sCtx.Assert().Error(err)
		sCtx.Assert().Zero(pages)
		sCtx.Assert().Nil(salads)
	})
}

func (suite *SaladSuite) TestSaladService_GetAllRatedByUser1(t provider.T) {
	t.Title("[Salad get all] successful get")
	t.Tags("salad", "get_all_rated_by_user")
	t.Parallel()

	t.WithNewStep("successful get", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		userId := uuid.New()
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

		repo := mocks.NewISaladRepository(t)
		repo.On("GetAllRatedByUser", ctx, userId, page).Return(
			expSalads, 1, nil,
		).Once()

		logger := utils.NewMockLogger()
		service := services.NewSaladService(repo, logger)

		sCtx.WithNewParameters("ctx", ctx, "request", page, userId)
		salads, pages, err := service.GetAllRatedByUser(ctx, userId, page)

		sCtx.Assert().NoError(err)
		sCtx.Assert().NotZero(pages)
		sCtx.Assert().Equal(salads, expSalads)
	})
}

func (suite *SaladSuite) TestSaladService_GetAllRatedByUser2(t provider.T) {
	t.Title("[Salad get all] repo error")
	t.Tags("salad", "get_all_rated_by_user")
	t.Parallel()

	t.WithNewStep("repo error", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		page := 1
		userId := uuid.New()

		repo := mocks.NewISaladRepository(t)
		repo.On("GetAllRatedByUser", ctx, userId, page).Return(
			nil, 1, fmt.Errorf("repo error"),
		).Once()

		logger := utils.NewMockLogger()
		service := services.NewSaladService(repo, logger)

		sCtx.WithNewParameters("ctx", ctx, "request", userId, page)
		salads, pages, err := service.GetAllRatedByUser(ctx, userId, page)

		sCtx.Assert().Error(err)
		sCtx.Assert().Zero(pages)
		sCtx.Assert().Nil(salads)
	})
}

func (suite *SaladSuite) TestSaladService_GetAllByUserId1(t provider.T) {
	t.Title("[Salad get all] successful get")
	t.Tags("salad", "get_all_by_user_id")
	t.Parallel()

	t.WithNewStep("successful get", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		userId := uuid.New()
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

		repo := mocks.NewISaladRepository(t)
		repo.On("GetAllByUserId", ctx, userId).Return(
			expSalads, nil,
		).Once()

		logger := utils.NewMockLogger()
		service := services.NewSaladService(repo, logger)

		sCtx.WithNewParameters("ctx", ctx, "request", userId)
		salads, err := service.GetAllByUserId(ctx, userId)

		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(salads, expSalads)
	})
}

func (suite *SaladSuite) TestSaladService_GetAllByUserId2(t provider.T) {
	t.Title("[Salad get all] repo error")
	t.Tags("salad", "get_all_by_user_id")
	t.Parallel()

	t.WithNewStep("repo error", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		userId := uuid.New()

		repo := mocks.NewISaladRepository(t)
		repo.On("GetAllByUserId", ctx, userId).Return(
			nil, fmt.Errorf("repo error"),
		).Once()

		logger := utils.NewMockLogger()
		service := services.NewSaladService(repo, logger)

		sCtx.WithNewParameters("ctx", ctx, "request", userId)
		salads, err := service.GetAllByUserId(ctx, userId)

		sCtx.Assert().Error(err)
		sCtx.Assert().Nil(salads)
	})
}

func (suite *SaladSuite) TestSaladService_GetById1(t provider.T) {
	t.Title("[Salad get by id] successful get")
	t.Tags("salad", "get_by_id")
	t.Parallel()

	t.WithNewStep("successful get", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		expSalad := utils.NewSaladBuilder().
			WithId(uuid.New()).
			WithAuthorId(uuid.New()).
			WithName("salad").
			WithDescription("description").
			ToDto()

		repo := mocks.NewISaladRepository(t)
		repo.On("GetById", ctx, expSalad.ID).Return(
			expSalad, nil,
		).Once()

		logger := utils.NewMockLogger()
		service := services.NewSaladService(repo, logger)

		sCtx.WithNewParameters("ctx", ctx, "request", expSalad.ID)
		Ingredient, err := service.GetById(ctx, expSalad.ID)

		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(Ingredient, expSalad)
	})
}

func (suite *SaladSuite) TestSaladService_GetById2(t provider.T) {
	t.Title("[Salad get by id] repo error")
	t.Tags("salad", "get_by_id")
	t.Parallel()

	t.WithNewStep("repo error", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		id := uuid.New()

		repo := mocks.NewISaladRepository(t)
		repo.On("GetById", ctx, id).Return(
			nil, fmt.Errorf("repo error"),
		).Once()

		logger := utils.NewMockLogger()
		service := services.NewSaladService(repo, logger)

		sCtx.WithNewParameters("ctx", ctx, "request", id)
		comment, err := service.GetById(ctx, id)

		sCtx.Assert().Error(err)
		sCtx.Assert().Nil(comment)
	})
}

func (suite *SaladSuite) TestSaladService_Update1(t provider.T) {
	t.Title("[Salad update] successfully updated")
	t.Tags("salad", "update")
	t.Parallel()

	t.WithNewStep("successfully updated", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		expSalad := utils.NewSaladBuilder().
			WithId(uuid.New()).
			WithAuthorId(uuid.New()).
			WithName("salad").
			WithDescription("description").
			ToDto()

		repo := mocks.NewISaladRepository(t)
		repo.On("Update", ctx, expSalad).Return(
			nil,
		).Once()

		logger := utils.NewMockLogger()
		service := services.NewSaladService(repo, logger)

		sCtx.WithNewParameters("ctx", ctx, "request", expSalad)
		err := service.Update(ctx, expSalad)

		sCtx.Assert().NoError(err)
	})
}

func (suite *SaladSuite) TestSaladService_Update2(t provider.T) {
	t.Title("[Salad update] repo error")
	t.Tags("salad", "update")
	t.Parallel()

	t.WithNewStep("repo error", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		expSalad := utils.NewSaladBuilder().
			WithId(uuid.New()).
			WithAuthorId(uuid.New()).
			WithName("salad").
			WithDescription("description").
			ToDto()

		repo := mocks.NewISaladRepository(t)
		repo.On("Update", ctx, expSalad).Return(
			fmt.Errorf("repo error"),
		).Once()

		logger := utils.NewMockLogger()
		service := services.NewSaladService(repo, logger)

		sCtx.WithNewParameters("ctx", ctx, "request", expSalad)
		err := service.Update(ctx, expSalad)

		sCtx.Assert().Error(err)
	})
}
