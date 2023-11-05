package main

import (
	"github.com/go-oauth2/oauth2/v4"
	"github.com/go-oauth2/oauth2/v4/server"
)

func NewServer(manager oauth2.Manager) *server.Server {
	srv := server.NewServer(server.NewConfig(), manager)
	srv.SetAllowedGrantType(oauth2.PasswordCredentials, oauth2.Refreshing)
	srv.SetAllowGetAccessRequest(true)

	srv.SetClientInfoHandler(server.ClientFormHandler)
	srv.SetInternalErrorHandler(InternalErrorHandler)
	srv.SetResponseTokenHandler(ResponseTokenHandler)

	srv.SetPasswordAuthorizationHandler(PasswordAuthorizationHandler)

	return srv
}
