package entity

import "github.com/gofrs/uuid/v5"

type Role struct {
	ID          uuid.UUID `json:"id"`
	Title       RoleTitle `json:"title"`
	Description string    `json:"description"`
}

type RoleTitle string

const (
	Admin    RoleTitle = "admin"
	Employee RoleTitle = "employee"
	Student  RoleTitle = "student"
	Enrollee RoleTitle = "enrollee"
)
