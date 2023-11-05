package oauth

import (
	"auth-service/internal/net"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/go-oauth2/oauth2/v4"
	"github.com/go-oauth2/oauth2/v4/errors"
	"github.com/go-oauth2/oauth2/v4/server"
)

type Server struct {
	*server.Server
	webapi net.User
}

func NewServer(manager oauth2.Manager, webapi net.User) *Server {
	srv := &Server{server.NewServer(server.NewConfig(), manager), webapi}

	srv.SetAllowedGrantType(oauth2.PasswordCredentials, oauth2.Refreshing)
	srv.SetAllowGetAccessRequest(true)

	srv.SetClientInfoHandler(server.ClientFormHandler)
	srv.SetInternalErrorHandler(srv.internalErrorHandler)
	srv.SetResponseTokenHandler(srv.responseTokenHandler)

	srv.SetPasswordAuthorizationHandler(srv.passwordAuthorizationHandler)

	return srv
}

func (s *Server) responseTokenHandler(w http.ResponseWriter, data map[string]interface{}, header http.Header, statusCode ...int) error {
	if rt, ok := data["refresh_token"]; ok {
		if rtStr, ok := rt.(string); ok {
			cookie := &http.Cookie{
				Name:     "refresh_token",
				Path:     "/token",
				Expires:  time.Now().Add(time.Hour * 24 * 7),
				Value:    rtStr,
				Secure:   true,
				HttpOnly: true,
				SameSite: http.SameSiteNoneMode,
			}
			http.SetCookie(w, cookie)
		}
	}

	w.Header().Set("Content-Type", "application/json")

	for key := range header {
		w.Header().Set(key, header.Get(key))
	}

	status := http.StatusOK
	if len(statusCode) > 0 && statusCode[0] > 0 {
		status = statusCode[0]
	}

	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}

func (s *Server) passwordAuthorizationHandler(ctx context.Context, clientID, username, password string) (userID string, err error) {
	resp, err := s.webapi.Authenticate(username, password)
	if err == nil {
		userID = resp.ID
	}
	return
}

func (s *Server) internalErrorHandler(err error) (re *errors.Response) {
	log.Println("Internal Error:", err.Error())
	return
}
