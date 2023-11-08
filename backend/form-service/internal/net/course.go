package net

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/gofrs/uuid/v5"
)

type course struct {
	url string
}

func NewCourse(url string) *course {
	return &course{url}
}

func (c *course) Append(userID uuid.UUID, courseID uuid.UUID) error {
	client := &http.Client{}
	url := fmt.Sprintf("%s/courses/%s/add_user?user_id=%s", c.url, courseID, userID)
	req, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		return err
	}

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

		return errors.New(respData["message"])
	}

	return nil
}
