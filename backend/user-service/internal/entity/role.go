package entity

import "github.com/gofrs/uuid/v5"

type Role struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
}
