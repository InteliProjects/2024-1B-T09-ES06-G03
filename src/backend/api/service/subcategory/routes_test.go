package subcategory

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Inteli-College/2024-1B-T09-ES06-G03/service/auth"
	"github.com/Inteli-College/2024-1B-T09-ES06-G03/types"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

type MockSubcategoryStore struct {
	GetSubcategoriesFunc           func() ([]types.Subcategory, error)
	GetSubcategoryByIDFunc         func(int) (*types.Subcategory, error)
	CreateSubcategoryFunc          func(*types.Subcategory) error
	UpdateSubCategoryFunc          func(int, *types.Subcategory) error
	DeleteSubcategoryFunc          func(int) error
	GetSubcategoriesByCategoryFunc func(int) ([]types.Subcategory, error)
}

func (m *MockSubcategoryStore) GetSubcategories() ([]types.Subcategory, error) {
	return m.GetSubcategoriesFunc()
}

func (m *MockSubcategoryStore) GetSubcategoryByID(id int) (*types.Subcategory, error) {
	return m.GetSubcategoryByIDFunc(id)
}

func (m *MockSubcategoryStore) CreateSubcategory(c *types.Subcategory) error {
	return m.CreateSubcategoryFunc(c)
}

func (m *MockSubcategoryStore) UpdateSubCategory(id int, c *types.Subcategory) error {
	return m.UpdateSubCategoryFunc(id, c)
}

func (m *MockSubcategoryStore) DeleteSubcategory(id int) error {
	return m.DeleteSubcategoryFunc(id)
}

func (m *MockSubcategoryStore) GetSubcategoriesByCategory(categoryID int) ([]types.Subcategory, error) {
	return m.GetSubcategoriesByCategoryFunc(categoryID)
}

func TestSubcategoryGetHandler(t *testing.T) {
	mockStore := &MockSubcategoryStore{
		GetSubcategoriesFunc: func() ([]types.Subcategory, error) {
			return []types.Subcategory{{ID: 1, Name: "Subcategory 1", CategoryID: 1}}, nil
		},
		GetSubcategoryByIDFunc: func(id int) (*types.Subcategory, error) {
			if id == 1 {
				return &types.Subcategory{ID: 1, Name: "Subcategory 1", CategoryID: 1}, nil
			}
			return nil, fmt.Errorf("not found")
		},
	}

	handler := NewHandler(mockStore, nil)
	router := mux.NewRouter()

	router.HandleFunc("/subcategories", auth.WithMockJWTAuth(handler.handleGetSubcategories)).Methods("GET")
	router.HandleFunc("/subcategories/{id}", auth.WithMockJWTAuth(handler.handleGetSubcategoryByID)).Methods("GET")

	t.Run("Get Subcategories", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/subcategories", nil)
		assert.NoError(t, err)
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusOK, rr.Code)
	})

	t.Run("Get Subcategory by ID - Existing", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/subcategories/1", nil)
		assert.NoError(t, err)
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusOK, rr.Code)
	})
}

func TestSubcategoryPostHandler(t *testing.T) {
	mockStore := &MockSubcategoryStore{
		CreateSubcategoryFunc: func(c *types.Subcategory) error {
			if c.Name == "" {
				return fmt.Errorf("invalid subcategory data")
			}
			return nil
		},
	}

	handler := NewHandler(mockStore, nil) // Assuming nil for simplicity, replace with actual userStore if needed
	router := mux.NewRouter()
	router.HandleFunc("/subcategories", auth.WithMockJWTAuth(handler.handlePostSubcategory)).Methods("POST")

	t.Run("Create Subcategory - Valid Data", func(t *testing.T) {
		subcategory := types.Subcategory{Name: "New Subcategory", CategoryID: 1}
		jsonBody, _ := json.Marshal(subcategory)
		req, err := http.NewRequest("POST", "/subcategories", bytes.NewBuffer(jsonBody))
		req.Header.Set("Content-Type", "application/json")
		assert.NoError(t, err)

		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusCreated, rr.Code, "should create the subcategory and return status created")
	})

	t.Run("Create Subcategory - Invalid Data", func(t *testing.T) {
		subcategory := types.Subcategory{Name: "", CategoryID: 1}
		jsonBody, _ := json.Marshal(subcategory)
		req, err := http.NewRequest("POST", "/subcategories", bytes.NewBuffer(jsonBody))
		req.Header.Set("Content-Type", "application/json")
		assert.NoError(t, err)

		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusBadRequest, rr.Code, "should fail to create the subcategory due to invalid data")
	})
}

func TestSubcategoryUpdateHandler(t *testing.T) {
	mockStore := &MockSubcategoryStore{
		UpdateSubCategoryFunc: func(id int, c *types.Subcategory) error {
			if id == 999 {
				return fmt.Errorf("subcategory not found")
			}
			if c.Name == "" {
				return fmt.Errorf("invalid subcategory data")
			}
			return nil
		},
	}

	handler := NewHandler(mockStore, nil)
	router := mux.NewRouter()
	router.HandleFunc("/subcategories/{id}", auth.WithMockJWTAuth(handler.handleUpdateSubcategory)).Methods("PUT")

	t.Run("Update Subcategory - Valid Data", func(t *testing.T) {
		updatedSubcategory := types.Subcategory{Name: "Updated Subcategory", CategoryID: 1}
		jsonBody, _ := json.Marshal(updatedSubcategory)
		req, err := http.NewRequest("PUT", "/subcategories/1", bytes.NewBuffer(jsonBody))
		req.Header.Set("Content-Type", "application/json")
		assert.NoError(t, err)

		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusOK, rr.Code, "should update the subcategory and return status OK")
	})

	t.Run("Update Subcategory - Invalid Data", func(t *testing.T) {
		updatedSubcategory := types.Subcategory{Name: "", CategoryID: 1}
		jsonBody, _ := json.Marshal(updatedSubcategory)
		req, err := http.NewRequest("PUT", "/subcategories/1", bytes.NewBuffer(jsonBody))
		req.Header.Set("Content-Type", "application/json")
		assert.NoError(t, err)

		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusBadRequest, rr.Code, "should fail to update the subcategory due to invalid data")
	})
}

func TestSubcategoryGetByCategoryHandler(t *testing.T) {
	mockStore := &MockSubcategoryStore{
		GetSubcategoriesByCategoryFunc: func(categoryID int) ([]types.Subcategory, error) {
			if categoryID == 1 {
				return []types.Subcategory{
					{ID: 1, Name: "Subcategory 1", CategoryID: 1},
					{ID: 2, Name: "Subcategory 2", CategoryID: 1},
				}, nil
			}
			return nil, fmt.Errorf("no subcategories found for this category")
		},
	}

	handler := NewHandler(mockStore, nil)
	router := mux.NewRouter()

	router.HandleFunc("/categories/{category_id}/subcategories", auth.WithMockJWTAuth(handler.handleGetSubcategoriesByCategory)).Methods("GET")

	t.Run("Get Subcategories by Existing Category", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/categories/1/subcategories", nil)
		assert.NoError(t, err)
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusOK, rr.Code, "should return status OK and subcategories list for existing category")
	})
}
