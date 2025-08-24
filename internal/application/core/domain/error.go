package domain

import "errors"

var (
	ErrNotFound           = errors.New("not found")
	ErrInvalidInput       = errors.New("invalid input")
	ErrNameIsRequired     = errors.New("name is required")
	ErrNameTooShort       = errors.New("name is too short (min 3 characters)")
	ErrNameTooLong        = errors.New("name is too long (max 100 characters)")
	ErrInvalidDescription = errors.New("description is too long (max 255 characters)")
)
