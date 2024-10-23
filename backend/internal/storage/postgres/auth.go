package postgres

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"ppo/domain"
)

type authRepository struct {
	//db *pgxpool.Pool
	db IPool
}

func NewAuthRepository(db IPool) domain.IAuthRepository {
	return &authRepository{
		db: db,
	}
}

func (r *authRepository) Register(ctx context.Context, authInfo *domain.User) (uuid.UUID, error) {
	query := `insert into saladRecipes.user(name, email, login, password) 
		values ($1, $2, $3, $4)
	returning id`

	id := uuid.Nil
	err := r.db.QueryRow(
		ctx,
		query,
		authInfo.Name,
		authInfo.Email.Address,
		authInfo.Username,
		authInfo.Password,
	).Scan(&id)
	if err != nil {
		return uuid.Nil, fmt.Errorf("user registration: %w", err)
	}
	return id, nil
}

func (r *authRepository) GetByUsername(ctx context.Context, username string) (*domain.UserAuth, error) {
	query := `select id, password, role
		from saladRecipes.user
		where login = $1`

	data := new(domain.UserAuth)
	err := r.db.QueryRow(
		ctx,
		query,
		username,
	).Scan(
		&data.ID,
		&data.HashedPass,
		&data.Role,
	)
	if err != nil {
		return nil, fmt.Errorf("getting user by username: %w", err)
	}
	return data, nil
}
