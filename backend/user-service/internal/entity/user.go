package entity

import (
	"github.com/gofrs/uuid/v5"
)

type User struct {
	ID       uuid.UUID `json:"id"`
	Email    *string   `json:"email"`
	Password *string   `json:"password"`

	FirstName  *string `json:"first_name"`
	LastName   *string `json:"last_name"`
	Patronymic *string `json:"patronymic"`
	Phone      *string `json:"phone"`
	Telegram   *string `json:"telegram"`
	Roles      []Role  `json:"roles,omitempty"`
}

func (e *User) Validate() error {
	if e.Email == nil {
		return &RequiredError{"email"}
	}

	if e.Password == nil {
		return &RequiredError{"password"}
	}

	if e.FirstName == nil {
		return &RequiredError{"first_name"}
	}

	if e.LastName == nil {
		return &RequiredError{"last_name"}
	}

	return nil
}
