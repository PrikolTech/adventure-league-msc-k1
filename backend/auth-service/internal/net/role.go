package net

import (
	"auth-service/internal/entity"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type role struct {
	client *http.Client
	url    string
}

func NewRole(url string) *role {
	client := &http.Client{}
	return &role{client, url}
}

func (r *role) GetByUser(userID string, token string) ([]entity.Role, error) {
	url := fmt.Sprintf("%s/role?user_id=%s", r.url, userID)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, ErrInternalNetwork
	}

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := r.client.Do(req)
	if err != nil {
		return nil, ErrInternalNetwork
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		respData := make(map[string]string)
		d := json.NewDecoder(resp.Body)
		err = d.Decode(&respData)
		if err != nil {
			return nil, ErrInternalNetwork
		}

		return nil, errors.New(respData["error"])
	}

	var respData []entity.Role
	d := json.NewDecoder(resp.Body)
	err = d.Decode(&respData)

	return respData, err
}
