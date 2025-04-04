package repository

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserRepository struct {
	db *pgxpool.Pool
}

func (r *UserRepository) Register(ctx context.Context, newUser *User) error {
	return withTx(r.db, ctx, func(tx pgx.Tx) error {
		query := `
			INSERT INTO users (id, username, email, password, created_at, updated_at)
			VALUES ($1, $2, $3, $4, $5, $6)
		`

		ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
		defer cancel()

		if _, err := tx.Exec(
			ctx,
			query,
			uuid.New(),
			newUser.Username,
			newUser.Email,
			newUser.Password,
			time.Now(),
			time.Now(),
		); err != nil {
			return err
		}

		return nil
	})
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (bool, error) {
	query := `
		SELECT id FROM users
		WHERE email = $1
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	row := r.db.QueryRow(ctx, query, email)

	var id uuid.UUID
	if err := row.Scan(&id); err != nil {
		switch {
		case errors.Is(err, pgx.ErrNoRows):
			return false, nil
		default:
			return false, err
		}
	}

	return true, nil
}
