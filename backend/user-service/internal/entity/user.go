package entity

import (
	"encoding/json"
	"time"

	"github.com/gofrs/uuid/v5"
)

type User struct {
	ID       uuid.UUID `json:"id"`
	Email    *string   `json:"email"`
	Password *string   `json:"password,omitempty"`

	FirstName  *string    `json:"first_name"`
	LastName   *string    `json:"last_name"`
	Patronymic *string    `json:"patronymic"`
	Birthdate  *time.Time `json:"birthdate"`
	Phone      *string    `json:"phone"`
	Telegram   *string    `json:"telegram"`
	Roles      []Role     `json:"roles"`
}

func (e User) MarshalJSON() ([]byte, error) {
	type user User
	u := user(e)
	u.Password = nil
	if u.Roles == nil {
		u.Roles = make([]Role, 0)
	}
	return json.Marshal(u)
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

	if e.Birthdate == nil {
		return &RequiredError{"birthdate"}
	}

	return nil
}
