package user

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/yann-fk-21/todo-platform/config"
	"github.com/yann-fk-21/todo-platform/service/auth"
	"github.com/yann-fk-21/todo-platform/types"
	"github.com/yann-fk-21/todo-platform/utils"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{ store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.loginHandler).Methods("POST")
	router.HandleFunc("/register", h.registerHandler).Methods("POST")
}

func (h *Handler)loginHandler(w http.ResponseWriter, r *http.Request) {
	var userPayload types.LoginUserPayload

	err := utils.ParseJson(r, &userPayload)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}

	err = utils.Validate.Struct(userPayload)
	if err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors))
		return
	}

	u, err := h.store.GetUserByEmail(userPayload.Email)
	if err != nil {
       utils.WriteError(w, http.StatusBadRequest, 
		fmt.Errorf("user with %v email not found", 
		userPayload.Email))
		return
	}

	if !auth.ComparedHashPassword(u.Password, []byte(userPayload.Password)) {
		utils.WriteError(w, http.StatusBadRequest, 
			fmt.Errorf("email or password is invalid"))
			return
	}

	secret := []byte(config.Envs.JWTSecret)
	token, err := auth.CreateJWT(secret, u.ID)

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJson(w, http.StatusOK, map[string]string{"token":token})
}

func(h *Handler) registerHandler(w http.ResponseWriter, r *http.Request) {
	var userPayload types.RegisterUserPayload

	err := utils.ParseJson(r, &userPayload)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}

	err = utils.Validate.Struct(userPayload)
	if err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors))
		return
	}

	_, err = h.store.GetUserByEmail(userPayload.Email)
	if err != nil {
       utils.WriteError(w, http.StatusBadRequest, 
		fmt.Errorf("user with %v email already exists", 
		userPayload.Email))
		return
	}

	hashPassword, err := auth.HashPassword(userPayload.Password)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
	}


	err = h.store.CreateUser(types.User{
		FirstName: userPayload.FirstName,
		LastName: userPayload.LastName,
		Email: userPayload.Email,
		Password: hashPassword,
		CreatedAt: time.Now(),
	})

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
	}

	utils.WriteJson(w, http.StatusCreated, nil)
	
}