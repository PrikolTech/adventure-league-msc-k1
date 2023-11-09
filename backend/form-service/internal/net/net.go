package net

import (
	"form-service/internal/entity"

	"github.com/gofrs/uuid/v5"
)

type User interface {
	Get(id uuid.UUID) (*entity.User, error)
	Create(data entity.User) (*entity.User, error)
}

type Role interface {
	Append(userID uuid.UUID, title string) error
}

type Course interface {
	Append(userID uuid.UUID, courseID uuid.UUID) error
}
