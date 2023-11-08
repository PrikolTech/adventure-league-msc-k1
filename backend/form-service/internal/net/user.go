package net

import (
	"form-service/internal/entity"
)

type user struct {
}

func (u *user) Create(data entity.User) (*entity.User, error)
