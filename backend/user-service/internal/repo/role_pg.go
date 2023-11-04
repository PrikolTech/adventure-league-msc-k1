package repo

import (
	"context"
	"user-service/internal/entity"
	"user-service/pkg/postgres"

	"github.com/gofrs/uuid/v5"
	"github.com/jackc/pgx/v5"
)

type RolePG struct {
	*postgres.Postgres
}

func NewRolePG(pg *postgres.Postgres) *RolePG {
	return &RolePG{pg}
}

func (r *RolePG) GetByUser(ctx context.Context, userID uuid.UUID) ([]entity.Role, error) {
	const query = `SELECT id, title, description FROM 
		(SELECT * FROM user_role WHERE user_id = $1) AS "user_role" JOIN "role" ON "user_role".role_id = role.id`

	rows, err := r.Pool.Query(ctx, query, userID)
	if err != nil {
		return nil, err
	}

	return pgx.CollectRows(rows, pgx.RowToStructByPos[entity.Role])
}

func (r *RolePG) List(ctx context.Context) ([]entity.Role, error) {
	const query = `SELECT id, title, description FROM "role"`

	rows, err := r.Pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	return pgx.CollectRows(rows, pgx.RowToStructByPos[entity.Role])
}
