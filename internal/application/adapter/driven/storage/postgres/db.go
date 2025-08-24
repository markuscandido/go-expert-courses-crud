package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

type Database struct {
	db *sql.DB
}

func NewDatabase(dsn string) (*Database, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("FATAL: %v", err)
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		if err := db.Close(); err != nil {
			log.Fatalf("FATAL: %v", err)
			return nil, fmt.Errorf("failed to close database: %w", err)
		}
		log.Fatalf("FATAL: %v", err)
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	return &Database{db: db}, nil
}

func (d *Database) Close() error {
	if d.db != nil {
		if err := d.db.Close(); err != nil {
			log.Fatalf("FATAL: %v", err)
			return fmt.Errorf("failed to close database: %w", err)
		}
	}
	return nil
}

func (d *Database) GetDB() *sql.DB {
	return d.db
}
