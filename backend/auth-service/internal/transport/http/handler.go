package http

import (
	"auth-service/internal/entity"
	"auth-service/internal/pkg/jwt"
	"auth-service/internal/service"
	"encoding/json"
	"net/http"
	"time"
)

type TokenHandler struct {
	oauth service.OAuth
}

func (h TokenHandler) getToken(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	gt := query.Get("grant_type")
	var (
		token *entity.Token
		err   error
	)

	clientID := query.Get("client_id")
	if gt == "password" {
		email := query.Get("username")
		password := query.Get("password")
		token, err = h.oauth.GenerateToken(clientID, email, password)
	} else if gt == "refresh_token" {
		data := ""
		if cookie, err := r.Cookie("refresh_token"); err != http.ErrNoCookie {
			data = cookie.Value
		}
		token, err = h.oauth.RefreshToken(clientID, data)
	}

	if err != nil {
		ErrorJSON(w, err.Error(), http.StatusBadRequest)
		return
	}

	cookie := &http.Cookie{
		Name:     "refresh_token",
		Path:     "/token",
		Expires:  time.Now().Add(token.RefreshExpiresIn),
		Value:    token.Refresh,
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
	}
	http.SetCookie(w, cookie)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	e := json.NewEncoder(w)
	e.Encode(token)
	return
}

func (h TokenHandler) revokeToken(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("refresh_token")
	if err != nil {
		ErrorJSON(w, "cookie is empty", http.StatusBadRequest)
		return
	}

	err = h.oauth.RemoveToken(cookie.Value)
	if err != nil {
		ErrorJSON(w, err.Error(), http.StatusBadRequest)
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

func (h TokenHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		h.getToken(w, r)
		return
	}

	if r.Method == http.MethodDelete {
		h.revokeToken(w, r)
		return
	}

	ErrorJSON(w, "method not allowed", http.StatusMethodNotAllowed)
}

type VerifyHandler struct {
	Parser jwt.Parser
}

func (h VerifyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	token := r.URL.Query().Get("access_token")
	claims, err := h.Parser.Parse(token)
	if err != nil {
		ErrorJSON(w, err.Error(), http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	e := json.NewEncoder(w)
	e.Encode(map[string]any{
		"id":    claims.Id,
		"roles": claims.Scopes,
	})
}
