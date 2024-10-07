package postgresql

import (
	"backend/src/internal/model"
	"backend/src/internal/model/dto"
	repositoryInterface "backend/src/internal/repository/interface"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) repositoryInterface.IUserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r UserRepository) Get(ctx context.Context, request *dto.GetUserRequest) (user *model.User, err error) {

	query := `    select id, 
				       login,
				       password,
				       role,
				       first_name,
				       second_name,
				       third_name
				from "user"
				where id = $1`

	user = new(model.User)
	//var startTime string = ""
	//var endTime string = ""

	err = r.db.QueryRow(
		ctx,
		query,
		request.Id,
	).Scan(
		&user.Id,
		&user.Login,
		&user.Role,
		&user.Password,
		&user.FirstName,
		&user.SecondName,
		&user.ThirdName,
	)

	if err != nil {
		return nil, fmt.Errorf("запрос не выполнен: %w", err)
	}

	return user, err
}

func (r UserRepository) GetByLogin(ctx context.Context, request *dto.GetUserByLoginRequest) (user *model.User, err error) {
	query := `
				select id, 
				       login,
				       password,
				       role,
				       first_name,
				       second_name,
				       third_name
				from "user" 
				where login = $1`

	user = new(model.User)

	err = r.db.QueryRow(
		ctx,
		query,
		request.Login,
	).Scan(
		&user.Id,
		&user.Login,
		&user.Password,
		&user.Role,
		&user.FirstName,
		&user.SecondName,
		&user.ThirdName,
	)

	if err != nil {
		return nil, fmt.Errorf("запрос не выполнен: %w", err)
	}

	return user, err
}

func (r UserRepository) Add(ctx context.Context, request *dto.AddUserRequest) (err error) {
	query := `insert into "user"(login, password, role, first_name, second_name, third_name) values ($1, $2, $3, $4, $5, $6)`

	_, err = r.db.Exec(
		ctx,
		query,
		request.Login,
		request.Password,
		request.Role,
		request.FirstName,
		request.SecondName,
		request.ThirdName,
	)
	if err != nil {
		return fmt.Errorf("добавление пользователя: %w", err)
	}

	return err
}

func (r UserRepository) Update(ctx context.Context, request *dto.UpdateUserRequest) (err error) {
	query := `
			update "user"
			set 
			    id = $1,
			    login = $2,
			    password = $3,
			    first_name = $4,
			    second_name = $5,
			    third_name = $6
			where id = $1`

	_, err = r.db.Exec(
		ctx,
		query,
		request.Id,
		request.Login,
		request.Password,
		request.FirstName,
		request.SecondName,
		request.ThirdName,
	)
	if err != nil {
		return fmt.Errorf("обновление информации о пользователе: %w", err)
	}

	return err
}

func (r UserRepository) Delete(ctx context.Context, request *dto.DeleteUserRequest) (err error) {
	query := `delete from "user" where id = $1`

	_, err = r.db.Exec(
		ctx,
		query,
		request.Id,
	)
	if err != nil {
		return fmt.Errorf("удаление пользователя по id: %w", err)
	}

	return nil
}
