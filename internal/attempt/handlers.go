package attempt

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type AttemptHandler struct {
	AttemptService *AttemptService
}

func (h *AttemptHandler) HandleGetAttempts(w http.ResponseWriter, r *http.Request) {

	attempts, err := h.AttemptService.GetAttempts(context.Background())
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

func (h *AttemptHandler) HandleGetAttemptByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(fmt.Sprintf("Error: %s", err.Error()))
		return
	}

	attempt, err := h.AttemptService.GetAttemptByID(context.Background(), int64(id))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(fmt.Sprintf("Error: %s", err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&attempt)
}

func (h *AttemptHandler) HandleGetAttemptByUserID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(fmt.Sprintf("Error: %s", err.Error()))
		return
	}

	attempt, err := h.AttemptService.GetAttemptsByUserID(context.Background(), int64(id))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(fmt.Sprintf("Error: %s", err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&attempt)
}

func (h *AttemptHandler) HandleCreateNewAttempt(w http.ResponseWriter, r *http.Request) {
	var attempt Attempt

	if err := json.NewDecoder(r.Body).Decode(&attempt); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(fmt.Sprintf("Error: %s", err.Error()))
		return
	}

	a, err := h.AttemptService.CreateAttempt(context.Background(), CreateAttemptParams{
		UserID: attempt.UserID,
		QuizID: attempt.QuizID,
		Answer: attempt.Answer,
	})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(fmt.Sprintf("Error: %s", err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(a)
}
