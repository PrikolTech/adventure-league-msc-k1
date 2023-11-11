package net

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/gofrs/uuid/v5"
)

type role struct {
	client *http.Client
	url    string
}

func NewRole(url string) *role {
	client := &http.Client{}
	return &role{client, url}
}

func (r *role) Append(userID uuid.UUID, title string, token string) error {
	reqData := new(bytes.Buffer)
	e := json.NewEncoder(reqData)
	e.Encode(map[string]string{
		"user_id": userID.String(),
		"title":   title,
	})

	url := fmt.Sprintf("%s/role/append", r.url)
	req, err := http.NewRequest(http.MethodPost, url, reqData)
	if err != nil {
		return ErrInternalNetwork
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := r.client.Do(req)
	if err != nil {
		return ErrInternalNetwork
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		respData := make(map[string]string)
		d := json.NewDecoder(resp.Body)
		err = d.Decode(&respData)
		if err != nil {
			return ErrInternalNetwork
		}

		return errors.New(respData["error"])
	}

	return nil
}
