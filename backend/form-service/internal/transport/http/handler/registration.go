package handler

import (
	"context"
	"encoding/json"
	"errors"
	"form-service/internal/entity"
	"form-service/internal/service"
	"net/http"
	"slices"

	"github.com/go-chi/chi/v5"
	"github.com/gofrs/uuid/v5"
)

type Registration struct {
	service service.Registration
}

func NewRegistration(service service.Registration) *Registration {
	return &Registration{service}
}

func (h *Registration) readRegistration(r *http.Request) (*entity.Registration, error) {
	data := new(entity.Registration)
	d := json.NewDecoder(r.Body)
	err := d.Decode(data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (h *Registration) create(w http.ResponseWriter, data *entity.Registration) {
	registration, err := h.service.Create(*data)
	if err != nil {
		ErrorJSON(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	e := json.NewEncoder(w)
	e.Encode(registration)
}

func (h *Registration) Create(w http.ResponseWriter, r *http.Request) {
	data, err := h.readRegistration(r)
	if err != nil {
		ErrorJSON(w, err.Error(), http.StatusBadRequest)
		return
	}

	data.UserID = nil
	h.create(w, data)
}

func (h *Registration) Append(w http.ResponseWriter, r *http.Request) {
	data, err := h.readRegistration(r)
	if err != nil {
		ErrorJSON(w, err.Error(), http.StatusBadRequest)
		return
	}

	if data.UserID == nil {
		ErrorJSON(w, "user_id must be non-nil", http.StatusBadRequest)
		return
	}

	if err := validateUserID(r.Context(), *data.UserID); err != nil {
		ErrorJSON(w, err.Error(), http.StatusForbidden)
		return
	}

	if err := validateRoles(r.Context(), entity.Enrollee, entity.Student); err != nil {
		ErrorJSON(w, err.Error(), http.StatusForbidden)
		return
	}

	h.create(w, data)
}

func (h *Registration) Get(w http.ResponseWriter, r *http.Request) {
	param := chi.URLParam(r, "id")
	id, err := uuid.FromString(param)
	if err != nil {
		DecodingError(w)
		return
	}

	registration, err := h.service.Get(id)
	if err != nil {
		ErrorJSON(w, err.Error(), http.StatusBadRequest)
		return
	}

	if registration.UserID == nil {
		if err := validateRoles(r.Context(), entity.Employee); err != nil {
			ErrorJSON(w, err.Error(), http.StatusForbidden)
			return
		}
	} else {
		if err := validateUserID(r.Context(), *registration.UserID); err != nil {
			ErrorJSON(w, err.Error(), http.StatusForbidden)
			return
		}

		if err = validateRoles(r.Context(), entity.Enrollee, entity.Student, entity.Employee); err != nil {
			ErrorJSON(w, err.Error(), http.StatusForbidden)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	e := json.NewEncoder(w)
	e.Encode(registration)
}

func (h *Registration) List(w http.ResponseWriter, r *http.Request) {
	var (
		forms []entity.Registration
		err   error
	)

	param := r.URL.Query().Get("user_id")
	if param == "" {
		if err := validateRoles(r.Context(), entity.Employee); err != nil {
			ErrorJSON(w, err.Error(), http.StatusForbidden)
			return
		}

		param = r.URL.Query().Get("course_id")
		if param == "" {
			forms, err = h.service.List()
		} else {
			var courseID uuid.UUID
			courseID, err = uuid.FromString(param)
			if err != nil {
				DecodingError(w)
				return
			}
			forms, err = h.service.GetByCourse(courseID)
		}
	} else {
		var userID uuid.UUID
		userID, err = uuid.FromString(param)
		if err != nil {
			DecodingError(w)
			return
		}

		if err := validateUserID(r.Context(), userID); err != nil {
			ErrorJSON(w, err.Error(), http.StatusForbidden)
			return
		}

		if err := validateRoles(r.Context(), entity.Enrollee, entity.Student, entity.Employee); err != nil {
			ErrorJSON(w, err.Error(), http.StatusForbidden)
			return
		}

		forms, err = h.service.GetByUser(userID)
	}

	if err != nil {
		ErrorJSON(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	e := json.NewEncoder(w)
	e.Encode(forms)
}

func (h *Registration) Update(w http.ResponseWriter, r *http.Request) {
	if err := validateRoles(r.Context(), entity.Employee); err != nil {
		ErrorJSON(w, err.Error(), http.StatusForbidden)
		return
	}

	param := chi.URLParam(r, "id")
	id, err := uuid.FromString(param)
	if err != nil {
		DecodingError(w)
		return
	}

	data, err := h.readRegistration(r)
	if err != nil {
		ErrorJSON(w, err.Error(), http.StatusBadRequest)
		return
	}

	data.ID = id

	ctx := r.Context().Value("token")
	form, err := h.service.Update(*data, ctx.(string))
	if err != nil {
		ErrorJSON(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	e := json.NewEncoder(w)
	e.Encode(form)
}

func validateUserID(ctx context.Context, id uuid.UUID) error {
	userID, _ := ctx.Value("userID").(uuid.UUID)

	err := validateRoles(ctx, entity.Employee)
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
