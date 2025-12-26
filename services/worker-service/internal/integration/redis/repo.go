package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type Repo struct {
	rdb *redis.Client
}

func NewRepo(addr string) *Repo {
	rdb := redis.NewClient(&redis.Options{
		Addr: addr,
	})
	return &Repo{rdb: rdb}
}

func (r *Repo) Close() error {
	return r.rdb.Close()
}

func (r *Repo) IncrTotal(ctx context.Context) error {
	return r.rdb.Incr(ctx, "orders:count").Err()
}

func (r *Repo) SaveLast(ctx context.Context, jsonStr string) error {
	return r.rdb.Set(ctx, "orders:last", jsonStr, 0).Err()
}

func (r *Repo) IncrByUser(ctx context.Context, userID string) error {
	key := "orders:by_user:" + userID
	return r.rdb.Incr(ctx, key).Err()
}
