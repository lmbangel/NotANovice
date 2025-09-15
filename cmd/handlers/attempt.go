package handlers

import (
	"encoding/json"
	"net/http"

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
	attempt := db.Attempt{}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(attempt)
}
