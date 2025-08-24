package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/markuscandido/go-expert-courses-crud/internal/application/core/domain"
	"github.com/markuscandido/go-expert-courses-crud/internal/application/core/domain/entity"
	"github.com/markuscandido/go-expert-courses-crud/internal/application/core/port"
)

type CourseRepository struct {
	db *sql.DB
}

func NewCourseRepository(db *sql.DB) port.ICourseRepository {
	return &CourseRepository{
		db: db,
	}
}

func (r *CourseRepository) Create(ctx context.Context, course *entity.Course) error {
	query := `
		INSERT INTO courses (id, name, description, category_id)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`

	err := r.db.QueryRowContext(
		ctx,
		query,
		course.ID,
		course.Name,
		course.Description,
		course.CategoryID,
	).Scan(&course.ID)

	return err
}

func (r *CourseRepository) FindById(ctx context.Context, id string) (*entity.Course, error) {
	var course entity.Course
	query := `SELECT id, name, description, category_id, created_at, updated_at, is_active FROM courses WHERE id = $1`

	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&course.ID,
		&course.Name,
		&course.Description,
		&course.CategoryID,
		&course.CreatedAt,
		&course.UpdatedAt,
		&course.IsActive,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &course, nil
}

func (r *CourseRepository) FindAll(ctx context.Context, limit, offset int) ([]*entity.Course, error) {
	query := `SELECT id, name, description, category_id, created_at, updated_at, is_active FROM courses`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := rows.Close(); err != nil {
			log.Fatalf("FATAL: %v", err)
		}
	}()

	var courses []*entity.Course
	for rows.Next() {
		var course entity.Course
		err := rows.Scan(
			&course.ID,
			&course.Name,
			&course.Description,
			&course.CategoryID,
			&course.CreatedAt,
			&course.UpdatedAt,
			&course.IsActive,
		)
		if err != nil {
			return nil, err
		}
		courses = append(courses, &course)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return courses, nil
}

func (r *CourseRepository) FindByCategoryId(ctx context.Context, categoryId string) ([]*entity.Course, error) {
	query := `SELECT id, name, description, category_id, created_at, updated_at, is_active FROM courses WHERE category_id = $1`
	rows, err := r.db.QueryContext(ctx, query, categoryId)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := rows.Close(); err != nil {
			log.Fatalf("FATAL: %v", err)
		}
	}()

	var courses []*entity.Course
	for rows.Next() {
		var course entity.Course
		err := rows.Scan(
			&course.ID,
			&course.Name,
			&course.Description,
			&course.CategoryID,
			&course.CreatedAt,
			&course.UpdatedAt,
			&course.IsActive,
		)
		if err != nil {
			return nil, err
		}
		courses = append(courses, &course)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return courses, nil
}

func (r *CourseRepository) Update(ctx context.Context, course *entity.Course) error {
	query := `
		UPDATE courses
		SET name = $1, description = $2, category_id = $3
		WHERE id = $4
		RETURNING id
	`

	err := r.db.QueryRowContext(
		ctx,
		query,
		course.Name,
		course.Description,
		course.CategoryID,
		course.ID,
	).Scan(&course.ID)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.ErrNotFound
		}
		return fmt.Errorf("failed to update course: %w", err)
	}

	return nil
}

func (r *CourseRepository) Delete(ctx context.Context, id string) error {
	query := `
		DELETE FROM courses
		WHERE id = $1
	`

	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete course: %w", err)
	}

	return nil
}

func (r *CourseRepository) Count(ctx context.Context) (int, error) {
	query := `SELECT COUNT(*) FROM courses`
	var count int
	err := r.db.QueryRowContext(ctx, query).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *CourseRepository) Exists(ctx context.Context, id string) (bool, error) {
	query := `SELECT EXISTS (SELECT 1 FROM courses WHERE id = $1)`
	var exists bool
	err := r.db.QueryRowContext(ctx, query, id).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}
