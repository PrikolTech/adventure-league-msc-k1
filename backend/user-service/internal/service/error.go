package service

import "errors"

var (
	ErrUserExists        = errors.New("user already exists")
	ErrUserNotExist      = errors.New("user does not exist")
	ErrUserIDInvalid     = errors.New("user id is invalid")
	ErrEmailInvalid      = errors.New("email is invalid")
	ErrPasswordInvalid   = errors.New("password is invalid")
	ErrPasswordIncorrect = errors.New("password is incorrect")
	ErrRoleExists        = errors.New("role already exists")
)
