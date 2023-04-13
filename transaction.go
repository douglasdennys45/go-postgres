package gopostgres

import (
	"context"
	"database/sql"
)

type PostgresTransctionInterface interface {
	OpenTransaction(ctx context.Context) (*sql.Tx, error)
	Commit(tx *sql.Tx) error
	Rollback(tx *sql.Tx) error
	GetTX() *sql.Tx
	GetContext() context.Context
}

type postgresTransaction struct{}

func NewPostgresTransaction() PostgresTransctionInterface {
	return &postgresTransaction{}
}

func (postgres *postgresTransaction) OpenTransaction(context context.Context) (*sql.Tx, error) {
	transaction, err := db.BeginTx(ctx, nil)
	tx = transaction
	ctx = context
	return tx, err
}

func (postgres *postgresTransaction) Commit(tx *sql.Tx) error {
	return tx.Commit()
}

func (postgres *postgresTransaction) Rollback(tx *sql.Tx) error {
	return tx.Rollback()
}

func (postgres *postgresTransaction) GetTX() *sql.Tx {
	return tx
}

func (postgres *postgresTransaction) GetContext() context.Context {
	return ctx
}
