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

type SaladTypeSuite struct {
	suite.Suite
}

func (suite *SaladTypeSuite) TestSaladTypeService_Create1(t provider.T) {
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

		repo := mocks.NewISaladTypeRepository(t)
		repo.On("Create", ctx, saladType).Return(
			nil,
		).Once()

		logger := utils.NewMockLogger()
		service := services.NewSaladTypeService(repo, logger)

		sCtx.WithNewParameters("ctx", ctx, "request", saladType)
		err := service.Create(ctx, saladType)

		sCtx.Assert().NoError(err)
	})
}

func (suite *SaladTypeSuite) TestSaladTypeService_Create2(t provider.T) {
	t.Title("[Salad type create] empty name")
	t.Tags("salad_type", "create")
	t.Parallel()

	t.WithNewStep("empty name", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		saladType := utils.NewSaladTypeBuilder().
			WithId(uuid.New()).
			WithName("").
			WithDescription("description").
			ToDto()

		repo := mocks.NewISaladTypeRepository(t)
		repo.On("Create", ctx, saladType).Return(
			nil,
		).Maybe()

		logger := utils.NewMockLogger()
		service := services.NewSaladTypeService(repo, logger)

		sCtx.WithNewParameters("ctx", ctx, "request", saladType)
		err := service.Create(ctx, saladType)

		sCtx.Assert().Error(err)
	})
}

func (suite *SaladTypeSuite) TestSaladTypeService_DeleteById1(t provider.T) {
	t.Title("[Salad type delete by id] successfully deleted")
	t.Tags("salad_type", "delete_by_id")
	t.Parallel()

	t.WithNewStep("successfully deleted", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		id := uuid.New()

		repo := mocks.NewISaladTypeRepository(t)
		repo.On("DeleteById", ctx, id).Return(
			nil,
		).Once()

		logger := utils.NewMockLogger()
		service := services.NewSaladTypeService(repo, logger)

		sCtx.WithNewParameters("ctx", ctx, "request", id)
		err := service.DeleteById(ctx, id)

		sCtx.Assert().NoError(err)
	})
}

func (suite *SaladTypeSuite) TestSaladTypeService_DeleteById2(t provider.T) {
	t.Title("[Salad type delete by id] repo error")
	t.Tags("salad_type", "delete_by_id")
	t.Parallel()

	t.WithNewStep("repo error", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		id := uuid.New()

		repo := mocks.NewISaladTypeRepository(t)
		repo.On("DeleteById", ctx, id).Return(
			fmt.Errorf("repo error"),
		).Once()

		logger := utils.NewMockLogger()
		service := services.NewSaladTypeService(repo, logger)

		sCtx.WithNewParameters("ctx", ctx, "request", id)
		err := service.DeleteById(ctx, id)

		sCtx.Assert().Error(err)
	})
}

func (suite *SaladTypeSuite) TestSaladTypeService_GetAll1(t provider.T) {
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

		repo := mocks.NewISaladTypeRepository(t)
		repo.On("GetAll", ctx, page).Return(
			expTypes, 1, nil,
		).Once()

		logger := utils.NewMockLogger()
		service := services.NewSaladTypeService(repo, logger)

		sCtx.WithNewParameters("ctx", ctx, "request", page)
		types, pages, err := service.GetAll(ctx, page)

		sCtx.Assert().NoError(err)
		sCtx.Assert().NotZero(pages)
		sCtx.Assert().Equal(types, expTypes)
	})
}

func (suite *SaladTypeSuite) TestSaladTypeService_GetAll2(t provider.T) {
	t.Title("[Salad type get all] repo error")
	t.Tags("salad_type", "get_all")
	t.Parallel()

	t.WithNewStep("repo error", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		page := 1

		repo := mocks.NewISaladTypeRepository(t)
		repo.On("GetAll", ctx, page).Return(
			nil, 1, fmt.Errorf("repo error"),
		).Once()

		logger := utils.NewMockLogger()
		service := services.NewSaladTypeService(repo, logger)

		sCtx.WithNewParameters("ctx", ctx, "request", page)
		types, pages, err := service.GetAll(ctx, page)

		sCtx.Assert().Error(err)
		sCtx.Assert().Zero(pages)
		sCtx.Assert().Nil(types)
	})
}

