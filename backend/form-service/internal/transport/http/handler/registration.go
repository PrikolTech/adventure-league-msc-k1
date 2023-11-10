package handler

import (
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

	if err := validateAccess(data.ID, r); err != nil {
		ErrorJSON(w, err.Error(), http.StatusUnauthorized)
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

	var userID uuid.UUID
	if registration.UserID == nil {
		userID = uuid.Nil
	} else {
		userID = *registration.UserID
	}

	if err := validateAccess(userID, r); err != nil {
		ErrorJSON(w, err.Error(), http.StatusUnauthorized)
		return
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
		if err := validateAccess(uuid.Nil, r); err != nil {
			ErrorJSON(w, err.Error(), http.StatusUnauthorized)
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

		if err := validateAccess(userID, r); err != nil {
			ErrorJSON(w, err.Error(), http.StatusUnauthorized)
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
	if err := validateAccess(uuid.Nil, r); err != nil {
		ErrorJSON(w, err.Error(), http.StatusUnauthorized)
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

func validateAccess(id uuid.UUID, r *http.Request) error {
	ctx := r.Context()
	userID, _ := ctx.Value("userID").(uuid.UUID)
	roles, _ := ctx.Value("roles").([]string)

	isAdmin := slices.Contains(roles, string(entity.Admin))

	if id != userID && !isAdmin {
		return errors.New("unauthorized")
	}

	return nil
}
