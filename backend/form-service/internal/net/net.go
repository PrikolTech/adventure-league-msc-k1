package net

import (
	"form-service/internal/entity"

	"github.com/gofrs/uuid/v5"
)

type User interface {
	Create(data entity.User) (*entity.User, error)
}

type Role interface {
	Append(userID uuid.UUID, roleID uuid.UUID) error
}

type Course interface {
	Append(userID uuid.UUID, courseID uuid.UUID) error
}
