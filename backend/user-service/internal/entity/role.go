package entity

import "github.com/gofrs/uuid/v5"

type Role struct {
	ID          uuid.UUID `json:"id"`
	Title       *string   `json:"title"`
	Description *string   `json:"description"`
}

func (e *Role) Validate() error {
	if e.Title == nil {
		return &RequiredError{"title"}
	}

	if e.Description == nil {
		return &RequiredError{"description"}
	}

	return nil
}