func (suite *SaladTypeSuite) TestSaladTypeService_GetById1(t provider.T) {
	t.Title("[Salad type get by id] successful get")
	t.Tags("salad_type", "get_by_id")
	t.Parallel()

	t.WithNewStep("successful get", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		expType := utils.NewSaladTypeBuilder().
			WithId(uuid.New()).
			WithName("type").
			WithDescription("description").
			ToDto()

		repo := mocks.NewISaladTypeRepository(t)
		repo.On("GetById", ctx, expType.ID).Return(
			expType, nil,
		).Once()

		logger := utils.NewMockLogger()
		service := services.NewSaladTypeService(repo, logger)

		sCtx.WithNewParameters("ctx", ctx, "request", expType.ID)
		saladType, err := service.GetById(ctx, expType.ID)

		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(saladType, expType)
	})
}

func (suite *SaladTypeSuite) TestSaladTypeService_GetById2(t provider.T) {
	t.Title("[Salad type get by id] repo error")
	t.Tags("salad_type", "get_by_id")
	t.Parallel()

	t.WithNewStep("repo error", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		id := uuid.New()

		repo := mocks.NewISaladTypeRepository(t)
		repo.On("GetById", ctx, id).Return(
			nil, fmt.Errorf("repo error"),
		).Once()

		logger := utils.NewMockLogger()
		service := services.NewSaladTypeService(repo, logger)

		sCtx.WithNewParameters("ctx", ctx, "request", id)
		saladType, err := service.GetById(ctx, id)

		sCtx.Assert().Error(err)
		sCtx.Assert().Nil(saladType)
	})
}

func (suite *SaladTypeSuite) TestSaladTypeService_Update1(t provider.T) {
	t.Title("[Salad type update] successfully updated")
	t.Tags("salad_type", "update")
	t.Parallel()

	t.WithNewStep("successfully updated", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		expType := utils.NewSaladTypeBuilder().
			WithId(uuid.New()).
			WithName("type").
			WithDescription("description").
			ToDto()

		repo := mocks.NewISaladTypeRepository(t)
		repo.On("Update", ctx, expType).Return(
			nil,
		).Once()

		logger := utils.NewMockLogger()
		service := services.NewSaladTypeService(repo, logger)

		sCtx.WithNewParameters("ctx", ctx, "request", expType)
		err := service.Update(ctx, expType)

		sCtx.Assert().NoError(err)
	})
}

func (suite *SaladTypeSuite) TestSaladTypeService_Update2(t provider.T) {
	t.Title("[Salad type update] repo error")
	t.Tags("salad_type", "update")
	t.Parallel()

	t.WithNewStep("repo error", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		expType := utils.NewSaladTypeBuilder().
			WithId(uuid.New()).
			WithName("type").
			WithDescription("description").
			ToDto()

		repo := mocks.NewISaladTypeRepository(t)
		repo.On("Update", ctx, expType).Return(
			fmt.Errorf("repo error"),
		).Once()

		logger := utils.NewMockLogger()
		service := services.NewSaladTypeService(repo, logger)

		sCtx.WithNewParameters("ctx", ctx, "request", expType)
		err := service.Update(ctx, expType)

		sCtx.Assert().Error(err)
	})
}

func (suite *SaladTypeSuite) TestSaladTypeService_GetAllBySaladId1(t provider.T) {
	t.Title("[Salad type get all] successful get")
	t.Tags("salad_type", "get_all_by_salad_id")
	t.Parallel()

	t.WithNewStep("successful get", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		id := uuid.New()
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

		repo := mocks.NewISaladTypeRepository(t)
		repo.On("GetAllBySaladId", ctx, id).Return(
			expTypes, nil,
		).Once()

		logger := utils.NewMockLogger()
		service := services.NewSaladTypeService(repo, logger)

		sCtx.WithNewParameters("ctx", ctx, "request", id)
		types, err := service.GetAllBySaladId(ctx, id)

		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(types, expTypes)
	})
}

