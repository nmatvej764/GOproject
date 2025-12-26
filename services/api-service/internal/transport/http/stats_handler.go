package http

import (
	"encoding/json"
	"net/http"

	"api-service/internal/usecase"
)

type StatsHandler struct {
	uc *usecase.StatsUseCase
}

func NewStatsHandler(uc *usecase.StatsUseCase) *StatsHandler {
	return &StatsHandler{uc: uc}
}

func (h *StatsHandler) GetStats(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("userId")
	if userID == "" {
		userID = "42" 
	}

	stats, err := h.uc.GetStats(r.Context(), userID)
	if err != nil {
		http.Error(w, "failed to get stats", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stats)
}
