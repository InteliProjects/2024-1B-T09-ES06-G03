package project

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/Inteli-College/2024-1B-T09-ES06-G03/project/service/auth"
	"github.com/Inteli-College/2024-1B-T09-ES06-G03/project/types"
	"github.com/Inteli-College/2024-1B-T09-ES06-G03/project/utils"
	"github.com/gorilla/mux"
)

// Handler é o handler para os endpoints de project
type Handler struct {
	store types.ProjectStore
}

// NewHandler cria um novo handler de project
func NewHandler(store types.ProjectStore) *Handler {
	return &Handler{store: store}
}

// RegisterRoutes registra as rotas de projects no roteador
func (h *Handler) RegisterRoutes(router *mux.Router) {
    // Protect the GET route for all projects to require JWT
    router.Handle("/projects", auth.JWTMiddleware(http.HandlerFunc(h.handleGetProjects))).Methods(http.MethodGet)

    // Protect the POST route for creating a project, requiring JWT
    router.Handle("/projects", auth.JWTMiddleware(http.HandlerFunc(h.handlePostProject))).Methods(http.MethodPost)

    // Route for getting projects of the logged-in user, already properly protected
    router.Handle("/projects/me", auth.JWTMiddleware(http.HandlerFunc(h.handleGetMyProjects))).Methods("GET")

    // Protect the GET by ID to ensure only authenticated users can access specific project details
    router.Handle("/projects/{id}", auth.JWTMiddleware(http.HandlerFunc(h.handleGetProjectByID))).Methods(http.MethodGet)

    // Protect the PUT for updating projects to require JWT
    router.Handle("/projects/{id}", auth.JWTMiddleware(http.HandlerFunc(h.handleUpdateProject))).Methods(http.MethodPut)

    // Protect the DELETE to require JWT
    router.Handle("/projects/{id}", auth.JWTMiddleware(http.HandlerFunc(h.handleDeleteProject))).Methods(http.MethodDelete)

    // Assuming you want to keep the GET by name public, no JWT middleware here
    router.HandleFunc("/projects/name/{name}", h.handleGetProjectsByName).Methods(http.MethodGet)

    // Protect the GET projects by CEO ID route to restrict access to authenticated users
    router.Handle("/projects/ceo/{ceoId}", auth.JWTMiddleware(http.HandlerFunc(h.handleGetProjectsByCeoID))).Methods(http.MethodGet)

    router.HandleFunc("/projects/predict", h.handlePredictProjects).Methods(http.MethodPost) // Rota modelo de recomendação
}


// handleGetProjects manipula as solicitações de obtenção de todos os projetos
// @Summary Obtém todos os projetos
// @Description Retorna uma lista de todos os projetos
// @Tags projects
// @Produce json
// @Success 200 {array} types.ProjectDetails
// @Failure 500 {object} utils.ErrorResponse
// @Router /projects [get]
func (h *Handler) handleGetProjects(w http.ResponseWriter, r *http.Request) {
	projects, err := h.store.GetProjects()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, projects)
}

// handleGetProjectByID manipula as solicitações de obtenção de um projeto por ID
// @Summary Obtém um projeto por ID
// @Description Retorna um projeto com base no ID fornecido
// @Tags projects
// @Produce json
// @Param id path int true "ID do projeto"
// @Success 200 {object} types.Project
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /projects/{id} [get]
func (h *Handler) handleGetProjectByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid project ID"))
		return
	}

	project, err := h.store.GetProjectByID(id)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	if project == nil {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("project not found"))
		return
	}

	utils.WriteJSON(w, http.StatusOK, project)
}

// handlePostProject manipula as solicitações de criação de um novo projeto
// @Summary Cria um novo projeto
// @Description Cria um novo projeto com base nos dados fornecidos
// @Tags projects
// @Accept json
// @Produce json
// @Param project body types.ProjectPayload true "Dados do projeto"
// @Success 201 {string} string "project created successfully"
// @Failure 400 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /projects [post]
func (h *Handler) handlePostProject(w http.ResponseWriter, r *http.Request) {
	var payload types.ProjectPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", err))
		return
	}

	project := types.Project{
		Name:          payload.Name,
		Description:   payload.Description,
		Status:        payload.Status,
		UserID:        payload.UserID,
		CategoryID:    payload.CategoryID,
		SubcategoryID: payload.SubcategoryID,
		CreatedAt:     types.CustomDate{Time: time.Now()},
		UpdatedAt:     types.CustomDate{Time: time.Now()},
		Photo:         payload.Photo,
		Local:         payload.Local, // Adicionado aqui
	}

	if err := h.store.CreateProject(&project); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, "project created successfully")
}

// handleUpdateProject manipula as solicitações de atualização de um projeto
// @Summary Atualiza um projeto existente
// @Description Atualiza um projeto existente com base nos dados fornecidos
// @Tags projects
// @Accept json
// @Produce json
// @Param id path int true "ID do projeto"
// @Param project body types.ProjectPayload true "Dados do projeto"
// @Success 200 {string} string "project updated successfully"
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /projects/{id} [put]
func (h *Handler) handleUpdateProject(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid project ID"))
		return
	}

	var payload types.ProjectPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", err))
		return
	}

	project := types.Project{
		ID:            id,
		Name:          payload.Name,
		Description:   payload.Description,
		Status:        payload.Status,
		UserID:        payload.UserID,
		CategoryID:    payload.CategoryID,
		SubcategoryID: payload.SubcategoryID,
		UpdatedAt:     types.CustomDate{Time: time.Now()},
		Photo:         payload.Photo,
		Local:         payload.Local, // Adicionado aqui
	}

	if err := h.store.UpdateProject(id, &project); err != nil {
		if err == sql.ErrNoRows {
			utils.WriteError(w, http.StatusNotFound, fmt.Errorf("project not found"))
		} else {
			utils.WriteError(w, http.StatusInternalServerError, err)
		}
		return
	}

	utils.WriteJSON(w, http.StatusOK, "project updated successfully")
}

