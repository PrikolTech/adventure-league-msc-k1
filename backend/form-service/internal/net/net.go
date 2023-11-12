package net

import (
	"form-service/internal/entity"

	"github.com/gofrs/uuid/v5"
)

type User interface {
	Exist(email string) (bool, error)
	Get(id uuid.UUID, token string) (*entity.User, error)
	Create(data entity.User) (*entity.User, error)
}

type Role interface {
	Append(userID uuid.UUID, title string, token string) error
}

type Course interface {
	Append(userID uuid.UUID, courseID uuid.UUID, token string) error
}

type Auth interface {
	Verify(token string) (string, []string, error)
}

type SMTP interface {
	SendPassword(password string, to string) error
}
