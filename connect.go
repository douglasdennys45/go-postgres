package gopostgres

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type PostgresConnectionInterface interface {
	Connect(driver, dsn string, pool int) error
}

type postgresConnection struct{}

func NewPostgresConnection() PostgresConnectionInterface {
	return &postgresConnection{}
}

func (postgres *postgresConnection) Connect(driver, dsn string, pool int) error {
	database, err := sql.Open(driver, dsn)
	if err != nil {
		return err
	}
	db = database
	db.SetConnMaxLifetime(1000)
	db.SetMaxOpenConns(pool)
	db.SetMaxIdleConns(pool)
	return nil
}
