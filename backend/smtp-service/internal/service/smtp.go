package service

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"
	"path/filepath"
	"strings"
)

const tmplPath = "templates"

type SenderOptions struct {
	Host     string
	Port     string
	Username string
	Password string
}

type sender struct {
	from string
	addr string
	auth smtp.Auth
}

func NewSender(opts *SenderOptions) *sender {
	auth := smtp.PlainAuth("", opts.Username, opts.Password, opts.Host)
	addr := fmt.Sprintf("%s:%s", opts.Host, opts.Port)
	return &sender{opts.Username, addr, auth}
}

type passwordData struct {
	Email    string
	Password string
}

func (s *sender) SendPassword(password string, to string) error {
	var body bytes.Buffer
	tmpl, err := template.ParseFiles(filepath.Join(tmplPath, "password.html"))
	if err != nil {
		return err
	}

	err = tmpl.Execute(&body, passwordData{to, password})
	if err != nil {
		return err
	}

	subject := "Заявка на регистрацию"
	return s.Send(subject, body.Bytes(), []string{to})
}

func (s *sender) Send(subject string, body []byte, to []string) error {
	var msg bytes.Buffer
	fmt.Fprintf(&msg, "From: %s\n", s.from)
	fmt.Fprintf(&msg, "To: %s\n", strings.Join(to, ","))
	fmt.Fprintf(&msg, "Subject: %s\n", subject)
	msg.WriteString("MIME-version: 1.0\n")
	msg.WriteString("Content-Type: text/html; charset=\"utf-8\"\n")
	msg.Write(body)

	err := smtp.SendMail(s.addr, s.auth, s.from, to, msg.Bytes())
	if err != nil {
		return err
	}

	return nil
}
