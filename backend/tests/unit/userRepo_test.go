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
	"ppo/internal/config"
	"ppo/internal/storage/postgres"
	"ppo/tests/utils"
)

type UserRepoSuite struct {
	suite.Suite
}

func (suite *UserRepoSuite) TestUserRepo_Create1(t provider.T) {
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

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatal(err)
		}
		defer mock.Close()
		mock.ExpectExec("insert").
			WithArgs(
				user.Name,
				user.Email.Address,
				user.Username,
				user.Password,
			).
			WillReturnResult(pgxmock.NewResult("insert", 1))

		repo := postgres.NewUserRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "request", user)
		err = repo.Create(ctx, user)

		sCtx.Assert().NoError(err)
	})
}

func (suite *UserRepoSuite) TestUserRepo_Create2(t provider.T) {
	t.Title("[User create] not inserted")
	t.Tags("user", "create")
	t.Parallel()

	t.WithNewStep("not inserted", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		user := utils.NewUserBuilder().
			WithId(uuid.New()).
			WithName("user").
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
		mock.ExpectQuery("insert").
			WithArgs(
				user.Name,
				user.Email.Address,
				user.Username,
				user.Password,
			).
			WillReturnError(fmt.Errorf("insert error"))

		repo := postgres.NewUserRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "request", user)
		err = repo.Create(ctx, user)

		sCtx.Assert().Error(err)
	})
}

func (suite *UserRepoSuite) TestUserRepo_GetById1(t provider.T) {
	t.Title("[User get by id] successful get")
	t.Tags("user", "get_by_id")
	t.Parallel()

	t.WithNewStep("successful get", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		user := utils.NewUserBuilder().
			WithId(uuid.New()).
			WithName("user").
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
		mock.ExpectQuery("select").
			WithArgs(
				user.ID,
			).
			WillReturnRows(
				pgxmock.NewRows([]string{"ID", "Name", "Email", "Username", "Password", "Role"}).
					AddRow(user.ID, user.Name, user.Email.Address, user.Username, user.Password, user.Role),
			)

		repo := postgres.NewUserRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "request", user.ID)
		rUser, err := repo.GetById(ctx, user.ID)

		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(user, rUser)
	})
}

func (suite *UserRepoSuite) TestUserRepo_GetById2(t provider.T) {
	t.Title("[User get by id] failed to get")
	t.Tags("user", "get_by_id")
	t.Parallel()

	t.WithNewStep("failed to get", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		id := uuid.New()

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatal(err)
		}
		defer mock.Close()
		mock.ExpectQuery("select").
			WithArgs(
				id,
			).
			WillReturnRows(
				pgxmock.NewRows([]string{"ID", "Name", "Email", "Username", "Password", "Role"}),
			)

		repo := postgres.NewUserRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "request", id)
		rUser, err := repo.GetById(ctx, id)

		sCtx.Assert().Error(err)
		sCtx.Assert().Nil(rUser)
	})
}

func (suite *UserRepoSuite) TestUserRepo_GetAll1(t provider.T) {
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

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatal(err)
		}
		defer mock.Close()
		mock.ExpectQuery("select").
			WithArgs(
				config.PageSize*(page-1),
				config.PageSize,
			).
			WillReturnRows(
				pgxmock.NewRows([]string{"ID", "Name", "Email", "Username", "Password", "Role"}).
					AddRow(expUsers[0].ID, expUsers[0].Name, expUsers[0].Email.Address, expUsers[0].Username, expUsers[0].Password, expUsers[0].Role).
					AddRow(expUsers[1].ID, expUsers[1].Name, expUsers[1].Email.Address, expUsers[1].Username, expUsers[1].Password, expUsers[1].Role),
			)

		repo := postgres.NewUserRepository(mock)

		sCtx.WithNewParameters("ctx", ctx)
		rUsers, err := repo.GetAll(ctx, page)

		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(rUsers, expUsers)
	})
}

func (suite *UserRepoSuite) TestUserRepo_GetAll2(t provider.T) {
	t.Title("[User get all] sql error")
	t.Tags("user", "get_all")
	t.Parallel()

	t.WithNewStep("sql error", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		page := 1

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatal(err)
		}
		defer mock.Close()
		mock.ExpectQuery("select").
			WithArgs(
				config.PageSize*(page-1),
				config.PageSize,
			).
			WillReturnError(fmt.Errorf("sql error"))

		repo := postgres.NewUserRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "request", page)
		rMeasurements, err := repo.GetAll(ctx, page)

		sCtx.Assert().Error(err)
		sCtx.Assert().Nil(rMeasurements)
	})
}

