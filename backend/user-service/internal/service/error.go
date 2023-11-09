package service

import "errors"

var (
	ErrUserExists        = errors.New("user already exists")
	ErrUserNotExist      = errors.New("user does not exist")
	ErrEmailInvalid      = errors.New("email is invalid")
	ErrPasswordInvalid   = errors.New("password is invalid")
	ErrPasswordIncorrect = errors.New("password is incorrect")
	ErrRoleExists        = errors.New("role already exists")
	ErrRoleNotExist      = errors.New("role does not exist")
)
