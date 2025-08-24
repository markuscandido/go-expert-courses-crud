package usecase

import (
	"context"

	"github.com/markuscandido/go-expert-courses-crud/internal/application/core/domain/dto"
	"github.com/markuscandido/go-expert-courses-crud/internal/application/core/port"
)

type ListCoursesUseCase struct {
	CourseRepository port.ICourseRepository
}

func NewListCoursesUseCase(courseRepository port.ICourseRepository) port.IListCoursesUseCase {
	return &ListCoursesUseCase{
		CourseRepository: courseRepository,
	}
}

func (uc *ListCoursesUseCase) Execute(ctx context.Context, limit, offset int) ([]*dto.CourseDTO, error) {
	courses, err := uc.CourseRepository.FindAll(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	var output []*dto.CourseDTO
	for _, course := range courses {
		output = append(output, &dto.CourseDTO{
			ID:          course.ID,
			Name:        course.Name,
			Description: course.Description,
			CategoryID:  course.CategoryID,
			IsActive:    course.IsActive,
			CreatedAt:   course.CreatedAt,
			UpdatedAt:   course.UpdatedAt,
		})
	}

	return output, nil
}
