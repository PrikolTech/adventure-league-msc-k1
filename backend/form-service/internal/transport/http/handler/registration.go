package handler

import (
	"form-service/internal/service"
)

type Registration struct {
	service service.Registration
}

func NewRegistration(service service.Registration) *Registration {
	return &Registration{service}
}
