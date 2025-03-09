package main

import (
	"log"

	"github.com/bredacoder/go-bank/internal/api"
	"github.com/bredacoder/go-bank/internal/storage"
)

func main() {
	store, err := storage.NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}

	if err := store.Init(); err != nil {
		log.Fatal(err)
	}

	server := api.NewAPIServer(":5000", store)
	server.Run()
}
