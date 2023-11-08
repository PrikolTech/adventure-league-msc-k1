package repo

import (
	"context"
	"form-service/internal/entity"
	"form-service/pkg/postgres"

	"github.com/gofrs/uuid/v5"
)

type RegistrationPG struct {
	*postgres.Postgres
}

func NewRegistrationPG(pg *postgres.Postgres) *RegistrationPG {
	return &RegistrationPG{pg}
}

func (r *RegistrationPG) Create(ctx context.Context, data entity.Registration) (*entity.Registration, error) {
	const query = `INSERT INTO registration
		(email, initiator_first_name, initiator_last_name, initiator_patronymic, birthdate, supervisor_first_name, supervisor_last_name, supervisor_patronymic, department, post, history, achievements, motivation, phone, telegram, status) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16) RETURNING *`

	var registration entity.Registration
	err := r.Pool.
		QueryRow(ctx, query, data.Email, data.Initiator.FirstName, data.Initiator.LastName, data.Initiator.Patronymic, data.Birthdate, data.Supervisor.FirstName, data.Supervisor.LastName, data.Supervisor.Patronymic, data.Department, data.Post, data.History, data.Achievements, data.Motivation, data.Phone, data.Telegram, data.Status).
		Scan(&registration.ID, &registration.Email, &registration.Initiator.FirstName, &registration.Initiator.LastName, &registration.Initiator.Patronymic, &registration.Birthdate, &registration.Supervisor.FirstName, &registration.Supervisor.LastName, &registration.Supervisor.Patronymic, &registration.Department, &registration.Post, &registration.History, &registration.Achievements, &registration.Motivation, &registration.Phone, &registration.Telegram, &registration.Status)

	return &registration, err
}

func (r *RegistrationPG) UpdateStatus(ctx context.Context, id uuid.UUID, status entity.Status) (*entity.Registration, error) {
	const query = `UPDATE registration SET status = $2 WHERE id = $1 RETURNING *`

	var registration entity.Registration
	err := r.Pool.
		QueryRow(ctx, query, id, status).
		Scan(&registration.ID, &registration.Email, &registration.Initiator.FirstName, &registration.Initiator.LastName, &registration.Initiator.Patronymic, &registration.Birthdate, &registration.Supervisor.FirstName, &registration.Supervisor.LastName, &registration.Supervisor.Patronymic, &registration.Department, &registration.Post, &registration.History, &registration.Achievements, &registration.Motivation, &registration.Phone, &registration.Telegram, &registration.Status)

	return &registration, err
}
