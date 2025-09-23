package leaderboard

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type LeaderBoardHandler struct {
	LeaderBoardService *LeaderBoardService
}

func (h *LeaderBoardHandler) HandleGetLeaderBoard(w http.ResponseWriter, r *http.Request) {

	l, err := h.LeaderBoardService.GetLeaderBoard(context.Background())
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"Message": err.Error()})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(l)
}

func (h *LeaderBoardHandler) HandleGetLeaderBoardUserID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"Message": err.Error()})
		return
	}

	l, err := h.LeaderBoardService.GetLeaderBoardByUserID(context.Background(), int64(id))
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"Message": "No record found"})
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"Message": err.Error()})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(l)
}
