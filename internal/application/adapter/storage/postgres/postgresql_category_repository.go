package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/markuscandido/go-expert-courses-crud/internal/application/core/domain"
	"github.com/markuscandido/go-expert-courses-crud/internal/application/core/domain/entity"
	"github.com/markuscandido/go-expert-courses-crud/internal/application/core/port"
)

type categoryRepository struct {
	db *sql.DB
}

// NewCategoryRepository cria uma nova instância do repositório PostgreSQL
func NewCategoryRepository(db *sql.DB) port.ICategoryRepository {
	return &categoryRepository{
		db: db,
	}
}

const (
	duplicateKeyError = "23505" // Código de erro para violação de chave única
)

func (r *categoryRepository) Create(ctx context.Context, category *entity.Category) error {
	query := `
		INSERT INTO categories (id, name, description)
		VALUES ($1, $2, $3)
		RETURNING id
	`

	err := r.db.QueryRowContext(
		ctx,
		query,
		category.ID,
		category.Name,
		category.Description,
	).Scan(&category.ID)

	if err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok && pgErr.Code == duplicateKeyError {
			return fmt.Errorf("category with this name already exists")
		}
		return fmt.Errorf("failed to create category: %w", err)
	}

	return nil
}

func (r *categoryRepository) FindById(ctx context.Context, id string) (*entity.Category, error) {
	query := `
		SELECT id, name, description, created_at, updated_at, is_active
		FROM categories
		WHERE id = $1 AND is_active = true
	`

	var category entity.Category
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&category.ID,
		&category.Name,
		&category.Description,
		&category.CreatedAt,
		&category.UpdatedAt,
		&category.IsActive,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("category not found: %w", domain.ErrNotFound)
		}
		return nil, fmt.Errorf("failed to find category: %w", err)
	}

	return &category, nil
}

func (r *categoryRepository) FindAll(ctx context.Context, limit, offset int) ([]*entity.Category, error) {
	query := `
		SELECT id, name, description, created_at, updated_at, is_active
		FROM categories
		WHERE is_active = true
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2
	`

	rows, err := r.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to list categories: %w", err)
	}
	defer func() {
		if err := rows.Close(); err != nil {
			log.Fatalf("FATAL: %v", err)
		}
	}()

	var categories []*entity.Category
	for rows.Next() {
		var category entity.Category
		if err := rows.Scan(
			&category.ID,
			&category.Name,
			&category.Description,
			&category.CreatedAt,
			&category.UpdatedAt,
			&category.IsActive,
		); err != nil {
			return nil, fmt.Errorf("failed to scan category: %w", err)
		}
		categories = append(categories, &category)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating categories: %w", err)
	}

	return categories, nil
}

func (r *categoryRepository) FindByCourseId(ctx context.Context, courseId string) (*entity.Category, error) {
	query := `
		SELECT c.id, c.name, c.description, c.created_at, c.updated_at, c.is_active
		FROM categories c
		INNER JOIN courses co ON c.id = co.category_id
		WHERE co.id = $1 AND c.is_active = true
	`

	var category entity.Category
	err := r.db.QueryRowContext(ctx, query, courseId).Scan(
		&category.ID,
		&category.Name,
		&category.Description,
		&category.CreatedAt,
		&category.UpdatedAt,
		&category.IsActive,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("category not found: %w", domain.ErrNotFound)
		}
		return nil, fmt.Errorf("failed to find category: %w", err)
	}

	return &category, nil
}

func (r *categoryRepository) Update(ctx context.Context, category *entity.Category) error {
	query := `
		UPDATE categories
		SET name = $1, description = $2
		WHERE id = $3 AND is_active = true
		RETURNING id
	`

	err := r.db.QueryRowContext(
		ctx,
		query,
		category.Name,
		category.Description,
		category.ID,
	).Scan(&category.ID)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.ErrNotFound
		}
		return fmt.Errorf("failed to update category: %w", err)
	}

	return nil
}

func (r *categoryRepository) Delete(ctx context.Context, id string) error {
	query := `
		DELETE FROM categories
		WHERE id = $1
	`

	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete category: %w", err)
	}

	return nil
}

func (r *categoryRepository) Count(ctx context.Context) (int, error) {
	query := `SELECT COUNT(*) FROM categories WHERE is_active = true`

	var count int
	if err := r.db.QueryRowContext(ctx, query).Scan(&count); err != nil {
		return 0, fmt.Errorf("failed to count categories: %w", err)
	}

	return count, nil
}

func (r *categoryRepository) Exists(ctx context.Context, id string) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM categories WHERE id = $1 AND is_active = true)`

	var exists bool
	if err := r.db.QueryRowContext(ctx, query, id).Scan(&exists); err != nil {
		return false, fmt.Errorf("failed to check if category exists: %w", err)
	}

	return exists, nil
}
