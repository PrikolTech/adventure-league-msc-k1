package net

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type auth struct {
	client *http.Client
	url    string
}

func NewAuth(url string) *auth {
	client := &http.Client{}
	return &auth{client, url}
}

type VerifyResponseBody struct {
	ID    string   `json:"id"`
	Roles []string `json:"roles"`
}

func (a *auth) Verify(token string) (string, []string, error) {
	url := fmt.Sprintf("%s/verify?access_token=%s", a.url, token)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", nil, ErrInternalNetwork
	}

	resp, err := a.client.Do(req)
	if err != nil {
		return "", nil, ErrInternalNetwork
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		respData := make(map[string]string)
		d := json.NewDecoder(resp.Body)
		err = d.Decode(&respData)
		if err != nil {
			return "", nil, ErrInternalNetwork
		}

		return "", nil, errors.New(respData["error"])
	}

	respData := new(VerifyResponseBody)
	d := json.NewDecoder(resp.Body)
	err = d.Decode(respData)
	if err != nil {
		return "", nil, err
	}

	return respData.ID, respData.Roles, nil
}
