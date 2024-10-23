//go:build unit_test

package unit_tests

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
	"github.com/pashagolub/pgxmock/v4"
	"net/mail"
	"ppo/domain"
	"ppo/internal/storage/postgres"
	"ppo/mocks"
	"ppo/services"
	"ppo/tests/utils"
)

type UserSuite struct {
	suite.Suite
}

func (suite *UserSuite) TestUserService_Create1(t provider.T) {
	t.Title("[User create] successfully created")
	t.Tags("user", "create")
	t.Parallel()

	t.WithNewStep("successfully created", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		user := utils.NewUserBuilder().
			WithId(uuid.New()).
			WithName("user").
			WithUsername("username").
			WithPassword("password").
			WithEmail("user@mail.com").
			WithRole("user").
			ToDto()

		repo := mocks.NewIUserRepository(t)
		repo.On("Create", ctx, user).Return(
			nil,
		).Once()

		logger := utils.NewMockLogger()
		service := services.NewUserService(repo, logger)

		sCtx.WithNewParameters("ctx", ctx, "request", user)
		err := service.Create(ctx, user)

		sCtx.Assert().NoError(err)
	})
}

func (suite *UserSuite) TestUserService_Create2(t provider.T) {
	t.Title("[User create] empty name")
	t.Tags("user", "create")
	t.Parallel()

	t.WithNewStep("empty name", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		user := utils.NewUserBuilder().
			WithId(uuid.New()).
			WithName("").
			WithUsername("username").
			WithPassword("password").
			WithEmail("user@mail.com").
			WithRole("user").
			ToDto()

		repo := mocks.NewIUserRepository(t)
		repo.On("Create", ctx, user).Return(
			nil,
		).Maybe()

		logger := utils.NewMockLogger()
		service := services.NewUserService(repo, logger)

		sCtx.WithNewParameters("ctx", ctx, "request", user)
		err := service.Create(ctx, user)

		sCtx.Assert().Error(err)
	})
}

func (suite *UserSuite) TestUserService_DeleteById1(t provider.T) {
	t.Title("[User delete by id] successfully deleted")
	t.Tags("user", "delete_by_id")
	t.Parallel()

	t.WithNewStep("successfully deleted", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		id := uuid.New()

		repo := mocks.NewIUserRepository(t)
		repo.On("DeleteById", ctx, id).Return(
			nil,
		).Once()

		logger := utils.NewMockLogger()
		service := services.NewUserService(repo, logger)

		sCtx.WithNewParameters("ctx", ctx, "request", id)
		err := service.DeleteById(ctx, id)

		sCtx.Assert().NoError(err)
	})
}

func (suite *UserSuite) TestUserService_DeleteById2(t provider.T) {
	t.Title("[User delete by id] repo error")
	t.Tags("user", "delete_by_id")
	t.Parallel()

	t.WithNewStep("repo error", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		id := uuid.New()

		repo := mocks.NewIUserRepository(t)
		repo.On("DeleteById", ctx, id).Return(
			fmt.Errorf("repo error"),
		).Once()

		logger := utils.NewMockLogger()
		service := services.NewUserService(repo, logger)

		sCtx.WithNewParameters("ctx", ctx, "request", id)
		err := service.DeleteById(ctx, id)

		sCtx.Assert().Error(err)
	})
}

func (suite *UserSuite) TestUserService_GetAll1(t provider.T) {
	t.Title("[User get all] successful get")
	t.Tags("user", "get_all")
	t.Parallel()

	t.WithNewStep("successful get", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		page := 1
		expUsers := []*domain.User{
			{
				ID:       uuid.UUID{1},
				Name:     "user1",
				Username: "user1",
				Password: "password",
				Email:    mail.Address{Address: "user@mail.com"},
				Role:     "user",
			},
			{
				ID:       uuid.UUID{2},
				Name:     "user2",
				Username: "user2",
				Password: "password",
				Email:    mail.Address{Address: "user2@mail.com"},
				Role:     "user",
			},
		}

		repo := mocks.NewIUserRepository(t)
		repo.On("GetAll", ctx, page).Return(
			expUsers, nil,
		).Once()

		logger := utils.NewMockLogger()
		service := services.NewUserService(repo, logger)

		sCtx.WithNewParameters("ctx", ctx, "request", page)
		users, err := service.GetAll(ctx, page)

		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(users, expUsers)
	})
}

