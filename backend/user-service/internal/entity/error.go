package entity

import "fmt"

type RequiredError struct {
	field string
}

func (e *RequiredError) Error() string {
	return fmt.Sprintf("%s is required", e.field)
}
