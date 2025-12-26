package main

import (
	"log"
	"net/http"

	"api-service/internal/integration/kafka"
	redisrepo "api-service/internal/integration/redis"
	transport "api-service/internal/transport/http"
	"api-service/internal/usecase"

	"github.com/go-chi/chi/v5"
)

func main() {
	// Kafka publisher
	publisher := kafka.NewPublisher([]string{"localhost:9092"}, "orders.events")
	defer publisher.Close()

	// Redis repo
	redisRepo := redisrepo.NewRepo("localhost:6379")
	defer redisRepo.Close()

	// Usecases
	orderUC := usecase.NewOrderUseCase(publisher)
	statsUC := usecase.NewStatsUseCase(redisRepo)

	// Handlers
	orderHandler := transport.NewHandler(orderUC)
	statsHandler := transport.NewStatsHandler(statsUC)

	// Router
	r := chi.NewRouter()
	r.Get("/health", orderHandler.Health)
	r.Post("/orders", orderHandler.CreateOrder)
	r.Get("/stats", statsHandler.GetStats)

	log.Println("api-service started on :8081")
	log.Fatal(http.ListenAndServe(":8081", r))
}
