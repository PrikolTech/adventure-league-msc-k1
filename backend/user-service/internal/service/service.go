package service

import (
	"user-service/internal/entity"

	"github.com/gofrs/uuid/v5"
)

type User interface {
	Authenticate(email string, password string) (uuid.UUID, []entity.Role, error)
	Create(data entity.User) (*entity.User, error)
	Get(id uuid.UUID) (*entity.User, error)
	Update(data entity.User) error
	Delete(id uuid.UUID) error
}

type Role interface {
	GetByUser(userID uuid.UUID) ([]entity.Role, error)
	List() ([]entity.Role, error)
}
