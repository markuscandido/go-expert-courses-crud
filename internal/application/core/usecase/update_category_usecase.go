package usecase

import (
	"context"
	"fmt"

	"github.com/markuscandido/go-expert-courses-crud/internal/application/core/domain"
	"github.com/markuscandido/go-expert-courses-crud/internal/application/core/domain/dto"
	"github.com/markuscandido/go-expert-courses-crud/internal/application/core/port"
)

type UpdateCategoryUseCase struct {
	CategoryRepository port.ICategoryRepository
}

func NewUpdateCategoryUseCase(categoryRepository port.ICategoryRepository) port.IUpdateCategoryUseCase {
	return &UpdateCategoryUseCase{
		CategoryRepository: categoryRepository,
	}
}

func (uc *UpdateCategoryUseCase) Execute(ctx context.Context, input dto.UpdateCategoryInputDTO) error {
	category, err := uc.CategoryRepository.FindById(ctx, input.ID)
	if err != nil {
		if err == domain.ErrNotFound {
			return fmt.Errorf("category with ID %s not found: %w", input.ID, err)
		}
		return fmt.Errorf("failed to get category: %w", err)
	}

	if !category.IsActive {
		return fmt.Errorf("category with ID %s is not active", input.ID)
	}

	category.Name = input.Name
	category.Description = input.Description

	return uc.CategoryRepository.Update(ctx, category)
}
