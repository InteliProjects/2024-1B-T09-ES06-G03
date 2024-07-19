package synergy

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

type MockSynergyStore struct {
	GetSynergiesFunc   func() ([]types.Synergy, error)
	GetSynergyByIDFunc func(int) (*types.Synergy, error)
	CreateSynergyFunc  func(*types.Synergy) error
	UpdateSynergyFunc  func(int, *types.Synergy) error
	DeleteSynergyFunc  func(int) error
}

func (m *MockSynergyStore) GetSynergies() ([]types.Synergy, error) {
	return m.GetSynergiesFunc()
}

func (m *MockSynergyStore) GetSynergyByID(id int) (*types.Synergy, error) {
	return m.GetSynergyByIDFunc(id)
}

func (m *MockSynergyStore) CreateSynergy(s *types.Synergy) error {
	return m.CreateSynergyFunc(s)
}

func (m *MockSynergyStore) UpdateSynergy(id int, s *types.Synergy) error {
	return m.UpdateSynergyFunc(id, s)
}

func (m *MockSynergyStore) DeleteSynergy(id int) error {
	return m.DeleteSynergyFunc(id)
}

func TestSynergyGetHandler(t *testing.T) {
	mockStore := &MockSynergyStore{
		GetSynergiesFunc: func() ([]types.Synergy, error) {
			return []types.Synergy{{ID: 1, Description: "Test Synergy"}}, nil
		},
	}

	handler := NewHandler(mockStore)
	router := mux.NewRouter()

	router.HandleFunc("/synergies", handler.handleGetSynergies).Methods("GET")

	t.Run("Get Synergies", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/synergies", nil)
		assert.NoError(t, err)
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusOK, rr.Code)
	})

	t.Run("Get Synergy by ID - Non-Existing", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/synergies/2", nil)
		assert.NoError(t, err)
		rr := httptest.NewRecorder()

		router.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusNotFound, rr.Code, "should return not found for non-existing synergy")
	})
}

func TestSynergyUpdateHandler(t *testing.T) {
	mockStore := &MockSynergyStore{
		UpdateSynergyFunc: func(id int, s *types.Synergy) error {
			if id == 1 {
				if s.Description == "" {
					return fmt.Errorf("invalid synergy data")
				}
				return nil
			}
			return fmt.Errorf("synergy not found")
		},
	}

	handler := NewHandler(mockStore)
	router := mux.NewRouter()

	router.HandleFunc("/synergies/{id}", handler.handleUpdateSynergy).Methods("PUT")

	t.Run("Update Synergy - Valid Data", func(t *testing.T) {
		updatedSynergy := types.Synergy{SourceProjectID: 1, TargetProjectID: 2, Status: "active", Type: "collaboration", Description: "Updated Synergy"}
		jsonBody, _ := json.Marshal(updatedSynergy)
		req, err := http.NewRequest("PUT", "/synergies/1", bytes.NewBuffer(jsonBody))
		req.Header.Set("Content-Type", "application/json")
		assert.NoError(t, err)

		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusOK, rr.Code, "should update the synergy and return status OK")
	})
}

func TestSynergyDeleteHandler(t *testing.T) {
	mockStore := &MockSynergyStore{
		DeleteSynergyFunc: func(id int) error {
			if id == 1 {
				return nil
			}
			return fmt.Errorf("synergy not found")
		},
	}

	handler := NewHandler(mockStore)
	router := mux.NewRouter()

	router.HandleFunc("/synergies/{id}", handler.handleDeleteSynergy).Methods("DELETE")

	t.Run("Delete Synergy - Existing", func(t *testing.T) {
		req, err := http.NewRequest("DELETE", "/synergies/1", nil)
		assert.NoError(t, err)
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code, "should delete the synergy and return status OK")
	})
}
