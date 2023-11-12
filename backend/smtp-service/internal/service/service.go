package service

type Sender interface {
	Send(content []byte, to []string) error
	SendPassword(password string, to string) error
}
