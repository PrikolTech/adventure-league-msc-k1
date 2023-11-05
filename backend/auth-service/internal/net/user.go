package net

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type user struct {
	url string
}

func NewUser(url string) *user {
	return &user{url}
}

type Role struct {
	ID          string
	Title       string
	Description string
}

type AuthenticateRequest struct {
	Email    string
	Password string
}

type AuthenticateResponse struct {
	ID    string
	Roles []Role
}

func (u *user) Authenticate(email string, password string) (*AuthenticateResponse, error) {
	reqData := new(bytes.Buffer)
	e := json.NewEncoder(reqData)
	e.Encode(AuthenticateRequest{email, password})

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPost, u.url+"/user/authenticate", reqData)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	respData := new(AuthenticateResponse)
	d := json.NewDecoder(resp.Body)
	err = d.Decode(respData)

	return respData, err
}
