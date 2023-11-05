package service

import (
	"user-service/internal/entity"

	"github.com/gofrs/uuid/v5"
)

type User interface {
	Authenticate(email string, password string) (uuid.UUID, []entity.Role, error)
	Create(data entity.User) (*entity.User, error)
	Get(id uuid.UUID) (*entity.User, error)
	Update(data entity.User) (*entity.User, error)
	Delete(id uuid.UUID) error
}

type Role interface {
	Create(data entity.Role) (*entity.Role, error)
	Append(userID uuid.UUID, roleID uuid.UUID) error
	Remove(userID uuid.UUID, roleID uuid.UUID) error
	RemoveAll(userID uuid.UUID) error
	GetByUser(userID uuid.UUID) ([]entity.Role, error)
	List() ([]entity.Role, error)
	Delete(id uuid.UUID) error
}
