package usecase

import (
	"context"
	"fmt"

	"github.com/markuscandido/go-expert-courses-crud/internal/application/core/domain"
	"github.com/markuscandido/go-expert-courses-crud/internal/application/core/domain/dto"
	"github.com/markuscandido/go-expert-courses-crud/internal/application/core/port"
)

type GetCategoryByIdUseCase struct {
	CategoryRepository port.ICategoryRepository
}

func NewGetCategoryByIdUseCase(categoryRepository port.ICategoryRepository) port.IGetCategoryByIdUseCase {
	return &GetCategoryByIdUseCase{
		CategoryRepository: categoryRepository,
	}
}

func (uc *GetCategoryByIdUseCase) Execute(ctx context.Context, id string) (*dto.CategoryDTO, error) {
	if id == "" {
		return nil, domain.ErrInvalidInput
	}

	category, err := uc.CategoryRepository.FindById(ctx, id)
	if err != nil {
		if err == domain.ErrNotFound {
			return nil, fmt.Errorf("category with ID %s not found: %w", id, err)
		}
		return nil, fmt.Errorf("failed to get category: %w", err)
	}

	if !category.IsActive {
		return nil, fmt.Errorf("category with ID %s is not active", id)
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
