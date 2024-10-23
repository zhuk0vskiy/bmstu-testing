//go:build unit_test

package unit_tests

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
	"github.com/stretchr/testify/mock"
	"ppo/domain"
	"ppo/mocks"
	"ppo/services"
	"ppo/services/dto"
	"ppo/tests/utils"
)

type SaladInteractorSuite struct {
	suite.Suite
}

func (suite *SaladInteractorSuite) TestSaladInteractorService_Create1(t provider.T) {
	t.Title("[Salad interactor create] successfully created")
	t.Tags("salad_interactor", "create")
	t.Parallel()

	t.WithNewStep("successfully created", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		salad := utils.NewSaladBuilder().
			WithId(uuid.New()).
			WithAuthorId(uuid.New()).
			WithName("salad").
			WithDescription("description").
			ToDto()

		saladService := mocks.NewISaladService(t)
		saladService.On("Create", ctx, salad).Return(
			uuid.New(), nil,
		)

		validator := mocks.NewIValidatorService(t)
		validator.On("Verify", ctx, mock.Anything).Return(
			nil,
		).Maybe()

		service := services.NewSaladInteractor(saladService, []domain.IValidatorService{
			validator,
		})

		sCtx.WithNewParameters("ctx", ctx, "request", salad)
		id, err := service.Create(ctx, salad)

		sCtx.Assert().NoError(err)
		sCtx.Assert().NotEqual(id, uuid.Nil)
	})
}

func (suite *SaladInteractorSuite) TestSaladInteractorService_Create2(t provider.T) {
	t.Title("[Salad interactor create] service error")
	t.Tags("salad_interactor", "create")
	t.Parallel()

	t.WithNewStep("service error", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		salad := utils.NewSaladBuilder().
			WithId(uuid.New()).
			WithAuthorId(uuid.New()).
			WithName("salad").
			WithDescription("description").
			ToDto()

		saladService := mocks.NewISaladService(t)
		saladService.On("Create", ctx, salad).Return(
			uuid.Nil, fmt.Errorf("repo error"),
		)

		validator := mocks.NewIValidatorService(t)
		validator.On("Verify", ctx, mock.Anything).Return(
			nil,
		).Maybe()

		service := services.NewSaladInteractor(saladService, []domain.IValidatorService{
			validator,
		})

		sCtx.WithNewParameters("ctx", ctx, "request", salad)
		id, err := service.Create(ctx, salad)

		sCtx.Assert().Error(err)
		sCtx.Assert().Equal(id, uuid.Nil)
	})
}

func (suite *SaladInteractorSuite) TestSaladInteractorService_DeleteById1(t provider.T) {
	t.Title("[Salad interactor delete by id] successfully deleted")
	t.Tags("salad_interactor", "delete_by_id")
	t.Parallel()

	t.WithNewStep("successfully deleted", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		id := uuid.New()

		saladService := mocks.NewISaladService(t)
		saladService.On("DeleteById", ctx, id).Return(
			nil,
		)

		validator := mocks.NewIValidatorService(t)
		validator.On("Verify", ctx, mock.Anything).Return(
			nil,
		).Maybe()

		service := services.NewSaladInteractor(saladService, []domain.IValidatorService{
			validator,
		})

		sCtx.WithNewParameters("ctx", ctx, "request", id)
		err := service.DeleteById(ctx, id)

		sCtx.Assert().NoError(err)
	})
}

func (suite *SaladInteractorSuite) TestSaladInteractorService_DeleteById2(t provider.T) {
	t.Title("[Salad interactor delete by id] service error")
	t.Tags("salad_interactor", "delete_by_id")
	t.Parallel()

	t.WithNewStep("service error", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		id := uuid.New()

		saladService := mocks.NewISaladService(t)
		saladService.On("DeleteById", ctx, id).Return(
			nil,
		)

		validator := mocks.NewIValidatorService(t)
		validator.On("Verify", ctx, mock.Anything).Return(
			nil,
		).Maybe()

		service := services.NewSaladInteractor(saladService, []domain.IValidatorService{
			validator,
		})

		sCtx.WithNewParameters("ctx", ctx, "request", id)
		err := service.DeleteById(ctx, id)

		sCtx.Assert().NoError(err)
	})
}

func (suite *SaladInteractorSuite) TestSaladInteractorService_GetAll1(t provider.T) {
	t.Title("[Salad interactor get all] successful get")
	t.Tags("salad_interactor", "get_all")
	t.Parallel()

	t.WithNewStep("successful get", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		page := 1
		filter := &dto.RecipeFilter{
			AvailableIngredients: nil,
			MinRate:              0,
			SaladTypes:           nil,
			Status:               4,
		}
		expSalads := []*domain.Salad{}

		saladService := mocks.NewISaladService(t)
		saladService.On("GetAll", ctx, filter, page).Return(
			expSalads, 1, nil,
		)

		validator := mocks.NewIValidatorService(t)
		validator.On("Verify", ctx, mock.Anything).Return(
			nil,
		).Maybe()

		service := services.NewSaladInteractor(saladService, []domain.IValidatorService{
			validator,
		})

		sCtx.WithNewParameters("ctx", ctx, "request", filter, page)
		salads, pages, err := service.GetAll(ctx, filter, page)

		sCtx.Assert().NoError(err)
		sCtx.Assert().NotZero(pages)
		sCtx.Assert().Equal(salads, expSalads)
	})
}

