package handler

import (
	"encoding/json"
	"net/http"
	"smtp-service/internal/service"
)

type password struct {
	sender service.Sender
}

func NewPassword(sender service.Sender) *password {
	return &password{sender}
}

func (p *password) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		MethodNotAllowed(w, r)
		return
	}

	var body struct {
		To       string `json:"to"`
		Password string `json:"password"`
	}

	d := json.NewDecoder(r.Body)
	err := d.Decode(&body)
	if err != nil {
		DecodingError(w)
		return
	}

	if body.Password == "" || body.To == "" {
		DecodingError(w)
		return
	}

	err = p.sender.SendPassword(body.Password, body.To)
	if err != nil {
		InternalServerError(w)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
