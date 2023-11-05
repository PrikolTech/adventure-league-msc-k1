package main

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-oauth2/oauth2/v4/server"
)

type TokenHandler struct {
	srv *server.Server
}

func (h TokenHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	data := r.FormValue("refresh_token") // important to parse multipart form
	if cookie, err := r.Cookie("refresh_token"); err != http.ErrNoCookie {
		data = cookie.Value
	} else if data != "" {
		data = ""
	}
	r.Form.Set("refresh_token", data)

	err := h.srv.HandleTokenRequest(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func ResponseTokenHandler(w http.ResponseWriter, data map[string]interface{}, header http.Header, statusCode ...int) error {
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

func PasswordAuthorizationHandler(ctx context.Context, clientID, username, password string) (userID string, err error) {
	if username == "123" && password == "111" {
		return "id", nil
	}
	return "", nil
}
