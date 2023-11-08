package service

import (
	"context"
	"form-service/internal/entity"
	"form-service/internal/net"
	"form-service/internal/repo"

	"github.com/gofrs/uuid/v5"
)

type NetServices struct {
	User   net.User
	Role   net.Role
	Course net.Course
}

type registration struct {
	repo repo.Registration
	net  NetServices
}

func NewRegistration(repo repo.Registration, net NetServices) *registration {
	return &registration{repo, net}
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
		if err = data.Status.Validate(); err != nil {
			return nil, err
		}

		switch *data.Status {
		case entity.Acccepted:
			// create user
			user, err := r.net.User.Create(entity.User{
				Email:      *registration.Email,
				Password:   "",
				FirstName:  *registration.Initiator.FirstName,
				LastName:   *registration.Initiator.LastName,
				Patronymic: *registration.Initiator.Patronymic,
				Birthdate:  *registration.Birthdate,
				Phone:      *registration.Phone,
				Telegram:   *registration.Telegram,
			})
			if err != nil {
				return nil, err
			}

			// update user id
			registration, err = r.repo.UpdateUser(context.Background(), registration.ID, user.ID)
			if err != nil {
				return nil, err
			}

			// append role
			err = r.net.Role.Append(user.ID, uuid.Nil)
			if err != nil {
				return nil, err
			}
		case entity.Approved:
			err := r.net.Course.Append(*registration.UserID, *registration.CourseID)
			if err != nil {
				return nil, err
			}
		}

		registration, err = r.repo.UpdateStatus(context.Background(), registration.ID, *data.Status)
		if err != nil {
			return nil, err
		}
	}

	return registration, nil
}
