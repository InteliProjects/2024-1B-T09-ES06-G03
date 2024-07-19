package category

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Inteli-College/2024-1B-T09-ES06-G03/service/auth"
	"github.com/Inteli-College/2024-1B-T09-ES06-G03/types"
	"github.com/Inteli-College/2024-1B-T09-ES06-G03/utils"
	"github.com/gorilla/mux"
)

type Handler struct {
	store     types.CategoryStore
	userStore types.UserStore
}

func NewHandler(store types.CategoryStore, userStore types.UserStore) *Handler {
	return &Handler{store: store, userStore: userStore}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/categories", auth.WithJWTAuth(h.handleGetCategories, h.userStore)).Methods(http.MethodGet)
	router.HandleFunc("/categories/{id}", auth.WithJWTAuth(h.handleGetCategoryByID, h.userStore)).Methods(http.MethodGet)
	router.HandleFunc("/categories", auth.WithJWTAuth(h.handlePostCategories, h.userStore)).Methods(http.MethodPost)
	router.HandleFunc("/categories/{id}", auth.WithJWTAuth(h.handleUpdateCategory, h.userStore)).Methods(http.MethodPut)
	router.HandleFunc("/categories/{id}", auth.WithJWTAuth(h.handleDeleteCategory, h.userStore)).Methods(http.MethodDelete)
}

// @Summary List all categories
// @Description Retrieves a list of all categories
// @Tags categories
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {array} types.Category
// @Failure 500 {string} string "Internal server error"
// @Router /categories [get]
func (h *Handler) handleGetCategories(w http.ResponseWriter, r *http.Request) {
	categories, err := h.store.GetCategories()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, categories)
}

// @Summary Create a category
// @Description Adds a new category to the system
// @Tags categories
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param category body types.CategoryPayload true "Category information"
// @Success 201 {string} string "category created successfully"
// @Failure 400 {string} string "Invalid payload"
// @Failure 500 {string} string "Internal server error"
// @Router /categories [post]
func (h *Handler) handlePostCategories(w http.ResponseWriter, r *http.Request) {
	var payload types.CategoryPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", payload))
		return
	}

	category := types.Category{Name: payload.Name}
	if err := h.store.CreateCategory(&category); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, "category created successfully")
}

// @Summary Update a category
// @Description Updates an existing category by ID
// @Tags categories
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Category ID"
// @Param category body types.CategoryPayload true "Category information"
// @Success 200 {string} string "category updated successfully"
// @Failure 404 {string} string "Category not found"
// @Failure 400 {string} string "Invalid category ID or payload"
// @Failure 500 {string} string "Internal server error"
// @Router /categories/{id} [put]
func (h *Handler) handleUpdateCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid category ID: %v", err))
		return
	}

	var payload types.CategoryPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", err))
		return
	}

	category := types.Category{ID: id, Name: payload.Name}
	err = h.store.UpdateCategory(id, &category)
	if err != nil {
		if err == sql.ErrNoRows {
			utils.WriteError(w, http.StatusNotFound, fmt.Errorf("category not found"))
		} else {
			utils.WriteError(w, http.StatusInternalServerError, err)
		}
		return
	}

	utils.WriteJSON(w, http.StatusOK, "category updated successfully")
}

// @Summary Delete a category
// @Description Deletes a category by ID
// @Tags categories
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Category ID"
// @Success 200 {string} string "category deleted successfully"
// @Failure 404 {string} string "Category not found"
// @Failure 400 {string} string "Invalid category ID"
// @Failure 500 {string} string "Internal server error"
// @Router /categories/{id} [delete]
func (h *Handler) handleDeleteCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid category ID"))
		return
	}

	if err := h.store.DeleteCategory(id); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, "category deleted successfully")
}

// @Summary Get a category by ID
// @Description Retrieves a category by its ID
// @Tags categories
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Category ID"
// @Success 200 {object} types.Category
// @Failure 404 {string} string "Category not found"
// @Failure 400 {string} string "Invalid category ID"
// @Failure 500 {string} string "Internal server error"
// @Router /categories/{id} [get]
func (h *Handler) handleGetCategoryByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid category ID"))
		return
	}

	category, err := h.store.GetCategoryByID(id)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	if category == nil {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("category not found"))
		return
	}

	utils.WriteJSON(w, http.StatusOK, category)
}
