package main

import (
	"context"

	"worker-service/internal/consumer"
	"worker-service/internal/integration/redis"
	"worker-service/internal/usecase"
)

func main() {
	ctx := context.Background()

	repo := redis.NewRepo("localhost:6379")
	defer repo.Close()

	processor := usecase.NewProcessor(repo)

	kafkaConsumer := consumer.NewKafkaConsumer([]string{"localhost:9092"}, "orders.events", "orders-worker-group", processor)
	defer kafkaConsumer.Close()

	kafkaConsumer.Run(ctx)
}
