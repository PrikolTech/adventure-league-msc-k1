package entity

import (
	"errors"
	"fmt"
)

type RequiredError struct {
	Field string
}

func (e *RequiredError) Error() string {
	return fmt.Sprintf("%s is required", e.Field)
}

var (
	ErrEmailInvalid = errors.New("email is invalid")
)