// handleDeleteProject manipula as solicitações de exclusão de um projeto
// @Summary Exclui um projeto existente
// @Description Exclui um projeto existente com base no ID fornecido
// @Tags projects
// @Produce json
// @Param id path int true "ID do projeto"
// @Success 200 {string} string "project deleted successfully"
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /projects/{id} [delete]
func (h *Handler) handleDeleteProject(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid project ID"))
		return
	}

	if err := h.store.DeleteProject(id); err != nil {
		if err == sql.ErrNoRows {
			utils.WriteError(w, http.StatusNotFound, fmt.Errorf("project not found"))
		} else {
			utils.WriteError(w, http.StatusInternalServerError, err)
		}
		return
	}

	utils.WriteJSON(w, http.StatusOK, "project deleted successfully")
}

func (h *Handler) handleGetProjectsByName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	projects, err := h.store.GetProjectsByName(name)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	if len(projects) == 0 {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("no projects found with the name %s", name))
		return
	}

	utils.WriteJSON(w, http.StatusOK, projects)
}

// handleGetProjectsByCeoID manipula as solicitações de obtenção de projetos por CEO ID
// @Summary Obtém todos os projetos de um CEO
// @Description Retorna uma lista de todos os projetos de um CEO com base no ID fornecido
// @Tags projects
// @Produce json
// @Param ceoId path int true "ID do CEO"
// @Success 200 {array} types.Project
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /projects/ceo/{ceoId} [get]
func (h *Handler) handleGetProjectsByCeoID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ceoId, err := strconv.Atoi(vars["ceoId"])
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid CEO ID"))
		return
	}

	projects, err := h.store.GetProjectsByCeoID(ceoId)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	if projects == nil {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("no projects found for the given CEO ID"))
		return
	}

	utils.WriteJSON(w, http.StatusOK, projects)
}

// handlePredictProjects manipula as solicitações para a rota de previsão de projetos
// @Summary Obtém previsões de projetos
// @Description Envia os dados do usuário para um endpoint externo e retorna a previsão
// @Tags projects
// @Accept json
// @Produce json
// @Param predictRequest body types.PredictRequest true "Dados para previsão"
// @Success 200 {array} types.Project
// @Failure 400 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /projects/predict [post]
func (h *Handler) handlePredictProjectsWithURL(w http.ResponseWriter, r *http.Request, predictURL string) {
	var predictRequest types.PredictRequest
	if err := json.NewDecoder(r.Body).Decode(&predictRequest); err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid request payload"))
		return
	}

	projects, err := h.store.GetProjects()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("failed to get projects: %v", err))
		return
	}

	evaluations, err := h.store.GetEvaluations()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("failed to get evaluations: %v", err))
		return
	}

	requestData := types.PredictRequest{
		UserID: predictRequest.UserID,
		Data:   evaluations,
	}

	requestBody, err := json.Marshal(requestData)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("failed to marshal request body: %v", err))
		return
	}

	req, err := http.NewRequest(http.MethodPost, predictURL, bytes.NewBuffer(requestBody))
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("failed to create request: %v", err))
		return
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("failed to send request: %v", err))
		return
	}
	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("failed to read response body: %v", err))
		return
	}

	var predictResponses []types.PredictResponse
	if err := json.Unmarshal(responseBody, &predictResponses); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("failed to unmarshal response body: %v", err))
		return
	}

	projectMap := make(map[int]types.ProjectDetails)
	for _, project := range projects {
		projectMap[project.ID] = project
	}

	var sortedProjects []types.ProjectDetails
	for _, predictResponse := range predictResponses {
		if project, found := projectMap[predictResponse.ID]; found {
			sortedProjects = append(sortedProjects, project)
		}
	}

	utils.WriteJSON(w, http.StatusOK, sortedProjects)
}

func (h *Handler) handlePredictProjects(w http.ResponseWriter, r *http.Request) {
	h.handlePredictProjectsWithURL(w, r, "http://localhost:5000/predict")
}

// handleGetMyProjects retorna projetos associados ao usuário autenticado
func (h *Handler) handleGetMyProjects(w http.ResponseWriter, r *http.Request) {
    userID, ok := r.Context().Value("userID").(int)
    if !ok {
        http.Error(w, "User ID not found", http.StatusForbidden)
        return
    }

    projects, err := h.store.GetProjectsByCeoID(userID)  // Supondo que você tenha essa função no seu ProjectStore
    if err != nil {
        utils.WriteError(w, http.StatusInternalServerError, err)
        return
    }

    utils.WriteJSON(w, http.StatusOK, projects)
}
// handleGetInterestedAndSynergies manipula as solicitações de obtenção dos interessados e sinergias de um projeto específico
// @Summary Obtém os interessados e sinergias de um projeto específico
// @Description Retorna uma lista de usuários interessados e sinergias associadas a um projeto específico
// @Tags projects
// @Produce json
// @Param id path int true "ID do projeto"
// @Success 200 {object} types.InterestedAndSynergiesResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /projects/{id}/interested [get]
func (h *Handler) handleGetInterestedAndSynergies(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	projectID, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid project ID"))
		return
	}

	response, err := h.store.GetInterestedAndSynergiesByProjectID(projectID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	if response == nil {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("no data found for the given project ID"))
		return
	}

	utils.WriteJSON(w, http.StatusOK, response)
}
