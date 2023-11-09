package net

import (
	"auth-service/internal/entity"
)

type User interface {
	Authenticate(email string, password string) (string, error)
}

type Role interface {
	GetByUser(userID string, token string) ([]entity.Role, error)
}
