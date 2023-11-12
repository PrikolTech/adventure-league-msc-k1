package service

type Sender interface {
	Send(subject string, body []byte, to []string) error
	SendPassword(password string, to string) error
}
