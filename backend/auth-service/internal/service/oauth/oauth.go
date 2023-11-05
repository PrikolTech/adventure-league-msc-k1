package oauth

import "auth-service/internal/webapi"

type OAuth struct {
	Server *Server
}

func New(user webapi.User, key []byte) (*OAuth, error) {
	manager, err := NewManager(key)
	if err != nil {
		return nil, err
	}

	server := NewServer(manager, user)

	return &OAuth{server}, nil
}
