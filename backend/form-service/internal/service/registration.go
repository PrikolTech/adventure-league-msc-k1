package service

import (
	"bytes"
	"context"
	cryptoRand "crypto/rand"
	"fmt"
	"form-service/internal/entity"
	"form-service/internal/net"
	"form-service/internal/repo"
	"math"
	"math/big"
	mathRand "math/rand"
	"slices"

	"github.com/gofrs/uuid/v5"
)

type NetServices struct {
	User   net.User
	Role   net.Role
	Course net.Course
	SMTP   net.SMTP
}

type registration struct {
	repo repo.Registration
	net  NetServices
}

func NewRegistration(repo repo.Registration, net NetServices) *registration {
	return &registration{repo, net}
}

func (r *registration) Create(data entity.Registration) (*entity.Registration, error) {
	if err := data.Validate(); err != nil {
		return nil, err
	}

	if data.UserID == nil {
		exists, err := r.net.User.Exist(*data.Email)
		if err != nil {
			return nil, err
		}

		if exists {
			return nil, ErrEmailOccupied
		}
	}

	return r.repo.Create(context.Background(), data)
}

func (r *registration) Get(id uuid.UUID) (*entity.Registration, error) {
	registration, err := r.repo.GetByID(context.Background(), id)
	if err != nil {
		return nil, ErrFormNotExist
	}

	return registration, nil
}

func (r *registration) GetByUser(userID uuid.UUID) ([]entity.Registration, error) {
	return r.repo.GetByUser(context.Background(), userID)
}

func (r *registration) GetByCourse(courseID uuid.UUID) ([]entity.Registration, error) {
	return r.repo.GetByCourse(context.Background(), courseID)
}

func (r *registration) List() ([]entity.Registration, error) {
	return r.repo.List(context.Background())
}

func (r *registration) Update(data entity.Registration, token string) (*entity.Registration, error) {
	registration, err := r.Get(data.ID)
	if err != nil {
		return nil, err
	}

	if data.Status != nil {
		if err = data.Status.Validate(); err != nil {
			return nil, err
		}

		status := *data.Status
		switch status {
		case entity.Acccepted, entity.Approved:
			fmt.Println("before getOrCreate")
			user, err := r.getOrCreateUser(registration, token)
			if err != nil {
				return nil, err
			}

			fmt.Println("before appendRole")
			err = r.appendRole(user, entity.Enrollee, token)
			if err != nil {
				return nil, err
			}

			if status == entity.Approved {
				fmt.Println("before CourseAppend")
				err = r.net.Course.Append(user.ID, *registration.CourseID, token)
				if err != nil {
					return nil, err
				}

				fmt.Println("before appendRole [x2]")
				err = r.appendRole(user, entity.Student, token)
				if err != nil {
					return nil, err
				}
			}
		}

		fmt.Println("before UpdateStatus")
		registration, err = r.repo.UpdateStatus(context.Background(), registration.ID, status)
		if err != nil {
			return nil, err
		}
	}

	return registration, nil
}

func (r *registration) createUser(data *entity.Registration) (*entity.User, error) {
	password, err := generatePassword()
	if err != nil {
		return nil, err
	}

	userData := entity.User{
		Email:     *data.Email,
		Password:  string(password),
		FirstName: *data.Initiator.FirstName,
		LastName:  *data.Initiator.LastName,
		Birthdate: *data.Birthdate,
	}

	if data.Initiator.Patronymic != nil {
		userData.Patronymic = *data.Initiator.Patronymic
	}

	if data.Phone != nil {
		userData.Phone = *data.Phone
	}

	if data.Telegram != nil {
		userData.Telegram = *data.Telegram
	}

	err = r.net.SMTP.SendPassword(userData.Password, userData.Email)
	if err != nil {
		return nil, err
	}

	return r.net.User.Create(userData)
}

func (r *registration) getOrCreateUser(data *entity.Registration, token string) (*entity.User, error) {
	if data.UserID == nil {
		user, err := r.createUser(data)
		if err != nil {
			return nil, err
		}

		_, err = r.repo.UpdateUser(context.Background(), data.ID, user.ID)
		if err != nil {
			return nil, err
		}

		return user, nil
	}
	return r.net.User.Get(*data.UserID, token)
}

func (r *registration) appendRole(user *entity.User, title entity.RoleTitle, token string) error {
	hasRole := slices.ContainsFunc(user.Roles, func(role entity.Role) bool {
		return role.Title == title
	})

	if hasRole {
		return nil
	}

	return r.net.Role.Append(user.ID, string(title), token)
}

const (
	lowerChars   = `abcdefghijklmnopqrstuvwxyz`
	upperChars   = `ABCDEFGHIJKLMNOPQRSTUVWXYZ`
	digitChars   = `0123456789`
	specialChars = `!"#$%&'()*+,-./:;<=>?@[\]^_{|}~`
)

var chars = []byte(lowerChars + upperChars + digitChars + specialChars)

func generatePassword() ([]byte, error) {
	buf := new(bytes.Buffer)
	buf.Grow(8)

	// lower
	val, err := cryptoRand.Int(cryptoRand.Reader, big.NewInt(int64(len(lowerChars))))
	if err != nil {
		return nil, ErrPasswordGeneration
	}
	buf.WriteByte(lowerChars[val.Int64()])

	// upper
	val, err = cryptoRand.Int(cryptoRand.Reader, big.NewInt(int64(len(upperChars))))
	if err != nil {
		return nil, ErrPasswordGeneration
	}
	buf.WriteByte(upperChars[val.Int64()])

	// digit
	val, err = cryptoRand.Int(cryptoRand.Reader, big.NewInt(int64(len(digitChars))))
	if err != nil {
		return nil, ErrPasswordGeneration
	}
	buf.WriteByte(digitChars[val.Int64()])

	// special
	val, err = cryptoRand.Int(cryptoRand.Reader, big.NewInt(int64(len(specialChars))))
	if err != nil {
		return nil, ErrPasswordGeneration
	}
	buf.WriteByte(specialChars[val.Int64()])

	// remaining
	maxInt := big.NewInt(int64(len(chars)))
	for i := 0; i < 4; i++ {
		val, err := cryptoRand.Int(cryptoRand.Reader, maxInt)
		if err != nil {
			return nil, ErrPasswordGeneration
		}
		buf.WriteByte(chars[val.Int64()])
	}

	// shuffle
	val, err = cryptoRand.Int(cryptoRand.Reader, big.NewInt(math.MaxInt64))
	if err != nil {
		return nil, ErrPasswordGeneration
	}

	src := mathRand.NewSource(val.Int64())
	r := mathRand.New(src)

	password := buf.Bytes()
	r.Shuffle(len(password), func(i, j int) {
		password[i], password[j] = password[j], password[i]
	})

	return password, nil
}
