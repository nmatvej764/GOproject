package usecase

import (
	"context"
	"encoding/json"
	"strconv"

	"api-service/internal/domain"
)

type StatsRepository interface {
	GetOrdersCount(ctx context.Context) (string, error)
	GetLastOrder(ctx context.Context) (string, error)
	GetOrdersByUser(ctx context.Context, userID string) (string, error)
}

type StatsResponse struct {
	OrdersCount int                       `json:"ordersCount"`
	UserCount   int                       `json:"userCount"`
	LastOrder   *domain.OrderCreatedEvent `json:"lastOrder,omitempty"`
}

type StatsUseCase struct {
	repo StatsRepository
}

func NewStatsUseCase(repo StatsRepository) *StatsUseCase {
	return &StatsUseCase{repo: repo}
}

func (u *StatsUseCase) GetStats(ctx context.Context, userID string) (StatsResponse, error) {
	res := StatsResponse{}

	countStr, err := u.repo.GetOrdersCount(ctx)
	if err == nil {
		res.OrdersCount, _ = strconv.Atoi(countStr)
	}

	userStr, err := u.repo.GetOrdersByUser(ctx, userID)
	if err == nil {
		res.UserCount, _ = strconv.Atoi(userStr)
	}

	lastStr, err := u.repo.GetLastOrder(ctx)
	if err == nil && lastStr != "" {
		var last domain.OrderCreatedEvent
		if json.Unmarshal([]byte(lastStr), &last) == nil {
			res.LastOrder = &last
		}
	}

	return res, nil
}
