package update

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

// Handler é o handler para os endpoints de updates
type Handler struct {
	store types.UpdateStore
}

// NewHandler cria um novo handler de updates
func NewHandler(store types.UpdateStore) *Handler {
	return &Handler{store: store}
}

// RegisterRoutes registra as rotas de updates no roteador
func (h *Handler) RegisterRoutes(router *mux.Router) {
    // Protect the GET route for all updates to require JWT
    router.Handle("/updates", auth.JWTMiddleware(http.HandlerFunc(h.handleGetUpdates))).Methods(http.MethodGet)

    // Protect the GET by ID route to ensure only authenticated users can access specific update details
    router.Handle("/updates/{id}", auth.JWTMiddleware(http.HandlerFunc(h.handleGetUpdateByID))).Methods(http.MethodGet)

    // Protect the POST route for creating an update, requiring JWT
    router.Handle("/updates", auth.JWTMiddleware(http.HandlerFunc(h.handlePostUpdate))).Methods(http.MethodPost)

    // Protect the PUT for updating an update to require JWT
    router.Handle("/updates/{id}", auth.JWTMiddleware(http.HandlerFunc(h.handleUpdateUpdate))).Methods(http.MethodPut)

    // Protect the DELETE to require JWT
    router.Handle("/updates/{id}", auth.JWTMiddleware(http.HandlerFunc(h.handleDeleteUpdate))).Methods(http.MethodDelete)

    // Protect the GET update by title route to restrict access to authenticated users
    router.Handle("/updates/title/{title}", auth.JWTMiddleware(http.HandlerFunc(h.handleGetUpdateByTitle))).Methods(http.MethodGet) // Nova rota
}


// handleGetUpdates manipula as solicitações de obtenção de todas as atualizações
// @Summary Obtém todas as atualizações
// @Description Retorna uma lista de todas as atualizações
// @Tags updates
// @Produce json
// @Success 200 {array} types.Update
// @Failure 500 {object} utils.ErrorResponse
// @Router /updates [get]
func (h *Handler) handleGetUpdates(w http.ResponseWriter, r *http.Request) {
	updates, err := h.store.GetUpdates()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, updates)
}

// handleGetUpdateByID manipula as solicitações de obtenção de uma atualização por ID
// @Summary Obtém uma atualização por ID
// @Description Retorna uma atualização com base no ID fornecido
// @Tags updates
// @Produce json
// @Param id path int true "ID da atualização"
// @Success 200 {object} types.Update
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /updates/{id} [get]
func (h *Handler) handleGetUpdateByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid update ID"))
		return
	}

	update, err := h.store.GetUpdateByID(id)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	if update == nil {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("update not found"))
		return
	}

	utils.WriteJSON(w, http.StatusOK, update)
}

// handlePostUpdate manipula as solicitações de criação de uma nova atualização
// @Summary Cria uma nova atualização
// @Description Cria uma nova atualização com base nos dados fornecidos
// @Tags updates
// @Accept json
// @Produce json
// @Param update body types.UpdatePayload true "Dados da atualização"
// @Success 201 {string} string "update created successfully"
// @Failure 400 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /updates [post]
func (h *Handler) handlePostUpdate(w http.ResponseWriter, r *http.Request) {
	var payload types.UpdatePayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", err))
		return
	}

	update := types.Update{
		Title:       payload.Title,
		Description: payload.Description,
		Date:        payload.Date,
		SynergyID:   payload.SynergyID,
	}

	if err := h.store.CreateUpdate(&update); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, "update created successfully")
}

// handleUpdateUpdate manipula as solicitações de atualização de uma atualização
// @Summary Atualiza uma atualização existente
// @Description Atualiza uma atualização existente com base nos dados fornecidos
// @Tags updates
// @Accept json
// @Produce json
// @Param id path int true "ID da atualização"
// @Param update body types.UpdatePayload true "Dados da atualização"
// @Success 200 {string} string "update updated successfully"
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /updates/{id} [put]
func (h *Handler) handleUpdateUpdate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid update ID"))
		return
	}

	var payload types.UpdatePayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", err))
		return
	}

	update := types.Update{
		ID:          id,
		Title:       payload.Title,
		Description: payload.Description,
		Date:        payload.Date,
		SynergyID:   payload.SynergyID,
	}

	if err := h.store.UpdateUpdate(id, &update); err != nil {
		if err == sql.ErrNoRows {
			utils.WriteError(w, http.StatusNotFound, fmt.Errorf("update not found"))
		} else {
			utils.WriteError(w, http.StatusInternalServerError, err)
		}
		return
	}

	utils.WriteJSON(w, http.StatusOK, "update updated successfully")
}

// handleDeleteUpdate manipula as solicitações de exclusão de uma atualização
// @Summary Exclui uma atualização existente
// @Description Exclui uma atualização existente com base no ID fornecido
// @Tags updates
// @Produce json
// @Param id path int true "ID da atualização"
// @Success 200 {string} string "update deleted successfully"
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /updates/{id} [delete]
func (h *Handler) handleDeleteUpdate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid update ID"))
		return
	}

	if err := h.store.DeleteUpdate(id); err != nil {
		if err == sql.ErrNoRows {
			utils.WriteError(w, http.StatusNotFound, fmt.Errorf("update not found"))
		} else {
			utils.WriteError(w, http.StatusInternalServerError, err)
		}
		return
	}

	utils.WriteJSON(w, http.StatusOK, "update deleted successfully")
}

// handleGetUpdateByTitle manipula as solicitações de obtenção de uma atualização por título
// @Summary Obtém uma atualização por título
// @Description Retorna uma atualização com base no título fornecido
// @Tags updates
// @Produce json
// @Param title path string true "Título da atualização"
// @Success 200 {array} types.Update
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /updates/title/{title} [get]
func (h *Handler) handleGetUpdateByTitle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]

	updates, err := h.store.GetUpdateByTitle(title)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	if len(updates) == 0 {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("update not found"))
		return
	}

	utils.WriteJSON(w, http.StatusOK, updates)
}

