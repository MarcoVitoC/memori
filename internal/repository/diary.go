package repository

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Diary struct {
	ID        uuid.UUID `json:"id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type DiaryRepository struct {
	db *pgxpool.Pool
}

func (r *DiaryRepository) Create(ctx context.Context, newDiary *Diary) error {
	return withTx(r.db, ctx, func(tx pgx.Tx) error {
		query := `
			INSERT INTO diaries (id, content, created_at, updated_at)
			VALUES ($1, $2, $3, $4)
		`

		ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
		defer cancel()

		if _, err := tx.Exec(
			ctx,
			query,
			uuid.New(),
			newDiary.Content,
			time.Now(),
			time.Now(),
		); err != nil {
			return err
		}

		return nil
	})
}