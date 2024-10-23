//go:build unit_test

package unit_tests

import (
	"context"
	"github.com/google/uuid"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
	"github.com/pashagolub/pgxmock/v4"
	"ppo/internal/storage/postgres"
	"ppo/tests/utils"
)

type AuthRepoSuite struct {
	suite.Suite
}

func (suite *AuthRepoSuite) TestAuthRepo_Register1(t provider.T) {
	t.Title("[Register] successfully inserted")
	t.Tags("auth", "register")
	t.Parallel()

	t.WithNewStep("successfully inserted", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		authInfo := utils.NewUserBuilder().
			WithId(uuid.New()).
			WithUsername("username").
			WithPassword("password").
			WithName("name").
			WithEmail("mail@mail.com").
			WithRole("user").
			ToDto()

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatal(err)
		}
		defer mock.Close()
		mock.ExpectQuery("insert into saladRecipes.user").
			WithArgs(
				authInfo.Name,
				authInfo.Email.Address,
				authInfo.Username,
				authInfo.Password,
			).
			WillReturnRows(pgxmock.NewRows([]string{"id"}).AddRow(uuid.New()))

		repo := postgres.NewAuthRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "request", authInfo)
		id, err := repo.Register(ctx, authInfo)

		sCtx.Assert().NoError(err)
		sCtx.Assert().NotEqual(id, uuid.Nil)
	})
}

func (suite *AuthRepoSuite) TestAuthRepo_Register2(t provider.T) {
	t.Title("[Register] not inserted")
	t.Tags("auth", "Register")
	t.Parallel()

	t.WithNewStep("not inserted", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		authInfo := utils.NewUserBuilder().
			WithId(uuid.New()).
			WithUsername("username").
			WithPassword("password").
			WithName("name").
			WithEmail("mail@mail.com").
			WithRole("user").
			ToDto()

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatal(err)
		}
		defer mock.Close()
		mock.ExpectQuery("insert into saladRecipes.user").
			WithArgs(
				authInfo.Name,
				authInfo.Email.Address,
				authInfo.Username,
				authInfo.Password,
			).
			WillReturnRows(pgxmock.NewRows([]string{"id"}))

		repo := postgres.NewAuthRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "request", authInfo)
		id, err := repo.Register(ctx, authInfo)

		sCtx.Assert().Error(err)
		sCtx.Assert().Equal(id, uuid.Nil)
	})
}

func (suite *AuthRepoSuite) TestAuthRepo_GetByUsername1(t provider.T) {
	t.Title("[Get] successful get")
	t.Tags("auth", "get_by_username")
	t.Parallel()

	t.WithNewStep("successfull get", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		authInfo := utils.NewUserBuilder().
			WithId(uuid.New()).
			WithUsername("username").
			WithPassword("password").
			WithName("name").
			WithEmail("mail@mail.com").
			WithRole("user").
			ToDto()

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatal(err)
		}
		defer mock.Close()
		mock.ExpectQuery("select").
			WithArgs(
				authInfo.Username,
			).
			WillReturnRows(pgxmock.NewRows([]string{"id", "password", "role"}).AddRow(uuid.New(), "password", "user"))

		repo := postgres.NewAuthRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "request", authInfo.Username)
		_, err = repo.GetByUsername(ctx, authInfo.Username)

		sCtx.Assert().NoError(err)
	})
}

func (suite *AuthRepoSuite) TestAuthRepo_GetByUsername2(t provider.T) {
	t.Title("[Get] not found")
	t.Tags("auth", "get_by_username")
	t.Parallel()

	t.WithNewStep("not found", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		authInfo := utils.NewUserBuilder().
			WithId(uuid.New()).
			WithUsername("username").
			WithPassword("password").
			WithName("name").
			WithEmail("mail@mail.com").
			WithRole("user").
			ToDto()

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatal(err)
		}
		defer mock.Close()
		mock.ExpectQuery("select").
			WithArgs(
				authInfo.Username,
			).
			WillReturnRows(pgxmock.NewRows([]string{"id", "password", "role"}))

		repo := postgres.NewAuthRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "request", authInfo.Username)
		_, err = repo.GetByUsername(ctx, authInfo.Username)

		sCtx.Assert().Error(err)
	})
}
