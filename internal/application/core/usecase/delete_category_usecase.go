package usecase

import (
	"context"
	"fmt"

	"github.com/markuscandido/go-expert-courses-crud/internal/application/core/domain"
	"github.com/markuscandido/go-expert-courses-crud/internal/application/core/port"
)

type DeleteCategoryUseCase struct {
	CategoryRepository port.ICategoryRepository
}

func NewDeleteCategoryUseCase(categoryRepository port.ICategoryRepository) port.IDeleteCategoryUseCase {
	return &DeleteCategoryUseCase{
		CategoryRepository: categoryRepository,
	}
}

func (uc *DeleteCategoryUseCase) Execute(ctx context.Context, id string) error {
	category, err := uc.CategoryRepository.FindById(ctx, id)
	if err != nil {
		if err == domain.ErrNotFound {
			return fmt.Errorf("category with ID %s not found: %w", id, err)
		}
		return fmt.Errorf("failed to get category: %w", err)
	}

	if !category.IsActive {
		return fmt.Errorf("category with ID %s is not active", id)
	}

	return uc.CategoryRepository.Delete(ctx, id)
}
