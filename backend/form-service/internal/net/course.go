package net

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/gofrs/uuid/v5"
)

type course struct {
	client *http.Client
	url    string
}

func NewCourse(url string) *course {
	client := &http.Client{}
	return &course{client, url}
}

func (c *course) Append(userID uuid.UUID, courseID uuid.UUID) error {
	url := fmt.Sprintf("%s/courses/%s/add_user?user_id=%s", c.url, courseID, userID)
	req, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		return ErrInternalNetwork
	}

	resp, err := c.client.Do(req)
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

		return errors.New(respData["message"])
	}

	return nil
}
