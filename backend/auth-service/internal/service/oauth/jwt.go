package oauth

import (
	"context"
	"crypto/ecdsa"
	"encoding/base64"
	"strings"

	"github.com/go-oauth2/oauth2/v4"
	"github.com/go-oauth2/oauth2/v4/generates"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

func NewJWTAccessGenerate(key []byte) (*JWTAccessGenerate, error) {
	signedKey, err := jwt.ParseECPrivateKeyFromPEM(key)
	if err != nil {
		return nil, err
	}

	return &JWTAccessGenerate{signedKey}, nil
}

type JWTAccessGenerate struct {
	SignedKey *ecdsa.PrivateKey
}

func (a *JWTAccessGenerate) Token(ctx context.Context, data *oauth2.GenerateBasic, isGenRefresh bool) (string, string, error) {
	claims := &generates.JWTAccessClaims{
		jwt.StandardClaims{
			Audience:  data.Client.GetID(),
			Subject:   data.UserID,
			ExpiresAt: data.TokenInfo.GetAccessCreateAt().Add(data.TokenInfo.GetAccessExpiresIn()).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES512, claims)

	access, err := token.SignedString(a.SignedKey)
	if err != nil {
		return "", "", err
	}

	refresh := ""
	if isGenRefresh {
		t := uuid.NewSHA1(uuid.Must(uuid.NewRandom()), []byte(access)).String()
		refresh = base64.URLEncoding.EncodeToString([]byte(t))
		refresh = strings.ToUpper(strings.TrimRight(refresh, "="))
	}

	return access, refresh, nil
}
