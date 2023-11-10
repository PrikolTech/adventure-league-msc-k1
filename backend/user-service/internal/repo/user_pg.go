package repo

import (
	"context"
	"user-service/internal/entity"
	"user-service/pkg/postgres"

	"github.com/gofrs/uuid/v5"
	"github.com/jackc/pgx/v5"
)

type UserPG struct {
	*postgres.Postgres
}

func NewUserPG(pg *postgres.Postgres) *UserPG {
	return &UserPG{pg}
}

func (u *UserPG) collectRow(row pgx.Row) (*entity.User, error) {
	var data entity.User
	err := row.Scan(&data.ID, &data.Email, &data.Password, &data.FirstName, &data.LastName, &data.Patronymic, &data.Birthdate, &data.Phone, &data.Telegram)
	if err == pgx.ErrNoRows {
		return nil, ErrNoRows
	}
	return &data, err
}

func (u *UserPG) collectRows(rows pgx.Rows) ([]entity.User, error) {
	return pgx.CollectRows(rows, func(row pgx.CollectableRow) (entity.User, error) {
		data, err := u.collectRow(row)
		return *data, err
	})
}

func (u *UserPG) Create(ctx context.Context, data entity.User) (*entity.User, error) {
	const query = `INSERT INTO "data"
		(email, password, first_name, last_name, patronymic, birthdate, phone, telegram) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING *`

	row := u.Pool.
		QueryRow(ctx, query, data.Email, data.Password, data.FirstName, data.LastName, data.Patronymic, data.Birthdate, data.Phone, data.Telegram)

	return u.collectRow(row)
}

func (u *UserPG) GetByID(ctx context.Context, id uuid.UUID) (*entity.User, error) {
	const query = `SELECT * FROM "data" WHERE id = $1`

	row := u.Pool.QueryRow(ctx, query, id)
	return u.collectRow(row)
}

func (u *UserPG) GetByEmail(ctx context.Context, email string) (*entity.User, error) {
	const query = `SELECT * FROM "data" WHERE email = $1`

	row := u.Pool.QueryRow(ctx, query, email)
	return u.collectRow(row)
}

func (u *UserPG) UpdatePassword(ctx context.Context, id uuid.UUID, password string) (*entity.User, error) {
	const query = `UPDATE "data" SET password = $2 WHERE id = $1 RETURNING *`

	row := u.Pool.QueryRow(ctx, query, id, password)
	return u.collectRow(row)
}

func (u *UserPG) UpdateContact(ctx context.Context, data entity.User) (*entity.User, error) {
	const query = `UPDATE "data" SET phone = $2, telegram = $3 WHERE id = $1 RETURNING *`

	row := u.Pool.QueryRow(ctx, query, data.ID, data.Phone, data.Telegram)
	return u.collectRow(row)
}

func (u *UserPG) DeleteByID(ctx context.Context, id uuid.UUID) error {
	const query = `DELETE FROM "data" WHERE id = $1`

	if _, err := u.Pool.Exec(ctx, query, id); err != nil {
		return err
	}

	return nil
}
