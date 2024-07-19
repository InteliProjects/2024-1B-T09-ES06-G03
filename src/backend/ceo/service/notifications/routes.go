package notifications

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

// Handler é o handler para os endpoints de notifications
type Handler struct {
	store types.NotificationStore
}

// NewHandler cria um novo handler de notifications
func NewHandler(store types.NotificationStore) *Handler {
	return &Handler{store: store}
}

// RegisterRoutes registra as rotas de notifications no roteador
func (h *Handler) RegisterRoutes(router *mux.Router) {
    // Protect the GET route for all notifications to require JWT
    router.Handle("/notifications", auth.JWTMiddleware(http.HandlerFunc(h.handleGetNotifications))).Methods(http.MethodGet)

    // Specific route for getting notifications for the logged-in user
    router.Handle("/notifications/me", auth.JWTMiddleware(http.HandlerFunc(h.handleGetMyNotifications))).Methods("GET")

    // Assuming POST for creating a notification does not necessarily require a user context
    router.HandleFunc("/notifications", h.handlePostNotification).Methods(http.MethodPost)

    // Protect the GET by ID to ensure only authenticated users can access
    router.Handle("/notifications/{id}", auth.JWTMiddleware(http.HandlerFunc(h.handleGetNotificationsByID))).Methods(http.MethodGet)

    // Protect the PUT for updating notifications to require JWT
    router.Handle("/notifications/{id}", auth.JWTMiddleware(http.HandlerFunc(h.handleUpdateNotification))).Methods(http.MethodPut)

    // Protect the DELETE to require JWT
    router.Handle("/notifications/{id}", auth.JWTMiddleware(http.HandlerFunc(h.handleDeleteNotification))).Methods(http.MethodDelete)

    // Protect the GET by title to ensure only authenticated users can access
    router.Handle("/notifications/title/{title}", auth.JWTMiddleware(http.HandlerFunc(h.handleGetNotificationByTitle))).Methods(http.MethodGet)

    // Protect the route to get notifications by user ID
    router.Handle("/notifications/user/{userId}", auth.JWTMiddleware(http.HandlerFunc(h.handleGetNotificationsByUser))).Methods(http.MethodGet)
}

// handleGetNotifications manipula as solicitações de obtenção de notificações
// @Summary Obtém todas as notificações
// @Description Retorna uma lista de todas as notificações
// @Tags notifications
// @Produce json
// @Success 200 {array} types.Notification
// @Failure 500 {object} utils.ErrorResponse
// @Router /notifications [get]
func (h *Handler) handleGetNotifications(w http.ResponseWriter, r *http.Request) {
	notifications, err := h.store.GetNotifications()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, notifications)
}

// handleGetNotificationsByID manipula as solicitações de obtenção de uma notificação por ID
// @Summary Obtém uma notificação por ID
// @Description Retorna uma notificação com base no ID fornecido
// @Tags notifications
// @Produce json
// @Param id path int true "ID da notificação"
// @Success 200 {object} types.Notification
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /notifications/{id} [get]
func (h *Handler) handleGetNotificationsByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid notification ID"))
		return
	}

	notification, err := h.store.GetNotificationByID(id)
	if err == sql.ErrNoRows {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("notification not found"))
		return
	} else if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, notification)
}

// handlePostNotification manipula as solicitações de criação de uma nova notificação
// @Summary Cria uma nova notificação
// @Description Cria uma nova notificação com base nos dados fornecidos
// @Tags notifications
// @Accept json
// @Produce json
// @Param notification body types.NotificationPayload true "Dados da notificação"
// @Success 201 {string} string "notification created successfully"
// @Failure 400 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /notifications [post]
func (h *Handler) handlePostNotification(w http.ResponseWriter, r *http.Request) {
	var payload types.NotificationPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", err))
		return
	}

	notification := types.Notification{
		ReceivedUserID:   payload.ReceivedUserID,
		SentUserID:       payload.SentUserID,
		ReceivedProjectID: payload.ReceivedProjectID,
		SentProjectID:     payload.SentProjectID,
		SynergyType:     payload.SynergyType,
		Type:            payload.Type,
		Title:           payload.Title,
		Message:         payload.Message,
		Status:          payload.Status,
		CreatedAt:       payload.CreatedAt.Time,
	}

	if err := h.store.CreateNotification(&notification); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, "notification created successfully")
}

