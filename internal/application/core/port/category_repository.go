package port

import (
	"context"

	"github.com/markuscandido/go-expert-courses-crud/internal/application/core/domain/entity"
)

type ICategoryRepository interface {
	Create(ctx context.Context, category *entity.Category) error
	FindById(ctx context.Context, id string) (*entity.Category, error)
	FindAll(ctx context.Context, limit, offset int) ([]*entity.Category, error)
	FindByCourseId(ctx context.Context, courseId string) (*entity.Category, error)
	Update(ctx context.Context, category *entity.Category) error
	Delete(ctx context.Context, id string) error
	Count(ctx context.Context) (int, error)
	Exists(ctx context.Context, id string) (bool, error)
}
