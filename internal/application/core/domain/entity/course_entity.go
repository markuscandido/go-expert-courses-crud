package entity

import (
	"time"

	"github.com/google/uuid"
	"github.com/markuscandido/go-expert-courses-crud/internal/application/core/domain"
)

type Course struct {
	ID          string
	Name        string
	Description *string
	CategoryID  string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	IsActive    bool
}

func NewCourse(name string, description *string, categoryID string) (*Course, error) {
	course := &Course{
		ID:          uuid.New().String(),
		Name:        name,
		Description: description,
		CategoryID:  categoryID,
	}

	if err := course.Validate(); err != nil {
		return nil, err
	}

	return course, nil
}

func (c *Course) Validate() error {
	if c.Name == "" {
		return domain.ErrNameIsRequired
	}
	if len(c.Name) < 3 {
		return domain.ErrNameTooShort
	}
	if len(c.Name) > 100 {
		return domain.ErrNameTooLong
	}
	if c.Description != nil && len(*c.Description) > 255 {
		return domain.ErrInvalidDescription
	}
	return nil
}

func (c *Course) Update(name string, description *string) error {
	c.Name = name
	c.Description = description
	return c.Validate()
}

func (c *Course) Activate() {
	c.IsActive = true
	c.UpdatedAt = time.Now().UTC()
}

func (c *Course) Deactivate() {
	c.IsActive = false
	c.UpdatedAt = time.Now().UTC()
}
