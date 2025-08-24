package usecase

import (
	"context"

	"github.com/markuscandido/go-expert-courses-crud/internal/application/core/domain/dto"
	"github.com/markuscandido/go-expert-courses-crud/internal/application/core/domain/entity"
	"github.com/markuscandido/go-expert-courses-crud/internal/application/core/port"
)

type CreateCategoryUseCase struct {
	CategoryRepository port.ICategoryRepository
}

func NewCreateCategoryUseCase(categoryRepository port.ICategoryRepository) port.ICreateCategoryUseCase {
	return &CreateCategoryUseCase{
		CategoryRepository: categoryRepository,
	}
}

func (uc *CreateCategoryUseCase) Execute(ctx context.Context, input dto.CreateCategoryInputDTO) (*dto.CategoryDTO, error) {
	category, err := entity.NewCategory(input.Name, input.Description)
	if err != nil {
		return nil, err
	}

	err = uc.CategoryRepository.Create(ctx, category)
	if err != nil {
		return nil, err
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
