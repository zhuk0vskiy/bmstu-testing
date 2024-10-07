package repository_test

import (
	"backend/src/internal/model/dto"
	"backend/src/internal/repository/interface/mocks"
	"backend/src/tests/utils"
	"context"
	"fmt"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
)

type UserRepositorySuite struct {
	suite.Suite

	mockRepo mocks.IUserRepository
}

//func (c *UserRepositorySuite) BeforeAll(t provider.T) {
//	t.Title("Init employee mock storage")
//	c.employeeMockStorage = *pool.IUserRepository(t)
//	t.Tags("fixture", "employee")
//}

func (suite *UserRepositorySuite) TestUserAdd01(t provider.T) {
	t.Title("[Add] successfully add")
	t.Tags("repository", "user", "add")
	t.Parallel()

	t.WithNewStep("successfully add", func(sCtx provider.StepCtx) {
		ctx := context.Background()

		request := utils.UserFabric{}.CorrectUserAdd()

		suite.mockRepo.On("Add", ctx, request).Return(
			nil,
		).Once()

		sCtx.WithNewParameters("ctx", ctx, "request", request)
		err := suite.mockRepo.Add(ctx, &dto.AddUserRequest{
			Login:      request.Login,
			Password:   request.Password,
			Role:       request.Role,
			FirstName:  request.FirstName,
			SecondName: request.SecondName,
			ThirdName:  request.ThirdName,
		})

		sCtx.Assert().NoError(err)
	})
}

func (suite *UserRepositorySuite) TestUserAdd02(t provider.T) {
	t.Title("[Add] successfully add")
	t.Tags("repository", "user", "add")
	t.Parallel()

	t.WithNewStep("successfully add", func(sCtx provider.StepCtx) {
		ctx := context.Background()

		request := utils.UserFabric{}.IncorrectUserAdd()

		suite.mockRepo.On("Add", ctx, request).Return(
			fmt.Errorf("empty login"),
		).Once()

		sCtx.WithNewParameters("ctx", ctx, "request", request)
		err := suite.mockRepo.Add(ctx, &dto.AddUserRequest{
			Login:      request.Login,
			Password:   request.Password,
			Role:       request.Role,
			FirstName:  request.FirstName,
			SecondName: request.SecondName,
			ThirdName:  request.ThirdName,
		})

		sCtx.Assert().Error(err)
	})
}

func (suite *UserRepositorySuite) TestUserGet01(t provider.T) {
	t.Title("[Add] success get")
	t.Tags("repository", "user", "get")
	t.Parallel()

	t.WithNewStep("success to get", func(sCtx provider.StepCtx) {
		ctx := context.Background()

		builder := utils.UserAuthBuilder{}
		user := builder.
			WithId(1).
			WithLogin("test").
			WithPassword("hashedPass123").
			WithRole("client").
			WithFirstName("1").
			WithSecondName("2").
			WithThirdName("3").
			ToDto()

		suite.mockRepo.On("Get", ctx, &dto.GetUserRequest{
			Id: user.Id,
		}).Return(
			user, nil,
		).Once()

		sCtx.WithNewParameters("ctx", ctx, "request", user)
		user, err := suite.mockRepo.Get(ctx, &dto.GetUserRequest{
			Id: user.Id,
		})

		sCtx.Assert().NoError(err)
		sCtx.Assert().NotNil(user)
	})
}

func (suite *UserRepositorySuite) TestUserGet02(t provider.T) {
	t.Title("[Add] failed get")
	t.Tags("repository", "user", "get")
	t.Parallel()

	t.WithNewStep("fail to get", func(sCtx provider.StepCtx) {
		ctx := context.Background()

		builder := utils.UserAuthBuilder{}
		user := builder.
			WithId(0).
			WithLogin("test").
			WithPassword("hashedPass123").
			WithRole("client").
			WithFirstName("1").
			WithSecondName("2").
			WithThirdName("3").
			ToDto()

		suite.mockRepo.On("Get", ctx, &dto.GetUserRequest{
			Id: user.Id,
		}).Return(
			nil, fmt.Errorf("incorrect id"),
		).Once()

		sCtx.WithNewParameters("ctx", ctx, "request", user)
		user, err := suite.mockRepo.Get(ctx, &dto.GetUserRequest{
			Id: user.Id,
		})

		sCtx.Assert().Error(err)
		sCtx.Assert().Nil(user)
	})
}
func (suite *UserRepositorySuite) TestUserDelete01(t provider.T) {
	t.Title("[Delete] success")
	t.Tags("repository", "user", "delete")
	t.Parallel()

	t.WithNewStep("successfully delete", func(sCtx provider.StepCtx) {
		ctx := context.Background()

		request := utils.UserFabric{Id: 1}.UserDelete()

		//repo := pool.IUserRepository{t}
		suite.mockRepo.On("Delete", ctx, &dto.DeleteUserRequest{
			Id: request.Id,
		}).Return(
			nil,
		).Once()

		sCtx.WithNewParameters("ctx", ctx, "request", request)
		err := suite.mockRepo.Delete(ctx, &dto.DeleteUserRequest{
			Id: request.Id,
		})

		sCtx.Assert().NoError(err)
	})
}

