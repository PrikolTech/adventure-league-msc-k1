package repo

import (
	"context"
	"form-service/internal/entity"
	"form-service/pkg/postgres"
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

	var user entity.Registration
	err := r.Pool.
		QueryRow(ctx, query, data.Email, data.Initiator.FirstName, data.Initiator.LastName, data.Initiator.Patronymic, data.Birthdate, data.Supervisor.FirstName, data.Supervisor.LastName, data.Supervisor.Patronymic, data.Department, data.Post, data.History, data.Achievements, data.Motivation, data.Phone, data.Telegram, data.Status).
		Scan(&user.ID, &user.Email, &user.Initiator.FirstName, &user.Initiator.LastName, &user.Initiator.Patronymic, &user.Birthdate, &user.Supervisor.FirstName, &user.Supervisor.LastName, &user.Supervisor.Patronymic, &data.Department, &data.Post, &data.History, &data.Achievements, &data.Motivation, &user.Phone, &user.Telegram, &data.Status)

	return &user, err
}
