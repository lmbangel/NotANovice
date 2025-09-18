package quiz

import (
	"context"
	"encoding/json"
	"net/http"
)

type QuizHandler struct {
	QuizService *QuizService
}

func (h *QuizHandler) HandleGetQuizes(w http.ResponseWriter, r *http.Request) {

	quizes, err := h.QuizService.GetQuizes(context.Background())
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"Message": err.Error()})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(quizes)
}
