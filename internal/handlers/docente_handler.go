package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/steveenacostapatino/tutorias-uleam-api/internal/middlewares"
	"github.com/steveenacostapatino/tutorias-uleam-api/internal/models"
	"github.com/steveenacostapatino/tutorias-uleam-api/internal/repositories"
	"github.com/steveenacostapatino/tutorias-uleam-api/internal/services"
)

type DocenteHandler struct {
	service services.DocenteServiceInterface
}

func NewDocenteHandler(service services.DocenteServiceInterface) *DocenteHandler {
	return &DocenteHandler{
		service: service,
	}
}

func (h *DocenteHandler) Routes() chi.Router {

	r := chi.NewRouter()

	r.Get("/", h.GetAll)
	r.Get("/{id}", h.GetByID)

	r.Group(func(protegida chi.Router) {

		protegida.Use(middlewares.RequireAuth)

		protegida.Post("/", h.Create)
		protegida.Put("/{id}", h.Update)
		protegida.Delete("/{id}", h.Delete)

	})

	return r
}

func (h *DocenteHandler) Create(w http.ResponseWriter, r *http.Request) {

	var docente models.Docente

	if err := json.NewDecoder(r.Body).Decode(&docente); err != nil {
		escribirDocenteError(w, http.StatusBadRequest, "JSON invalido")
		return
	}

	if err := h.service.Create(r.Context(), &docente); err != nil {
		manejarErrorDocente(w, err)
		return
	}

	escribirDocenteJSON(w, http.StatusCreated, docente)
}

func (h *DocenteHandler) GetAll(w http.ResponseWriter, r *http.Request) {

	docentes, err := h.service.GetAll(r.Context())

	if err != nil {
		escribirDocenteError(w, http.StatusInternalServerError, "No se pudieron obtener los docentes")
		return
	}

	escribirDocenteJSON(w, http.StatusOK, docentes)
}

func (h *DocenteHandler) GetByID(w http.ResponseWriter, r *http.Request) {

	id, err := obtenerDocenteID(r)

	if err != nil {
		escribirDocenteError(w, http.StatusBadRequest, "ID invalido")
		return
	}

	docente, err := h.service.GetByID(r.Context(), id)

	if err != nil {
		manejarErrorDocente(w, err)
		return
	}

	escribirDocenteJSON(w, http.StatusOK, docente)
}

func (h *DocenteHandler) Update(w http.ResponseWriter, r *http.Request) {

	id, err := obtenerDocenteID(r)

	if err != nil {
		escribirDocenteError(w, http.StatusBadRequest, "ID invalido")
		return
	}

	var docente models.Docente

	if err := json.NewDecoder(r.Body).Decode(&docente); err != nil {
		escribirDocenteError(w, http.StatusBadRequest, "JSON invalido")
		return
	}

	if err := h.service.Update(r.Context(), id, &docente); err != nil {
		manejarErrorDocente(w, err)
		return
	}

	escribirDocenteJSON(w, http.StatusOK, docente)
}

func (h *DocenteHandler) Delete(w http.ResponseWriter, r *http.Request) {

	id, err := obtenerDocenteID(r)

	if err != nil {
		escribirDocenteError(w, http.StatusBadRequest, "ID invalido")
		return
	}

	if err := h.service.Delete(r.Context(), id); err != nil {
		manejarErrorDocente(w, err)
		return
	}

	escribirDocenteJSON(w, http.StatusOK, map[string]string{
		"mensaje": "Docente eliminado correctamente",
	})
}

func obtenerDocenteID(r *http.Request) (uint, error) {

	idParam := chi.URLParam(r, "id")

	id, err := strconv.ParseUint(idParam, 10, 64)

	return uint(id), err
}

func escribirDocenteJSON(w http.ResponseWriter, status int, data any) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(data)
}

func escribirDocenteError(w http.ResponseWriter, status int, mensaje string) {

	escribirDocenteJSON(w, status, map[string]string{
		"error": mensaje,
	})
}

func manejarErrorDocente(w http.ResponseWriter, err error) {

	if errors.Is(err, services.ErrDatosInvalidos) {

		escribirDocenteError(w, http.StatusBadRequest, "Datos invalidos")
		return

	}

	if errors.Is(err, repositories.ErrDocenteNoEncontrado) {

		escribirDocenteError(w, http.StatusNotFound, "Docente no encontrado")
		return

	}

	escribirDocenteError(w, http.StatusInternalServerError, "Error interno del servidor")
}
