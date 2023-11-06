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

func (r *RolePG) Create(ctx context.Context, data entity.Role) (*entity.Role, error) {
	const query = `INSERT INTO "role"
		(title, description)
		VALUES ($1, $2) RETURNING *`

	var role entity.Role
	err := r.Pool.
		QueryRow(ctx, query, data.Title, data.Description).
		Scan(&role.ID, &role.Title, &role.Description)

	return &role, err
}

func (r *RolePG) CreateUserRole(ctx context.Context, userID uuid.UUID, roleID uuid.UUID) error {
	const query = `INSERT INTO "data_role"
		(role_id, data_id) 
		VALUES ($1, $2)`

	_, err := r.Pool.Exec(ctx, query, roleID, userID)
	return err
}

func (r *RolePG) DeleteUserRole(ctx context.Context, userID uuid.UUID, roleID uuid.UUID) error {
	const query = `DELETE FROM "data_role" WHERE data_id = $1 AND role_id = $2`

	_, err := r.Pool.Exec(ctx, query, userID, roleID)
	return err
}

func (r *RolePG) DeleteUserRoles(ctx context.Context, userID uuid.UUID) error {
	const query = `DELETE FROM "data_role" WHERE data_id = $1`

	_, err := r.Pool.Exec(ctx, query, userID)
	return err
}

func (r *RolePG) GetByID(ctx context.Context, id uuid.UUID) (*entity.Role, error) {
	const query = `SELECT * FROM "role" WHERE id = $1`

	var role entity.Role
	err := r.Pool.
		QueryRow(ctx, query, id).
		Scan(&role.ID, &role.Title, &role.Description)

	return &role, err
}

func (r *RolePG) GetByTitle(ctx context.Context, title string) (*entity.Role, error) {
	const query = `SELECT * FROM "role" WHERE title = $1`

	var role entity.Role
	err := r.Pool.
		QueryRow(ctx, query, title).
		Scan(&role.ID, &role.Title, &role.Description)

	return &role, err
}

func (r *RolePG) GetByUser(ctx context.Context, userID uuid.UUID) ([]entity.Role, error) {
	const query = `SELECT id, title, description FROM 
		(SELECT * FROM "data_role" WHERE data_id = $1) AS "data_role" JOIN "role" ON "data_role".role_id = role.id`

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

func (r *RolePG) DeleteByID(ctx context.Context, id uuid.UUID) error {
	const query = `DELETE FROM "role" WHERE id = $1`

	if _, err := r.Pool.Exec(ctx, query, id); err != nil {
		return err
	}

	return nil
}
