package user

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Inteli-College/2024-1B-T09-ES06-G03/config"
	"github.com/Inteli-College/2024-1B-T09-ES06-G03/service/auth"
	"github.com/Inteli-College/2024-1B-T09-ES06-G03/types"
	"github.com/Inteli-College/2024-1B-T09-ES06-G03/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleRegister).Methods("POST")
	router.HandleFunc("/user/{id}", h.handleGetCeoByID).Methods("GET")

	authMiddleware := auth.WithJWTAuth(h.handleGetLoggedCeo, h.store)
	router.Handle("/me", authMiddleware).Methods("GET")
}

// handleLogin autentica um usuário
// @Summary Logar usuário
// @Description Autentica o usuário usando email e senha e retorna um JWT
// @Tags user
// @Accept json
// @Produce json
// @Param user body types.LoginUserPayload true "Informações de Login"
// @Success 200 {object} map[string]string "token: JWT token"
// @Failure 400 {object} string "erro de validação ou login falho"
// @Router /login [post]
func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
	// get json payload
	var user types.LoginUserPayload
	if err := utils.ParseJSON(r, &user); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// validate payload
	if err := utils.Validate.Struct(user); err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", err.(validator.ValidationErrors)))
		return
	}

	u, err := h.store.GetUserByEmail(user.Email)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("not found, invalid email or password"))
		return
	}

	if !auth.ComparePasswords(u.Password, []byte(user.Password)) {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("not found, invalid email or password"))
		return
	}

	secret := []byte(config.Envs.JWTSecret)
	token, err := auth.CreateJWT(secret, u.ID)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]string{"token": token})
}

// handleRegister registra um novo usuário
// @Summary Registrar novo usuário
// @Description Registra um novo usuário no sistema e retorna mensagem de sucesso
// @Tags user
// @Accept json
// @Produce json
// @Param user body types.RegisterUserPayload true "Informações de Registro do Usuário"
// @Success 201 {string} string "Usuário criado com sucesso"
// @Failure 400 {object} string "erro de validação ou email já existente"
// @Failure 500 {object} string "erro interno no servidor"
// @Router /register [post]
func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	var user types.RegisterUserPayload
	if err := utils.ParseJSON(r, &user); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(user); err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", err.(validator.ValidationErrors)))
		return
	}

	existingUser, err := h.store.GetUserByEmail(user.Email)
	if existingUser != nil || err != nil {
		if existingUser != nil {
			utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user with email %s already exists", user.Email))
		} else {
			utils.WriteError(w, http.StatusInternalServerError, err)
		}
		return
	}

	hashedPassword, err := auth.HashedPassword(user.Password)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	err = h.store.CreateUser(types.User{
		Name:        user.Name,
		Email:       user.Email,
		Password:    hashedPassword,
		Company:     user.Company,
		Instagram:   user.Instagram,
		Linkedin:    user.Linkedin,
		Photo:       user.Photo,
		Description: user.Description,
	})
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, "User created successfully")
}

// handleGetCeoByID retorna todas as informações de um CEO (exceto senha)
// @Summary Obter informações do CEO por ID
// @Description Retorna todas as informações de um CEO, exceto a senha
// @Tags user
// @Produce json
// @Param id path int true "ID do CEO"
// @Success 200 {object} types.User
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /ceo/{id} [get]
func (h *Handler) handleGetCeoByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid CEO ID"))
		return
	}

	user, err := h.store.GetUserByID(id)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	if user == nil {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("CEO not found"))
		return
	}

	// user.Password = "" // Remover a senha da resposta
	utils.WriteJSON(w, http.StatusOK, user)
}

func (h *Handler) handleGetLoggedCeo(w http.ResponseWriter, r *http.Request) {
    userID, ok := r.Context().Value(auth.UserKey).(int)
    if !ok {
        utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf("unauthorized access"))
        return
    }

    user, err := h.store.GetUserByID(userID)
    if err != nil {
        utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("failed to get user: %v", err))
        return
    }

    if user == nil {
        utils.WriteError(w, http.StatusNotFound, fmt.Errorf("user not found"))
        return
    }

    // Opcionalmente, remover a senha da resposta
    user.Password = ""

    utils.WriteJSON(w, http.StatusOK, user)
}
