package dto

import "time"

type (
	CreateCategoryInputDTO struct {
		Name        string  `json:"name" validate:"required,min=3,max=100"`
		Description *string `json:"description" validate:"max=255"`
	}

	CategoryDTO struct {
		ID          string    `json:"id"`
		Name        string    `json:"name"`
		Description *string   `json:"description"`
		IsActive    bool      `json:"is_active"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
	}

	ListCategoriesOutputDTO struct {
		Items  []CategoryDTO `json:"items"`
		Total  int           `json:"total"`
		Limit  int           `json:"limit"`
		Offset int           `json:"offset"`
	}

	UpdateCategoryInputDTO struct {
		ID          string  `json:"-"` // ID da categoria a ser atualizada
		Name        string  `json:"name,omitempty" validate:"omitempty,min=3,max=100"`
		Description *string `json:"description,omitempty" validate:"omitempty,max=255"`
		IsActive    *bool   `json:"is_active,omitempty"`
	}
)
