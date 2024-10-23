//go:build integration_test

package integration_tests

import (
	"context"
	"github.com/google/uuid"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
	"ppo/domain"
	"ppo/internal/storage/postgres"
	"ppo/services/dto"
)

type ITSaladSuite struct {
	suite.Suite
	repo domain.ISaladRepository
}

func (s *ITSaladSuite) BeforeAll(t provider.T) {
	t.Title("init test repository")
	s.repo = postgres.NewSaladRepository(testDbInstance)
	t.Tags("salad")
}

func (s *ITSaladSuite) Test_RecipeStepRepo_GetAll1(t provider.T) {
	t.Title("[Get all] success, filter by ingredients and types")
	t.Tags("integration_test", "postgres", "get_all")
	t.Parallel()

	t.WithNewStep("Success, filter by ingredients and types", func(sCtx provider.StepCtx) {
		ctx := context.TODO()

		page := 1
		filter := new(dto.RecipeFilter)
		filter.AvailableIngredients = make([]uuid.UUID, 2)
		filter.AvailableIngredients[0], _ = uuid.Parse("f1fc4bfc-799c-4471-a971-1bb00f7dd30a")
		filter.AvailableIngredients[1], _ = uuid.Parse("01000000-0000-0000-0000-000000000000")
		filter.Status = dto.PublishedSaladStatus
		filter.SaladTypes = make([]uuid.UUID, 1)
		filter.SaladTypes[0], _ = uuid.Parse("7e17866b-2b97-4d2b-b399-42ceeebd5480")

		saladId, _ := uuid.Parse("fbabc2aa-cd4a-42b0-b68d-d3cf67fba06f")
		saladId2, _ := uuid.Parse("01000000-0000-0000-0000-000000000000")
		expected := make([]*domain.Salad, 2)
		expected[0] = &domain.Salad{
			ID:          saladId,
			AuthorID:    uuid.Nil,
			Name:        "цезарь",
			Description: "",
		}
		expected[1] = &domain.Salad{
			ID:          saladId2,
			AuthorID:    uuid.Nil,
			Name:        "овощной",
			Description: "",
		}

		sCtx.WithNewParameters("ctx", ctx, "request", page, filter)
		salads, _, err := s.repo.GetAll(ctx, filter, page)

		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(salads, expected)
	})
}

func (s *ITSaladSuite) Test_RecipeStepRepo_GetAll2(t provider.T) {
	t.Title("[Get all] success, filter by ingredients")
	t.Tags("integration_test", "postgres", "get_all")
	t.Parallel()

	t.WithNewStep("Success, filter by ingredients", func(sCtx provider.StepCtx) {
		ctx := context.TODO()

		page := 1
		filter := new(dto.RecipeFilter)
		filter.AvailableIngredients = make([]uuid.UUID, 2)
		filter.AvailableIngredients[0], _ = uuid.Parse("f1fc4bfc-799c-4471-a971-1bb00f7dd30a")
		filter.AvailableIngredients[1], _ = uuid.Parse("02000000-0000-0000-0000-000000000000")
		filter.Status = dto.PublishedSaladStatus

		saladId, _ := uuid.Parse("fbabc2aa-cd4a-42b0-b68d-d3cf67fba06f")
		saladId2, _ := uuid.Parse("03000000-0000-0000-0000-000000000000")
		expected := make([]*domain.Salad, 2)
		expected[0] = &domain.Salad{
			ID:          saladId,
			AuthorID:    uuid.Nil,
			Name:        "цезарь",
			Description: "",
		}
		expected[1] = &domain.Salad{
			ID:          saladId2,
			AuthorID:    uuid.Nil,
			Name:        "сельдь под шубой",
			Description: "",
		}

		sCtx.WithNewParameters("ctx", ctx, "request", page, filter)
		salads, _, err := s.repo.GetAll(ctx, filter, page)

		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(salads, expected)
	})
}

func (s *ITSaladSuite) Test_RecipeStepRepo_GetAll3(t provider.T) {
	t.Title("[Get all] success, filter by types")
	t.Tags("integration_test", "postgres", "get_all")
	t.Parallel()

	t.WithNewStep("Success, filter by types", func(sCtx provider.StepCtx) {
		ctx := context.TODO()

		page := 1
		filter := new(dto.RecipeFilter)
		filter.Status = dto.PublishedSaladStatus
		filter.SaladTypes = make([]uuid.UUID, 1)
		filter.SaladTypes[0], _ = uuid.Parse("01000000-0000-0000-0000-000000000000")

		saladId, _ := uuid.Parse("02000000-0000-0000-0000-000000000000")
		saladId2, _ := uuid.Parse("01000000-0000-0000-0000-000000000000")
		expected := make([]*domain.Salad, 2)
		expected[1] = &domain.Salad{
			ID:          saladId,
			AuthorID:    uuid.Nil,
			Name:        "сезонный",
			Description: "",
		}
		expected[0] = &domain.Salad{
			ID:          saladId2,
			AuthorID:    uuid.Nil,
			Name:        "овощной",
			Description: "",
		}

		sCtx.WithNewParameters("ctx", ctx, "request", page, filter)
		salads, _, err := s.repo.GetAll(ctx, filter, page)

		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(salads, expected)
	})
}

func (s *ITSaladSuite) Test_RecipeStepRepo_GetAll4(t provider.T) {
	t.Title("[Get all] success, empty filter")
	t.Tags("integration_test", "postgres", "get_all")
	t.Parallel()

	t.WithNewStep("Success, empty filter", func(sCtx provider.StepCtx) {
		ctx := context.TODO()

		page := 1
		filter := new(dto.RecipeFilter)
		filter.Status = dto.PublishedSaladStatus

		saladId, _ := uuid.Parse("fbabc2aa-cd4a-42b0-b68d-d3cf67fba06f")
		saladId2, _ := uuid.Parse("01000000-0000-0000-0000-000000000000")
		saladId3, _ := uuid.Parse("02000000-0000-0000-0000-000000000000")
		saladId4, _ := uuid.Parse("03000000-0000-0000-0000-000000000000")
		saladId5, _ := uuid.Parse("04000000-0000-0000-0000-000000000000")
		expected := make([]*domain.Salad, 5)
		expected[0] = &domain.Salad{
			ID:          saladId,
			AuthorID:    uuid.Nil,
			Name:        "цезарь",
			Description: "",
		}
		expected[1] = &domain.Salad{
			ID:          saladId2,
			AuthorID:    uuid.Nil,
			Name:        "овощной",
			Description: "",
		}
		expected[2] = &domain.Salad{
			ID:          saladId4,
			AuthorID:    uuid.Nil,
			Name:        "сельдь под шубой",
			Description: "",
		}
		expected[3] = &domain.Salad{
			ID:          saladId3,
			AuthorID:    uuid.Nil,
			Name:        "сезонный",
			Description: "",
		}
		expected[4] = &domain.Salad{
			ID:          saladId5,
			AuthorID:    uuid.Nil,
			Name:        "греческий",
			Description: "",
		}

		sCtx.WithNewParameters("ctx", ctx, "request", page, filter)
		salads, _, err := s.repo.GetAll(ctx, filter, page)

		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(salads, expected)
	})
}
