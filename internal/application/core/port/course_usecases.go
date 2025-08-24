package port

import (
	"context"

	"github.com/markuscandido/go-expert-courses-crud/internal/application/core/domain/dto"
)

type (
	ICreateCourseUseCase interface {
		Execute(ctx context.Context, input dto.CreateCourseInputDTO) (*dto.CourseDTO, error)
	}

	IListCoursesUseCase interface {
		Execute(ctx context.Context, limit, offset int) ([]*dto.CourseDTO, error)
	}

	IGetCourseByIdUseCase interface {
		Execute(ctx context.Context, id string) (*dto.CourseDTO, error)
	}

	IGetCategoryByCourseIdUseCase interface {
		Execute(ctx context.Context, courseID string) (*dto.CategoryDTO, error)
	}

	IListCoursesByCategoryIdUseCase interface {
		Execute(ctx context.Context, categoryID string) ([]*dto.CourseDTO, error)
	}
)
