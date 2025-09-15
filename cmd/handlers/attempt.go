package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/lmbangel/_novice/internal/attempt"
	"github.com/lmbangel/_novice/internal/db"
)

func HandleGetAttempts(w http.ResponseWriter, r *http.Request) {
	attempts, err := attempt.GetAttempts()
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]string{
				"message": "Attempts  not found, click sign up to create account.",
			})
			return
		}
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(attempts)

}

func HandleGetAttemptByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(fmt.Sprintf("Error: %s", err.Error()))
		return
	}

	attempt, err := attempt.GetAttemptByID(int64(id))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(fmt.Sprintf("Error: %s", err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&attempt)
}

func HandleCreateNewAttempt(w http.ResponseWriter, r *http.Request) {
	var a db.Attempt

	if err := json.NewDecoder(r.Body).Decode(&a); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"Message": err.Error(),
		})
		return
	}

	att, err := attempt.CreateAttempt(a)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"Message": err.Error(),
		})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&att)
}
