package handler

import (
	"encoding/json"
	"net/http"
	"user-service/internal/service"

	"github.com/go-chi/chi/v5"
	"github.com/gofrs/uuid/v5"
)

type Role struct {
	service service.Role
}

func NewRole(service service.Role) *Role {
	return &Role{service}
}

func (role *Role) Get(w http.ResponseWriter, r *http.Request) {
	param := chi.URLParam(r, "userID")
	userID, err := uuid.FromString(param)
	if err != nil {
		DecodingError(w)
		return
	}

	roles, err := role.service.GetByUser(userID)
	if err != nil {
		ErrorJSON(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	e := json.NewEncoder(w)
	e.Encode(roles)
}

func (role *Role) List(w http.ResponseWriter, r *http.Request) {
	roles, err := role.service.List()
	if err != nil {
		ErrorJSON(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	e := json.NewEncoder(w)
	e.Encode(roles)
}
