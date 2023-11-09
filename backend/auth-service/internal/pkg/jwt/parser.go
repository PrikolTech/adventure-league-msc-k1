package jwt

import (
	"errors"
	"os"

	"github.com/golang-jwt/jwt"
)

type Parser interface {
	Parse(tokenString string) (*jwt.StandardClaims, error)
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

func (p *parser) Parse(tokenString string) (*jwt.StandardClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(t *jwt.Token) (any, error) {
		return p.key, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*jwt.StandardClaims); ok {
		return claims, nil
	}

	return nil, ErrClaimsInvalid
}
