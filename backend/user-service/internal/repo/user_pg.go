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

func (u *UserPG) Create(ctx context.Context, data entity.User) (*entity.User, error)
func (u *UserPG) GetByID(ctx context.Context, id uuid.UUID) (*entity.User, error)
func (u *UserPG) GetByName(ctx context.Context, name string) (*entity.User, error)
func (u *UserPG) UpdatePassword(ctx context.Context, id uuid.UUID, password []byte) error
func (u *UserPG) DeleteByID(ctx context.Context, id uuid.UUID) error
