package category

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Inteli-College/2024-1B-T09-ES06-G03/service/auth"
	"github.com/Inteli-College/2024-1B-T09-ES06-G03/types" // Make sure this is the correct import path
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

type MockCategoryStore struct {
	GetCategoriesFunc   func() ([]types.Category, error)
	GetCategoryByIDFunc func(int) (*types.Category, error)
	CreateCategoryFunc  func(*types.Category) error
	UpdateCategoryFunc  func(int, *types.Category) error
	DeleteCategoryFunc  func(int) error
}

func (m *MockCategoryStore) GetCategories() ([]types.Category, error) {
	return m.GetCategoriesFunc()
}

func (m *MockCategoryStore) GetCategoryByID(id int) (*types.Category, error) {
	return m.GetCategoryByIDFunc(id)
}

func (m *MockCategoryStore) CreateCategory(c *types.Category) error {
	return m.CreateCategoryFunc(c)
}

func (m *MockCategoryStore) UpdateCategory(id int, c *types.Category) error {
	return m.UpdateCategoryFunc(id, c)
}

func (m *MockCategoryStore) DeleteCategory(id int) error {
	return m.DeleteCategoryFunc(id)
}

func TestCategoryGetHandler(t *testing.T) {
	mockStore := &MockCategoryStore{
		GetCategoriesFunc: func() ([]types.Category, error) {
			return []types.Category{{ID: 1, Name: "Test"}}, nil
		},
	}

	handler := NewHandler(mockStore, nil)
	router := mux.NewRouter()

	router.HandleFunc("/categories", auth.WithMockJWTAuth(handler.handleGetCategories)).Methods("GET")

	t.Run("Get Categories", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/categories", nil)

		assert.NoError(t, err)
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusOK, rr.Code)
	})

	t.Run("Get Category by ID - Non-Existing", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/categories/2", nil)
		assert.NoError(t, err)
		rr := httptest.NewRecorder()

		router.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusNotFound, rr.Code, "should return not found for non-existing category")
	})
}

func TestCategoryPostHandler(t *testing.T) {
	mockStore := &MockCategoryStore{
		CreateCategoryFunc: func(c *types.Category) error {
			if c.Name == "" {
				return fmt.Errorf("invalid category data")
			}
			return nil // Simulate successful creation
		},
	}

	handler := NewHandler(mockStore, nil)
	router := mux.NewRouter()

	// Using mock JWT Auth for simplifying test setup
	router.HandleFunc("/categories", auth.WithMockJWTAuth(handler.handlePostCategories)).Methods("POST")

	t.Run("Create Category - Valid Data", func(t *testing.T) {
		newCategory := types.Category{Name: "New Category"}
		jsonBody, _ := json.Marshal(newCategory)
		req, err := http.NewRequest("POST", "/categories", bytes.NewBuffer(jsonBody))
		req.Header.Set("Content-Type", "application/json")
		assert.NoError(t, err)

		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusCreated, rr.Code, "should create the category and return status created")
	})

	t.Run("Create Category - Invalid Data", func(t *testing.T) {
		newCategory := types.Category{Name: ""}
		jsonBody, _ := json.Marshal(newCategory)
		req, err := http.NewRequest("POST", "/categories", bytes.NewBuffer(jsonBody))
		req.Header.Set("Content-Type", "application/json")
		assert.NoError(t, err)

		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusBadRequest, rr.Code, "should fail to create the category due to invalid data")
	})
}

func TestCategoryUpdateHandler(t *testing.T) {
	mockStore := &MockCategoryStore{
		UpdateCategoryFunc: func(id int, c *types.Category) error {
			if id == 1 {
				if c.Name == "" {
					return fmt.Errorf("invalid category data")
				}
				return nil
			}
			return fmt.Errorf("category not found")
		},
	}

	handler := NewHandler(mockStore, nil)
	router := mux.NewRouter()

	router.HandleFunc("/categories/{id}", auth.WithMockJWTAuth(handler.handleUpdateCategory)).Methods("PUT")

	t.Run("Update Category - Valid Data", func(t *testing.T) {
		updatedCategory := types.Category{Name: "Updated Category"}
		jsonBody, _ := json.Marshal(updatedCategory)
		req, err := http.NewRequest("PUT", "/categories/1", bytes.NewBuffer(jsonBody))
		req.Header.Set("Content-Type", "application/json")
		assert.NoError(t, err)

		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusOK, rr.Code, "should update the category and return status OK")
	})

	t.Run("Update Category - Invalid Data", func(t *testing.T) {
		updatedCategory := types.Category{Name: ""}
		jsonBody, _ := json.Marshal(updatedCategory)
		req, err := http.NewRequest("PUT", "/categories/1", bytes.NewBuffer(jsonBody))
		req.Header.Set("Content-Type", "application/json")
		assert.NoError(t, err)

		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusBadRequest, rr.Code, "should fail to update the category due to invalid data")
	})

}

func TestCategoryDeleteHandler(t *testing.T) {
	mockStore := &MockCategoryStore{
		DeleteCategoryFunc: func(id int) error {
			if id == 1 {
				return nil
			}
			return fmt.Errorf("category not found")
		},
	}

	handler := NewHandler(mockStore, nil)
	router := mux.NewRouter()

	router.HandleFunc("/categories/{id}", auth.WithMockJWTAuth(handler.handleDeleteCategory)).Methods("DELETE")

	t.Run("Delete Category - Existing", func(t *testing.T) {
		req, err := http.NewRequest("DELETE", "/categories/1", nil)
		assert.NoError(t, err)
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code, "should delete the category and return status OK")
	})
}
