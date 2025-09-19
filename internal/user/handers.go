package user

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type UserHandler struct {
	UserService *UserService
}

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	u_name := r.URL.Query().Get("username")
	if u_name != "" {
		u, err := h.UserService.GetUserByUsername(context.Background(), u_name)
		if err != nil {
			if err.Error() == "sql: no rows in result set" {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusNotFound)
				json.NewEncoder(w).Encode(map[string]string{"Message": "User not found"})
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"Message": err.Error()})
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(u)
	} else {
		u, err := h.UserService.GetUsers(context.Background())
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"Message": err.Error()})
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(u)
	}
}

func (h *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	u, err := h.UserService.GetUserByID(context.Background(), int64(id))
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]string{"Message": "User not found"})
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"Message": err.Error()})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(u)
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	
}
