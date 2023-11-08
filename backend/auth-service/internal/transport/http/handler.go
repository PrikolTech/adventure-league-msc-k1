package http

import (
	"auth-service/internal/service/oauth"
	"context"
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

		err = h.Server.Manager.RemoveRefreshToken(context.Background(), cookie.Value)
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
