package integration

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Project struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	Status        string `json:"status"`
	UserID        int    `json:"user_id"`
	SubcategoryID int    `json:"subcategory_id"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
	Photo         string `json:"photo"`
}

type ProjectPayload struct {
	Name          string `json:"name"`
	Description   string `json:"description"`
	Status        string `json:"status"`
	UserID        int    `json:"user_id"`
	SubcategoryID int    `json:"subcategory_id"`
	Photo         string `json:"photo"`
}

type Synergy struct {
	ID              int    `json:"id"`
	SourceProjectID int    `json:"source_project_id"`
	TargetProjectID int    `json:"target_project_id"`
	Status          string `json:"status"`
	Type            string `json:"type"`
	Description     string `json:"description"`
}

type SynergyPayload struct {
	SourceProjectID int    `json:"source_project_id"`
	TargetProjectID int    `json:"target_project_id"`
	Status          string `json:"status"`
	Type            string `json:"type"`
	Description     string `json:"description"`
}

type Update struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Date        string `json:"date"`
	CreatedAt   string `json:"created_at"`
	SynergyID   int    `json:"synergy_id"`
}

type UpdatePayload struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Date        string `json:"date"`
	SynergyID   int    `json:"synergy_id"`
}

func TestGetProjects(t *testing.T) {
	resp, err := http.Get("http://localhost:8082/projects/v1/projects")
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	defer resp.Body.Close()
	var projects []Project
	err = json.NewDecoder(resp.Body).Decode(&projects)
	assert.NoError(t, err)

	assert.NotEmpty(t, projects)
}

func TestGetProjectByID(t *testing.T) {
	// Define o ID do projeto que será buscado
	projectID := 1

	// Recupera o projeto pelo ID
	resp, err := http.Get("http://localhost:8082/projects/v1/projects/" + strconv.Itoa(projectID))
	assert.NoError(t, err)

	defer resp.Body.Close()
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var project Project
	err = json.NewDecoder(resp.Body).Decode(&project)
	assert.NoError(t, err)
	assert.Equal(t, projectID, project.ID)

	// Adicione assertivas para outros campos, se necessário
	assert.NotEmpty(t, project.Name)
	assert.NotEmpty(t, project.Description)
	assert.NotEmpty(t, project.Status)
	assert.NotZero(t, project.UserID)
	assert.NotZero(t, project.SubcategoryID)
	assert.NotEmpty(t, project.Photo)
}

func TestGetProjectsByCeoID(t *testing.T) {
	ceoID := 1 

	resp, err := http.Get("http://localhost:8082/projects/v1/projects/ceo/" + strconv.Itoa(ceoID))
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	defer resp.Body.Close()
	var projects []Project
	err = json.NewDecoder(resp.Body).Decode(&projects)
	assert.NoError(t, err)

	assert.NotEmpty(t, projects)
	for _, project := range projects {
		assert.Equal(t, ceoID, project.UserID)
	}
}

func TestCreateProject(t *testing.T) {
	payload := ProjectPayload{
		Name:          "Test Project",
		Description:   "This is a test project",
		Status:        "Planejamento",
		UserID:        1,
		SubcategoryID: 1,
		Photo:         "test_photo_url",
	}

	body, err := json.Marshal(payload)
	assert.NoError(t, err)

	resp, err := http.Post("http://localhost:8082/projects/v1/projects", "application/json", bytes.NewBuffer(body))
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	defer resp.Body.Close()

	// Ajuste para lidar com a resposta como uma string simples
	var responseMessage string
	err = json.NewDecoder(resp.Body).Decode(&responseMessage)
	assert.NoError(t, err)
	assert.Contains(t, responseMessage, "project created successfully")
}

func TestUpdateProject(t *testing.T) {
	// Define o ID do projeto que será atualizado
	projectID := 1

	// Dados atualizados do projeto
	payload := ProjectPayload{
		Name:          "Updated Project",
		Description:   "Updated Description",
		Status:        "Desenvolvimento",
		UserID:        1,
		SubcategoryID: 1,
		Photo:         "updated_photo_url",
	}

	body, err := json.Marshal(payload)
	assert.NoError(t, err)

	req, err := http.NewRequest(http.MethodPut, "http://localhost:8082/projects/v1/projects/"+strconv.Itoa(projectID), bytes.NewBuffer(body))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	defer resp.Body.Close()

	// Verifica se a resposta contém a mensagem de sucesso
	var responseMessage string
	err = json.NewDecoder(resp.Body).Decode(&responseMessage)
	assert.NoError(t, err)
	assert.Contains(t, responseMessage, "project updated successfully")

	// Recupera o projeto atualizado e verifica se as mudanças foram aplicadas
	resp, err = http.Get("http://localhost:8082/projects/v1/projects/" + strconv.Itoa(projectID))
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	defer resp.Body.Close()
	var updatedProject Project
	err = json.NewDecoder(resp.Body).Decode(&updatedProject)
	assert.NoError(t, err)
	assert.Equal(t, payload.Name, updatedProject.Name)
	assert.Equal(t, payload.Description, updatedProject.Description)
	assert.Equal(t, payload.Status, updatedProject.Status)
	assert.Equal(t, payload.Photo, updatedProject.Photo)
}

func TestDeleteProject(t *testing.T) {
	// Criar um novo projeto para garantir que ele exista
	projectName := "Project to Delete"
	payload := ProjectPayload{
		Name:          projectName,
		Description:   "Description of Project to Delete",
		Status:        "Planejamento",
		UserID:        1,
		SubcategoryID: 1,
		Photo:         "photo_to_delete",
	}
	body, err := json.Marshal(payload)
	assert.NoError(t, err)

	resp, err := http.Post("http://localhost:8082/projects/v1/projects", "application/json", bytes.NewBuffer(body))
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	defer resp.Body.Close()

	// Buscar o projeto pelo nome para obter o ID
	resp, err = http.Get("http://localhost:8082/projects/v1/projects/name/" + projectName)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	defer resp.Body.Close()
	var projects []Project
	err = json.NewDecoder(resp.Body).Decode(&projects)
	assert.NoError(t, err)
	assert.NotEmpty(t, projects)

	createdProject := projects[0]

	// Deletar o projeto pelo ID
	req, err := http.NewRequest(http.MethodDelete, "http://localhost:8082/projects/v1/projects/"+strconv.Itoa(createdProject.ID), nil)
	assert.NoError(t, err)

	client := &http.Client{}
	resp, err = client.Do(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	defer resp.Body.Close()

	var responseMessage string
	err = json.NewDecoder(resp.Body).Decode(&responseMessage)
	assert.NoError(t, err)
	assert.Contains(t, responseMessage, "project deleted successfully")
}

func TestGetSynergies(t *testing.T) {
	resp, err := http.Get("http://localhost:8082/projects/v1/synergies")
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	defer resp.Body.Close()
	var synergies []Synergy
	err = json.NewDecoder(resp.Body).Decode(&synergies)
	assert.NoError(t, err)

	assert.NotEmpty(t, synergies)
}

func TestGetSynergyByID(t *testing.T) {
	synergyID := 1

	resp, err := http.Get("http://localhost:8082/projects/v1/synergies/" + strconv.Itoa(synergyID))
	assert.NoError(t, err)

	defer resp.Body.Close()
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var synergy Synergy
	err = json.NewDecoder(resp.Body).Decode(&synergy)
	assert.NoError(t, err)
	assert.Equal(t, synergyID, synergy.ID)

	assert.NotEmpty(t, synergy.Description)
	assert.NotEmpty(t, synergy.Status)
	assert.NotZero(t, synergy.SourceProjectID)
	assert.NotZero(t, synergy.TargetProjectID)
	assert.NotEmpty(t, synergy.Type)
}

func TestCreateSynergy(t *testing.T) {
	payload := SynergyPayload{
		SourceProjectID: 1,
		TargetProjectID: 2,
		Status:          "Solicitado",
		Type:            "Aprendizagem",
		Description:     "This is a test synergy",
	}

	body, err := json.Marshal(payload)
	assert.NoError(t, err)

	resp, err := http.Post("http://localhost:8082/projects/v1/synergies", "application/json", bytes.NewBuffer(body))
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	defer resp.Body.Close()

	var responseMessage string
	err = json.NewDecoder(resp.Body).Decode(&responseMessage)
	assert.NoError(t, err)
	assert.Contains(t, responseMessage, "synergy created successfully")
}

func TestUpdateSynergy(t *testing.T) {
	synergyID := 1

	payload := SynergyPayload{
		SourceProjectID: 1,
		TargetProjectID: 2,
		Status:          "Solicitado",
		Type:            "Aprendizagem",
		Description:     "This is a test synergy",
	}

	body, err := json.Marshal(payload)
	assert.NoError(t, err)

	req, err := http.NewRequest(http.MethodPut, "http://localhost:8082/projects/v1/synergies/"+strconv.Itoa(synergyID), bytes.NewBuffer(body))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	defer resp.Body.Close()

	var responseMessage string
	err = json.NewDecoder(resp.Body).Decode(&responseMessage)
	assert.NoError(t, err)
	assert.Contains(t, responseMessage, "synergy updated successfully")

	resp, err = http.Get("http://localhost:8082/projects/v1/synergies/" + strconv.Itoa(synergyID))
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	defer resp.Body.Close()
	var updatedProject Synergy
	err = json.NewDecoder(resp.Body).Decode(&updatedProject)
	assert.NoError(t, err)
	assert.Equal(t, payload.Description, updatedProject.Description)
	assert.Equal(t, payload.Status, updatedProject.Status)
	assert.Equal(t, payload.Type, updatedProject.Type)
	assert.Equal(t, payload.SourceProjectID, updatedProject.SourceProjectID)
	assert.Equal(t, payload.TargetProjectID, updatedProject.TargetProjectID)
}

func TestDeleteSynergy(t *testing.T) {
	synergyDescription := "Synergy to Delete"
	payload := SynergyPayload{
		SourceProjectID: 1,
		TargetProjectID: 2,
		Status:          "Solicitado",
		Type:            "Aprendizagem",
		Description:     synergyDescription,
	}
	body, err := json.Marshal(payload)
	assert.NoError(t, err)

	resp, err := http.Post("http://localhost:8082/projects/v1/synergies", "application/json", bytes.NewBuffer(body))
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	defer resp.Body.Close()

	resp, err = http.Get("http://localhost:8082/projects/v1/synergies/description/" + synergyDescription)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	defer resp.Body.Close()
	var synergies []Synergy
	err = json.NewDecoder(resp.Body).Decode(&synergies)
	assert.NoError(t, err)
	assert.NotEmpty(t, synergies)

	createdSynergy := synergies[0]

	req, err := http.NewRequest(http.MethodDelete, "http://localhost:8082/projects/v1/synergies/"+strconv.Itoa(createdSynergy.ID), nil)
	assert.NoError(t, err)

	client := &http.Client{}
	resp, err = client.Do(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	defer resp.Body.Close()
}

func TestGetUpdates(t *testing.T) {
	resp, err := http.Get("http://localhost:8082/projects/v1/updates")
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	defer resp.Body.Close()
	var updates []Update
	err = json.NewDecoder(resp.Body).Decode(&updates)
	assert.NoError(t, err)

	assert.NotEmpty(t, updates)
}

func TestGetUpdateByID(t *testing.T) {
	updateID := 1

	resp, err := http.Get("http://localhost:8082/projects/v1/updates/" + strconv.Itoa(updateID))
	assert.NoError(t, err)

	defer resp.Body.Close()
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var update Update
	err = json.NewDecoder(resp.Body).Decode(&update)
	assert.NoError(t, err)
	assert.Equal(t, updateID, update.ID)

	assert.NotEmpty(t, update.Title)
	assert.NotEmpty(t, update.Description)
	assert.NotEmpty(t, update.Date)
	assert.NotZero(t, update.SynergyID)
}

func TestCreateUpdate(t *testing.T) {
	payload := UpdatePayload{
		Title:       "Test Update",
		Description: "This is a test update",
		Date:        "2024-06-04",
		SynergyID:   1,
	}

	body, err := json.Marshal(payload)
	assert.NoError(t, err)

	resp, err := http.Post("http://localhost:8082/projects/v1/updates", "application/json", bytes.NewBuffer(body))
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	defer resp.Body.Close()

	var responseMessage string
	err = json.NewDecoder(resp.Body).Decode(&responseMessage)
	assert.NoError(t, err)
	assert.Contains(t, responseMessage, "update created successfully")
}

func TestUpdateUpdate(t *testing.T) {
	// Criar uma nova atualização para garantir que ela exista
	payload := UpdatePayload{
		Title:       "Test Update",
		Description: "This is a test update",
		Date:        "2024-06-05",
		SynergyID:   1,
	}

	body, err := json.Marshal(payload)
	assert.NoError(t, err)

	resp, err := http.Post("http://localhost:8082/projects/v1/updates", "application/json", bytes.NewBuffer(body))
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	defer resp.Body.Close()

	// Buscar a atualização pelo título para obter o ID
	resp, err = http.Get("http://localhost:8082/projects/v1/updates/title/" + payload.Title)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	defer resp.Body.Close()
	var updates []Update
	err = json.NewDecoder(resp.Body).Decode(&updates)
	assert.NoError(t, err)
	assert.NotEmpty(t, updates)

	createdUpdate := updates[0]

	// Atualizar a atualização criada
	updatedPayload := UpdatePayload{
		Title:       "Updated Test Update",
		Description: "This is an updated test update",
		Date:        "2024-06-06",
		SynergyID:   1,
	}

	updatedBody, err := json.Marshal(updatedPayload)
	assert.NoError(t, err)

	req, err := http.NewRequest(http.MethodPut, "http://localhost:8082/projects/v1/updates/"+strconv.Itoa(createdUpdate.ID), bytes.NewBuffer(updatedBody))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err = client.Do(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	defer resp.Body.Close()

	var responseMessage string
	err = json.NewDecoder(resp.Body).Decode(&responseMessage)
	assert.NoError(t, err)
	assert.Contains(t, responseMessage, "update updated successfully")

	// Recuperar a atualização e verificar as mudanças
	resp, err = http.Get("http://localhost:8082/projects/v1/updates/" + strconv.Itoa(createdUpdate.ID))
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	defer resp.Body.Close()
	var updatedUpdate Update
	err = json.NewDecoder(resp.Body).Decode(&updatedUpdate)
	assert.NoError(t, err)

	assert.Equal(t, updatedPayload.Title, updatedUpdate.Title)
	assert.Equal(t, updatedPayload.Description, updatedUpdate.Description)
	assert.Equal(t, updatedPayload.Date+"T00:00:00Z", updatedUpdate.Date)
	assert.Equal(t, updatedPayload.SynergyID, updatedUpdate.SynergyID)
}
