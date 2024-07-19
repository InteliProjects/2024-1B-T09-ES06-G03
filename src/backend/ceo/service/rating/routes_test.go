package rating

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

type MockRatingStore struct {
	GetRatingsFunc    func() ([]types.Rating, error)
	GetRatingByIDFunc func(int) (*types.Rating, error)
	CreateRatingFunc  func(*types.Rating) error
	UpdateRatingFunc  func(int, *types.Rating) error
	DeleteRatingFunc  func(int) error
}

func (m *MockRatingStore) GetRatings() ([]types.Rating, error) {
	return m.GetRatingsFunc()
}

func (m *MockRatingStore) GetRatingByID(id int) (*types.Rating, error) {
	return m.GetRatingByIDFunc(id)
}

func (m *MockRatingStore) CreateRating(r *types.Rating) error {
	return m.CreateRatingFunc(r)
}

func (m *MockRatingStore) UpdateRating(id int, r *types.Rating) error {
	return m.UpdateRatingFunc(id, r)
}

func (m *MockRatingStore) DeleteRating(id int) error {
	return m.DeleteRatingFunc(id)
}

func TestRatingGetHandler(t *testing.T) {
	mockStore := &MockRatingStore{
		GetRatingsFunc: func() ([]types.Rating, error) {
			return []types.Rating{{ID: 1, Level: "5"}}, nil
		},
	}

	handler := NewHandler(mockStore)
	router := mux.NewRouter()

	router.HandleFunc("/ratings", handler.handleGetRatings).Methods("GET")

	t.Run("Get Ratings", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/ratings", nil)

		assert.NoError(t, err)
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusOK, rr.Code)
	})

	t.Run("Get Rating by ID - Non-Existing", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/ratings/2", nil)
		assert.NoError(t, err)
		rr := httptest.NewRecorder()

		router.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusNotFound, rr.Code, "should return not found for non-existing rating")
	})
}

func TestRatingPostHandler(t *testing.T) {
	mockStore := &MockRatingStore{
		CreateRatingFunc: func(r *types.Rating) error {
			if r.Level == "" {
				return fmt.Errorf("invalid rating data")
			}
			return nil // Simulate successful creation
		},
	}

	handler := NewHandler(mockStore)
	router := mux.NewRouter()

	router.HandleFunc("/ratings", handler.handlePostRating).Methods("POST")

	t.Run("Create Rating - Invalid Data", func(t *testing.T) {
		newRating := types.Rating{Level: ""}
		jsonBody, _ := json.Marshal(newRating)
		req, err := http.NewRequest("POST", "/ratings", bytes.NewBuffer(jsonBody))
		req.Header.Set("Content-Type", "application/json")
		assert.NoError(t, err)

		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusBadRequest, rr.Code, "should fail to create the rating due to invalid data")
	})
}

func TestRatingUpdateHandler(t *testing.T) {
	mockStore := &MockRatingStore{
		UpdateRatingFunc: func(id int, r *types.Rating) error {
			if id == 1 {
				if r.Level == "" {
					return fmt.Errorf("invalid rating data")
				}
				return nil
			}
			return fmt.Errorf("rating not found")
		},
	}

	handler := NewHandler(mockStore)
	router := mux.NewRouter()

	router.HandleFunc("/ratings/{id}", handler.handleUpdateRating).Methods("PUT")

	t.Run("Update Rating - Invalid Data", func(t *testing.T) {
		updatedRating := types.Rating{Level: ""}
		jsonBody, _ := json.Marshal(updatedRating)
		req, err := http.NewRequest("PUT", "/ratings/1", bytes.NewBuffer(jsonBody))
		req.Header.Set("Content-Type", "application/json")
		assert.NoError(t, err)

		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusBadRequest, rr.Code, "should fail to update the rating due to invalid data")
	})
}

func TestRatingDeleteHandler(t *testing.T) {
	mockStore := &MockRatingStore{
		DeleteRatingFunc: func(id int) error {
			if id == 1 {
				return nil
			}
			return fmt.Errorf("rating not found")
		},
	}

	handler := NewHandler(mockStore)
	router := mux.NewRouter()

	router.HandleFunc("/ratings/{id}", handler.handleDeleteRating).Methods("DELETE")

	t.Run("Delete Rating - Existing", func(t *testing.T) {
		req, err := http.NewRequest("DELETE", "/ratings/1", nil)
		assert.NoError(t, err)
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code, "should delete the rating and return status OK")
	})
}
