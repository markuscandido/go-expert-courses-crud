package memory

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"sync"
	"time"

	"github.com/markuscandido/go-expert-courses-crud/internal/application/core/domain"
	"github.com/markuscandido/go-expert-courses-crud/internal/application/core/domain/entity"
	"github.com/markuscandido/go-expert-courses-crud/internal/application/core/port"
)

type categoryRepository struct {
	categories map[string]*entity.Category
	courses    map[string]*entity.Course
	mu         sync.RWMutex
}

func NewCategoryRepository() port.ICategoryRepository {
	return &categoryRepository{
		categories: make(map[string]*entity.Category),
		courses:    make(map[string]*entity.Course),
	}
}

func (r *categoryRepository) Create(ctx context.Context, category *entity.Category) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.categories[category.ID] = category
	return nil
}

func (r *categoryRepository) FindById(ctx context.Context, id string) (*entity.Category, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	category, exists := r.categories[id]
	if !exists || !category.IsActive {
		return nil, errors.New("category not found")
	}

	return &entity.Category{
		ID:          category.ID,
		Name:        category.Name,
		Description: category.Description,
		CreatedAt:   category.CreatedAt,
		UpdatedAt:   category.UpdatedAt,
		IsActive:    category.IsActive,
	}, nil
}

func (r *categoryRepository) FindAll(ctx context.Context, limit, offset int) ([]*entity.Category, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var result []*entity.Category
	for _, category := range r.categories {
		if category.IsActive {
			result = append(result, category)
		}
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].CreatedAt.After(result[j].CreatedAt)
	})

	start, end := paginate(len(result), limit, offset)
	return result[start:end], nil
}

func (r *categoryRepository) FindByCourseId(ctx context.Context, courseId string) (*entity.Category, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	course, exists := r.courses[courseId]
	if !exists {
		return nil, fmt.Errorf("course with ID %s not found: %w", courseId, domain.ErrNotFound)
	}

	category, exists := r.categories[course.CategoryID]
	if !exists || !category.IsActive {
		return nil, fmt.Errorf("category not found for course %s: %w", courseId, domain.ErrNotFound)
	}

	return &entity.Category{
		ID:          category.ID,
		Name:        category.Name,
		Description: category.Description,
		IsActive:    category.IsActive,
		CreatedAt:   category.CreatedAt,
		UpdatedAt:   category.UpdatedAt,
	}, nil
}

func (r *categoryRepository) Update(ctx context.Context, category *entity.Category) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.categories[category.ID]; !exists {
		return errors.New("category not found")
	}

	r.categories[category.ID] = category
	return nil
}

func (r *categoryRepository) Delete(ctx context.Context, id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	category, exists := r.categories[id]
	if !exists || !category.IsActive {
		return errors.New("category not found")
	}

	category.IsActive = false
	category.UpdatedAt = time.Now()
	return nil
}

func (r *categoryRepository) Count(ctx context.Context) (int, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	count := 0
	for _, category := range r.categories {
		if category.IsActive {
			count++
		}
	}

	return count, nil
}

func (r *categoryRepository) Exists(ctx context.Context, id string) (bool, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	category, exists := r.categories[id]
	return exists && category.IsActive, nil
}

func paginate(total, limit, offset int) (int, int) {
	if limit <= 0 {
		limit = 10
	}

	if offset < 0 {
		offset = 0
	}

	start := min(offset, total)

	end := min(start+limit, total)

	return start, end
}
