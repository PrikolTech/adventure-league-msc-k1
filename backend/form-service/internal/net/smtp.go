package net

import "net/http"

type smtp struct {
	client *http.Client
	url    string
}

func NewSMTP(url string) *smtp {
	client := &http.Client{}
	return &smtp{client, url}
}

func (s *smtp) SendPassword(password string, to string) {

}
