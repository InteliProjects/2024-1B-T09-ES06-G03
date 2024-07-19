package healthcheck

import (
	"net/http"

	"github.com/Inteli-College/2024-1B-T09-ES06-G03/config"
	"github.com/Inteli-College/2024-1B-T09-ES06-G03/utils"
	"github.com/gorilla/mux"
)

// Handler é o handler para os endpoints de healthcheck
type Handler struct{}

// NewHandler cria um novo handler de healthcheck
func NewHandler() *Handler {
	return &Handler{}
}

// HealthcheckHandler manipula as solicitações de verificação de integridade
// @Summary Verifica o status do serviço
// @Description Retorna o status atual do serviço, ambiente, versão e nome do serviço
// @Tags healthcheck
// @Produce json
// @Success 200 {object} map[string]string
// @Router /healthcheck [get]
func (h *Handler) HealthcheckHandler(w http.ResponseWriter, r *http.Request) {
	health := map[string]string{
		"status":  "available",
		"env":     config.Envs.Env,
		"version": config.Envs.Version,
		"service": config.Envs.Service,
	}

	utils.WriteJSON(w, http.StatusOK, health)
}

// RegisterRoutes registra as rotas de healthcheck no roteador
func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/healthcheck", h.HealthcheckHandler).Methods(http.MethodGet)
}
