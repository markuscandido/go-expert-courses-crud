package usecase

import (
	"context"
	"fmt"

	"github.com/markuscandido/go-expert-courses-crud/internal/application/core/domain"
	"github.com/markuscandido/go-expert-courses-crud/internal/application/core/domain/dto"
	"github.com/markuscandido/go-expert-courses-crud/internal/application/core/port"
)

type GetCategoryByCourseIdUseCase struct {
	CategoryRepository port.ICategoryRepository
}

func NewGetCategoryByCourseIdUseCase(categoryRepository port.ICategoryRepository) port.IGetCategoryByCourseIdUseCase {
	return &GetCategoryByCourseIdUseCase{
		CategoryRepository: categoryRepository,
	}
}

func (uc *GetCategoryByCourseIdUseCase) Execute(ctx context.Context, courseID string) (*dto.CategoryDTO, error) {
	if courseID == "" {
		return nil, domain.ErrInvalidInput
	}

	category, err := uc.CategoryRepository.FindByCourseId(ctx, courseID)
	if err != nil {
		if err == domain.ErrNotFound {
			return nil, fmt.Errorf("category for course ID %s not found: %w", courseID, err)
		}
		return nil, fmt.Errorf("failed to get category by course ID: %w", err)
	}

	if !category.IsActive {
		return nil, fmt.Errorf("category for course ID %s is not active", courseID)
	}

	return &dto.CategoryDTO{
		ID:          category.ID,
		Name:        category.Name,
		Description: category.Description,
		IsActive:    category.IsActive,
		CreatedAt:   category.CreatedAt,
		UpdatedAt:   category.UpdatedAt,
	}, nil
}
