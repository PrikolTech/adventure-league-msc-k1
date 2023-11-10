package handler

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"slices"
	"user-service/internal/entity"
	"user-service/internal/service"

	"github.com/go-chi/chi/v5"
	"github.com/gofrs/uuid/v5"
)

type User struct {
	service service.User
}

func NewUser(service service.User) *User {
	return &User{service}
}

type AuthenticateRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthenticateResponse struct {
	ID    uuid.UUID     `json:"id"`
	Roles []entity.Role `json:"roles"`
}

func (h *User) Authenticate(w http.ResponseWriter, r *http.Request) {
	data := new(AuthenticateRequest)
	d := json.NewDecoder(r.Body)
	err := d.Decode(data)
	if err != nil {
		DecodingError(w)
		return
	}

	id, roles, err := h.service.Authenticate(data.Email, data.Password)
	if err != nil {
		ErrorJSON(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	e := json.NewEncoder(w)
	e.Encode(AuthenticateResponse{
		ID:    id,
		Roles: roles,
	})
}

func (h *User) Exist(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")
	if email == "" {
		DecodingError(w)
		return
	}

	exists, err := h.service.Exist(email)
	if err != nil {
		InternalServerError(w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	e := json.NewEncoder(w)
	e.Encode(map[string]bool{
		"exists": exists,
	})
}

func (h *User) Create(w http.ResponseWriter, r *http.Request) {
	var data entity.User
	d := json.NewDecoder(r.Body)
	err := d.Decode(&data)
	if err != nil {
		DecodingError(w)
		return
	}

	user, err := h.service.Create(data)
	if err != nil {
		ErrorJSON(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	e := json.NewEncoder(w)
	e.Encode(user)
}

func (h *User) Get(w http.ResponseWriter, r *http.Request) {
	param := chi.URLParam(r, "id")
	id, err := uuid.FromString(param)
	if err != nil {
		DecodingError(w)
		return
	}

	user, err := h.service.Get(id)
	if err != nil {
		ErrorJSON(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	e := json.NewEncoder(w)
	e.Encode(user)
}

func (h *User) List(w http.ResponseWriter, r *http.Request) {
	if err := validateRoles(r.Context()); err != nil {
		ErrorJSON(w, err.Error(), http.StatusForbidden)
		return
	}

	users, err := h.service.List()
	if err != nil {
		ErrorJSON(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	e := json.NewEncoder(w)
	e.Encode(users)
}

func (h *User) Update(w http.ResponseWriter, r *http.Request) {
	param := chi.URLParam(r, "id")
	id, err := uuid.FromString(param)
	if err != nil {
		DecodingError(w)
		return
	}

	if err := validateUserID(r.Context(), id); err != nil {
		ErrorJSON(w, err.Error(), http.StatusForbidden)
		return
	}

	if err := validateRoles(r.Context(), entity.Enrollee, entity.Student); err != nil {
		ErrorJSON(w, err.Error(), http.StatusForbidden)
		return
	}

	var data entity.User
	d := json.NewDecoder(r.Body)
	err = d.Decode(&data)
	if err != nil {
		DecodingError(w)
		return
	}

	data.ID = id
	user, err := h.service.Update(data)
	if err != nil {
		ErrorJSON(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	e := json.NewEncoder(w)
	e.Encode(user)
}

func (h *User) Delete(w http.ResponseWriter, r *http.Request) {
	param := chi.URLParam(r, "id")
	id, err := uuid.FromString(param)
	if err != nil {
		DecodingError(w)
		return
	}

	if err := validateUserID(r.Context(), id); err != nil {
		ErrorJSON(w, err.Error(), http.StatusForbidden)
		return
	}

	if err := validateRoles(r.Context(), entity.Enrollee, entity.Student); err != nil {
		ErrorJSON(w, err.Error(), http.StatusForbidden)
		return
	}

	err = h.service.Delete(id)
	if err != nil {
		ErrorJSON(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func validateUserID(ctx context.Context, id uuid.UUID, roles ...entity.RoleTitle) error {
	userID, _ := ctx.Value("userID").(uuid.UUID)

	err := validateRoles(ctx, roles...)
	if userID != id && err != nil {
		return errors.New("you pick the wrong house")
	}

	return nil
}

func validateRoles(ctx context.Context, roles ...entity.RoleTitle) error {
	userRoles, _ := ctx.Value("roles").([]string)
	for _, role := range append(roles, entity.Admin) {
		if slices.Contains(userRoles, string(role)) {
			return nil
		}
	}

	return errors.New("you have no power here")
}
