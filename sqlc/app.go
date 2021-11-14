package main

import (
	"context"
	"database/sql"

	"github.com/inari111/sandbox/sqlc/db"
)

func main() {

}

func createUser(ctx context.Context) error {
	postgresDB, err := sql.Open("postgres", "")
	if err != nil {
		return err
	}

	queries := db.New(postgresDB)

	tx, err := postgresDB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	args := db.CreateParams{
		ID:   "00000000-0000-0000-0000-000000000000",
		Name: "inari111",
	}
	if _, err := queries.WithTx(tx).Create(ctx, args); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
