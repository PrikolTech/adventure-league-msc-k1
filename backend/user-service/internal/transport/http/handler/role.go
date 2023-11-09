package handler

import (
	"encoding/json"
	"net/http"
	"user-service/internal/entity"
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

func (h *Role) Create(w http.ResponseWriter, r *http.Request) {
	var data entity.Role
	d := json.NewDecoder(r.Body)
	err := d.Decode(&data)
	if err != nil {
		DecodingError(w)
		return
	}

	role, err := h.service.Create(data)
	if err != nil {
		ErrorJSON(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	e := json.NewEncoder(w)
	e.Encode(role)
}

type AppendRequestBody struct {
	UserID uuid.UUID  `json:"user_id"`
	RoleID *uuid.UUID `json:"role_id"`
	Title  string     `json:"title"`
}

func (h *Role) Append(w http.ResponseWriter, r *http.Request) {
	data := new(AppendRequestBody)
	d := json.NewDecoder(r.Body)
	err := d.Decode(data)
	if err != nil {
		DecodingError(w)
		return
	}

	if data.RoleID != nil {
		err = h.service.AppendByID(data.UserID, *data.RoleID)
	} else if data.Title != "" {
		err = h.service.AppendByTitle(data.UserID, data.Title)
	} else {
		DecodingError(w)
		return
	}

	if err != nil {
		ErrorJSON(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *Role) Remove(w http.ResponseWriter, r *http.Request) {
	userParam := r.URL.Query().Get("user_id")
	userID, err := uuid.FromString(userParam)
	if err != nil {
		DecodingError(w)
		return
	}

	id := r.URL.Query().Get("role_id")
	title := r.URL.Query().Get("title")

	if id != "" {
		var roleID uuid.UUID
		roleID, err = uuid.FromString(id)
		if err != nil {
			DecodingError(w)
			return
		}
		err = h.service.RemoveByID(userID, roleID)
	} else if title != "" {
		err = h.service.RemoveByTitle(userID, title)
	} else {
		err = h.service.RemoveAll(userID)
	}

	if err != nil {
		ErrorJSON(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *Role) Get(w http.ResponseWriter, r *http.Request) {
	param := r.URL.Query().Get("user_id")
	var (
		roles []entity.Role
		err   error
	)

	if param == "" {
		roles, err = h.service.List()
	} else {
		var userID uuid.UUID
		userID, err = uuid.FromString(param)
		if err != nil {
			DecodingError(w)
			return
		}

		roles, err = h.service.GetByUser(userID)
	}

	if err != nil {
		ErrorJSON(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	e := json.NewEncoder(w)
	e.Encode(roles)
}

func (h *Role) Delete(w http.ResponseWriter, r *http.Request) {
	param := chi.URLParam(r, "id")
	id, err := uuid.FromString(param)
	if err != nil {
		DecodingError(w)
		return
	}

	err = h.service.Delete(id)
	if err != nil {
		ErrorJSON(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
