package usecase

import (
	"context"
	"encoding/json"
	"time"

	"api-service/internal/domain"

	"github.com/google/uuid"
)

type EventPublisher interface {
	Publish(ctx context.Context, key string, value []byte) error
}

type OrderUseCase struct {
	publisher EventPublisher
}

func NewOrderUseCase(publisher EventPublisher) *OrderUseCase {
	return &OrderUseCase{publisher: publisher}
}

func (u *OrderUseCase) CreateOrder(ctx context.Context, req domain.CreateOrderRequest) (string, error) {
	orderID := uuid.New().String()

	event := domain.OrderCreatedEvent{
		EventType: "OrderCreated",
		OrderID:   orderID,
		UserID:    req.UserID,
		Total:     req.Total,
		CreatedAt: time.Now(),
	}

	data, err := json.Marshal(event)
	if err != nil {
		return "", err
	}

	err = u.publisher.Publish(ctx, orderID, data)
	if err != nil {
		return "", err
	}

	return orderID, nil
}
