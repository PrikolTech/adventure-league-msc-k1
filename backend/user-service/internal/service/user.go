package service

import (
	"context"
	"net/mail"
	"strings"
	"user-service/internal/entity"
	"user-service/internal/repo"

	"github.com/gofrs/uuid/v5"
	"golang.org/x/crypto/bcrypt"
)

type user struct {
	repo repo.User
}

func NewUser(repo repo.User) *user {
	return &user{repo}
}

func (u *user) Authenticate(email string, password string) (uuid.UUID, error) {
	user, err := u.repo.GetByEmail(context.Background(), email)
	if err != nil {
		return uuid.Nil, ErrUserNotExist
	}

	if err := verifyPassword(*user.Password, password); err != nil {
		return uuid.Nil, ErrPasswordIncorrect
	}

	return user.ID, nil
}

func (u *user) Create(data entity.User) (*entity.User, error) {
	_, err := u.repo.GetByEmail(context.Background(), *data.Email)
	if err == nil {
		return nil, ErrUserExists
	}

	if _, err := mail.ParseAddress(*data.Email); err != nil {
		return nil, ErrEmailInvalid
	}

	password, err := hashPassword(*data.Password)
	if err != nil {
		return nil, err
	}

	data.Password = &password
	return u.repo.Create(context.Background(), data)
}

func (u *user) Get(id uuid.UUID) (*entity.User, error) {
	user, err := u.repo.GetByID(context.Background(), id)
	if err != nil {
		return nil, ErrUserNotExist
	}

	return user, nil
}

func (u *user) Update(data entity.User) error {
	user, err := u.repo.GetByID(context.Background(), data.ID)
	if err != nil {
		return ErrUserNotExist
	}

	if data.Password != nil {
		password, err := hashPassword(*data.Password)
		if err != nil {
			return err
		}

		err = u.repo.UpdatePassword(context.Background(), user.ID, password)
		if err != nil {
			return err
		}
	}

	//
	return nil
}

func (u *user) Delete(id uuid.UUID) error {
	user, err := u.Get(id)
	if err != nil {
		return err
	}

	return u.repo.DeleteByID(context.Background(), user.ID)
}

const (
	lowerChars   = `abcdefghijklmnopqrstuvwxyz`
	upperChars   = `ABCDEFGHIJKLMNOPQRSTUVWXYZ`
	digitChars   = `0123456789`
	specialChars = ` !"#$%&'()*+,-./:;<=>?@[\]^_{|}~`
)

func validatePassword(password string) error {
	if length := len(password); length < 8 || length > 72 {
		return ErrPasswordInvalid
	}

	var lower, upper, digit, special bool

	for _, r := range password {
		switch {
		case strings.ContainsRune(lowerChars, r):
			lower = true
		case strings.ContainsRune(upperChars, r):
			upper = true
		case strings.ContainsRune(digitChars, r):
			digit = true
		case strings.ContainsRune(specialChars, r):
			special = true
		}

		if lower && upper && digit && special {
			return nil
		}
	}

	return ErrPasswordInvalid
}

func hashPassword(password string) (string, error) {
	if err := validatePassword(password); err != nil {
		return "", err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func verifyPassword(hashedPassword string, password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		return ErrPasswordIncorrect
	}

	return nil
}
