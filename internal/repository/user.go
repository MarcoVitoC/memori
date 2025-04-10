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

func (r *UserRepository) GetByEmail(ctx context.Context, email string) (*User, error) {
	query := `
		SELECT * FROM users
		WHERE email = $1
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	row := r.db.QueryRow(ctx, query, email)

	user := new(User)
	if err := row.Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	); err != nil {
		switch {
		case errors.Is(err, pgx.ErrNoRows):
			return nil, nil
		default:
			return nil, err
		}
	}

	return user, nil
}
