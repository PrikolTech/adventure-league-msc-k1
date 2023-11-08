package service

import (
	"bytes"
	"context"
	cryptoRand "crypto/rand"
	"form-service/internal/entity"
	"form-service/internal/net"
	"form-service/internal/repo"
	"math"
	"math/big"
	mathRand "math/rand"

	"github.com/gofrs/uuid/v5"
)

type NetServices struct {
	User   net.User
	Role   net.Role
	Course net.Course
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

func (r *registration) Update(data entity.Registration) (*entity.Registration, error) {
	registration, err := r.Get(data.ID)
	if err != nil {
		return nil, err
	}

	if data.Status != nil {
		if err = data.Status.Validate(); err != nil {
			return nil, err
		}

		switch *data.Status {
		case entity.Acccepted:
			// create user
			password, err := generatePassword()
			if err != nil {
				return nil, err
			}

			userData := entity.User{
				Email:     *registration.Email,
				Password:  string(password),
				FirstName: *registration.Initiator.FirstName,
				LastName:  *registration.Initiator.LastName,
				Birthdate: *registration.Birthdate,
			}

			if registration.Initiator.Patronymic != nil {
				userData.Patronymic = *registration.Initiator.Patronymic
			}

			if registration.Phone != nil {
				userData.Phone = *registration.Phone
			}

			if registration.Telegram != nil {
				userData.Telegram = *registration.Telegram
			}

			user, err := r.net.User.Create(userData)
			if err != nil {
				return nil, err
			}

			// update user id
			registration, err = r.repo.UpdateUser(context.Background(), registration.ID, user.ID)
			if err != nil {
				return nil, err
			}

			// append role
			err = r.net.Role.Append(user.ID, uuid.FromStringOrNil("11e8cd39-e2c5-4193-a61f-de11a3af67aa"))
			if err != nil {
				return nil, err
			}
		case entity.Approved:
			err := r.net.Course.Append(*registration.UserID, *registration.CourseID)
			if err != nil {
				return nil, err
			}
		}

		registration, err = r.repo.UpdateStatus(context.Background(), registration.ID, *data.Status)
		if err != nil {
			return nil, err
		}
	}

	return registration, nil
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
