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

func (r *Repo) GetOrdersCount(ctx context.Context) (string, error) {
	return r.rdb.Get(ctx, "orders:count").Result()
}

func (r *Repo) GetLastOrder(ctx context.Context) (string, error) {
	return r.rdb.Get(ctx, "orders:last").Result()
}

func (r *Repo) GetOrdersByUser(ctx context.Context, userID string) (string, error) {
	key := "orders:by_user:" + userID
	return r.rdb.Get(ctx, key).Result()
}
