package jwt

import (
	"errors"
	"os"

	"github.com/golang-jwt/jwt"
)

type Claims struct {
	jwt.StandardClaims
	Roles []string `json:"roles"`
}

type Parser interface {
	Parse(tokenString string) (*Claims, error)
}

var ErrClaimsInvalid = errors.New("claims is invalid")

type parser struct {
	key any
}

func NewParser(key any) *parser {
	return &parser{key}
}

func NewParserFromFile(path string) (*parser, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	key, err := jwt.ParseECPublicKeyFromPEM(data)
	if err != nil {
		return nil, err
	}

	return NewParser(key), nil
}

func (p *parser) Parse(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(t *jwt.Token) (any, error) {
		return p.key, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok {
		return claims, nil
	}

	return nil, ErrClaimsInvalid
}
