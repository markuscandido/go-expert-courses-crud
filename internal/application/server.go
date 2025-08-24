package application

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/markuscandido/go-expert-courses-crud/internal/application/adapter/config"
	"github.com/markuscandido/go-expert-courses-crud/internal/application/adapter/storage/postgres"
	"github.com/markuscandido/go-expert-courses-crud/internal/application/core/port"
	"github.com/markuscandido/go-expert-courses-crud/internal/application/core/usecase"
)

type UseCases struct {
	CreateCategoryUseCase          port.ICreateCategoryUseCase
	CreateCourseUseCase            port.ICreateCourseUseCase
	ListCategoriesUseCase          port.IListCategoriesUseCase
	ListCoursesUseCase             port.IListCoursesUseCase
	GetCategoryByCourseIdUseCase   port.IGetCategoryByCourseIdUseCase
	ListCoursesByCategoryIdUseCase port.IListCoursesByCategoryIdUseCase
}

func Setup() (*sql.DB, *config.Config, *UseCases, error) {
	log.Println("Loading configuration...")
	cfg, err := config.LoadConfig()
	if err != nil {
		return nil, nil, nil, fmt.Errorf("error loading configuration: %v", err)
	}
	log.Println("Configuration loaded successfully.")

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)

	log.Println("Connecting to database...")
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to open database: %v", err)
	}
	log.Println("Database connection established.")

	if err := runMigrations(dsn); err != nil {
		return nil, nil, nil, fmt.Errorf("failed to run migrations: %v", err)
	}

	// Inicializa os repositórios usando a conexão existente
	categoryRepo := postgres.NewCategoryRepository(db)
	courseRepo := postgres.NewCourseRepository(db)

	useCases := &UseCases{
		CreateCategoryUseCase:          usecase.NewCreateCategoryUseCase(categoryRepo),
		CreateCourseUseCase:            usecase.NewCreateCourseUseCase(courseRepo),
		ListCategoriesUseCase:          usecase.NewListCategoriesUseCase(categoryRepo),
		ListCoursesUseCase:             usecase.NewListCoursesUseCase(courseRepo),
		GetCategoryByCourseIdUseCase:   usecase.NewGetCategoryByCourseIdUseCase(categoryRepo),
		ListCoursesByCategoryIdUseCase: usecase.NewListCoursesByCategoryIdUseCase(courseRepo),
	}
	return db, cfg, useCases, nil
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
