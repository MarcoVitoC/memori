package database

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Init(ctx context.Context, conn string, maxConns int, minConns int) (*pgxpool.Pool, error) {
	cfg, err := pgxpool.ParseConfig(conn)
	if err != nil {
		return nil, err
	}
	
	cfg.MaxConns = int32(maxConns)
	cfg.MinConns = int32(minConns)
	
	db, err := pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(ctx, 5 * time.Second)
	defer cancel()

	if err := db.Ping(ctx); err != nil {
		return nil, err
	}

	log.Println("INFO: database connected successfully!")
	return db, nil
}