func (suite *SaladInteractorSuite) TestSaladInteractorService_GetAll2(t provider.T) {
	t.Title("[Salad interactor get all] service error")
	t.Tags("salad_interactor", "get_all")
	t.Parallel()

	t.WithNewStep("service error", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		page := 1
		filter := &dto.RecipeFilter{
			AvailableIngredients: nil,
			MinRate:              0,
			SaladTypes:           nil,
			Status:               4,
		}

		saladService := mocks.NewISaladService(t)
		saladService.On("GetAll", ctx, filter, page).Return(
			nil, 0, fmt.Errorf("repo error"),
		)

		validator := mocks.NewIValidatorService(t)
		validator.On("Verify", ctx, mock.Anything).Return(
			nil,
		).Maybe()

		service := services.NewSaladInteractor(saladService, []domain.IValidatorService{
			validator,
		})

		sCtx.WithNewParameters("ctx", ctx, "request", filter, page)
		salads, pages, err := service.GetAll(ctx, filter, page)

		sCtx.Assert().Error(err)
		sCtx.Assert().Zero(pages)
		sCtx.Assert().Nil(salads)
	})
}

func (suite *SaladInteractorSuite) TestSaladInteractorService_GetAllRatedByUser1(t provider.T) {
	t.Title("[Salad interactor get all rated by user] successful get")
	t.Tags("salad_interactor", "get_all_rated_by_user")
	t.Parallel()

	t.WithNewStep("successful get", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		page := 1
		userId := uuid.New()
		expSalads := []*domain.Salad{}

		saladService := mocks.NewISaladService(t)
		saladService.On("GetAllRatedByUser", ctx, userId, page).Return(
			expSalads, 1, nil,
		)

		validator := mocks.NewIValidatorService(t)
		validator.On("Verify", ctx, mock.Anything).Return(
			nil,
		).Maybe()

		service := services.NewSaladInteractor(saladService, []domain.IValidatorService{
			validator,
		})

		sCtx.WithNewParameters("ctx", ctx, "request", userId, page)
		salads, pages, err := service.GetAllRatedByUser(ctx, userId, page)

		sCtx.Assert().NoError(err)
		sCtx.Assert().NotZero(pages)
		sCtx.Assert().Equal(salads, expSalads)
	})
}

func (suite *SaladInteractorSuite) TestSaladInteractorService_GetAllRatedByUser2(t provider.T) {
	t.Title("[Salad interactor get all rated by user] service error")
	t.Tags("salad_interactor", "get_all_rated_by_user")
	t.Parallel()

	t.WithNewStep("service error", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		page := 1
		userId := uuid.New()

		saladService := mocks.NewISaladService(t)
		saladService.On("GetAllRatedByUser", ctx, userId, page).Return(
			nil, 0, nil,
		)

		validator := mocks.NewIValidatorService(t)
		validator.On("Verify", ctx, mock.Anything).Return(
			nil,
		).Maybe()

		service := services.NewSaladInteractor(saladService, []domain.IValidatorService{
			validator,
		})

		sCtx.WithNewParameters("ctx", ctx, "request", userId, page)
		salads, pages, err := service.GetAllRatedByUser(ctx, userId, page)

		sCtx.Assert().NoError(err)
		sCtx.Assert().Zero(pages)
		sCtx.Assert().Nil(salads)
	})
}

func (suite *SaladInteractorSuite) TestSaladInteractorService_GetAllByUserId1(t provider.T) {
	t.Title("[Salad interactor get all rated by user] successful get")
	t.Tags("salad_interactor", "get_all_rated_by_user")
	t.Parallel()

	t.WithNewStep("successful get", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		userId := uuid.New()
		expSalads := []*domain.Salad{}

		saladService := mocks.NewISaladService(t)
		saladService.On("GetAllByUserId", ctx, userId).Return(
			expSalads, nil,
		)

		validator := mocks.NewIValidatorService(t)
		validator.On("Verify", ctx, mock.Anything).Return(
			nil,
		).Maybe()

		service := services.NewSaladInteractor(saladService, []domain.IValidatorService{
			validator,
		})

		sCtx.WithNewParameters("ctx", ctx, "request", userId)
		salads, err := service.GetAllByUserId(ctx, userId)

		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(salads, expSalads)
	})
}

