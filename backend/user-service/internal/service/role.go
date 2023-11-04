package service

import (
	"user-service/internal/entity"
	"user-service/internal/repo"

	"github.com/gofrs/uuid/v5"
)

type role struct {
	repo repo.Role
}

func NewRole(repo repo.Role) *role {
	return &role{repo}
}

func (r *role) GetByUser(userID uuid.UUID) ([]entity.Role, error)
func (r *role) List() ([]entity.Role, error)
