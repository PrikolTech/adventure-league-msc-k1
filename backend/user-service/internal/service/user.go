package service

import (
	"context"
	"user-service/internal/entity"
	"user-service/internal/repo"

	uuid "github.com/jackc/pgx-gofrs-uuid"
)

type user struct {
	repo repo.User
}

func NewUser(repo repo.User) *user {
	return &user{repo}
}

func (u *user) Authenticate(ctx context.Context, data entity.User) (uuid.UUID, error)
func (u *user) Create(ctx context.Context, data entity.User) (*entity.User, error)
func (u *user) Get(ctx context.Context, id uuid.UUID) (*entity.User, error)
func (u *user) Update(ctx context.Context, data entity.User) error
func (u *user) Delete(ctx context.Context, id uuid.UUID) error
