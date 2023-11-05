package net

type User interface {
	Authenticate(email string, password string) (*AuthenticateResponse, error)
}
