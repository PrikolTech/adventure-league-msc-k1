package repo

import (
	"context"
	"user-service/internal/entity"
	"user-service/pkg/postgres"

	"github.com/gofrs/uuid/v5"
)

type UserPG struct {
	*postgres.Postgres
}

func NewUserPG(pg *postgres.Postgres) *UserPG {
	return &UserPG{pg}
}

func (u *UserPG) Create(ctx context.Context, data entity.User) (*entity.User, error) {
	const query = `INSERT INTO "data"
			(email, password, first_name, last_name, patronymic, phone, telegram) 
			VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING *`

	var user entity.User
	err := u.Pool.
		QueryRow(ctx, query, data.Email, data.Password, data.FirstName, data.LastName, data.Patronymic, data.Phone, data.Telegram).
		Scan(&user.ID, &user.Email, &user.Password, &user.FirstName, &user.LastName, &user.Patronymic, &user.Phone, &user.Telegram)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *UserPG) GetByID(ctx context.Context, id uuid.UUID) (*entity.User, error) {
	const query = `SELECT * FROM "data" WHERE id = $1`

	var user entity.User
	err := u.Pool.
		QueryRow(ctx, query, id).
		Scan(&user.ID, &user.Email, &user.Password, &user.FirstName, &user.LastName, &user.Patronymic, &user.Phone, &user.Telegram)

	return &user, err
}

func (u *UserPG) GetByName(ctx context.Context, name string) (*entity.User, error) {
	const query = `SELECT * FROM "data" WHERE name = $1`

	var user entity.User
	err := u.Pool.
		QueryRow(ctx, query, name).
		Scan(&user.ID, &user.Email, &user.Password, &user.FirstName, &user.LastName, &user.Patronymic, &user.Phone, &user.Telegram)

	return &user, err
}

func (u *UserPG) UpdatePassword(ctx context.Context, id uuid.UUID, password []byte) error {
	const query = `UPDATE "data" SET password = $2 WHERE id = $1`

	if _, err := u.Pool.Exec(ctx, query, id, password); err != nil {
		return err
	}

	return nil
}

func (u *UserPG) DeleteByID(ctx context.Context, id uuid.UUID) error {
	const query = `DELETE FROM "data" WHERE id = $1`

	if _, err := u.Pool.Exec(ctx, query, id); err != nil {
		return err
	}

	return nil
}
