package net

type Auth interface {
	Verify(token string) (string, []string, error)
}
