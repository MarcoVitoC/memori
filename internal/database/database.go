package database

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type DBConfig struct {
	Conn string
	MaxConns int
	MinConns int
}

func NewDBConfig(conn string, maxConns int, minConns int) *DBConfig {
	return &DBConfig{
		Conn: conn,
		MaxConns: maxConns,
		MinConns: minConns,
	}
}

func Init(ctx context.Context, logger *zap.SugaredLogger, dbCfg *DBConfig) (*pgxpool.Pool, error) {
	cfg, err := pgxpool.ParseConfig(dbCfg.Conn)
	if err != nil {
		return nil, err
	}
	
	cfg.MaxConns = int32(dbCfg.MaxConns)
	cfg.MinConns = int32(dbCfg.MinConns)
	
	db, err := pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(ctx, 5 * time.Second)
	defer cancel()

	if err := db.Ping(ctx); err != nil {
		return nil, err
	}

	logger.Info("Database connected successfully")
	return db, nil
}