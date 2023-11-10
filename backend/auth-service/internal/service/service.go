package service

import (
	"auth-service/internal/entity"
)

type OAuth interface {
	GenerateToken(clientID string, email string, password string) (*entity.Token, error)
	RefreshToken(clientID string, token string) (*entity.Token, error)
	RemoveToken(token string) error
}
