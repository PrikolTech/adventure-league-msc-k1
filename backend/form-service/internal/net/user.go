package net

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"form-service/internal/entity"
	"net/http"

	"github.com/gofrs/uuid/v5"
)

type user struct {
	client *http.Client
	url    string
}

func NewUser(url string) *user {
	client := &http.Client{}
	return &user{client, url}
}

func (u *user) Exist(email string) (bool, error) {
	url := fmt.Sprintf("%s/user?email=%s", u.url, email)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return false, ErrInternalNetwork
	}

	resp, err := u.client.Do(req)
	if err != nil {
		return false, ErrInternalNetwork
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		respData := make(map[string]string)
		d := json.NewDecoder(resp.Body)
		err = d.Decode(&respData)
		if err != nil {
			return false, ErrInternalNetwork
		}

		return false, errors.New(respData["error"])
	}

	respData := make(map[string]bool)
	d := json.NewDecoder(resp.Body)
	err = d.Decode(&respData)
	if err != nil {
		return false, ErrInternalNetwork
	}

	exists, ok := respData["exists"]
	if ok {
		return exists, nil
	}

	return false, ErrInternalNetwork
}

func (u *user) Get(id uuid.UUID, token string) (*entity.User, error) {
	url := fmt.Sprintf("%s/user/%s", u.url, id)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+token)
	resp, err := u.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
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

func (u *user) Create(data entity.User) (*entity.User, error) {
	reqData := new(bytes.Buffer)
	e := json.NewEncoder(reqData)
	e.Encode(data)

	req, err := http.NewRequest(http.MethodPost, u.url+"/user", reqData)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := u.client.Do(req)
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
