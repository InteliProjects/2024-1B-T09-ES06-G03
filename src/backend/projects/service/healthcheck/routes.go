package healthcheck

import (
	"net/http"

	"github.com/Inteli-College/2024-1B-T09-ES06-G03/project/config"
	"github.com/Inteli-College/2024-1B-T09-ES06-G03/project/utils"
	"github.com/gorilla/mux"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) HealthcheckHandler(w http.ResponseWriter, r *http.Request) {
	health := map[string]string{
		"status":  "available",
		"env":     config.Envs.Env,
		"version": config.Envs.Version,
		"service": config.Envs.Service,
	}

	utils.WriteJSON(w, http.StatusOK, health)
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/healthcheck", h.HealthcheckHandler).Methods(http.MethodGet)
}
