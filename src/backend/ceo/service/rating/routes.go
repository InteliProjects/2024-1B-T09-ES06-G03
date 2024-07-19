package rating

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Inteli-College/2024-1B-T09-ES06-G03/ceo/service/auth"
	"github.com/Inteli-College/2024-1B-T09-ES06-G03/ceo/types"
	"github.com/Inteli-College/2024-1B-T09-ES06-G03/ceo/utils"
	"github.com/gorilla/mux"
)

// Handler é o handler para os endpoints de ratings
type Handler struct {
	store types.RatingStore
}

// NewHandler cria um novo handler de ratings
func NewHandler(store types.RatingStore) *Handler {
	return &Handler{store: store}
}

// RegisterRoutes registra as rotas de ratings no roteador
func (h *Handler) RegisterRoutes(router *mux.Router) {
    // Protect the GET route for all ratings to require JWT
    router.Handle("/ratings", auth.JWTMiddleware(http.HandlerFunc(h.handleGetRatings))).Methods(http.MethodGet)

    // Route for getting ratings of the logged-in user, already properly protected
    router.Handle("/ratings/me", auth.JWTMiddleware(http.HandlerFunc(h.handleGetMyRatings))).Methods("GET")

    // Assuming POST for creating a rating is intended for authenticated users
    router.Handle("/ratings", auth.JWTMiddleware(http.HandlerFunc(h.handlePostRating))).Methods(http.MethodPost)

    // Protect the GET by ID to ensure only authenticated users can access
    router.Handle("/ratings/{id}", auth.JWTMiddleware(http.HandlerFunc(h.handleGetRatingByID))).Methods(http.MethodGet)

    // Protect the PUT for updating ratings to require JWT
    router.Handle("/ratings/{id}", auth.JWTMiddleware(http.HandlerFunc(h.handleUpdateRating))).Methods(http.MethodPut)

    // Protect the DELETE to require JWT
    router.Handle("/ratings/{id}", auth.JWTMiddleware(http.HandlerFunc(h.handleDeleteRating))).Methods(http.MethodDelete)
}

// handleGetRatings manipula as solicitações de obtenção de todas as avaliações
// @Summary Obtém todas as avaliações
// @Description Retorna uma lista de todas as avaliações
// @Tags ratings
// @Produce json
// @Success 200 {array} types.Rating
// @Failure 500 {object} utils.ErrorResponse
// @Router /ratings [get]
func (h *Handler) handleGetRatings(w http.ResponseWriter, r *http.Request) {
	ratings, err := h.store.GetRatings()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, ratings)
}

// handleGetRatingByID manipula as solicitações de obtenção de uma avaliação por ID
// @Summary Obtém uma avaliação por ID
// @Description Retorna uma avaliação com base no ID fornecido
// @Tags ratings
// @Produce json
// @Param id path int true "ID da avaliação"
// @Success 200 {object} types.Rating
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /ratings/{id} [get]
func (h *Handler) handleGetRatingByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid rating ID"))
		return
	}

	rating, err := h.store.GetRatingByID(id)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	if rating == nil {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("rating not found"))
		return
	}

	utils.WriteJSON(w, http.StatusOK, rating)
}

// handlePostRating manipula as solicitações de criação de uma nova avaliação
// @Summary Cria uma nova avaliação
// @Description Cria uma nova avaliação com base nos dados fornecidos
// @Tags ratings
// @Accept json
// @Produce json
// @Param rating body types.RatingPayload true "Dados da avaliação"
// @Success 201 {string} string "rating created successfully"
// @Failure 400 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /ratings [post]
func (h *Handler) handlePostRating(w http.ResponseWriter, r *http.Request) {
	var payload types.RatingPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", err))
		return
	}

	rating := types.Rating{
		Date:      payload.Date.Time,
		Level:     payload.Level,
		UserID:    payload.UserID,
		ProjectID: payload.ProjectID,
	}

	if err := h.store.CreateRating(&rating); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, "rating created successfully")
}

// handleUpdateRating manipula as solicitações de atualização de uma avaliação
// @Summary Atualiza uma avaliação existente
// @Description Atualiza uma avaliação existente com base nos dados fornecidos
// @Tags ratings
// @Accept json
// @Produce json
// @Param id path int true "ID da avaliação"
// @Param rating body types.RatingPayload true "Dados da avaliação"
// @Success 200 {string} string "rating updated successfully"
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /ratings/{id} [put]
func (h *Handler) handleUpdateRating(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid rating ID"))
		return
	}

	var payload types.RatingPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", err))
		return
	}

	rating := types.Rating{
		Date:      payload.Date.Time,
		Level:     payload.Level,
		UserID:    payload.UserID,
		ProjectID: payload.ProjectID,
	}

	if err := h.store.UpdateRating(id, &rating); err != nil {
		if err == sql.ErrNoRows {
			utils.WriteError(w, http.StatusNotFound, fmt.Errorf("rating not found"))
		} else {
			utils.WriteError(w, http.StatusInternalServerError, err)
		}
		return
	}

	utils.WriteJSON(w, http.StatusOK, "rating updated successfully")
}

// handleDeleteRating manipula as solicitações de exclusão de uma avaliação
// @Summary Exclui uma avaliação existente
// @Description Exclui uma avaliação existente com base no ID fornecido
// @Tags ratings
// @Produce json
// @Param id path int true "ID da avaliação"
// @Success 200 {string} string "rating deleted successfully"
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /ratings/{id} [delete]
func (h *Handler) handleDeleteRating(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid rating ID"))
		return
	}

	if err := h.store.DeleteRating(id); err != nil {
		if err == sql.ErrNoRows {
			utils.WriteError(w, http.StatusNotFound, fmt.Errorf("rating not found"))
		} else {
			utils.WriteError(w, http.StatusInternalServerError, err)
		}
		return
	}

	utils.WriteJSON(w, http.StatusOK, "rating deleted successfully")
}

func (h *Handler) handleGetMyRatings(w http.ResponseWriter, r *http.Request) {
    userID, ok := r.Context().Value("userID").(int)
    if !ok {
        http.Error(w, "User ID not found", http.StatusForbidden)
        return
    }

    projects, err := h.store.GetRatingsByUserID(userID)  
    if err != nil {
        utils.WriteError(w, http.StatusInternalServerError, err)
        return
    }

    utils.WriteJSON(w, http.StatusOK, projects)
}
