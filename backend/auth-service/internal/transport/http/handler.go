package http

import (
	"auth-service/internal/net"
	"auth-service/internal/pkg/jwt"
	"auth-service/internal/service/oauth"
	"context"
	"encoding/json"
	"net/http"
	"time"
)

type TokenHandler struct {
	Server *oauth.Server
}

func (h TokenHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		data := r.FormValue("refresh_token") // important to parse multipart form
		if cookie, err := r.Cookie("refresh_token"); err != http.ErrNoCookie {
			data = cookie.Value
		} else if data != "" {
			data = ""
		}
		r.Form.Set("refresh_token", data)

		err := h.Server.HandleTokenRequest(w, r)
		if err != nil {
			ErrorJSON(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	if r.Method == http.MethodDelete {
		cookie, err := r.Cookie("refresh_token")
		if err != nil {
			ErrorJSON(w, "cookie is empty", http.StatusBadRequest)
			return
		}

		err = h.Server.Manager.RemoveRefreshToken(context.Background(), cookie.Value)
		if err != nil {
			ErrorJSON(w, err.Error(), http.StatusInternalServerError)
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

type VerifyHandler struct {
	Role   net.Role
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

	id := claims.Subject
	roles, err := h.Role.GetByUser(id, token)
	if err != nil {
		ErrorJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	e := json.NewEncoder(w)
	e.Encode(map[string]any{
		"id":    id,
		"roles": roles,
	})
}
