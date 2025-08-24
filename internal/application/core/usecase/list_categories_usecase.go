package usecase

import (
	"context"
	"fmt"

	"github.com/markuscandido/go-expert-courses-crud/internal/application/core/domain/dto"
	"github.com/markuscandido/go-expert-courses-crud/internal/application/core/port"
)

type ListCategoriesUseCase struct {
	CategoryRepository port.ICategoryRepository
}

func NewListCategoriesUseCase(categoryRepository port.ICategoryRepository) port.IListCategoriesUseCase {
	return &ListCategoriesUseCase{
		CategoryRepository: categoryRepository,
	}
}

func (uc *ListCategoriesUseCase) Execute(ctx context.Context, limit, offset int) (*dto.ListCategoriesOutputDTO, error) {
	if limit <= 0 {
		limit = 10
	}
	if offset < 0 {
		offset = 0
	}

	categories, err := uc.CategoryRepository.FindAll(ctx, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to list categories: %w", err)
	}

	total, err := uc.CategoryRepository.Count(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to count categories: %w", err)
	}

	items := make([]dto.CategoryDTO, 0, len(categories))
	for _, category := range categories {
		items = append(items, dto.CategoryDTO{
			ID:          category.ID,
			Name:        category.Name,
			Description: category.Description,
			IsActive:    category.IsActive,
			CreatedAt:   category.CreatedAt,
			UpdatedAt:   category.UpdatedAt,
		})
	}

	return &dto.ListCategoriesOutputDTO{
		Items:  items,
		Total:  total,
		Limit:  limit,
		Offset: offset,
	}, nil
}