func (suite *SaladTypeSuite) TestSaladTypeService_GetAllBySaladId2(t provider.T) {
	t.Title("[Salad type get all] repo error")
	t.Tags("salad_type", "get_all_by_salad_id")
	t.Parallel()

	t.WithNewStep("repo error", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		id := uuid.New()

		repo := mocks.NewISaladTypeRepository(t)
		repo.On("GetAllBySaladId", ctx, id).Return(
			nil, fmt.Errorf("repo error"),
		).Once()

		logger := utils.NewMockLogger()
		service := services.NewSaladTypeService(repo, logger)

		sCtx.WithNewParameters("ctx", ctx, "request", id)
		types, err := service.GetAllBySaladId(ctx, id)

		sCtx.Assert().Error(err)
		sCtx.Assert().Nil(types)
	})
}

func (suite *SaladTypeSuite) TestSaladTypeService_Link1(t provider.T) {
	t.Title("[Salad type link] successfully linked")
	t.Tags("salad_type", "link")
	t.Parallel()

	t.WithNewStep("successfully linked", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		saladId := uuid.New()
		saladTypeId := uuid.New()

		repo := mocks.NewISaladTypeRepository(t)
		repo.On("Link", ctx, saladId, saladTypeId).Return(
			nil,
		).Once()

		logger := utils.NewMockLogger()
		service := services.NewSaladTypeService(repo, logger)

		sCtx.WithNewParameters("ctx", ctx, "request", saladId, saladTypeId)
		err := service.Link(ctx, saladId, saladTypeId)

		sCtx.Assert().NoError(err)
	})
}

func (suite *SaladTypeSuite) TestSaladTypeService_Link2(t provider.T) {
	t.Title("[Salad type link] repo error")
	t.Tags("salad_type", "link")
	t.Parallel()

	t.WithNewStep("successfully linked", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		saladId := uuid.New()
		saladTypeId := uuid.New()

		repo := mocks.NewISaladTypeRepository(t)
		repo.On("Link", ctx, saladId, saladTypeId).Return(
			fmt.Errorf("repo error"),
		).Once()

		logger := utils.NewMockLogger()
		service := services.NewSaladTypeService(repo, logger)

		sCtx.WithNewParameters("ctx", ctx, "request", saladId, saladTypeId)
		err := service.Link(ctx, saladId, saladTypeId)

		sCtx.Assert().Error(err)
	})
}

func (suite *IngredientSuite) TestSaladTypeService_Unlink1(t provider.T) {
	t.Title("[Salad type unlink] successfully unlinked")
	t.Tags("salad_type", "unlink")
	t.Parallel()

	t.WithNewStep("successfully unlinked", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		saladId := uuid.New()
		saladTypeId := uuid.New()

		repo := mocks.NewISaladTypeRepository(t)
		repo.On("Unlink", ctx, saladId, saladTypeId).Return(
			nil,
		).Once()

		logger := utils.NewMockLogger()
		service := services.NewSaladTypeService(repo, logger)

		sCtx.WithNewParameters("ctx", ctx, "request", saladId, saladTypeId)
		err := service.Unlink(ctx, saladId, saladTypeId)

		sCtx.Assert().NoError(err)
	})
}

func (suite *IngredientSuite) TestSaladTypeService_Unlink2(t provider.T) {
	t.Title("[Salad type unlink] successfully unlinked")
	t.Tags("salad_type", "unlink")
	t.Parallel()

	t.WithNewStep("successfully unlinked", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		saladId := uuid.New()
		saladTypeId := uuid.New()

		repo := mocks.NewISaladTypeRepository(t)
		repo.On("Unlink", ctx, saladId, saladTypeId).Return(
			fmt.Errorf("repo error"),
		).Once()

		logger := utils.NewMockLogger()
		service := services.NewSaladTypeService(repo, logger)

		sCtx.WithNewParameters("ctx", ctx, "request", saladId, saladTypeId)
		err := service.Unlink(ctx, saladId, saladTypeId)

		sCtx.Assert().Error(err)
	})
}
