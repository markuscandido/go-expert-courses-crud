package entity

import (
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/markuscandido/go-expert-courses-crud/internal/application/core/domain"
)

type Category struct {
	ID          string
	Name        string
	Description *string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	IsActive    bool
}

func NewCategory(name string, description *string) (*Category, error) {
	category := &Category{
		ID:          uuid.New().String(),
		Name:        strings.TrimSpace(name),
		Description: description,
	}

	if err := category.Validate(); err != nil {
		return nil, err
	}

	return category, nil
}

func (c *Category) Validate() error {
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

func (c *Category) Update(name string, description *string) error {
	c.Name = strings.TrimSpace(name)
	c.Description = description
	return c.Validate()
}

func (c *Category) Activate() {
	c.IsActive = true
	c.UpdatedAt = time.Now().UTC()
}

func (c *Category) Deactivate() {
	c.IsActive = false
	c.UpdatedAt = time.Now().UTC()
}
