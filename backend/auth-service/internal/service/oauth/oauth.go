package oauth

import "auth-service/internal/net"

type OAuth struct {
	Server *Server
}

func New(user net.User, opts ManagerOptions) (*OAuth, error) {
	manager, err := NewManager(opts)
	if err != nil {
		return nil, err
	}

	server := NewServer(manager, user)

	return &OAuth{server}, nil
}
