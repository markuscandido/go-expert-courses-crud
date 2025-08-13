package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
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

	if err := runMigrations(dsn); err != nil {
		log.Fatalf("FATAL: failed to run migrations: %v", err)
	}

	categoryDb := database.NewCategory(db)
	courseDb := database.NewCourse(db)

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{CategoryDB: categoryDb, CourseDB: courseDb}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("GraphQL playground available at http://localhost:%s/", cfg.Port)
	log.Printf("Starting server on port %s...", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, nil); err != nil {
		log.Fatalf("FATAL: could not start server: %v", err)
	}
}

func runMigrations(dsn string) error {
	var migrationsPath string
	var possiblePaths = []string{
		filepath.Join(filepath.Dir(os.Args[0]), "..", "sql", "migrations"),
		filepath.Join(".", "sql", "migrations"),
		filepath.Join(filepath.Dir(filepath.Dir(filepath.Dir(os.Args[0]))), "sql", "migrations"),
	}

	for _, path := range possiblePaths {
		absPath, err := filepath.Abs(path)
		if err != nil {
			continue
		}
		if _, err := os.Stat(filepath.Join(absPath, "000001_create_categories_table.up.sql")); err == nil {
			migrationsPath = absPath
			break
		}
	}

	if migrationsPath == "" {
		return fmt.Errorf("could not find migrations directory. Tried: %v", possiblePaths)
	}

	migrationsURL := fmt.Sprintf("file://%s", filepath.ToSlash(migrationsPath))
	log.Printf("Using migrations from: %s", migrationsURL)

	m, err := migrate.New(migrationsURL, dsn)
	if err != nil {
		return fmt.Errorf("failed to create migrate instance: %v", err)
	}

	log.Println("Running database migrations...")
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("failed to apply migrations: %v", err)
	}

	log.Println("Database migrations completed successfully.")
	return nil
}
