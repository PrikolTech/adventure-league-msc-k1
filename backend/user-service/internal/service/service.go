package service

import (
	"user-service/internal/entity"

	"github.com/gofrs/uuid/v5"
)

type User interface {
	Authenticate(email string, password string) (uuid.UUID, []entity.Role, error)
	Create(data entity.User) (*entity.User, error)
	Get(id uuid.UUID) (*entity.User, error)
	Exist(email string) (bool, error)
	Update(data entity.User) (*entity.User, error)
	Delete(id uuid.UUID) error
}

type Role interface {
	Create(data entity.Role) (*entity.Role, error)
	AppendByID(userID uuid.UUID, roleID uuid.UUID) error
	AppendByTitle(userID uuid.UUID, title string) error
	RemoveByID(userID uuid.UUID, roleID uuid.UUID) error
	RemoveByTitle(userID uuid.UUID, title string) error
	RemoveAll(userID uuid.UUID) error
	GetByUser(userID uuid.UUID) ([]entity.Role, error)
	List() ([]entity.Role, error)
	Delete(id uuid.UUID) error
}
