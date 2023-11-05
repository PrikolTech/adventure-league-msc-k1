package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/go-oauth2/oauth2/v4/errors"
	"github.com/go-oauth2/oauth2/v4/server"
)

type TokenHandler struct {
	srv *server.Server
}

func (h TokenHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
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
		return
	}

	if r.Method == http.MethodDelete {
		cookie, err := r.Cookie("refresh_token")
		if err != nil {
			http.Error(w, "cookie is empty", http.StatusBadRequest)
			return
		}

		err = h.srv.Manager.RemoveRefreshToken(context.Background(), cookie.Value)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		cookie = &http.Cookie{
			Name:    "refresh_token",
			Expires: time.Unix(0, 0),
		}

		http.SetCookie(w, cookie)
		w.WriteHeader(http.StatusNoContent)
		return
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
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

func InternalErrorHandler(err error) (re *errors.Response) {
	log.Println("Internal Error:", err.Error())
	return
}