func (suite *UserRepoSuite) TestUserRepo_Update1(t provider.T) {
	t.Title("[User update] successfully updated")
	t.Tags("user", "update")
	t.Parallel()

	t.WithNewStep("successfully updated", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		user := utils.NewUserBuilder().
			WithId(uuid.New()).
			WithName("user").
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
				user.Name,
				user.Email.Address,
				user.Password,
				user.Role,
			).
			WillReturnResult(pgxmock.NewResult("update", 1))

		repo := postgres.NewUserRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "request", user)
		err = repo.Update(ctx, user)

		sCtx.Assert().NoError(err)
	})
}

func (suite *UserRepoSuite) TestUserRepo_Update2(t provider.T) {
	t.Title("[User update] failed to update")
	t.Tags("user", "update")
	t.Parallel()

	t.WithNewStep("failed to updated", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		user := utils.NewUserBuilder().
			WithId(uuid.New()).
			WithName("user").
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
				user.Name,
				user.Email.Address,
				user.Password,
				user.Role,
			).
			WillReturnError(fmt.Errorf("sql error"))

		repo := postgres.NewUserRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "request", user)
		err = repo.Update(ctx, user)

		sCtx.Assert().Error(err)
	})
}

func (suite *UserRepoSuite) TestUserRepo_DeleteById1(t provider.T) {
	t.Title("[User delete by id] successfully deleted")
	t.Tags("user", "delete_by_id")
	t.Parallel()

	t.WithNewStep("successfully deleted", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		id := uuid.New()

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatal(err)
		}
		defer mock.Close()
		mock.ExpectExec("delete").
			WithArgs(
				id,
			).
			WillReturnResult(pgxmock.NewResult("delete", 1))

		repo := postgres.NewUserRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "request", id)
		err = repo.DeleteById(ctx, id)

		sCtx.Assert().NoError(err)
	})
}

func (suite *UserRepoSuite) TestUserRepo_DeleteById2(t provider.T) {
	t.Title("[User delete by id] sql error")
	t.Tags("user", "delete_by_id")
	t.Parallel()

	t.WithNewStep("sql error", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		id := uuid.New()

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatal(err)
		}
		defer mock.Close()
		mock.ExpectExec("delete").
			WithArgs(
				id,
			).
			WillReturnError(fmt.Errorf("sql error"))

		repo := postgres.NewUserRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "request", id)
		err = repo.DeleteById(ctx, id)

		sCtx.Assert().Error(err)
	})
}

func (suite *UserRepoSuite) TestUserRepo_GetByUsername1(t provider.T) {
	t.Title("[User get by username] successful get")
	t.Tags("user", "get_by_username")
	t.Parallel()

	t.WithNewStep("successful get", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		user := utils.NewUserBuilder().
			WithId(uuid.New()).
			WithName("user").
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
		mock.ExpectQuery("select").
			WithArgs(
				user.Username,
			).
			WillReturnRows(
				pgxmock.NewRows([]string{"ID", "Name", "Email", "Username", "Password", "Role"}).
					AddRow(user.ID, user.Name, user.Email.Address, user.Username, user.Password, user.Role),
			)

		repo := postgres.NewUserRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "request", user.Username)
		rUser, err := repo.GetByUsername(ctx, user.Username)

		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(user, rUser)
	})
}

func (suite *UserRepoSuite) TestUserRepo_GetByUsername2(t provider.T) {
	t.Title("[User get by username] failed to get")
	t.Tags("user", "get_by_username")
	t.Parallel()

	t.WithNewStep("failed to get", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		username := "username"

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatal(err)
		}
		defer mock.Close()
		mock.ExpectQuery("select").
			WithArgs(
				username,
			).
			WillReturnRows(
				pgxmock.NewRows([]string{"ID", "Name", "Email", "Username", "Password", "Role"}),
			)

		repo := postgres.NewUserRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "request", username)
		rUser, err := repo.GetByUsername(ctx, username)

		sCtx.Assert().Error(err)
		sCtx.Assert().Nil(rUser)
	})
}
