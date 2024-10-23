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

type AuthSuite struct {
	suite.Suite
	JwtKey string
}

func (suite *AuthSuite) TestAuthService_Login1(t provider.T) {
	t.Title("[Login] successfully signed in")
	t.Tags("auth", "login")
	t.Parallel()

	t.WithNewStep("successfully signed in", func(sCtx provider.StepCtx) {
		ctx := context.TODO()

		fabric := utils.AuthInfoObjectMother{}
		userAuth := fabric.Default()

		repo := mocks.NewIAuthRepository(t)
		repo.On("GetByUsername", ctx, userAuth.Username).Return(
			userAuth, nil,
		).Once()

		crypto := mocks.NewIHashCrypto(t)
		crypto.On("CheckPasswordHash", userAuth.Password, userAuth.HashedPass).Return(
			true,
		).Once()

		logger := utils.NewMockLogger()
		service := services.NewAuthService(repo, logger, crypto, suite.JwtKey)

		sCtx.WithNewParameters("ctx", ctx, "request", userAuth)
		token, err := service.Login(ctx, userAuth)

		sCtx.Assert().NoError(err)
		sCtx.Assert().NotEmpty(token)
	})
}

func (suite *AuthSuite) TestAuthService_Login2(t provider.T) {
	t.Title("[Login] username not found")
	t.Tags("auth", "login")
	t.Parallel()

	t.WithNewStep("username not found", func(sCtx provider.StepCtx) {
		ctx := context.TODO()

		fabric := utils.AuthInfoObjectMother{}
		userAuth := fabric.UsernameNotFound()

		repo := mocks.NewIAuthRepository(t)
		repo.On("GetByUsername", ctx, userAuth.Username).
			Return(nil, fmt.Errorf("getting user err"))

		crypto := mocks.NewIHashCrypto(t)
		logger := utils.NewMockLogger()
		service := services.NewAuthService(repo, logger, crypto, suite.JwtKey)

		sCtx.WithNewParameters("ctx", ctx, "request", userAuth)
		token, err := service.Login(ctx, userAuth)

		sCtx.Assert().Error(err)
		sCtx.Assert().Empty(token)
	})
}

func (suite *AuthSuite) TestAuthService_Register1(t provider.T) {
	t.Title("[Register] successfully signed up")
	t.Tags("auth", "register")
	t.Parallel()

	t.WithNewStep("successfully signed up", func(sCtx provider.StepCtx) {
		ctx := context.TODO()

		builder := utils.UserBuilder{}
		userAuth := builder.
			WithName("test").
			WithUsername("test123").
			WithPassword("pass123").
			WithEmail("test@mail.ru").
			ToDto()

		repo := mocks.NewIAuthRepository(t)
		repo.On("Register", ctx, userAuth).Return(
			uuid.New(), nil,
		).Once()

		crypto := mocks.NewIHashCrypto(t)
		crypto.On("GenerateHashPass", userAuth.Password).Return(
			"hashedPass123", nil,
		).Once()

		logger := utils.NewMockLogger()
		service := services.NewAuthService(repo, logger, crypto, suite.JwtKey)

		sCtx.WithNewParameters("ctx", ctx, "request", userAuth)
		id, err := service.Register(ctx, userAuth)

		sCtx.Assert().NoError(err)
		sCtx.Assert().NotNil(id)
	})
}

func (suite *AuthSuite) TestAuthService_Register2(t provider.T) {
	t.Title("[Register] empty name")
	t.Tags("auth", "register")
	t.Parallel()

	t.WithNewStep("empty name", func(sCtx provider.StepCtx) {
		ctx := context.TODO()

		builder := utils.UserBuilder{}
		userAuth := builder.
			WithName("").
			WithUsername("test123").
			WithPassword("pass123").
			WithEmail("test@mail.ru").
			ToDto()

		repo := mocks.NewIAuthRepository(t)
		repo.On("Register", ctx, userAuth).Return(
			uuid.New(), nil,
		).Maybe()

		crypto := mocks.NewIHashCrypto(t)
		crypto.On("GenerateHashPass", userAuth.Password).Return(
			"hashedPass123", nil,
		).Maybe()

		logger := utils.NewMockLogger()
		service := services.NewAuthService(repo, logger, crypto, suite.JwtKey)

		sCtx.WithNewParameters("ctx", ctx, "request", userAuth)
		token, err := service.Register(ctx, userAuth)

		sCtx.Assert().Error(err)
		sCtx.Assert().Empty(token)
	})
}

func (suite *AuthSuite) TestAuthService_Register3(t provider.T) {
	t.Title("[Register] successfully signed up")
	t.Tags("auth", "register")
	t.Parallel()

	t.WithNewStep("successfully signed up", func(sCtx provider.StepCtx) {
		ctx := context.TODO()

		builder := utils.UserBuilder{}
		userAuth := builder.
			WithName("test").
			WithUsername("test123").
			WithPassword("pass123").
			WithEmail("test@mail.ru").
			ToDto()

		repo := mocks.NewIAuthRepository(t)
		repo.On("Register", ctx, userAuth).Return(
			uuid.New(), nil,
		).Once()

		crypto := mocks.NewIHashCrypto(t)
		crypto.On("GenerateHashPass", userAuth.Password).Return(
			"hashedPass123", nil,
		).Once()

		logger := utils.NewMockLogger()
		service := services.NewAuthService(repo, logger, crypto, suite.JwtKey)

		sCtx.WithNewParameters("ctx", ctx, "request", userAuth)
		id, err := service.Register(ctx, userAuth)

		sCtx.Assert().NoError(err)
		sCtx.Assert().NotNil(id)
	})
}
