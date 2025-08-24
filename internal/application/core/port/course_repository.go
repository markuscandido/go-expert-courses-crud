package port

import (
	"context"

	"github.com/markuscandido/go-expert-courses-crud/internal/application/core/domain/entity"
)

type ICourseRepository interface {
	Create(ctx context.Context, course *entity.Course) error
	FindById(ctx context.Context, id string) (*entity.Course, error)
	FindAll(ctx context.Context, limit, offset int) ([]*entity.Course, error)
	FindByCategoryId(ctx context.Context, categoryId string) ([]*entity.Course, error)
	Update(ctx context.Context, course *entity.Course) error
	Delete(ctx context.Context, id string) error
	Count(ctx context.Context) (int, error)
	Exists(ctx context.Context, id string) (bool, error)
}
