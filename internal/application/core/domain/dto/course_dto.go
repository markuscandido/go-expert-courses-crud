package dto

import "time"

type (
	CreateCourseInputDTO struct {
		Name        string
		Description *string
		CategoryID  string
	}

	CourseDTO struct {
		ID          string
		Name        string
		Description *string
		CategoryID  string
		IsActive    bool      `json:"is_active"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
	}
)
