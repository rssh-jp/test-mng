package domain

import (
	"errors"
)

var (
	ErrNotFound = errors.New("Not found")
	ErrInvalid  = errors.New("Invalid")
)
