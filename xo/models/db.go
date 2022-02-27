package models

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
)

type db struct {
	DB *sqlx.DB
}

func NewDB() DB {
	d, err := connectDB()
	if err != nil {
		panic(err)
	}

	return &db{
		DB: d,
	}
}

func connectDB() (*sqlx.DB, error) {
	db, err := sqlx.Open(
		"postgres",
		fmt.Sprintf(
			"postgres://%s:%s@%s:%s/%s?sslmode=disable",
			os.Getenv("PGUSER"),
			os.Getenv("PGPASSWORD"),
			os.Getenv("PGHOST"),
			os.Getenv("PGPORT"),
			os.Getenv("PGDATABASE"),
		),
	)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (d *db) ExecContext(ctx context.Context, s string, i ...interface{}) (sql.Result, error) {
	return d.DB.ExecContext(ctx, s, i...)
}

func (d *db) QueryContext(ctx context.Context, s string, i ...interface{}) (*sql.Rows, error) {
	rows, err := d.DB.QueryxContext(ctx, s, i...)
	if err != nil {
		return nil, err
	}
	return rows.Rows, nil
}

func (d *db) QueryRowContext(ctx context.Context, s string, i ...interface{}) *sql.Row {
	return d.DB.QueryRowContext(ctx, s, i...)
}
