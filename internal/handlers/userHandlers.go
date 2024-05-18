package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/ovijeet26/golang-rest-api/internal/middlewares"
	"github.com/ovijeet26/golang-rest-api/internal/models"
	"github.com/ovijeet26/golang-rest-api/internal/services"
)

type UserHandler struct {
	Service *services.UserService
}

func NewUserHandler(service *services.UserService) *UserHandler {
	return &UserHandler{service}
}

func (h *UserHandler) Register(router *mux.Router) {
	router.HandleFunc("/users", h.GetAllUsers).Methods("GET")
	router.HandleFunc("/users", h.CreateUser).Methods("POST")
	router.HandleFunc("/login", h.Login).Methods("POST")

	router.Use(middlewares.JWTAuthMiddleware)
}

func (h *UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.Service.GetAllUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(users)
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.Service.CreateUser(user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var credentials models.User
	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	token, err := h.Service.Authenticate(credentials.Username, credentials.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
