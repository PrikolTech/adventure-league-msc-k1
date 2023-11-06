package service

import (
	"context"
	"errors"
	"user-service/internal/entity"
	"user-service/internal/repo"

	"github.com/gofrs/uuid/v5"
)

type role struct {
	repo repo.Role
}

func NewRole(repo repo.Role) *role {
	return &role{repo}
}

func (r *role) Create(data entity.Role) (*entity.Role, error) {
	if err := data.Validate(); err != nil {
		return nil, err
	}

	_, err := r.repo.GetByTitle(context.Background(), *data.Title)
	if err == nil {
		return nil, ErrRoleExists
	}

	return r.repo.Create(context.Background(), data)
}

func (r *role) Append(userID uuid.UUID, roleID uuid.UUID) error {
	err := r.repo.CreateUserRole(context.Background(), userID, roleID)
	if err != nil {
		return errors.New("invalid user/role id")
	}
	return nil
}

func (r *role) Remove(userID uuid.UUID, roleID uuid.UUID) error {
	return r.repo.DeleteUserRole(context.Background(), userID, roleID)
}

func (r *role) RemoveAll(userID uuid.UUID) error {
	return r.repo.DeleteUserRoles(context.Background(), userID)
}

func (r *role) GetByUser(userID uuid.UUID) ([]entity.Role, error) {
	return r.repo.GetByUser(context.Background(), userID)
}

func (r *role) List() ([]entity.Role, error) {
	return r.repo.List(context.Background())
}

func (r *role) Delete(id uuid.UUID) error {
	role, err := r.repo.GetByID(context.Background(), id)
	if err != nil {
		return err
	}

	return r.repo.DeleteByID(context.Background(), role.ID)
}
