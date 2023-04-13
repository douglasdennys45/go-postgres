package repositories

import (
	"context"
	"database/sql"
)

type PostgresRepositoryInterface interface {
	Insert(context.Context, string, ...any) (sql.Result, error)
	Query(context.Context, string, ...any) (sql.Result, error)
}

type postgresRepository struct {
	*sql.Tx
}

func NewPostgresRepository(tx *sql.Tx) PostgresRepositoryInterface {
	return &postgresRepository{tx}
}

func (postgres *postgresRepository) Insert(ctx context.Context, query string, args ...any) (sql.Result, error) {
	result, err := postgres.ExecContext(ctx, query, args)
	return result, err
}

func (postgres *postgresRepository) Query(ctx context.Context, query string, args ...any) (sql.Result, error) {
	err := postgres.QueryRowContext(ctx, query, args).Scan()
	return nil, err
}