func (suite *UserRepositorySuite) TestUserDelete02(t provider.T) {
	t.Title("[Delete] success")
	t.Tags("repository", "user", "delete")
	t.Parallel()

	t.WithNewStep("successfully delete", func(sCtx provider.StepCtx) {
		ctx := context.Background()

		request := utils.UserFabric{Id: 0}.UserDelete()

		//repo := pool.IUserRepository{t}
		suite.mockRepo.On("Delete", ctx, &dto.DeleteUserRequest{
			Id: request.Id,
		}).Return(
			fmt.Errorf("incorrect id"),
		).Once()

		sCtx.WithNewParameters("ctx", ctx, "request", request)
		err := suite.mockRepo.Delete(ctx, &dto.DeleteUserRequest{
			Id: request.Id,
		})

		sCtx.Assert().Error(err)
	})
}

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
//		repo := new(pool.IUserRepository)
//		repo.On("GetByLogin", ctx, &dto.GetUserByLoginRequest{Login: userAuth.Login}).
//			Return(nil, fmt.Errorf("getting user err"))
//
//		crypto := pool.NewIHashCrypto(t)
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
//
//func (suite *AuthSuite) TestAuthSignIn01(t provider.T) {
//	t.Title("[Register] successfully signed up")
//	t.Tags("auth", "register")
//	t.Parallel()
//
//	t.WithNewStep("successfully signed up", func(sCtx provider.StepCtx) {
//		ctx := context.Background()
//
//		builder := utils.UserAuthBuilder{}
//		userAuth := builder.
//			WithLogin("test").
//			WithPassword("hashedPass123").
//			WithRole("client").
//			WithFirstName("1").
//			WithSecondName("2").
//			WithThirdName("3").
//			ToDto()
//
//		repo := new(pool.IUserRepository)
//		repo.On("Add", ctx, &dto.AddUserRequest{
//			Login:      userAuth.Login,
//			Password:   userAuth.Password,
//			Role:       userAuth.Role,
//			FirstName:  userAuth.FirstName,
//			SecondName: userAuth.SecondName,
//			ThirdName:  userAuth.ThirdName,
//		}).Return(nil).Once()
//
//		crypto := pool.NewIHashCrypto(t)
//		crypto.On("GenerateHashPass", userAuth.Password).Return(
//			"hashedPass123", nil,
//		).Once()
//
//		logger := utils.NewMockLogger()
//		service := impl.NewAuthService(logger, repo, crypto, suite.JwtKey)
//
//		sCtx.WithNewParameters("ctx", ctx, "request", userAuth)
//		err := service.SignIn(&dto.SignInRequest{
//			Login:      userAuth.Login,
//			Password:   userAuth.Password,
//			FirstName:  userAuth.FirstName,
//			SecondName: userAuth.SecondName,
//			ThirdName:  userAuth.ThirdName,
//		})
//
//		sCtx.Assert().NoError(err)
//	})
//}
//
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
//		repo := new(pool.IUserRepository)
//		repo.On("Add", ctx, requestAdd).Return(nil).Once()
//
//		crypto := pool.NewIHashCrypto(t)
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
//
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
//		repo := new(pool.IUserRepository)
//
//		crypto := pool.NewIHashCrypto(t)
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
//
////
////func (suite *AuthSuite) TestAuthService_Register2(t provider.T) {
////	t.Title("[Register] empty name")
////	t.Tags("auth", "register")
////	t.Parallel()
////
////	t.WithNewStep("empty name", func(sCtx provider.StepCtx) {
////		ctx := context.TODO()
////
////		builder := utils.UserBuilder{}
////		userAuth := builder.
////			WithName("").
////			WithLogin("test123").
////			WithPassword("pass123").
////			WithEmail("test@mail.ru").
////			ToDto()
////
////		repo := pool.NewIAuthRepository(t)
////		repo.On("Register", ctx, userAuth).Return(
////			uuid.New(), nil,
////		).Maybe()
////
////		crypto := pool.NewIHashCrypto(t)
////		crypto.On("GenerateHashPass", userAuth.Password).Return(
////			"hashedPass123", nil,
////		).Maybe()
////
////		logger := utils.NewMockLogger()
////		service_test := services.NewAuthService(repo, logger, crypto, suite.JwtKey)
////
////		sCtx.WithNewParameters("ctx", ctx, "request", userAuth)
////		token, err := service_test.Register(ctx, userAuth)
////
////		sCtx.Assert().Error(err)
////		sCtx.Assert().Empty(token)
////	})
////}
