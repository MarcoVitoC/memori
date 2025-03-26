package internal

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func InitDB(conn string) (*pgxpool.Pool, error) {
	db, err := pgxpool.New(context.Background(), conn)
	if err != nil {
		return nil, err
	}

	log.Println("INFO: database connected successfully!")
	return db, nil
}