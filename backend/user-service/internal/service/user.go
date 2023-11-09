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
	userRepo repo.User
	roleRepo repo.Role
}

func NewUser(userRepo repo.User, roleRepo repo.Role) *user {
	return &user{userRepo, roleRepo}
}

func (u *user) Authenticate(email string, password string) (uuid.UUID, []entity.Role, error) {
	user, err := u.userRepo.GetByEmail(context.Background(), email)
	if err != nil {
		return uuid.Nil, nil, ErrUserNotExist
	}

	if err := verifyPassword(*user.Password, password); err != nil {
		return uuid.Nil, nil, ErrPasswordIncorrect
	}

	roles, err := u.roleRepo.GetByUser(context.Background(), user.ID)
	if err != nil {
		return uuid.Nil, nil, err
	}

	return user.ID, roles, nil
}

func (u *user) Create(data entity.User) (*entity.User, error) {
	if err := data.Validate(); err != nil {
		return nil, err
	}

	_, err := u.userRepo.GetByEmail(context.Background(), *data.Email)
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
	user, err := u.userRepo.Create(context.Background(), data)
	if err != nil {
		return nil, err
	}

	user.Roles = make([]entity.Role, 0)
	return user, nil
}

func (u *user) Get(id uuid.UUID) (*entity.User, error) {
	user, err := u.userRepo.GetByID(context.Background(), id)
	if err != nil {
		return nil, ErrUserNotExist
	}

	roles, err := u.roleRepo.GetByUser(context.Background(), user.ID)
	if err != nil {
		return nil, err
	}

	user.Roles = roles
	return user, nil
}

func (u *user) Exist(email string) (bool, error) {
	_, err := u.userRepo.GetByEmail(context.Background(), email)
	if err == nil {
		return true, nil
	}

	if err == repo.ErrNoRows {
		return false, nil
	}

	return false, err
}

func (u *user) Update(data entity.User) (*entity.User, error) {
	user, err := u.Get(data.ID)
	if err != nil {
		return nil, err
	}

	// password
	if data.Password != nil {
		password, err := hashPassword(*data.Password)
		if err != nil {
			return nil, err
		}

		user, err = u.userRepo.UpdatePassword(context.Background(), user.ID, password)
		if err != nil {
			return nil, err
		}
	}

	// contact info
	if data.Phone != nil || data.Telegram != nil {
		if data.Phone == nil {
			data.Phone = user.Phone
		}

		if data.Telegram == nil {
			data.Telegram = user.Telegram
		}

		user, err = u.userRepo.UpdateContact(context.Background(), data)
		if err != nil {
			return nil, err
		}
	}

	return user, nil
}

func (u *user) Delete(id uuid.UUID) error {
	user, err := u.Get(id)
	if err != nil {
		return err
	}

	return u.userRepo.DeleteByID(context.Background(), user.ID)
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
