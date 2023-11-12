package repo

import (
	"context"
	"form-service/internal/entity"

	"github.com/gofrs/uuid/v5"
)

type Registration interface {
	Create(ctx context.Context, data entity.Registration) (*entity.Registration, error)
	GetByID(ctx context.Context, id uuid.UUID) (*entity.Registration, error)
	GetByUser(ctx context.Context, userID uuid.UUID) ([]entity.Registration, error)
	GetByCourse(ctx context.Context, courseID uuid.UUID) ([]entity.Registration, error)
	List(ctx context.Context) ([]entity.Registration, error)
	UpdateStatus(ctx context.Context, id uuid.UUID, status entity.Status) (*entity.Registration, error)
	UpdateUser(ctx context.Context, id uuid.UUID, userID uuid.UUID) (*entity.Registration, error)
}
