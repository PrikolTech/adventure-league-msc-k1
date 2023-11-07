package repo

import (
	"context"
	"form-service/internal/entity"
)

type Registration interface {
	Create(ctx context.Context, data entity.Registration) (*entity.Registration, error)
}
