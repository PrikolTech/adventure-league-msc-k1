package service

import (
	"form-service/internal/repo"
)

type registration struct {
	repo repo.Registration
}

func NewRegistration(repo repo.Registration) *registration {
	return &registration{repo}
}
