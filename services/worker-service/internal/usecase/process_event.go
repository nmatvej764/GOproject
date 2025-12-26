package usecase

import (
	"context"

	"worker-service/internal/domain"
)

type StatsRepository interface {
	IncrTotal(ctx context.Context) error
	SaveLast(ctx context.Context, jsonStr string) error
	IncrByUser(ctx context.Context, userID string) error
}

type Processor struct {
	repo StatsRepository
}

func NewProcessor(repo StatsRepository) *Processor { 
	return &Processor{repo: repo}
}

func (p *Processor) Process(ctx context.Context, event domain.OrderCreatedEvent, rawJSON string) error {
	if err := p.repo.IncrTotal(ctx); err != nil {
		return err
	}
	if err := p.repo.SaveLast(ctx, rawJSON); err != nil {
		return err
	}
	if err := p.repo.IncrByUser(ctx, event.UserID); err != nil {
		return err
	}
	return nil
}
