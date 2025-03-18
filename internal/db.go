package internal

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func New() (*pgxpool.Pool, error) {
	db, err := pgxpool.New(context.Background(), "")
	if err != nil {
		return nil, err
	}

	return db, nil
}