package service

import (
	"context"
	"user-service/internal/entity"
	"user-service/internal/repo"

	uuid "github.com/jackc/pgx-gofrs-uuid"
)

type role struct {
	repo repo.Role
}

func NewRole(repo repo.Role) *role {
	return &role{repo}
}

func (r *role) GetByUser(ctx context.Context, userID uuid.UUID) ([]entity.Role, error)
func (r *role) List(ctx context.Context) ([]entity.Role, error)
