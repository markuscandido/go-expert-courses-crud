package dto

import (
	"time"
)

type (
	CourseResponse struct {
		ID          string    `json:"id"`
		Name        string    `json:"name"`
		Description *string   `json:"description,omitempty"`
		CategoryID  string    `json:"categoryID"`
		IsActive    bool      `json:"isActive"`
		CreatedAt   time.Time `json:"createdAt"`
		UpdatedAt   time.Time `json:"updatedAt"`
	}
)