func (suite *SaladInteractorSuite) TestSaladInteractorService_GetAllByUserId2(t provider.T) {
	t.Title("[Salad interactor get all rated by user] service error")
	t.Tags("salad_interactor", "get_all_rated_by_user")
	t.Parallel()

	t.WithNewStep("service error", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		userId := uuid.New()

		saladService := mocks.NewISaladService(t)
		saladService.On("GetAllByUserId", ctx, userId).Return(
			nil, nil,
		)

		validator := mocks.NewIValidatorService(t)
		validator.On("Verify", ctx, mock.Anything).Return(
			nil,
		).Maybe()

		service := services.NewSaladInteractor(saladService, []domain.IValidatorService{
			validator,
		})

		sCtx.WithNewParameters("ctx", ctx, "request", userId)
		salads, err := service.GetAllByUserId(ctx, userId)

		sCtx.Assert().NoError(err)
		sCtx.Assert().Nil(salads)
	})
}

func (suite *SaladInteractorSuite) TestSaladInteractorService_GetById1(t provider.T) {
	t.Title("[Salad interactor create] successful get")
	t.Tags("salad_interactor", "get_by_id")
	t.Parallel()

	t.WithNewStep("successful get", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		expSalad := utils.NewSaladBuilder().
			WithId(uuid.New()).
			WithAuthorId(uuid.New()).
			WithName("salad").
			WithDescription("description").
			ToDto()

		saladService := mocks.NewISaladService(t)
		saladService.On("GetById", ctx, expSalad.ID).Return(
			expSalad, nil,
		)

		validator := mocks.NewIValidatorService(t)
		validator.On("Verify", ctx, mock.Anything).Return(
			nil,
		).Maybe()

		service := services.NewSaladInteractor(saladService, []domain.IValidatorService{
			validator,
		})

		sCtx.WithNewParameters("ctx", ctx, "request", expSalad.ID)
		salad, err := service.GetById(ctx, expSalad.ID)

		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(salad, expSalad)
	})
}

func (suite *SaladInteractorSuite) TestSaladInteractorService_GetById2(t provider.T) {
	t.Title("[Salad interactor create] successful get")
	t.Tags("salad_interactor", "get_by_id")
	t.Parallel()

	t.WithNewStep("successful get", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		id := uuid.New()

		saladService := mocks.NewISaladService(t)
		saladService.On("GetById", ctx, id).Return(
			nil, fmt.Errorf("ervice error"),
		)

		validator := mocks.NewIValidatorService(t)
		validator.On("Verify", ctx, mock.Anything).Return(
			nil,
		).Maybe()

		service := services.NewSaladInteractor(saladService, []domain.IValidatorService{
			validator,
		})

		sCtx.WithNewParameters("ctx", ctx, "request", id)
		salad, err := service.GetById(ctx, id)

		sCtx.Assert().Error(err)
		sCtx.Assert().Nil(salad)
	})
}

func (suite *SaladInteractorSuite) TestSaladInteractorService_Update1(t provider.T) {
	t.Title("[Salad interactor create] successfully updated")
	t.Tags("salad_interactor", "update")
	t.Parallel()

	t.WithNewStep("successfully updated", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		expSalad := utils.NewSaladBuilder().
			WithId(uuid.New()).
			WithAuthorId(uuid.New()).
			WithName("salad").
			WithDescription("description").
			ToDto()

		saladService := mocks.NewISaladService(t)
		saladService.On("Update", ctx, expSalad).Return(
			nil,
		)

		validator := mocks.NewIValidatorService(t)
		validator.On("Verify", ctx, mock.Anything).Return(
			nil,
		).Maybe()

		service := services.NewSaladInteractor(saladService, []domain.IValidatorService{
			validator,
		})

		sCtx.WithNewParameters("ctx", ctx, "request", expSalad)
		err := service.Update(ctx, expSalad)

		sCtx.Assert().NoError(err)
	})
}

func (suite *SaladInteractorSuite) TestSaladInteractorService_Update2(t provider.T) {
	t.Title("[Salad interactor create] serviece error")
	t.Tags("salad_interactor", "update")
	t.Parallel()

	t.WithNewStep("service error", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		expSalad := utils.NewSaladBuilder().
			WithId(uuid.New()).
			WithAuthorId(uuid.New()).
			WithName("salad").
			WithDescription("description").
			ToDto()

		saladService := mocks.NewISaladService(t)
		saladService.On("Update", ctx, expSalad).Return(
			fmt.Errorf("service error"),
		)

		validator := mocks.NewIValidatorService(t)
		validator.On("Verify", ctx, mock.Anything).Return(
			nil,
		).Maybe()

		service := services.NewSaladInteractor(saladService, []domain.IValidatorService{
			validator,
		})

		sCtx.WithNewParameters("ctx", ctx, "request", expSalad)
		err := service.Update(ctx, expSalad)

		sCtx.Assert().Error(err)
	})
}
