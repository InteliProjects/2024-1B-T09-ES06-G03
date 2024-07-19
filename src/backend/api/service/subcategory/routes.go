package subcategory

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Inteli-College/2024-1B-T09-ES06-G03/service/auth"
	"github.com/Inteli-College/2024-1B-T09-ES06-G03/types"
	"github.com/Inteli-College/2024-1B-T09-ES06-G03/utils"
	"github.com/gorilla/mux"
)

type Handler struct {
	store     types.SubcategoryStore
	userStore types.UserStore
}

func NewHandler(store types.SubcategoryStore, userStore types.UserStore) *Handler {
	return &Handler{store: store, userStore: userStore}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/subcategories", auth.WithJWTAuth(h.handleGetSubcategories, h.userStore)).Methods(http.MethodGet)
	router.HandleFunc("/subcategories/{id}", auth.WithJWTAuth(h.handleGetSubcategoryByID, h.userStore)).Methods(http.MethodGet)
	router.HandleFunc("/subcategories", auth.WithJWTAuth(h.handlePostSubcategory, h.userStore)).Methods(http.MethodPost)
	router.HandleFunc("/subcategories/{id}", auth.WithJWTAuth(h.handleUpdateSubcategory, h.userStore)).Methods(http.MethodPut)
	router.HandleFunc("/subcategories/{id}", auth.WithJWTAuth(h.handleDeleteSubcategory, h.userStore)).Methods(http.MethodDelete)
	router.HandleFunc("/categories/{category_id}/subcategories", auth.WithJWTAuth(h.handleGetSubcategoriesByCategory, h.userStore)).Methods(http.MethodGet)
}

// @Summary List all subcategories
// @Description Retrieves a list of all subcategories
// @Tags subcategories
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {array} types.Subcategory
// @Failure 500 {string} string "Internal server error"
// @Router /subcategories [get]
func (h *Handler) handleGetSubcategories(w http.ResponseWriter, r *http.Request) {
	subcategories, err := h.store.GetSubcategories()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, subcategories)
}

// @Summary Get a subcategory by ID
// @Description Retrieves a subcategory by its ID
// @Tags subcategories
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Subcategory ID"
// @Success 200 {object} types.Subcategory
// @Failure 404 {string} string "Subcategory not found"
// @Failure 400 {string} string "Invalid subcategory ID"
// @Failure 500 {string} string "Internal server error"
// @Router /subcategories/{id} [get]
func (h *Handler) handleGetSubcategoryByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid subcategory ID"))
		return
	}

	subcategory, err := h.store.GetSubcategoryByID(id)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	if subcategory == nil {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("subcategory not found"))
		return
	}

	utils.WriteJSON(w, http.StatusOK, subcategory)
}

// @Summary Get subcategories by category ID
// @Description Retrieves subcategories associated with a category ID
// @Tags subcategories
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param category_id path int true "Category ID"
// @Success 200 {array} types.Subcategory
// @Failure 404 {string} string "Category not found"
// @Failure 400 {string} string "Invalid category ID"
// @Failure 500 {string} string "Internal server error"
// @Router /categories/{category_id}/subcategories [get]
func (h *Handler) handleGetSubcategoriesByCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	categoryID, err := strconv.Atoi(vars["category_id"])
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid category ID"))
		return
	}

	subcategories, err := h.store.GetSubcategoriesByCategory(categoryID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, subcategories)
}

// @Summary Create a subcategory
// @Description Adds a new subcategory to the system
// @Tags subcategories
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param subcategory body types.SubcategoryPayload true "Subcategory information"
// @Success 201 {string} string "subcategory created successfully"
// @Failure 400 {string} string "Invalid payload"
// @Failure 500 {string} string "Internal server error"
// @Router /subcategories [post]
func (h *Handler) handlePostSubcategory(w http.ResponseWriter, r *http.Request) {
	var payload types.SubcategoryPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload"))
		return
	}

	subcategory := types.Subcategory{Name: payload.Name, CategoryID: payload.CategoryID}
	if err := h.store.CreateSubcategory(&subcategory); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, "subcategory created successfully")
}

// @Summary Update a subcategory
// @Description Updates an existing subcategory by ID
// @Tags subcategories
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Subcategory ID"
// @Param subcategory body types.SubcategoryPayload true "Subcategory information"
// @Success 200 {string} string "subcategory updated successfully"
// @Failure 404 {string} string "Subcategory not found"
// @Failure 400 {string} string "Invalid subcategory ID or payload"
// @Failure 500 {string} string "Internal server error"
// @Router /subcategories/{id} [put]
func (h *Handler) handleUpdateSubcategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid subcategory ID"))
		return
	}

	var payload types.SubcategoryPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload"))
		return
	}

	subcategory := types.Subcategory{ID: id, Name: payload.Name}
	if err := h.store.UpdateSubCategory(id, &subcategory); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, "subcategory updated successfully")
}

// @Summary Delete a subcategory
// @Description Deletes a subcategory by ID
// @Tags subcategories
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Subcategory ID"
// @Success 200 {string} string "subcategory deleted successfully"
// @Failure 404 {string} string "Subcategory not found"
// @Failure 400 {string} string "Invalid subcategory ID"
// @Failure 500 {string} string "Internal server error"
// @Router /subcategories/{id} [delete]
func (h *Handler) handleDeleteSubcategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid subcategory ID"))
		return
	}

	if err := h.store.DeleteSubcategory(id); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, "subcategory deleted successfully")
}
