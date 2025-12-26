package kafka

import (
	"context"

	kafkago "github.com/segmentio/kafka-go"
)

type Publisher struct {
	writer *kafkago.Writer
}

func NewPublisher(brokers []string, topic string) *Publisher {
	writer := kafkago.NewWriter(kafkago.WriterConfig{
		Brokers:  brokers,
		Topic:    topic,
		Balancer: &kafkago.LeastBytes{},
	})

	return &Publisher{writer: writer}
}

func (p *Publisher) Publish(ctx context.Context, key string, value []byte) error {
	return p.writer.WriteMessages(ctx, kafkago.Message{
		Key:   []byte(key),
		Value: value,
	})
}

func (p *Publisher) Close() error {
	return p.writer.Close()
}
