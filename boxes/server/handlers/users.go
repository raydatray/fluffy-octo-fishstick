package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/raydatray/fluffy-octo-fishstick/boxes/server/ent"
	"github.com/raydatray/fluffy-octo-fishstick/boxes/server/ent/user"
)

type UserHandler struct {
	client *ent.Client
}

func NewUserHandler(client *ent.Client) *UserHandler {
	return &UserHandler{client: client}
}

type CreateUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role,omitempty"`
}

type UserResponse struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

type UpdateUserRequest struct {
	Email    *string `json:"email,omitempty"`
	Password *string `json:"password,omitempty"`
	Role     *string `json:"role,omitempty"`
}

// HandleUsers handles /api/users endpoint
func (h *UserHandler) HandleUsers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.createUser(w, r)
	case http.MethodGet:
		h.getAllUsers(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// HandleUserByID handles /api/users/{id} endpoint
func (h *UserHandler) HandleUserByID(w http.ResponseWriter, r *http.Request) {
	// Extract ID from URL path
	path := strings.TrimPrefix(r.URL.Path, "/api/users/")
	if path == "" {
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(path)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		h.getUserByID(w, r, id)
	case http.MethodPut:
		h.updateUser(w, r, id)
	case http.MethodDelete:
		h.deleteUser(w, r, id)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// createUser creates a new user
func (h *UserHandler) createUser(w http.ResponseWriter, r *http.Request) {
	var req CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Validate required fields
	if req.Email == "" {
		http.Error(w, "Email is required", http.StatusBadRequest)
		return
	}
	if req.Password == "" {
		http.Error(w, "Password is required", http.StatusBadRequest)
		return
	}

	// Set default role if not provided
	if req.Role == "" {
		req.Role = "STUDENT"
	}

	// Validate role
	validRoles := map[string]bool{"PROFESSOR": true, "TA": true, "STUDENT": true}
	if !validRoles[req.Role] {
		http.Error(w, "Invalid role. Must be PROFESSOR, TA, or STUDENT", http.StatusBadRequest)
		return
	}

	// Create user
	u, err := h.client.User.
		Create().
		SetEmail(req.Email).
		SetPassword(req.Password).
		SetRole(user.Role(req.Role)).
		Save(context.Background())

	if err != nil {
		if ent.IsConstraintError(err) {
			http.Error(w, "Email already exists", http.StatusConflict)
			return
		}
		http.Error(w, fmt.Sprintf("Failed to create user: %v", err), http.StatusInternalServerError)
		return
	}

	response := UserResponse{
		ID:    u.ID,
		Email: u.Email,
		Role:  string(u.Role),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

// getAllUsers retrieves all users
func (h *UserHandler) getAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.client.User.Query().All(context.Background())
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to retrieve users: %v", err), http.StatusInternalServerError)
		return
	}

	var response []UserResponse
	for _, u := range users {
		response = append(response, UserResponse{
			ID:    u.ID,
			Email: u.Email,
			Role:  string(u.Role),
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// getUserByID retrieves a user by ID
func (h *UserHandler) getUserByID(w http.ResponseWriter, r *http.Request, id int) {
	u, err := h.client.User.Get(context.Background(), id)
	if err != nil {
		if ent.IsNotFound(err) {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}
		http.Error(w, fmt.Sprintf("Failed to retrieve user: %v", err), http.StatusInternalServerError)
		return
	}

	response := UserResponse{
		ID:    u.ID,
		Email: u.Email,
		Role:  string(u.Role),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// updateUser updates a user by ID
func (h *UserHandler) updateUser(w http.ResponseWriter, r *http.Request, id int) {
	var req UpdateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Check if user exists
	exists, err := h.client.User.Query().Where(user.ID(id)).Exist(context.Background())
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to check user existence: %v", err), http.StatusInternalServerError)
		return
	}
	if !exists {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Build update query
	updateQuery := h.client.User.UpdateOneID(id)

	if req.Email != nil {
		updateQuery = updateQuery.SetEmail(*req.Email)
	}
	if req.Password != nil {
		updateQuery = updateQuery.SetPassword(*req.Password)
	}
	if req.Role != nil {
		// Validate role
		validRoles := map[string]bool{"PROFESSOR": true, "TA": true, "STUDENT": true}
		if !validRoles[*req.Role] {
			http.Error(w, "Invalid role. Must be PROFESSOR, TA, or STUDENT", http.StatusBadRequest)
			return
		}
		updateQuery = updateQuery.SetRole(user.Role(*req.Role))
	}

	// Execute update
	u, err := updateQuery.Save(context.Background())
	if err != nil {
		if ent.IsConstraintError(err) {
			http.Error(w, "Email already exists", http.StatusConflict)
			return
		}
		http.Error(w, fmt.Sprintf("Failed to update user: %v", err), http.StatusInternalServerError)
		return
	}

	response := UserResponse{
		ID:    u.ID,
		Email: u.Email,
		Role:  string(u.Role),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// deleteUser deletes a user by ID
func (h *UserHandler) deleteUser(w http.ResponseWriter, r *http.Request, id int) {
	// Check if user exists
	exists, err := h.client.User.Query().Where(user.ID(id)).Exist(context.Background())
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to check user existence: %v", err), http.StatusInternalServerError)
		return
	}
	if !exists {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Delete user
	err = h.client.User.DeleteOneID(id).Exec(context.Background())
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to delete user: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
