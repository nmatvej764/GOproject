package http

import (
	"encoding/json"
	"net/http"

	"api-service/internal/domain"
	"api-service/internal/usecase"
)

type Handler struct {
	uc *usecase.OrderUseCase
}

func NewHandler(uc *usecase.OrderUseCase) *Handler {
	return &Handler{uc: uc}
}

func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}

func (h *Handler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var req domain.CreateOrderRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	orderID, err := h.uc.CreateOrder(r.Context(), req)
	if err != nil {
		http.Error(w, "failed to create order", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(map[string]string{
		"orderId": orderID,
		"status":  "accepted",
	})
}
