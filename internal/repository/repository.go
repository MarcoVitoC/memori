package repository

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

const QueryTimeoutDuration = 5 * time.Second

type Repository struct {
	Diary interface {
		Create(ctx context.Context, diary *Diary) error
	}
}

func NewRepository(db *pgxpool.Pool) Repository {
	return Repository{
		Diary: &DiaryRepository{db},
	}
}

func withTx(db *pgxpool.Pool, ctx context.Context, fn func(tx pgx.Tx) error) error {
	tx, err := db.Begin(ctx)
	if err != nil {
		return err
	}

	if err := fn(tx); err != nil {
		tx.Rollback(ctx)
		return err
	}

	return tx.Commit(ctx)
}