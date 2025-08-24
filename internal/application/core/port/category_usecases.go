package port

import (
	"context"

	"github.com/markuscandido/go-expert-courses-crud/internal/application/core/domain/dto"
)

type (
	ICreateCategoryUseCase interface {
		Execute(ctx context.Context, input dto.CreateCategoryInputDTO) (*dto.CategoryDTO, error)
	}

	IListCategoriesUseCase interface {
		Execute(ctx context.Context, limit, offset int) (*dto.ListCategoriesOutputDTO, error)
	}

	IGetCategoryByIdUseCase interface {
		Execute(ctx context.Context, id string) (*dto.CategoryDTO, error)
	}

	IUpdateCategoryUseCase interface {
		Execute(ctx context.Context, input dto.UpdateCategoryInputDTO) error
	}

	IDeleteCategoryUseCase interface {
		Execute(ctx context.Context, id string) error
	}
)
