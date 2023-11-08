package net

import (
	"bytes"
	"encoding/json"
	"errors"
	"form-service/internal/entity"
	"net/http"
)

type user struct {
	url string
}

func NewUser(url string) *user {
	return &user{url}
}

func (u *user) Create(data entity.User) (*entity.User, error) {
	reqData := new(bytes.Buffer)
	e := json.NewEncoder(reqData)
	e.Encode(data)

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPost, u.url+"/user", reqData)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		respData := make(map[string]string)
		d := json.NewDecoder(resp.Body)
		err = d.Decode(&respData)
		if err != nil {
			return nil, err
		}

		return nil, errors.New(respData["error"])
	}

	respData := new(entity.User)
	d := json.NewDecoder(resp.Body)
	err = d.Decode(respData)
	return respData, err
}
