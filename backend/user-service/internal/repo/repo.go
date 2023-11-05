package repo

import (
	"context"
	"user-service/internal/entity"

	"github.com/gofrs/uuid/v5"
)

type User interface {
	Create(ctx context.Context, data entity.User) (*entity.User, error)
	GetByID(ctx context.Context, id uuid.UUID) (*entity.User, error)
	GetByEmail(ctx context.Context, email string) (*entity.User, error)
	UpdatePassword(ctx context.Context, id uuid.UUID, password string) error
	DeleteByID(ctx context.Context, id uuid.UUID) error
}

type Role interface {
	Create(ctx context.Context, data entity.Role) (*entity.Role, error)
	CreateUserRole(ctx context.Context, userID uuid.UUID, roleID uuid.UUID) error
	DeleteUserRole(ctx context.Context, userID uuid.UUID, roleID uuid.UUID) error
	DeleteUserRoles(ctx context.Context, userID uuid.UUID) error
	GetByID(ctx context.Context, id uuid.UUID) (*entity.Role, error)
	GetByTitle(ctx context.Context, title string) (*entity.Role, error)
	GetByUser(ctx context.Context, userID uuid.UUID) ([]entity.Role, error)
	List(ctx context.Context) ([]entity.Role, error)
	DeleteByID(ctx context.Context, id uuid.UUID) error
}
