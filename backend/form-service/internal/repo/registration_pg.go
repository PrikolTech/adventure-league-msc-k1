package repo

import (
	"form-service/pkg/postgres"
)

type RegistrationPG struct {
	*postgres.Postgres
}

func NewRegistrationPG(pg *postgres.Postgres) *RegistrationPG {
	return &RegistrationPG{pg}
}
