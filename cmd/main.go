package main

import (
	"context"
	"log"

	"github.com/MarcoVitoC/memori/internal/database"
	"github.com/MarcoVitoC/memori/internal/env"
	"github.com/MarcoVitoC/memori/internal/handler"
	"github.com/joho/godotenv"
)

type dbConfig struct {
	url string
	maxConns int
	minConns int
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("ERROR: failed to load .env file with error ", err)
	}

	ctx := context.Background()
	dbCfg := dbConfig{
		url: env.GetString("DB_URL", "postgres://postgres:password@localhost:5432/memori"),
		maxConns: env.GetInt("DB_MAX_CONNS", 10),
		minConns: env.GetInt("DB_MIN_CONNS", 2),
	}

	db, err := database.Init(ctx, dbCfg.url, dbCfg.maxConns, dbCfg.minConns)
	if err != nil {
		log.Fatal("ERROR: failed to connect to database with error ", err)
	}
	defer db.Close()

	server := handler.Server{
		Addr: "localhost:8080",
		DB: db,
	}

	mux := server.Mount()
	log.Fatal(server.Run(mux))
}