package gopostgres

import (
	"context"
	"database/sql"
)

var (
	db  *sql.DB
	tx  *sql.Tx
	ctx context.Context
)
