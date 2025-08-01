package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	_ "github.com/lib/pq"
	"github.com/markuscandido/go-expert-graphql/graph"
	"github.com/markuscandido/go-expert-graphql/internal/config"
	"github.com/markuscandido/go-expert-graphql/internal/database"
)

func main() {
	log.Println("Starting application...")

	log.Println("Loading configuration...")
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("FATAL: error loading configuration: %v", err)
	}
	log.Println("Configuration loaded successfully.")

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)

	log.Println("Connecting to database...")
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("FATAL: failed to open database: %v", err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Printf("ERROR: failed to close database: %v", err)
		}
	}()
	log.Println("Database connection established.")

	categoryDb := database.NewCategory(db)
	//courseDb := database.NewCourse(db)

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{CategoryDB: categoryDb}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("GraphQL playground available at http://localhost:%s/", cfg.Port)
	log.Printf("Starting server on port %s...", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, nil); err != nil {
		log.Fatalf("FATAL: could not start server: %v", err)
	}
}
