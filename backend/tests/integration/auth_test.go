//go:build integration_test

package integration_tests

import (
	"context"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
	"ppo/domain"
	"ppo/internal/storage/postgres"
	"ppo/tests/utils"
)

type ITAuthSuite struct {
	suite.Suite
	repo domain.IAuthRepository
}

func (s *ITAuthSuite) BeforeAll(t provider.T) {
	t.Title("init test repository")
	s.repo = postgres.NewAuthRepository(testDbInstance)
	t.Tags("auth")
}

func (s *ITAuthSuite) Test_AuthRepo_Register1(t provider.T) {
	t.Title("[Register] successfully signed up")
	t.Tags("integration_test", "postgres", "register")
	t.Parallel()

	t.WithNewStep("Successfully signed up", func(sCtx provider.StepCtx) {
		ctx := context.TODO()

		builder := utils.UserBuilder{}
		userAuth := builder.
			WithName("testingUser").
			WithUsername("testingUser").
			WithPassword("testingUser").
			WithEmail("test@mail.ru").
			ToDto()

		sCtx.WithNewParameters("ctx", ctx, "request", userAuth)
		_, err := s.repo.Register(ctx, userAuth)

		sCtx.Assert().NoError(err)
	})
}

func (s *ITAuthSuite) Test_AuthRepo_Register2(t provider.T) {
	t.Title("[Register] username isn't unique")
	t.Tags("integration_test", "postgres", "register")
	t.Parallel()

	t.WithNewStep("Username isn't unique", func(sCtx provider.StepCtx) {
		ctx := context.TODO()

		builder := utils.UserBuilder{}
		userAuth := builder.
			WithName("testingUser").
			WithUsername("anotherUsername").
			WithPassword("testingUser").
			WithEmail("anotherTest@mail.ru").
			ToDto()

		sCtx.WithNewParameters("ctx", ctx, "request", userAuth)
		_, err := s.repo.Register(ctx, userAuth)

		sCtx.Assert().Error(err)
	})
}

func (s *ITAuthSuite) Test_AuthRepo_Register3(t provider.T) {
	t.Title("[Register] email isn't unique")
	t.Tags("integration_test", "postgres", "register")
	t.Parallel()

	t.WithNewStep("Email isn't unique", func(sCtx provider.StepCtx) {
		ctx := context.TODO()

		builder := utils.UserBuilder{}
		userAuth := builder.
			WithName("testingUser").
			WithUsername("anotherTestingUser").
			WithPassword("testingUser").
			WithEmail("existingMail@mail.ru").
			ToDto()

		sCtx.WithNewParameters("ctx", ctx, "request", userAuth)
		_, err := s.repo.Register(ctx, userAuth)

		sCtx.Assert().Error(err)
	})
}

func (s *ITAuthSuite) Test_AuthRepo_GetByUsername1(t provider.T) {
	t.Title("[Get by username] success")
	t.Tags("integration_test", "postgres", "get_by_username")
	t.Parallel()

	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		ctx := context.TODO()

		username := "anotherUsername"

		sCtx.WithNewParameters("ctx", ctx, "request", username)
		_, err := s.repo.GetByUsername(ctx, username)

		sCtx.Assert().NoError(err)
	})
}

func (s *ITAuthSuite) Test_AuthRepo_GetByUsername2(t provider.T) {
	t.Title("[Get by username] username not found")
	t.Tags("integration_test", "postgres", "get_by_username")
	t.Parallel()

	t.WithNewStep("Username not found", func(sCtx provider.StepCtx) {
		ctx := context.TODO()

		username := "testingUserNotFound"

		sCtx.WithNewParameters("ctx", ctx, "request", username)
		_, err := s.repo.GetByUsername(ctx, username)

		sCtx.Assert().Error(err)
	})
}
