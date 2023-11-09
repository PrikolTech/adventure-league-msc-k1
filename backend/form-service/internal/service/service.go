package service

import (
	"form-service/internal/entity"

	"github.com/gofrs/uuid/v5"
)

type Registration interface {
	Create(data entity.Registration) (*entity.Registration, error)
	Get(id uuid.UUID) (*entity.Registration, error)
	GetByUser(userID uuid.UUID) ([]entity.Registration, error)
	GetByCourse(courseID uuid.UUID) ([]entity.Registration, error)
	List() ([]entity.Registration, error)
	Update(data entity.Registration) (*entity.Registration, error)
}
