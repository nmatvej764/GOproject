package domain

import "time"

type OrderCreatedEvent struct {
	EventType string    `json:"eventType"`
	OrderID   string    `json:"orderId"`
	UserID    string    `json:"userId"`
	Total     float64   `json:"total"`
	CreatedAt time.Time `json:"createdAt"`
}
