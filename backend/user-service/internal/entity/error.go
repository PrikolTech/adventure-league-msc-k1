package entity

import "fmt"

type RequiredError struct {
	Field string
}

func (e *RequiredError) Error() string {
	return fmt.Sprintf("%s is required", e.Field)
}
