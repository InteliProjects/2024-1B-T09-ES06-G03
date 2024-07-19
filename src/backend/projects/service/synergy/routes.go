package synergy

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Inteli-College/2024-1B-T09-ES06-G03/project/service/auth"
	"github.com/Inteli-College/2024-1B-T09-ES06-G03/project/types"
	"github.com/Inteli-College/2024-1B-T09-ES06-G03/project/utils"
	"github.com/gorilla/mux"
)

// Handler é o handler para os endpoints de synergy
type Handler struct {
	store types.SynergyStore
}

// NewHandler cria um novo handler de synergy
func NewHandler(store types.SynergyStore) *Handler {
	return &Handler{store: store}
}

// RegisterRoutes registra as rotas de synergy no roteador
func (h *Handler) RegisterRoutes(router *mux.Router) {
    // Protect the GET route for all synergies to require JWT
    router.Handle("/synergies", auth.JWTMiddleware(http.HandlerFunc(h.handleGetSynergies))).Methods(http.MethodGet)

    // Protect the POST route for creating a synergy, requiring JWT
    router.Handle("/synergies", auth.JWTMiddleware(http.HandlerFunc(h.handlePostSynergy))).Methods(http.MethodPost)

    // Protect the GET by ID to ensure only authenticated users can access specific synergy details
    router.Handle("/synergies/{id}", auth.JWTMiddleware(http.HandlerFunc(h.handleGetSynergyByID))).Methods(http.MethodGet)

    // Protect the PUT for updating synergies to require JWT
    router.Handle("/synergies/{id}", auth.JWTMiddleware(http.HandlerFunc(h.handleUpdateSynergy))).Methods(http.MethodPut)

    // Protect the DELETE to require JWT
    router.Handle("/synergies/{id}", auth.JWTMiddleware(http.HandlerFunc(h.handleDeleteSynergy))).Methods(http.MethodDelete)

    // Protect the GET synergy by description route to restrict access to authenticated users
    router.Handle("/synergies/description/{description}", auth.JWTMiddleware(http.HandlerFunc(h.handleGetSynergyByDescription))).Methods(http.MethodGet) // Nova rota
}


// handleGetSynergies manipula as solicitações de obtenção de todas as sinergias
// @Summary Obtém todas as sinergias
// @Description Retorna uma lista de todas as sinergias, incluindo detalhes dos projetos
// @Tags synergies
// @Produce json
// @Success 200 {array} types.DetailedSynergy
// @Failure 500 {object} utils.ErrorResponse
// @Router /synergies [get]
func (h *Handler) handleGetSynergies(w http.ResponseWriter, r *http.Request) {
	synergies, err := h.store.GetSynergies()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, synergies)
}

// handleGetSynergyByID manipula as solicitações de obtenção de uma sinergia por ID
// @Summary Obtém uma sinergia por ID
// @Description Retorna uma sinergia com base no ID fornecido, incluindo detalhes dos projetos
// @Tags synergies
// @Produce json
// @Param id path int true "ID da sinergia"
// @Success 200 {object} types.DetailedSynergy
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /synergies/{id} [get]
func (h *Handler) handleGetSynergyByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid synergy ID"))
		return
	}

	synergy, err := h.store.GetSynergyByID(id)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	if synergy == nil {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("synergy not found"))
		return
	}

	utils.WriteJSON(w, http.StatusOK, synergy)
}

// handlePostSynergy manipula as solicitações de criação de uma nova sinergia
// @Summary Cria uma nova sinergia
// @Description Cria uma nova sinergia com base nos dados fornecidos
// @Tags synergies
// @Accept json
// @Produce json
// @Param synergy body types.SynergyPayload true "Dados da sinergia"
// @Success 201 {string} string "synergy created successfully"
// @Failure 400 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /synergies [post]
func (h *Handler) handlePostSynergy(w http.ResponseWriter, r *http.Request) {
	var payload types.SynergyPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", err))
		return
	}

	synergy := types.Synergy{
		SourceProjectID: payload.SourceProjectID,
		TargetProjectID: payload.TargetProjectID,
		Status:          payload.Status,
		Type:            payload.Type,
		Description:     payload.Description,
	}

	if err := h.store.CreateSynergy(&synergy); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, "synergy created successfully")
}

// handleUpdateSynergy manipula as solicitações de atualização de uma sinergia
// @Summary Atualiza uma sinergia existente
// @Description Atualiza uma sinergia existente com base nos dados fornecidos
// @Tags synergies
// @Accept json
// @Produce json
// @Param id path int true "ID da sinergia"
// @Param synergy body types.SynergyPayload true "Dados da sinergia"
// @Success 200 {string} string "synergy updated successfully"
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /synergies/{id} [put]
func (h *Handler) handleUpdateSynergy(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid synergy ID"))
		return
	}

	var payload types.SynergyPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", err))
		return
	}

	synergy := types.Synergy{
		ID:              id,
		SourceProjectID: payload.SourceProjectID,
		TargetProjectID: payload.TargetProjectID,
		Status:          payload.Status,
		Type:            payload.Type,
		Description:     payload.Description,
	}

	if err := h.store.UpdateSynergy(id, &synergy); err != nil {
		if err == sql.ErrNoRows {
			utils.WriteError(w, http.StatusNotFound, fmt.Errorf("synergy not found"))
		} else {
			utils.WriteError(w, http.StatusInternalServerError, err)
		}
		return
	}

	utils.WriteJSON(w, http.StatusOK, "synergy updated successfully")
}

// handleDeleteSynergy manipula as solicitações de exclusão de uma sinergia
// @Summary Exclui uma sinergia existente
// @Description Exclui uma sinergia existente com base no ID fornecido
// @Tags synergies
// @Produce json
// @Param id path int true "ID da sinergia"
// @Success 200 {string} string "synergy deleted successfully"
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /synergies/{id} [delete]
func (h *Handler) handleDeleteSynergy(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid synergy ID"))
		return
	}

	if err := h.store.DeleteSynergy(id); err != nil {
		if err == sql.ErrNoRows {
			utils.WriteError(w, http.StatusNotFound, fmt.Errorf("synergy not found"))
		} else {
			utils.WriteError(w, http.StatusInternalServerError, err)
		}
		return
	}

	utils.WriteJSON(w, http.StatusOK, "synergy deleted successfully")
}

func (h *Handler) handleGetSynergyByDescription(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	description := vars["description"]

	synergies, err := h.store.GetSynergiesByDescription(description)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, synergies)
}
