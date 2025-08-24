package usecase

import (
	"context"
	"fmt"

	"github.com/markuscandido/go-expert-courses-crud/internal/application/core/domain"
	"github.com/markuscandido/go-expert-courses-crud/internal/application/core/domain/dto"
	"github.com/markuscandido/go-expert-courses-crud/internal/application/core/port"
)

type ListCoursesByCategoryIdUseCase struct {
	CourseRepository port.ICourseRepository
}

func NewListCoursesByCategoryIdUseCase(courseRepository port.ICourseRepository) port.IListCoursesByCategoryIdUseCase {
	return &ListCoursesByCategoryIdUseCase{
		CourseRepository: courseRepository,
	}
}

func (uc *ListCoursesByCategoryIdUseCase) Execute(ctx context.Context, categoryID string) ([]*dto.CourseDTO, error) {
	if categoryID == "" {
		return nil, domain.ErrInvalidInput
	}

	courses, err := uc.CourseRepository.FindByCategoryId(ctx, categoryID)
	if err != nil {
		if err == domain.ErrNotFound {
			return nil, fmt.Errorf("no courses found for category ID %s: %w", categoryID, err)
		}
		return nil, fmt.Errorf("failed to list courses by category ID: %w", err)
	}

	var result []*dto.CourseDTO
	for _, course := range courses {
		result = append(result, &dto.CourseDTO{
			ID:          course.ID,
			Name:        course.Name,
			Description: course.Description,
			CategoryID:  course.CategoryID,
			IsActive:    course.IsActive,
			CreatedAt:   course.CreatedAt,
			UpdatedAt:   course.UpdatedAt,
		})
	}

	return result, nil
}
