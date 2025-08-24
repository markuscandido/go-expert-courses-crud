package memory

import (
	"context"
	"sort"
	"sync"
	"time"

	"github.com/markuscandido/go-expert-courses-crud/internal/application/core/domain/entity"
	"github.com/markuscandido/go-expert-courses-crud/internal/application/core/port"
)

type courseRepository struct {
	mu      sync.RWMutex
	courses map[string]*entity.Course
}

func NewCourseRepository() port.ICourseRepository {
	return &courseRepository{
		courses: make(map[string]*entity.Course),
	}
}

func (r *courseRepository) Create(ctx context.Context, course *entity.Course) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if course.CreatedAt.IsZero() {
		course.CreatedAt = time.Now()
	}

	r.courses[course.ID] = course
	return nil
}

func (r *courseRepository) FindById(ctx context.Context, id string) (*entity.Course, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	course, exists := r.courses[id]
	if !exists {
		return nil, nil
	}

	return &entity.Course{
		ID:          course.ID,
		Name:        course.Name,
		Description: course.Description,
		CategoryID:  course.CategoryID,
		CreatedAt:   course.CreatedAt,
	}, nil
}

func (r *courseRepository) FindAll(ctx context.Context, limit, offset int) ([]*entity.Course, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var result []*entity.Course
	for _, course := range r.courses {
		result = append(result, &entity.Course{
			ID:          course.ID,
			Name:        course.Name,
			Description: course.Description,
			CategoryID:  course.CategoryID,
			CreatedAt:   course.CreatedAt,
		})
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].CreatedAt.After(result[j].CreatedAt)
	})

	start, end := paginate(len(result), limit, offset)
	return result[start:end], nil
}

func (r *courseRepository) FindByCategoryId(ctx context.Context, categoryId string) ([]*entity.Course, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var result []*entity.Course
	for _, course := range r.courses {
		if course.CategoryID == categoryId {
			result = append(result, &entity.Course{
				ID:          course.ID,
				Name:        course.Name,
				Description: course.Description,
				CategoryID:  course.CategoryID,
				CreatedAt:   course.CreatedAt,
			})
		}
	}

	return result, nil
}

func (r *courseRepository) Delete(ctx context.Context, id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	delete(r.courses, id)
	return nil
}

func (r *courseRepository) Update(ctx context.Context, course *entity.Course) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.courses[course.ID]; !exists {
		return nil
	}

	r.courses[course.ID] = course
	return nil
}

func (r *courseRepository) Count(ctx context.Context) (int, error) {
	return len(r.courses), nil
}

func (r *courseRepository) Exists(ctx context.Context, id string) (bool, error) {
	_, exists := r.courses[id]
	return exists, nil
}
