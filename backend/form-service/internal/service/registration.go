package service

import (
	"context"
	"form-service/internal/entity"
	"form-service/internal/repo"

	"github.com/gofrs/uuid/v5"
)

type registration struct {
	repo repo.Registration
}

func NewRegistration(repo repo.Registration) *registration {
	return &registration{repo}
}

func (r *registration) Create(data entity.Registration) (*entity.Registration, error) {
	if err := data.Validate(); err != nil {
		return nil, err
	}

	return r.repo.Create(context.Background(), data)
}

func (r *registration) Get(id uuid.UUID) (*entity.Registration, error) {
	registration, err := r.repo.GetByID(context.Background(), id)
	if err != nil {
		return nil, ErrFormNotExist
	}

	return registration, nil
}

func (r *registration) GetByUser(userID uuid.UUID) ([]entity.Registration, error) {
	return r.repo.GetByUser(context.Background(), userID)
}

func (r *registration) GetByCourse(courseID uuid.UUID) ([]entity.Registration, error) {
	return r.repo.GetByCourse(context.Background(), courseID)
}

func (r *registration) List() ([]entity.Registration, error) {
	return r.repo.List(context.Background())
}

func (r *registration) Update(data entity.Registration) (*entity.Registration, error) {
	registration, err := r.Get(data.ID)
	if err != nil {
		return nil, err
	}

	if data.Status != nil {
		registration, err = r.repo.UpdateStatus(context.Background(), registration.ID, *data.Status)
		if err != nil {
			return nil, err
		}
	}

	return registration, nil
}
