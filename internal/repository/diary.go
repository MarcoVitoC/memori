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

func (r *DiaryRepository) GetAll(ctx context.Context) ([]Diary, error) {
	query := `SELECT * FROM diaries`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	var diaries []Diary
	for rows.Next() {
		var diary Diary

		if err := rows.Scan(
			&diary.ID,
			&diary.Content,
			&diary.CreatedAt,
			&diary.UpdatedAt,
		); err != nil {
			return nil, err
		}

		diaries = append(diaries, diary)
	}

	return diaries, nil
}

func (r *DiaryRepository) GetById(ctx context.Context, id string) (Diary, error) {
	query := `
		SELECT * FROM diaries
		WHERE id = $1
	`

	row := r.db.QueryRow(ctx, query, id)
	var diary Diary
	if err := row.Scan(
		&diary.ID,
		&diary.Content,
		&diary.CreatedAt,
		&diary.UpdatedAt,
	); err != nil {
		return Diary{}, err
	}

	return diary, nil
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

func (r *DiaryRepository) Update(ctx context.Context, id string, updatedDiary *Diary) error {
	return withTx(r.db, ctx, func(tx pgx.Tx) error {
		query :=`
			UPDATE diaries
			SET 
				content = $1,
				updated_at = $2
			WHERE id = $3
		`

		ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
		defer cancel()

		if _, err := tx.Exec(
			ctx,
			query,
			updatedDiary.Content,
			time.Now(),
			id,
		); err != nil {
			return err
		}

		return nil
	})
}

func (r DiaryRepository) Delete(ctx context.Context, id string) error {
	return withTx(r.db, ctx, func(tx pgx.Tx) error {
		query := `
			DELETE FROM diaries
			WHERE id = $1
		`

		ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
		defer cancel()

		if _, err := tx.Exec(ctx, query, id); err != nil {
			return err
		}

		return nil
	})
}
