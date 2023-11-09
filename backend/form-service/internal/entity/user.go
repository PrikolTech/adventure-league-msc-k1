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

type Role struct {
	ID          uuid.UUID `json:"id"`
	Title       RoleTitle `json:"title"`
	Description string    `json:"description"`
}

type RoleTitle string

const (
	Enrollee RoleTitle = "enrollee"
	Student  RoleTitle = "student"
)
