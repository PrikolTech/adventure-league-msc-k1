package net

import (
	"auth-service/internal/entity"
)

type User interface {
	Authenticate(email string, password string) (string, []entity.Role, error)
}
