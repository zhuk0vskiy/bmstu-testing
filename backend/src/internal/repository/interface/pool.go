package _interface

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

//go:generate go run github.com/vektra/mockery/v2@v2.42.1 --name=IPool
type IPool interface {
	Close()
	Exec(query string, args ...interface{}) (pgconn.CommandTag, error)
	Query(query string, args ...interface{}) (pgx.Rows, error)
	QueryRow(query string, args ...interface{}) pgx.Row
	Begin(ctx context.Context) (pgx.Tx, error)
	BeginTx(ctx context.Context, opts pgx.TxOptions) (pgx.Tx, error)
}