func (suite *UserSuite) TestUserService_GetAll2(t provider.T) {
	t.Title("[User get all] repo error")
	t.Tags("user", "get_all")
	t.Parallel()

	t.WithNewStep("repo error", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		page := 1

		repo := mocks.NewIUserRepository(t)
		repo.On("GetAll", ctx, page).Return(
			nil, fmt.Errorf("repo error"),
		).Once()

		logger := utils.NewMockLogger()
		service := services.NewUserService(repo, logger)

		sCtx.WithNewParameters("ctx", ctx, "request", page)
		users, err := service.GetAll(ctx, page)

		sCtx.Assert().Error(err)
		sCtx.Assert().Nil(users)
	})
}

func (suite *UserSuite) TestUserService_GetById1(t provider.T) {
	t.Title("[User get by id] successful get")
	t.Tags("user", "get_by_id")
	t.Parallel()

	t.WithNewStep("successful get", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		expUser := utils.NewUserBuilder().
			WithId(uuid.New()).
			WithName("user").
			WithUsername("username").
			WithPassword("password").
			WithEmail("user@mail.com").
			WithRole("user").
			ToDto()

		repo := mocks.NewIUserRepository(t)
		repo.On("GetById", ctx, expUser.ID).Return(
			expUser, nil,
		).Once()

		logger := utils.NewMockLogger()
		service := services.NewUserService(repo, logger)

		sCtx.WithNewParameters("ctx", ctx, "request", expUser.ID)
		user, err := service.GetById(ctx, expUser.ID)

		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(user, expUser)
	})
}

func (suite *UserSuite) TestUserService_GetById2(t provider.T) {
	t.Title("[User get by id] repo error")
	t.Tags("user", "get_by_id")
	t.Parallel()

	t.WithNewStep("repo error", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		id := uuid.New()

		repo := mocks.NewIUserRepository(t)
		repo.On("GetById", ctx, id).Return(
			nil, fmt.Errorf("repo error"),
		).Once()

		logger := utils.NewMockLogger()
		service := services.NewUserService(repo, logger)

		sCtx.WithNewParameters("ctx", ctx, "request", id)
		comment, err := service.GetById(ctx, id)

		sCtx.Assert().Error(err)
		sCtx.Assert().Nil(comment)
	})
}

func (suite *UserSuite) TestUserService_Update1(t provider.T) {
	t.Title("[User update] successfully updated")
	t.Tags("user", "update")
	t.Parallel()

	t.WithNewStep("successfully updated", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		expUser := utils.NewUserBuilder().
			WithId(uuid.New()).
			WithName("user").
			WithUsername("username").
			WithPassword("password").
			WithEmail("user@mail.com").
			WithRole("user").
			ToDto()

		repo := mocks.NewIUserRepository(t)
		repo.On("Update", ctx, expUser).Return(
			nil,
		).Once()

		logger := utils.NewMockLogger()
		service := services.NewUserService(repo, logger)

		sCtx.WithNewParameters("ctx", ctx, "request", expUser)
		err := service.Update(ctx, expUser)

		sCtx.Assert().NoError(err)
	})
}

func (suite *UserSuite) TestUserService_Update2(t provider.T) {
	t.Title("[User update] repo error")
	t.Tags("user", "update")
	t.Parallel()

	t.WithNewStep("repo error", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		expUser := utils.NewUserBuilder().
			WithId(uuid.New()).
			WithName("user").
			WithUsername("username").
			WithPassword("password").
			WithEmail("user@mail.com").
			WithRole("user").
			ToDto()

		repo := mocks.NewIUserRepository(t)
		repo.On("Update", ctx, expUser).Return(
			fmt.Errorf("repo error"),
		).Once()

		logger := utils.NewMockLogger()
		service := services.NewUserService(repo, logger)

		sCtx.WithNewParameters("ctx", ctx, "request", expUser)
		err := service.Update(ctx, expUser)

		sCtx.Assert().Error(err)
	})
}

