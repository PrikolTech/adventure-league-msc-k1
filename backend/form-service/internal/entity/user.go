package entity

import (
	"time"

	"github.com/gofrs/uuid/v5"
)

type User struct {
	ID       uuid.UUID `json:"id"`
	Email    string    `json:"email"`
	Password string    `json:"password"`

	FirstName  string    `json:"first_name"`
	LastName   string    `json:"last_name"`
	Patronymic string    `json:"patronymic"`
	Birthdate  time.Time `json:"birthdate"`
	Phone      string    `json:"phone"`
	Telegram   string    `json:"telegram"`

	Roles []Role `json:"roles"`
}
