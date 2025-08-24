package dto

import (
	"time"
)

type (
	CategoryResponse struct {
		ID          string    `json:"id"`
		Name        string    `json:"name"`
		Description *string   `json:"description,omitempty"`
		IsActive    bool      `json:"isActive"`
		CreatedAt   time.Time `json:"createdAt"`
		UpdatedAt   time.Time `json:"updatedAt"`
	}
)
