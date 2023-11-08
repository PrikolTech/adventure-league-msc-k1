package handler

import (
	"encoding/json"
	"form-service/internal/entity"
	"form-service/internal/service"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/gofrs/uuid/v5"
)

type Registration struct {
	service service.Registration
}

func NewRegistration(service service.Registration) *Registration {
	return &Registration{service}
}

func (h *Registration) Create(w http.ResponseWriter, r *http.Request) {
	var data entity.Registration
	d := json.NewDecoder(r.Body)
	err := d.Decode(&data)
	if err != nil {
		DecodingError(w)
		return
	}

	registration, err := h.service.Create(data)
	if err != nil {
		ErrorJSON(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	e := json.NewEncoder(w)
	e.Encode(registration)
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
	param := chi.URLParam(r, "id")
	id, err := uuid.FromString(param)
	if err != nil {
		DecodingError(w)
		return
	}

	var data entity.Registration
	d := json.NewDecoder(r.Body)
	err = d.Decode(&data)
	if err != nil {
		DecodingError(w)
		return
	}

	data.ID = id
	form, err := h.service.Update(data)
	if err != nil {
		ErrorJSON(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	e := json.NewEncoder(w)
	e.Encode(form)
}
