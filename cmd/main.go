package main

import (
	"log"
	"os"

	"github.com/MarcoVitoC/memori/internal"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("ERROR: failed to load .env file with error ", err)
	}

	db, err := internal.InitDB(os.Getenv("DB_URL"))
	if err != nil {
		log.Fatal("ERROR: failed to connect to database with error ", err)
	}
	defer db.Close()

	server := internal.Server{
		Addr: "localhost:8080",
		DB: db,
	}

	mux := server.Mount()
	log.Fatal(server.Run(mux))
}