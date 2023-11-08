package repo

import (
	"context"
	"form-service/internal/entity"
	"form-service/pkg/postgres"

	"github.com/gofrs/uuid/v5"
	"github.com/jackc/pgx/v5"
)

type RegistrationPG struct {
	*postgres.Postgres
}

func NewRegistrationPG(pg *postgres.Postgres) *RegistrationPG {
	return &RegistrationPG{pg}
}

func (r *RegistrationPG) collectRow(row pgx.Row) (*entity.Registration, error) {
	var data entity.Registration
	err := row.Scan(&data.ID, &data.Email, &data.Initiator.FirstName, &data.Initiator.LastName, &data.Initiator.Patronymic, &data.Birthdate, &data.Supervisor.FirstName, &data.Supervisor.LastName, &data.Supervisor.Patronymic, &data.Department, &data.Post, &data.History, &data.Achievements, &data.Motivation, &data.Phone, &data.Telegram, &data.Status, &data.UserID, &data.CourseID)
	return &data, err
}

func (r *RegistrationPG) collectRows(rows pgx.Rows) ([]entity.Registration, error) {
	return pgx.CollectRows(rows, func(row pgx.CollectableRow) (entity.Registration, error) {
		data, err := r.collectRow(row)
		return *data, err
	})
}

func (r *RegistrationPG) Create(ctx context.Context, data entity.Registration) (*entity.Registration, error) {
	const query = `INSERT INTO registration
		(email, initiator_first_name, initiator_last_name, initiator_patronymic, birthdate, supervisor_first_name, supervisor_last_name, supervisor_patronymic, department, post, history, achievements, motivation, phone, telegram, user_id, course_id)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17) RETURNING *`

	row := r.Pool.
		QueryRow(ctx, query, data.Email, data.Initiator.FirstName, data.Initiator.LastName, data.Initiator.Patronymic, data.Birthdate, data.Supervisor.FirstName, data.Supervisor.LastName, data.Supervisor.Patronymic, data.Department, data.Post, data.History, data.Achievements, data.Motivation, data.Phone, data.Telegram, data.UserID, data.CourseID)

	return r.collectRow(row)
}

func (r *RegistrationPG) GetByID(ctx context.Context, id uuid.UUID) (*entity.Registration, error) {
	const query = `SELECT * FROM registration WHERE id = $1`

	row := r.Pool.QueryRow(ctx, query, id)
	return r.collectRow(row)
}

func (r *RegistrationPG) GetByUser(ctx context.Context, userID uuid.UUID) ([]entity.Registration, error) {
	const query = `SELECT * FROM registration WHERE user_id = $1`

	rows, err := r.Pool.Query(ctx, query, userID)
	if err != nil {
		return nil, err
	}

	return r.collectRows(rows)
}

func (r *RegistrationPG) GetByCourse(ctx context.Context, courseID uuid.UUID) ([]entity.Registration, error) {
	const query = `SELECT * FROM registration WHERE course_id = $1`

	rows, err := r.Pool.Query(ctx, query, courseID)
	if err != nil {
		return nil, err
	}

	return r.collectRows(rows)
}

func (r *RegistrationPG) List(ctx context.Context) ([]entity.Registration, error) {
	const query = `SELECT * FROM registration`

	rows, err := r.Pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	return r.collectRows(rows)
}

func (r *RegistrationPG) UpdateStatus(ctx context.Context, id uuid.UUID, status entity.Status) (*entity.Registration, error) {
	const query = `UPDATE registration SET status = $2 WHERE id = $1 RETURNING *`

	row := r.Pool.QueryRow(ctx, query, id, status)
	return r.collectRow(row)
}

func (r *RegistrationPG) UpdateUser(ctx context.Context, id uuid.UUID, userID uuid.UUID) (*entity.Registration, error) {
	const query = `UPDATE registration SET user_id = $2 WHERE id = $1 RETURNING *`

	row := r.Pool.QueryRow(ctx, query, id, userID)
	return r.collectRow(row)
}
