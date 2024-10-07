package tests

import (
	"backend/src/internal/model/dto"
	"backend/src/internal/repository/impl/postgresql"
	"backend/src/internal/service/impl"
	"backend/src/pkg/base"
	"backend/src/tests/utils"
	"context"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
)

type AuthSuite struct {
	suite.Suite

	JwtKey string
}

//func (suite *AuthSuite) TestAuthLogin01(t provider.T) {
//	t.Title("[Login] successfully signed in")
//	t.Tags("auth", "login")
//	t.Parallel()
//
//	t.WithNewStep("successfully signed in", func(sCtx provider.StepCtx) {
//		ctx := context.Background()
//
//		builder := utils.UserAuthBuilder{}
//		userAuth := builder.
//			WithLogin("username").
//			WithPassword("password").
//			ToDto()
//
//		//repo := pool.IUserRepository{t}
//		//repo := new(mocks.IUserRepository)
//		//repo.On("GetByLogin", ctx, &dto.GetUserByLoginRequest{Login: userAuth.Login}).Return(
//		//	userAuth, nil,
//		//).Once()
//		repo := postgresql.NewUserRepository(testDbInstance)
//
//		crypto := mocks.NewIHashCrypto(t)
//
//		crypto.On("CheckPasswordHash", userAuth.Password, mock.Anything).Return(true)
//
//		logger := utils.NewMockLogger()
//		service := impl.NewAuthService(logger, repo, crypto, suite.JwtKey)
//
//		sCtx.WithNewParameters("ctx", ctx, "request", userAuth)
//		token, err := service.LogIn(&dto.LogInRequest{Login: userAuth.Login, Password: userAuth.Password})
//
//		sCtx.Assert().NoError(err)
//		sCtx.Assert().NotEmpty(token)
//	})
//}
//
//func (suite *AuthSuite) TestAuthLogin02(t provider.T) {
//	t.Title("[Login] username not found")
//	t.Tags("auth", "login")
//	t.Parallel()
//
//	t.WithNewStep("username not found", func(sCtx provider.StepCtx) {
//		ctx := context.Background()
//
//		builder := utils.UserAuthBuilder{}
//		userAuth := builder.
//			WithLogin("notFound").
//			WithPassword("pass").
//			ToDto()
//
//		repo := new(mocks.IUserRepository)
//		repo.On("GetByLogin", ctx, &dto.GetUserByLoginRequest{Login: userAuth.Login}).
//			Return(nil, fmt.Errorf("getting user err"))
//
//		crypto := mocks.NewIHashCrypto(t)
//		logger := utils.NewMockLogger()
//		service := impl.NewAuthService(logger, repo, crypto, suite.JwtKey)
//
//		sCtx.WithNewParameters("ctx", ctx, "request", userAuth)
//		token, err := service.LogIn(&dto.LogInRequest{Login: userAuth.Login, Password: userAuth.Password})
//
//		sCtx.Assert().Error(err)
//		sCtx.Assert().Empty(token)
//	})
//}

func (suite *AuthSuite) TestAuthSignIn01(t provider.T) {
	t.Title("[Register] successfully signed up")
	t.Tags("auth", "register")
	t.Parallel()

	t.WithNewStep("successfully signed up", func(sCtx provider.StepCtx) {
		ctx := context.Background()

		builder := utils.UserAuthBuilder{}
		userAuth := builder.
			WithLogin("test").
			WithPassword("hashedPass123").
			WithRole("client").
			WithFirstName("1").
			WithSecondName("2").
			WithThirdName("3").
			ToDto()

		//repo := new(mocks.IUserRepository)
		repo := postgresql.NewUserRepository(testDbInstance)
		//repo.On("Add", ctx, &dto.AddUserRequest{
		//	Login:      userAuth.Login,
		//	Password:   userAuth.Password,
		//	Role:       userAuth.Role,
		//	FirstName:  userAuth.FirstName,
		//	SecondName: userAuth.SecondName,
		//	ThirdName:  userAuth.ThirdName,
		//}).Return(nil).Once()

		//crypto := mocks.NewIHashCrypto(t)
		crypto := base.NewHashCrypto()

		//crypto.On("GenerateHashPass", userAuth.Password).Return(
		//	"hashedPass123", nil,
		//).Once()

		logger := utils.NewMockLogger()
		service := impl.NewAuthService(logger, repo, crypto, suite.JwtKey)

		sCtx.WithNewParameters("ctx", ctx, "request", userAuth)
		err := service.SignIn(&dto.SignInRequest{
			Login:      userAuth.Login,
			Password:   userAuth.Password,
			FirstName:  userAuth.FirstName,
			SecondName: userAuth.SecondName,
			ThirdName:  userAuth.ThirdName,
		})

		sCtx.Assert().NoError(err)
	})
}

//func (suite *AuthSuite) TestAuthSignIn02(t provider.T) {
//	t.Title("[Register] successfully signed up")
//	t.Tags("auth", "register")
//	t.Parallel()
//
//	t.WithNewStep("successfully signed up", func(sCtx provider.StepCtx) {
//		ctx := context.Background()
//
//		requestAdd, requestSignIn := utils.SignInFabric{Id: 1}.CorrectUserSignIn()
//
//		repo := new(mocks.IUserRepository)
//		repo.On("Add", ctx, requestAdd).Return(nil).Once()
//
//		crypto := mocks.NewIHashCrypto(t)
//		crypto.On("GenerateHashPass", requestSignIn.Password).Return(
//			"hashedPass123", nil,
//		).Once()
//
//		logger := utils.NewMockLogger()
//		service := impl.NewAuthService(logger, repo, crypto, suite.JwtKey)
//
//		sCtx.WithNewParameters("ctx", ctx, "request", requestSignIn)
//		err := service.SignIn(requestSignIn)
//
//		sCtx.Assert().NoError(err)
//	})
//}

//func (suite *AuthSuite) TestAuthSignIn03(t provider.T) {
//	t.Title("[SignIn] failed to signed up")
//	t.Tags("auth", "signIn")
//	t.Parallel()
//
//	t.WithNewStep("failed to signed up", func(sCtx provider.StepCtx) {
//		ctx := context.Background()
//
//		_, requestSignIn := utils.SignInFabric{Id: 1}.IncorrectUserSignIn()
//
//		repo := new(mocks.IUserRepository)
//
//		crypto := mocks.NewIHashCrypto(t)
//
//		logger := utils.NewMockLogger()
//		service := impl.NewAuthService(logger, repo, crypto, suite.JwtKey)
//
//		sCtx.WithNewParameters("ctx", ctx, "request", requestSignIn)
//		err := service.SignIn(requestSignIn)
//
//		sCtx.Assert().Error(err)
//	})
//}
