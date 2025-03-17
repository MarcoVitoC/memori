package main

import (
	"log"

	"github.com/MarcoVitoC/memori/internal"
)

func main() {
	server := internal.Server{
		Addr: "localhost:8080",
	}

	mux := server.Mount()
	log.Fatal(server.Run(mux))
}