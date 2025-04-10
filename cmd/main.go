package main

import (
	"context"
	"fmt"

	"github.com/MarcoVitoC/memori/internal/auth"
	"github.com/MarcoVitoC/memori/internal/database"
	"github.com/MarcoVitoC/memori/internal/env"
	"github.com/MarcoVitoC/memori/internal/handler"
	"github.com/MarcoVitoC/memori/pkg"
	"github.com/joho/godotenv"
)

func main() {
	logger := pkg.NewLogger()

	err := godotenv.Load()
	if err != nil {
		logger.Fatalw("Failed to load .env", "error", err)
	}

	ctx := context.Background()

	dbCfg := database.NewDBConfig(
		fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
			env.GetString("DB_USERNAME", "postgres"),
			env.GetString("DB_PASSWORD", ""),
			env.GetString("DB_HOST", "localhost"),
			env.GetString("DB_PORT", "5432"),
			env.GetString("DB_DATABASE", "memori")),
		env.GetInt("DB_MAX_CONNS", 10),
		env.GetInt("DB_MIN_CONNS", 2),
	)

	db, err := database.Init(ctx, logger, dbCfg)
	if err != nil {
		logger.Fatalw("Failed to connect to database", "error", err)
	}
	defer db.Close()

	authenticator := auth.NewAuthenticator(env.GetString("JWT_PRIVATE_KEY", "PRIVATE_KEY"))

	server := handler.Server{
		Addr: env.GetString("APP_PORT", ":8080"),
		DB: db,
		Authenticator: authenticator,
	}

	mux := server.Mount()
	logger.Fatal(server.Run(logger, mux))
}