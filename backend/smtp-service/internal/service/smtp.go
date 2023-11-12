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

func (s *sender) SendPassword(password string, to string) error {
	var body bytes.Buffer
	tmpl, err := template.ParseFiles(filepath.Join(tmplPath, "index.html"))
	if err != nil {
		return err
	}

	err = tmpl.Execute(&body, nil)
	if err != nil {
		return err
	}

	return s.Send(body.Bytes(), []string{to})
}

func (s *sender) Send(content []byte, to []string) error {
	var body bytes.Buffer
	fmt.Fprintf(&body, "From: %s\n", s.from)
	fmt.Fprintf(&body, "To: %s\n", strings.Join(to, ","))
	body.WriteString("Subject: This is a test subject\n")
	body.WriteString("MIME-version: 1.0\n")
	body.WriteString("Content-Type: text/html; charset=\"utf-8\"\n")
	body.Write(content)

	err := smtp.SendMail(s.addr, s.auth, s.from, to, body.Bytes())
	if err != nil {
		return err
	}

	return nil
}
