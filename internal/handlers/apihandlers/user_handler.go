package apihandlers

import (
	"net/http"
	"encoding/json"
	"strconv"

	"github.com/qazpalm/gig-agg/internal/store"
	"github.com/qazpalm/gig-agg/internal/auth"
)

type newUserRequest struct {
	Email    	string `json:"email"`
	Username 	string `json:"username"`
	Password 	string `json:"password"`
	IsAdmin   	bool   `json:"is_admin"`
}

type getUsersBody struct {
	Count  int `json:"count"`
	Offset int `json:"offset"`
}

type UserHandler struct {
	store store.UserStore
	auth  *auth.UserAuthManager
}

func NewUserHandler(store store.UserStore, authManager *auth.UserAuthManager) *UserHandler {
	return &UserHandler{store: store, auth: authManager}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	req := newUserRequest{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if req.Email == "" || req.Username == "" || req.Password == "" {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	user, err := h.auth.RegisterUser(req.Email, req.Username, req.Password)
	if err != nil {
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	if req.IsAdmin {
		user.IsAdmin = true
		if err := h.store.UpdateUser(&user); err != nil {
			http.Error(w, "Error updating user", http.StatusInternalServerError)
			return
		}
	}
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	idPath := r.PathValue("id")
	id, err := strconv.Atoi(idPath)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	user, err := h.store.GetUser(id)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}
}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	idPath := r.PathValue("id")
	id, err := strconv.Atoi(idPath)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	req := newUserRequest{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	user, err := h.store.GetUser(id)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	if req.Email != "" {
		user.Email = req.Email
	}
	if req.Username != "" {
		user.Username = req.Username
	}
	if req.Password != "" {
		user.PasswordHash = req.Password
	}

	err = h.store.UpdateUser(user)
	if err != nil {
		http.Error(w, "Error updating user", http.StatusInternalServerError)
		return
	}
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	idPath := r.PathValue("id")
	id, err := strconv.Atoi(idPath)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	err = h.store.DeleteUser(id)
	if err != nil {
		http.Error(w, "Error deleting user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	body := &getUsersBody{}
	err := json.NewDecoder(r.Body).Decode(body)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	users := []*models.User{}
	if body.Count <= 0 && body.Offset < 0 {
		users, err = h.store.GetAllUsers()	
	} else {
		users, err = h.store.GetUsers(body.Count, body.Offset)
	}
	if err != nil {
		http.Error(w, "Error retrieving users", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(users)
	if err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}
}
