package consumer

import (
	"context"
	"encoding/json"
	"log"

	"worker-service/internal/domain"
	"worker-service/internal/usecase"

	"github.com/segmentio/kafka-go"
)

type KafkaConsumer struct {
	reader    *kafka.Reader
	processor *usecase.Processor
}

func NewKafkaConsumer(brokers []string, topic, groupID string, processor *usecase.Processor) *KafkaConsumer {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{brokers[0]},
		Topic:    topic,
		GroupID:  groupID,
		MinBytes: 1,
		MaxBytes: 10e6,
	})
	return &KafkaConsumer{
		reader:    reader,
		processor: processor,
	}
}

func (c *KafkaConsumer) Close() error {
	return c.reader.Close()
}

func (c *KafkaConsumer) Run(ctx context.Context) {
	log.Println("worker-service started. Listening topic: orders.events")

	for {
		msg, err := c.reader.ReadMessage(ctx)
		if err != nil {
			log.Println("error reading kafka message:", err)
			continue
		}

		var event domain.OrderCreatedEvent
		if err := json.Unmarshal(msg.Value, &event); err != nil {
			log.Println("bad JSON:", err)
			continue
		}

		err = c.processor.Process(ctx, event, string(msg.Value))
		if err != nil {
			log.Println("failed process event:", err)
			continue
		}

		log.Println("processed order:", event.OrderID, "user:", event.UserID, "total:", event.Total)
	}
}
