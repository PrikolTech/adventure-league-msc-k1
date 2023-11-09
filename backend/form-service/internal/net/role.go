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
	url string
}

func NewRole(url string) *role {
	return &role{url}
}

func (r *role) Append(userID uuid.UUID, title string) error {
	reqData := new(bytes.Buffer)
	e := json.NewEncoder(reqData)
	e.Encode(map[string]string{
		"user_id": userID.String(),
		"title":   title,
	})

	client := &http.Client{}
	url := fmt.Sprintf("%s/role/append", r.url)
	req, err := http.NewRequest(http.MethodPost, url, reqData)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		respData := make(map[string]string)
		d := json.NewDecoder(resp.Body)
		err = d.Decode(&respData)
		if err != nil {
			return err
		}

		return errors.New(respData["error"])
	}

	return nil
}
