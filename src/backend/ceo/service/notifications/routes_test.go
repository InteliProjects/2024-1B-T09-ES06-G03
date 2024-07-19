package notifications

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Inteli-College/2024-1B-T09-ES06-G03/ceo/types"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

type MockNotificationStore struct {
	GetNotificationsFunc    func() ([]types.Notification, error)
	GetNotificationByIDFunc func(int) (*types.Notification, error)
	CreateNotificationFunc  func(*types.Notification) error
	UpdateNotificationFunc  func(int, *types.Notification) error
	DeleteNotificationFunc  func(int) error
}

func (m *MockNotificationStore) GetNotifications() ([]types.Notification, error) {
	return m.GetNotificationsFunc()
}

func (m *MockNotificationStore) GetNotificationByID(id int) (*types.Notification, error) {
	return m.GetNotificationByIDFunc(id)
}

func (m *MockNotificationStore) CreateNotification(c *types.Notification) error {
	return m.CreateNotificationFunc(c)
}

func (m *MockNotificationStore) UpdateNotification(id int, c *types.Notification) error {
	return m.UpdateNotificationFunc(id, c)
}

func (m *MockNotificationStore) DeleteNotification(id int) error {
	return m.DeleteNotificationFunc(id)
}

func TestNotificationGetHandler(t *testing.T) {
	mockStore := &MockNotificationStore{
		GetNotificationsFunc: func() ([]types.Notification, error) {
			return []types.Notification{{ID: 1, Title: "Test"}}, nil
		},
	}

	handler := NewHandler(mockStore)
	router := mux.NewRouter()

	router.HandleFunc("/notifications", handler.handleGetNotifications).Methods("GET")

	t.Run("Get Notifications", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/notifications", nil)

		assert.NoError(t, err)
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusOK, rr.Code)
	})

	t.Run("Get Notification by ID - Non-Existing", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/notifications/2", nil)
		assert.NoError(t, err)
		rr := httptest.NewRecorder()

		router.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusNotFound, rr.Code, "should return not found for non-existing notification")
	})
}

func TestNotificationPostHandler(t *testing.T) {
	mockStore := &MockNotificationStore{
		CreateNotificationFunc: func(c *types.Notification) error {
			if c.Title == "" {
				return fmt.Errorf("invalid notification data")
			}
			return nil // Simulate successful creation
		},
	}

	handler := NewHandler(mockStore)
	router := mux.NewRouter()

	router.HandleFunc("/notifications", handler.handlePostNotification).Methods("POST")

	t.Run("Create Notification - Invalid Data", func(t *testing.T) {
		newNotification := types.Notification{Title: ""}
		jsonBody, _ := json.Marshal(newNotification)
		req, err := http.NewRequest("POST", "/notifications", bytes.NewBuffer(jsonBody))
		req.Header.Set("Content-Type", "application/json")
		assert.NoError(t, err)

		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusBadRequest, rr.Code, "should fail to create the notification due to invalid data")
	})
}

func TestNotificationUpdateHandler(t *testing.T) {
	mockStore := &MockNotificationStore{
		UpdateNotificationFunc: func(id int, c *types.Notification) error {
			if id == 1 {
				if c.Title == "" {
					return fmt.Errorf("invalid notification data")
				}
				return nil
			}
			return fmt.Errorf("notification not found")
		},
	}

	handler := NewHandler(mockStore)
	router := mux.NewRouter()

	router.HandleFunc("/notifications/{id}", handler.handleUpdateNotification).Methods("PUT")

	t.Run("Update Notification - Invalid Data", func(t *testing.T) {
		updatedNotification := types.Notification{Title: ""}
		jsonBody, _ := json.Marshal(updatedNotification)
		req, err := http.NewRequest("PUT", "/notifications/1", bytes.NewBuffer(jsonBody))
		req.Header.Set("Content-Type", "application/json")
		assert.NoError(t, err)

		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusBadRequest, rr.Code, "should fail to update the notification due to invalid data")
	})
}

func TestNotificationDeleteHandler(t *testing.T) {
	mockStore := &MockNotificationStore{
		DeleteNotificationFunc: func(id int) error {
			if id == 1 {
				return nil
			}
			return fmt.Errorf("notification not found")
		},
	}

	handler := NewHandler(mockStore)
	router := mux.NewRouter()

	router.HandleFunc("/notifications/{id}", handler.handleDeleteNotification).Methods("DELETE")

	t.Run("Delete Notification - Existing", func(t *testing.T) {
		req, err := http.NewRequest("DELETE", "/notifications/1", nil)
		assert.NoError(t, err)
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code, "should delete the notification and return status OK")
	})
}
