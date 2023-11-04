package repo

import (
	"context"
	"user-service/internal/entity"
	"user-service/pkg/postgres"

	"github.com/gofrs/uuid/v5"
)

type RolePG struct {
	*postgres.Postgres
}

func NewRolePG(pg *postgres.Postgres) *RolePG {
	return &RolePG{pg}
}

func (r *RolePG) GetByUser(ctx context.Context, userID uuid.UUID) ([]entity.Role, error)
func (r *RolePG) List(ctx context.Context) ([]entity.Role, error)
