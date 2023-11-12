package service

import "errors"

var (
	ErrFormNotExist       = errors.New("form does not exist")
	ErrPasswordGeneration = errors.New("failed to generate password")
	ErrEmailOccupied      = errors.New("email is occupied by another user")
)
