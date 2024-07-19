package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Inteli-College/2024-1B-T09-ES06-G03/types"
	"github.com/gorilla/mux"
)

func TestUserServiceHandlers(t *testing.T) {
	userStore := &mockUserStore{
		users: make(map[string]*types.User),
	}
	handler := NewHandler(userStore)

	t.Run("should fail if the user payload is invalid", func(t *testing.T) {
		payload := types.RegisterUserPayload{
			Name:      "user",
			Email:     "invalid",
			Password:  "short",
			Company:   "120jonsdjdv",
			Instagram: "notcheguers",
			Linkedin:  "notcheguers2",
		}
		marshalled, _ := json.Marshal(payload)

		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/register", handler.handleRegister)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("expected status code %d, got %d", http.StatusBadRequest, rr.Code)
		}
	})

	t.Run("should correctly register the user", func(t *testing.T) {
		payload := types.RegisterUserPayload{
			Name:        "Jane Doe",
			Email:       "jane.doe@example.com",
			Password:    "password123",
			Company:     "Doe Enterprises",
			Instagram:   "janedoeig",
			Linkedin:    "janedoein",
			Photo:       "janedoe.jpg",
			Description: "Software Engineer",
		}
		marshalled, _ := json.Marshal(payload)

		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/register", handler.handleRegister)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusCreated {
			t.Errorf("expected status code %d, got %d", http.StatusCreated, rr.Code)
		}
	})
}

type mockUserStore struct {
	users map[string]*types.User
}

func (m *mockUserStore) GetUserByEmail(email string) (*types.User, error) {
	if user, ok := m.users[email]; ok {
		return user, nil
	}
	return nil, nil // Simula nenhum usuário encontrado corretamente
}

func (m *mockUserStore) GetUserByID(id int) (*types.User, error) {
	// Implementação opcional se necessário
	return nil, nil
}

func (m *mockUserStore) CreateUser(user types.User) error {
	if _, exists := m.users[user.Email]; exists {
		return fmt.Errorf("user already exists")
	}
	m.users[user.Email] = &user
	return nil
}