// handleUpdateNotification manipula as solicitações de atualização de uma notificação
// @Summary Atualiza uma notificação existente
// @Description Atualiza uma notificação existente com base nos dados fornecidos
// @Tags notifications
// @Accept json
// @Produce json
// @Param id path int true "ID da notificação"
// @Param notification body types.NotificationPayload true "Dados da notificação"
// @Success 200 {string} string "notification updated successfully"
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /notifications/{id} [put]
func (h *Handler) handleUpdateNotification(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid notification ID"))
		return
	}

	var payload types.NotificationPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", err))
		return
	}

	notification := types.Notification{
		ID:               id,
		ReceivedUserID:   payload.ReceivedUserID,
		SentUserID:       payload.SentUserID,
		ReceivedProjectID: payload.ReceivedProjectID,
		SentProjectID:     payload.SentProjectID,
		SynergyType:     payload.SynergyType,
		Type:            payload.Type,
		Title:           payload.Title,
		Message:         payload.Message,
		Status:          payload.Status,
		CreatedAt:       payload.CreatedAt.Time,
	}

	if err := h.store.UpdateNotification(id, &notification); err != nil {
		if err == sql.ErrNoRows {
			utils.WriteError(w, http.StatusNotFound, fmt.Errorf("notification not found"))
		} else {
			utils.WriteError(w, http.StatusInternalServerError, err)
		}
		return
	}

	utils.WriteJSON(w, http.StatusOK, "notification updated successfully")
}

// handleDeleteNotification manipula as solicitações de exclusão de uma notificação
// @Summary Exclui uma notificação existente
// @Description Exclui uma notificação existente com base no ID fornecido
// @Tags notifications
// @Produce json
// @Param id path int true "ID da notificação"
// @Success 200 {string} string "notification deleted successfully"
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /notifications/{id} [delete]
func (h *Handler) handleDeleteNotification(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid notification ID"))
		return
	}

	err = h.store.DeleteNotification(id)
	if err == sql.ErrNoRows {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("notification not found"))
		return
	} else if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, "notification deleted successfully")
}

// handleGetNotificationByTitle manipula as solicitações de obtenção de uma notificação por título
// @Summary Obtém uma notificação por título
// @Description Retorna uma notificação com base no título fornecido
// @Tags notifications
// @Produce json
// @Param title path string true "Título da notificação"
// @Success 200 {object} types.Notification
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /notifications/title/{title} [get]
func (h *Handler) handleGetNotificationByTitle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]

	notification, err := h.store.GetNotificationByTitle(title)
	if err == sql.ErrNoRows {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("notification not found"))
		return
	} else if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, notification)
}

// handleGetNotificationsByUser manipula as solicitações de obtenção de notificações por usuário
// @Summary Obtém notificações por usuário
// @Description Retorna uma lista de notificações de um usuário específico, incluindo detalhes do projeto, subcategoria e categoria
// @Tags notifications
// @Produce json
// @Param userId path int true "ID do usuário"
// @Success 200 {array} types.DetailedNotification
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /notifications/user/{userId} [get]
func (h *Handler) handleGetNotificationsByUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId, err := strconv.Atoi(vars["userId"])
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid user ID"))
		return
	}

	notifications, err := h.store.GetNotificationsByUser(userId)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	if len(notifications) == 0 {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("no notifications found for the given user ID"))
		return
	}

	utils.WriteJSON(w, http.StatusOK, notifications)
}

func (h *Handler) handleGetMyNotifications(w http.ResponseWriter, r *http.Request) {
    userID, ok := r.Context().Value("userID").(int)
    if !ok {
        http.Error(w, "User ID not found", http.StatusForbidden)
        return
    }

    projects, err := h.store.GetNotificationsByUser(userID)  
    if err != nil {
        utils.WriteError(w, http.StatusInternalServerError, err)
        return
    }

    utils.WriteJSON(w, http.StatusOK, projects)
}
