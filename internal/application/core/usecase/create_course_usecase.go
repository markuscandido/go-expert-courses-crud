package usecase

import (
	"context"

	"github.com/markuscandido/go-expert-courses-crud/internal/application/core/domain/dto"
	"github.com/markuscandido/go-expert-courses-crud/internal/application/core/domain/entity"
	"github.com/markuscandido/go-expert-courses-crud/internal/application/core/port"
)

type CreateCourseUseCase struct {
	CourseRepository port.ICourseRepository
}

func NewCreateCourseUseCase(courseRepository port.ICourseRepository) port.ICreateCourseUseCase {
	return &CreateCourseUseCase{
		CourseRepository: courseRepository,
	}
}

func (uc *CreateCourseUseCase) Execute(ctx context.Context, input dto.CreateCourseInputDTO) (*dto.CourseDTO, error) {
	course, err := entity.NewCourse(input.Name, input.Description, input.CategoryID)
	if err != nil {
		return nil, err
	}

	err = uc.CourseRepository.Create(ctx, course)
	if err != nil {
		return nil, err
	}

	return &dto.CourseDTO{
		ID:          course.ID,
		Name:        course.Name,
		Description: course.Description,
		CategoryID:  course.CategoryID,
		IsActive:    course.IsActive,
		CreatedAt:   course.CreatedAt,
		UpdatedAt:   course.UpdatedAt,
	}, nil
}
