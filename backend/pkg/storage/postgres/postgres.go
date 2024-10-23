package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"ppo/internal/config"
)

func NewConn(ctx context.Context, cfg *config.Config) (*pgxpool.Pool, error) {
	connStr := fmt.Sprintf("%s://%s:%s@%s:%d/%s", cfg.Database.Driver, cfg.Database.User,
		cfg.Database.Password, cfg.Database.Host, cfg.Database.Port, cfg.Database.DBName)

	pool, err := pgxpool.New(ctx, connStr)
	if err != nil {
		return nil, fmt.Errorf("connecting to database: %w", err)
	}

	err = pool.Ping(ctx)
	if err != nil {
		return nil, fmt.Errorf("pinging database: %w", err)
	}
	return pool, nil
}
