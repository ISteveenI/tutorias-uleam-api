package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"github.com/steveenacostapatino/tutorias-uleam-api/internal/storage"
)

type SesionHandler struct {
	storage *storage.SesionStorage
}

func NewSesionHandler(storage *storage.SesionStorage) *SesionHandler {
	return &SesionHandler{
		storage: storage,
	}
}

func (h *SesionHandler) Routes() chi.Router {
	r := chi.NewRouter()

	r.Post("/", h.Create)
	r.Get("/", h.GetAll)
	r.Get("/{id}", h.GetByID)
	r.Put("/{id}", h.Update)
	r.Delete("/{id}", h.Delete)

	return r
}

func obtenerSesionID(r *http.Request) (int, error) {
	idParam := chi.URLParam(r, "id")
	return strconv.Atoi(idParam)
}

func escribirSesionJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func escribirSesionError(w http.ResponseWriter, status int, mensaje string) {
	respuesta := map[string]string{
		"error": mensaje,
	}

	escribirSesionJSON(w, status, respuesta)
}

func (h *SesionHandler) Create(w http.ResponseWriter, r *http.Request) {
	escribirSesionError(w, http.StatusNotImplemented, "Endpoint Create pendiente")
}

func (h *SesionHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	escribirSesionError(w, http.StatusNotImplemented, "Endpoint GetAll pendiente")
}

func (h *SesionHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	escribirSesionError(w, http.StatusNotImplemented, "Endpoint GetByID pendiente")
}

func (h *SesionHandler) Update(w http.ResponseWriter, r *http.Request) {
	escribirSesionError(w, http.StatusNotImplemented, "Endpoint Update pendiente")
}

func (h *SesionHandler) Delete(w http.ResponseWriter, r *http.Request) {
	escribirSesionError(w, http.StatusNotImplemented, "Endpoint Delete pendiente")
}
