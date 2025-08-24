package main

import (
	"log"

	"github.com/markuscandido/go-expert-courses-crud/internal/application"
	"github.com/markuscandido/go-expert-courses-crud/internal/application/adapter/driving/graphql"
)

func main() {
	db, cfg, useCases, err := application.Setup()
	if err != nil {
		log.Fatalf("FATAL: %v", err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Printf("ERROR: failed to close database: %v", err)
		}
	}()

	graphql.StartGraphQLServer(cfg, useCases)
}
