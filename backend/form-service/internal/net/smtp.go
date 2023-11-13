package net

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type smtp struct {
	client *http.Client
	url    string
}

func NewSMTP(url string) *smtp {
	client := &http.Client{}
	return &smtp{client, url}
}

func (s *smtp) SendPassword(password string, to string) error {
	reqData := new(bytes.Buffer)
	e := json.NewEncoder(reqData)
	e.Encode(map[string]string{
		"password": password,
		"to":       to,
	})

	url := fmt.Sprintf("%s/password", s.url)
	req, err := http.NewRequest(http.MethodPost, url, reqData)
	if err != nil {
		return ErrInternalNetwork
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := s.client.Do(req)
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
