package entity

import (
	"time"

	"github.com/gofrs/uuid/v5"
)

type FullName struct {
	FirstName  *string `json:"first_name"`
	LastName   *string `json:"last_name"`
	Patronymic *string `json:"patronymic"`
}

func (f *FullName) Validate() error {
	if f.FirstName == nil {
		return &RequiredError{"first_name"}
	}

	if f.LastName == nil {
		return &RequiredError{"last_name"}
	}
	return nil
}

type Registration struct {
	ID    uuid.UUID `json:"id"`
	Email *string   `json:"email"`

	Initiator    FullName   `json:"initiator"`
	Birthdate    *time.Time `json:"birthdate"`
	Supervisor   FullName   `json:"supervisor"`
	Department   *string    `json:"department"`
	Post         *string    `json:"post"`
	History      *string    `json:"history"`
	Achievements *string    `json:"achievements"`
	Motivation   *string    `json:"motivation"`

	Telegram *string `json:"telegram"`
	Phone    *string `json:"phone"`
}

func (f *Registration) Validate() error {
	if err := f.Initiator.Validate(); err != nil {
		return err
	}

	if err := f.Supervisor.Validate(); err != nil {
		return err
	}

	if f.Email == nil {
		return &RequiredError{"email"}
	}

	if f.Birthdate == nil {
		return &RequiredError{"birthdate"}
	}

	if f.Department == nil {
		return &RequiredError{"department"}
	}

	if f.Post == nil {
		return &RequiredError{"post"}
	}

	if f.History == nil {
		return &RequiredError{"history"}
	}

	if f.Achievements == nil {
		return &RequiredError{"achievements"}
	}

	if f.Motivation == nil {
		return &RequiredError{"motivation"}
	}

	return nil
}
