package integration

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type Notification struct {
	ID               int       `json:"id"`
	UserID           int       `json:"user_id"`
	RelatedProjectID int       `json:"related_project_id"`
	RelatedSynergyID *int      `json:"related_synergy_id"`
	Type             string    `json:"type"`
	Title            string    `json:"title"`
	Message          string    `json:"message"`
	Status           bool      `json:"status"`
	CreatedAt        time.Time `json:"created_at"`
}

type NotificationPayload struct {
	UserID           int    `json:"user_id"`
	RelatedProjectID int    `json:"related_project_id"`
	RelatedSynergyID *int   `json:"related_synergy_id"`
	Type             string `json:"type"`
	Title            string `json:"title"`
	Message          string `json:"message"`
	Status           bool   `json:"status"`
	CreatedAt        string `json:"created_at"`
}

type Rating struct {
	ID        int       `json:"id"`
	Date      time.Time `json:"date"`
	Level     string    `json:"level"`
	UserID    int       `json:"user_id"`
	ProjectID int       `json:"project_id"`
}

type RatingPayload struct {
	Date      string `json:"date"`
	Level     string `json:"level"`
	UserID    int    `json:"user_id"`
	ProjectID int    `json:"project_id"`
}

func TestGetNotification(t *testing.T) {
	resp, err := http.Get("http://localhost:8081/ceo/v1/notifications")

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	defer resp.Body.Close()
	var notifications []Notification
	err = json.NewDecoder(resp.Body).Decode(&notifications)
	assert.NoError(t, err)

	assert.NotEmpty(t, notifications)
}

func TestGetNotificationById(t *testing.T) {
	notificationsId := 1

	resp, err := http.Get("http://localhost:8081/ceo/v1/notifications/" + strconv.Itoa(notificationsId))
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	defer resp.Body.Close()

	var notification Notification
	err = json.NewDecoder(resp.Body).Decode(&notification)
	assert.NoError(t, err)
	assert.Equal(t, notificationsId, notification.ID)

	// assert other fields
	assert.NotZero(t, notification.UserID)
	assert.NotEmpty(t, notification.Type)
	assert.NotEmpty(t, notification.Title)
	assert.NotEmpty(t, notification.Message)
	assert.NotZero(t, notification.CreatedAt)
}

func TestCreateNotification(t *testing.T) {
	synergyID := 1

	payload := NotificationPayload{
		UserID:           1,
		RelatedProjectID: 2,
		RelatedSynergyID: &synergyID, // Correção aqui
		Type:             "Solicitação",
		Title:            "Teste de Notificação",
		Message:          "Esta é uma mensagem de teste para a notificação",
		Status:           false,
		CreatedAt:        "2024-06-05", // Correção aqui
	}

	body, err := json.Marshal(payload)
	assert.NoError(t, err)

	resp, err := http.Post("http://localhost:8081/ceo/v1/notifications", "application/json", bytes.NewBuffer(body))
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, resp.StatusCode)
}

func TestUpdateNotification(t *testing.T) {
	synergyID := 1
	notificationId := 1

	payload := NotificationPayload{
		UserID:           1,
		RelatedProjectID: 2,
		RelatedSynergyID: &synergyID, // Correção aqui
		Type:             "Solicitação",
		Title:            "Teste de Notificação",
		Message:          "Esta é uma mensagem de teste para a notificação",
		Status:           false,
		CreatedAt:        "2024-06-05", // Correção aqui
	}

	body, err := json.Marshal(payload)
	assert.NoError(t, err)

	req, err := http.NewRequest(http.MethodPut, "http://localhost:8081/ceo/v1/notifications/"+strconv.Itoa(notificationId), bytes.NewBuffer(body))
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
	assert.Contains(t, responseMessage, "notification updated successfully")

	resp, err = http.Get("http://localhost:8081/ceo/v1/notifications/" + strconv.Itoa(notificationId))
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	defer resp.Body.Close()
	var updatedNotification Notification
	err = json.NewDecoder(resp.Body).Decode(&updatedNotification)
	assert.NoError(t, err)
	assert.Equal(t, payload.Status, updatedNotification.Status)
	assert.Equal(t, payload.Type, updatedNotification.Type)
	assert.Equal(t, payload.Title, updatedNotification.Title)
	assert.Equal(t, payload.UserID, updatedNotification.UserID)
	assert.Equal(t, payload.Message, updatedNotification.Message)
	assert.Equal(t, payload.RelatedProjectID, updatedNotification.RelatedProjectID)
	assert.Equal(t, payload.RelatedSynergyID, updatedNotification.RelatedSynergyID)
}

