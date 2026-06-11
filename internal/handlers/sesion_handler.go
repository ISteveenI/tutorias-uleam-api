package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"github.com/steveenacostapatino/tutorias-uleam-api/internal/models"
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
func validarSesion(sesion models.SesionTutoria) bool {
	if sesion.SolicitudID == 0 {
		return false
	}

	if sesion.DisponibilidadID == 0 {
		return false
	}

	if sesion.DocenteID == 0 {
		return false
	}

	if sesion.EstudianteID == 0 {
		return false
	}

	if sesion.Fecha == "" {
		return false
	}

	if sesion.HoraInicio == "" {
		return false
	}

	if sesion.HoraFin == "" {
		return false
	}

	if sesion.Estado == "" {
		return false
	}

	return true
}
func (h *SesionHandler) Create(w http.ResponseWriter, r *http.Request) {
	var sesion models.SesionTutoria

	err := json.NewDecoder(r.Body).Decode(&sesion)
	if err != nil {
		escribirSesionError(w, http.StatusBadRequest, "JSON invalido")
		return
	}

	if !validarSesion(sesion) {
		escribirSesionError(w, http.StatusBadRequest, "Faltan campos obligatorios")
		return
	}

	sesionCreada := h.storage.Create(sesion)

	escribirSesionJSON(w, http.StatusCreated, sesionCreada)
}

func (h *SesionHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	sesiones := h.storage.GetAll()

	escribirSesionJSON(w, http.StatusOK, sesiones)
}
func (h *SesionHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := obtenerSesionID(r)
	if err != nil {
		escribirSesionError(w, http.StatusBadRequest, "ID invalido")
		return
	}

	sesion, err := h.storage.GetByID(id)
	if err != nil {
		escribirSesionError(w, http.StatusNotFound, "Sesion no encontrada")
		return
	}

	escribirSesionJSON(w, http.StatusOK, sesion)
}

func (h *SesionHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := obtenerSesionID(r)
	if err != nil {
		escribirSesionError(w, http.StatusBadRequest, "ID invalido")
		return
	}

	var sesion models.SesionTutoria

	err = json.NewDecoder(r.Body).Decode(&sesion)
	if err != nil {
		escribirSesionError(w, http.StatusBadRequest, "JSON invalido")
		return
	}

	if !validarSesion(sesion) {
		escribirSesionError(w, http.StatusBadRequest, "Faltan campos obligatorios")
		return
	}

	sesionActualizada, err := h.storage.Update(id, sesion)
	if err != nil {
		escribirSesionError(w, http.StatusNotFound, "Sesion no encontrada")
		return
	}

	escribirSesionJSON(w, http.StatusOK, sesionActualizada)
}

func (h *SesionHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := obtenerSesionID(r)
	if err != nil {
		escribirSesionError(w, http.StatusBadRequest, "ID invalido")
		return
	}

	err = h.storage.Delete(id)
	if err != nil {
		escribirSesionError(w, http.StatusNotFound, "Sesion no encontrada")
		return
	}

	respuesta := map[string]string{
		"mensaje": "Sesion eliminada correctamente",
	}

	escribirSesionJSON(w, http.StatusOK, respuesta)
}