func (suite *UserSuite) TestUserService_GetByUserName1(t provider.T) {
	t.Title("[User get by id] successful get")
	t.Tags("user", "get_by_id")
	t.Parallel()

	t.WithNewStep("successful get", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		expUser := utils.NewUserBuilder().
			WithId(uuid.New()).
			WithName("user").
			WithUsername("username").
			WithPassword("password").
			WithEmail("user@mail.com").
			WithRole("user").
			ToDto()

		repo := mocks.NewIUserRepository(t)
		repo.On("GetByUsername", ctx, expUser.Username).Return(
			expUser, nil,
		).Once()

		logger := utils.NewMockLogger()
		service := services.NewUserService(repo, logger)

		sCtx.WithNewParameters("ctx", ctx, "request", expUser.Username)
		user, err := service.GetByUsername(ctx, expUser.Username)

		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(user, expUser)
	})
}

func (suite *UserSuite) TestUserService_GetByUsername2(t provider.T) {
	t.Title("[User get by id] repo error")
	t.Tags("user", "get_by_id")
	t.Parallel()

	t.WithNewStep("repo error", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		username := "username"

		repo := mocks.NewIUserRepository(t)
		repo.On("GetByUsername", ctx, username).Return(
			nil, fmt.Errorf("repo error"),
		).Once()

		logger := utils.NewMockLogger()
		service := services.NewUserService(repo, logger)

		sCtx.WithNewParameters("ctx", ctx, "request", username)
		user, err := service.GetByUsername(ctx, username)

		sCtx.Assert().Error(err)
		sCtx.Assert().Nil(user)
	})
}

func (suite *UserSuite) TestUserService_Update3(t provider.T) {
	t.Title("[User update] empty name")
	t.Tags("user", "update")
	t.Parallel()

	t.WithNewStep("empty name", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		expUser := utils.NewUserBuilder().
			WithId(uuid.New()).
			WithName("").
			WithUsername("username").
			WithPassword("password").
			WithEmail("user@mail.com").
			WithRole("user").
			ToDto()

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatal(err)
		}
		defer mock.Close()
		mock.ExpectExec("update").
			WithArgs(
				expUser.Name,
				expUser.Email.Address,
				expUser.Password,
				expUser.Role,
			).
			WillReturnResult(pgxmock.NewResult("update", 1)).
			Maybe()

		repo := postgres.NewUserRepository(mock)

		logger := utils.NewMockLogger()
		service := services.NewUserService(repo, logger)

		sCtx.WithNewParameters("ctx", ctx, "request", expUser)
		err = service.Update(ctx, expUser)

		sCtx.Assert().Error(err)
	})
}

func (suite *UserSuite) TestUserService_Update4(t provider.T) {
	t.Title("[User update] successfully updated")
	t.Tags("user", "update")
	t.Parallel()

	t.WithNewStep("successfully updated", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		expUser := utils.NewUserBuilder().
			WithId(uuid.New()).
			WithName("name").
			WithUsername("username").
			WithPassword("password").
			WithEmail("user@mail.com").
			WithRole("user").
			ToDto()

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatal(err)
		}
		defer mock.Close()
		mock.ExpectExec("update").
			WithArgs(
				expUser.Name,
				expUser.Email.Address,
				expUser.Password,
				expUser.Role,
			).
			WillReturnResult(pgxmock.NewResult("update", 1))

		repo := postgres.NewUserRepository(mock)

		logger := utils.NewMockLogger()
		service := services.NewUserService(repo, logger)

		sCtx.WithNewParameters("ctx", ctx, "request", expUser)
		err = service.Update(ctx, expUser)
		sCtx.Assert().NoError(err)
	})
}