func TestDeleteNotification(t *testing.T) {
	synergyID := 1
	notificationTitle := "Notificação para Deletar"
	payload := NotificationPayload{
		UserID:           1,
		RelatedProjectID: 2,
		RelatedSynergyID: &synergyID,
		Type:             "Solicitação",
		Title:            notificationTitle,
		Message:          "Esta é uma mensagem de teste para deletar",
		Status:           false,
		CreatedAt:        "2024-06-05",
	}
	body, err := json.Marshal(payload)
	assert.NoError(t, err)

	resp, err := http.Post("http://localhost:8081/ceo/v1/notifications", "application/json", bytes.NewBuffer(body))
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	defer resp.Body.Close()

	resp, err = http.Get("http://localhost:8081/ceo/v1/notifications/title/" + notificationTitle)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	defer resp.Body.Close()
	var notification Notification
	err = json.NewDecoder(resp.Body).Decode(&notification)
	assert.NoError(t, err)

	req, err := http.NewRequest(http.MethodDelete, "http://localhost:8081/ceo/v1/notifications/"+strconv.Itoa(notification.ID), nil)
	assert.NoError(t, err)

	client := &http.Client{}
	resp, err = client.Do(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	defer resp.Body.Close()

	var responseMessage string
	err = json.NewDecoder(resp.Body).Decode(&responseMessage)
	assert.NoError(t, err)
	assert.Contains(t, responseMessage, "notification deleted successfully")
}

func TestGetRatings(t *testing.T) {
	resp, err := http.Get("http://localhost:8081/ceo/v1/ratings")
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	defer resp.Body.Close()
	var ratings []Rating
	err = json.NewDecoder(resp.Body).Decode(&ratings)
	assert.NoError(t, err)

	assert.NotEmpty(t, ratings)
}

func TestGetRatingById(t *testing.T) {
	ratingId := 1

	resp, err := http.Get("http://localhost:8081/ceo/v1/ratings/" + strconv.Itoa(ratingId))
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	defer resp.Body.Close()

	var rating Rating
	err = json.NewDecoder(resp.Body).Decode(&rating)
	assert.NoError(t, err)
	assert.Equal(t, ratingId, rating.ID)

	assert.NotZero(t, rating.Date)
	assert.NotEmpty(t, rating.Level)
	assert.NotZero(t, rating.UserID)
	assert.NotZero(t, rating.ProjectID)
}

func TestCreateRating(t *testing.T) {
	payload := RatingPayload{
		Date:      "2024-06-05",
		Level:     "3",
		UserID:    1,
		ProjectID: 2,
	}

	body, err := json.Marshal(payload)
	assert.NoError(t, err)

	resp, err := http.Post("http://localhost:8081/ceo/v1/ratings", "application/json", bytes.NewBuffer(body))
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	defer resp.Body.Close()
}

func TestUpdateRating(t *testing.T) {
	ratingID := 1
	payload := RatingPayload{
		Date:      "2024-06-10",
		Level:     "4",
		UserID:    1,
		ProjectID: 2,
	}

	body, err := json.Marshal(payload)
	assert.NoError(t, err)

	req, err := http.NewRequest(http.MethodPut, "http://localhost:8081/ceo/v1/ratings/"+strconv.Itoa(ratingID), bytes.NewBuffer(body))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	defer resp.Body.Close()
}

func TestDeleteRating(t *testing.T) {
	payload := RatingPayload{
		Date:      "2024-06-05",
		Level:     "3",
		UserID:    1,
		ProjectID: 2,
	}
	body, err := json.Marshal(payload)
	assert.NoError(t, err)

	resp, err := http.Post("http://localhost:8081/ceo/v1/ratings", "application/json", bytes.NewBuffer(body))
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	defer resp.Body.Close()

	resp, err = http.Get("http://localhost:8081/ceo/v1/ratings?level=3")
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	defer resp.Body.Close()
	var ratings []Rating
	err = json.NewDecoder(resp.Body).Decode(&ratings)
	assert.NoError(t, err)
	assert.NotEmpty(t, ratings)

	createdRating := ratings[0]

	req, err := http.NewRequest(http.MethodDelete, "http://localhost:8081/ceo/v1/ratings/"+strconv.Itoa(createdRating.ID), nil)
	assert.NoError(t, err)

	client := &http.Client{}
	resp, err = client.Do(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	defer resp.Body.Close()

	var responseMessage string
	err = json.NewDecoder(resp.Body).Decode(&responseMessage)
	assert.NoError(t, err)
	assert.Contains(t, responseMessage, "rating deleted successfully")
}

