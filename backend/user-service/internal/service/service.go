package service

import (
	"context"
	"user-service/internal/entity"

	uuid "github.com/jackc/pgx-gofrs-uuid"
)

type User interface {
	Authenticate(ctx context.Context, data entity.User) (uuid.UUID, error)
	Create(ctx context.Context, data entity.User) (*entity.User, error)
	Get(ctx context.Context, id uuid.UUID) (*entity.User, error)
	Update(ctx context.Context, data entity.User) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type Role interface {
	GetByUser(ctx context.Context, userID uuid.UUID) ([]entity.Role, error)
	List(ctx context.Context) ([]entity.Role, error)
}
