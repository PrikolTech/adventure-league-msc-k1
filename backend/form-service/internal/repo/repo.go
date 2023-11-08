package repo

import (
	"context"
	"form-service/internal/entity"

	"github.com/gofrs/uuid/v5"
)

type Registration interface {
	Create(ctx context.Context, data entity.Registration) (*entity.Registration, error)
	UpdateStatus(ctx context.Context, id uuid.UUID, status entity.Status) (*entity.Registration, error)
}
