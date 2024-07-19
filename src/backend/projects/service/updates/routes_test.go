package update

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Inteli-College/2024-1B-T09-ES06-G03/project/types"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

type MockUpdateStore struct {
	GetUpdatesFunc    func() ([]types.Update, error)
	GetUpdateByIDFunc func(int) (*types.Update, error)
	CreateUpdateFunc  func(*types.Update) error
	UpdateUpdateFunc  func(int, *types.Update) error
	DeleteUpdateFunc  func(int) error
}

func (m *MockUpdateStore) GetUpdates() ([]types.Update, error) {
	return m.GetUpdatesFunc()
}

func (m *MockUpdateStore) GetUpdateByID(id int) (*types.Update, error) {
	return m.GetUpdateByIDFunc(id)
}

func (m *MockUpdateStore) CreateUpdate(u *types.Update) error {
	return m.CreateUpdateFunc(u)
}

func (m *MockUpdateStore) UpdateUpdate(id int, u *types.Update) error {
	return m.UpdateUpdateFunc(id, u)
}

func (m *MockUpdateStore) DeleteUpdate(id int) error {
	return m.DeleteUpdateFunc(id)
}

func TestUpdateGetHandler(t *testing.T) {
	mockStore := &MockUpdateStore{
		GetUpdatesFunc: func() ([]types.Update, error) {
			return []types.Update{{ID: 1, Title: "Test Update"}}, nil
		},
	}

	handler := NewHandler(mockStore)
	router := mux.NewRouter()

	router.HandleFunc("/updates", handler.handleGetUpdates).Methods("GET")

	t.Run("Get Updates", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/updates", nil)
		assert.NoError(t, err)
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusOK, rr.Code)
	})

	t.Run("Get Update by ID - Non-Existing", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/updates/2", nil)
		assert.NoError(t, err)
		rr := httptest.NewRecorder()

		router.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusNotFound, rr.Code, "should return not found for non-existing update")
	})
}

func TestUpdateUpdateHandler(t *testing.T) {
	mockStore := &MockUpdateStore{
		UpdateUpdateFunc: func(id int, u *types.Update) error {
			if id == 1 {
				if u.Title == "" {
					return fmt.Errorf("invalid update data")
				}
				return nil
			}
			return fmt.Errorf("update not found")
		},
	}

	handler := NewHandler(mockStore)
	router := mux.NewRouter()

	router.HandleFunc("/updates/{id}", handler.handleUpdateUpdate).Methods("PUT")

	t.Run("Update Update - Valid Data", func(t *testing.T) {
		updatedUpdate := types.Update{Title: "Updated Update", Description: "Updated Description", Date: "2023-01-01", SynergyID: 1}
		jsonBody, _ := json.Marshal(updatedUpdate)
		req, err := http.NewRequest("PUT", "/updates/1", bytes.NewBuffer(jsonBody))
		req.Header.Set("Content-Type", "application/json")
		assert.NoError(t, err)

		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusOK, rr.Code, "should update the update and return status OK")
	})
}

func TestUpdateDeleteHandler(t *testing.T) {
	mockStore := &MockUpdateStore{
		DeleteUpdateFunc: func(id int) error {
			if id == 1 {
				return nil
			}
			return fmt.Errorf("update not found")
		},
	}

	handler := NewHandler(mockStore)
	router := mux.NewRouter()

	router.HandleFunc("/updates/{id}", handler.handleDeleteUpdate).Methods("DELETE")

	t.Run("Delete Update - Existing", func(t *testing.T) {
		req, err := http.NewRequest("DELETE", "/updates/1", nil)
		assert.NoError(t, err)
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code, "should delete the update and return status OK")
	})
}
