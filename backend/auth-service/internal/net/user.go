package net

import (
	"auth-service/internal/entity"
	"bytes"
	"encoding/json"
	"net/http"
)

type user struct {
	client *http.Client
	url    string
}

func NewUser(url string) *user {
	client := &http.Client{}
	return &user{client, url}
}

type AuthenticateRequestBody struct {
	Email    string
	Password string
}

type AuthenticateResponseBody struct {
	ID    string
	Roles []entity.Role
}

func (u *user) Authenticate(email string, password string) (string, []entity.Role, error) {
	reqData := new(bytes.Buffer)
	e := json.NewEncoder(reqData)
	e.Encode(AuthenticateRequestBody{email, password})

	req, err := http.NewRequest(http.MethodPost, u.url+"/user/authenticate", reqData)
	if err != nil {
		return "", nil, ErrInternalNetwork
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := u.client.Do(req)
	if err != nil {
		return "", nil, ErrInternalNetwork
	}

	defer resp.Body.Close()

	respData := new(AuthenticateResponseBody)
	d := json.NewDecoder(resp.Body)
	err = d.Decode(respData)

	return respData.ID, respData.Roles, err
}
