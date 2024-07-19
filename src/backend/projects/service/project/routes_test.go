package project

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Inteli-College/2024-1B-T09-ES06-G03/project/types"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

type MockProjectStore struct {
	GetProjectsFunc        func() ([]types.Project, error)
	GetProjectByIDFunc     func(int) (*types.Project, error)
	CreateProjectFunc      func(*types.Project) error
	UpdateProjectFunc      func(int, *types.Project) error
	DeleteProjectFunc      func(int) error
	GetProjectsByNameFunc  func(string) ([]types.Project, error)
	GetProjectsByCeoIDFunc func(int) ([]types.Project, error)
	GetEvaluationsFunc     func() ([]types.Evaluation, error)
}

func (m *MockProjectStore) GetProjects() ([]types.Project, error) {
	return m.GetProjectsFunc()
}

func (m *MockProjectStore) GetProjectByID(id int) (*types.Project, error) {
	return m.GetProjectByIDFunc(id)
}

func (m *MockProjectStore) CreateProject(p *types.Project) error {
	return m.CreateProjectFunc(p)
}

func (m *MockProjectStore) UpdateProject(id int, p *types.Project) error {
	return m.UpdateProjectFunc(id, p)
}

func (m *MockProjectStore) DeleteProject(id int) error {
	return m.DeleteProjectFunc(id)
}

func (m *MockProjectStore) GetProjectsByName(name string) ([]types.Project, error) {
	return m.GetProjectsByNameFunc(name)
}

func (m *MockProjectStore) GetProjectsByCeoID(ceoId int) ([]types.Project, error) {
	return m.GetProjectsByCeoIDFunc(ceoId)
}

func (m *MockProjectStore) GetEvaluations() ([]types.Evaluation, error) {
	return m.GetEvaluationsFunc()
}

func TestProjectGetHandler(t *testing.T) {
	mockStore := &MockProjectStore{
		GetProjectsFunc: func() ([]types.Project, error) {
			return []types.Project{{ID: 1, Name: "Test Project", CategoryID: 1, SubcategoryID: 1}}, nil
		},
	}

	handler := NewHandler(mockStore)
	router := mux.NewRouter()

	router.HandleFunc("/projects", handler.handleGetProjects).Methods("GET")

	t.Run("Get Projects", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/projects", nil)
		assert.NoError(t, err)
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusOK, rr.Code)
	})

	t.Run("Get Project by ID - Non-Existing", func(t *testing.T) {
		mockStore.GetProjectByIDFunc = func(id int) (*types.Project, error) {
			return nil, fmt.Errorf("project not found")
		}
		req, err := http.NewRequest("GET", "/projects/2", nil)
		assert.NoError(t, err)
		rr := httptest.NewRecorder()

		router.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusNotFound, rr.Code, "should return not found for non-existing project")
	})
}

func TestProjectUpdateHandler(t *testing.T) {
	mockStore := &MockProjectStore{
		UpdateProjectFunc: func(id int, p *types.Project) error {
			if id == 1 {
				if p.Name == "" {
					return fmt.Errorf("invalid project data")
				}
				return nil
			}
			return fmt.Errorf("project not found")
		},
	}

	handler := NewHandler(mockStore)
	router := mux.NewRouter()

	router.HandleFunc("/projects/{id}", handler.handleUpdateProject).Methods("PUT")

	t.Run("Update Project - Valid Data", func(t *testing.T) {
		updatedProject := types.Project{Name: "Updated Project", Description: "Updated Description", Status: "active", UserID: 1, CategoryID: 1, SubcategoryID: 1, Photo: "photo.png"}
		jsonBody, _ := json.Marshal(updatedProject)
		req, err := http.NewRequest("PUT", "/projects/1", bytes.NewBuffer(jsonBody))
		req.Header.Set("Content-Type", "application/json")
		assert.NoError(t, err)

		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusOK, rr.Code, "should update the project and return status OK")
	})
}

func TestProjectDeleteHandler(t *testing.T) {
	mockStore := &MockProjectStore{
		DeleteProjectFunc: func(id int) error {
			if id == 1 {
				return nil
			}
			return fmt.Errorf("project not found")
		},
	}

	handler := NewHandler(mockStore)
	router := mux.NewRouter()

	router.HandleFunc("/projects/{id}", handler.handleDeleteProject).Methods("DELETE")

	t.Run("Delete Project - Existing", func(t *testing.T) {
		req, err := http.NewRequest("DELETE", "/projects/1", nil)
		assert.NoError(t, err)
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code, "should delete the project and return status OK")
	})
}

func TestPredictProjectsHandler(t *testing.T) {
	mockStore := &MockProjectStore{
		GetProjectsFunc: func() ([]types.Project, error) {
			return []types.Project{
				{ID: 1, Name: "Project 1", CreatedAt: types.CustomDate{Time: time.Now()}, UpdatedAt: types.CustomDate{Time: time.Now()}},
				{ID: 2, Name: "Project 2", CreatedAt: types.CustomDate{Time: time.Now()}, UpdatedAt: types.CustomDate{Time: time.Now()}},
				{ID: 3, Name: "Project 3", CreatedAt: types.CustomDate{Time: time.Now()}, UpdatedAt: types.CustomDate{Time: time.Now()}},
			}, nil
		},
		GetEvaluationsFunc: func() ([]types.Evaluation, error) {
			return []types.Evaluation{
				{IDProponente: 1, IDProjeto: 1, Avaliacao: 2},
				{IDProponente: 2, IDProjeto: 2, Avaliacao: 2},
				{IDProponente: 3, IDProjeto: 1, Avaliacao: -2},
				{IDProponente: 3, IDProjeto: 3, Avaliacao: 2},
			}, nil
		},
	}

	handler := NewHandler(mockStore)
	router := mux.NewRouter()

	t.Run("Predict Projects - Valid Request", func(t *testing.T) {
		// Mock response from external service
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			response := []types.PredictResponse{
				{ID: 3},
				{ID: 2},
			}
			json.NewEncoder(w).Encode(response)
		}))
		defer server.Close()

		router.HandleFunc("/projects/predict", func(w http.ResponseWriter, r *http.Request) {
			handler.handlePredictProjectsWithURL(w, r, server.URL)
		}).Methods("POST")

		predictRequest := types.PredictRequest{
			UserID: 1,
		}
		jsonBody, _ := json.Marshal(predictRequest)
		req, err := http.NewRequest("POST", "/projects/predict", bytes.NewBuffer(jsonBody))
		req.Header.Set("Content-Type", "application/json")
		assert.NoError(t, err)

		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code, "should return status OK")
		var sortedProjects []types.Project
		err = json.NewDecoder(rr.Body).Decode(&sortedProjects)
		assert.NoError(t, err)

		// Validate the structure of the response
		assert.NotNil(t, sortedProjects, "response should not be nil")
		assert.GreaterOrEqual(t, len(sortedProjects), 1, "response should contain at least one project")

		for _, project := range sortedProjects {
			assert.NotZero(t, project.ID, "project ID should not be zero")
			assert.NotEmpty(t, project.Name, "project Name should not be empty")
		}
	})
}
