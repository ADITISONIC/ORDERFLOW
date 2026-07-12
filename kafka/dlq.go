package kafka

import (
	"context"

	kafkago "github.com/segmentio/kafka-go"
)

var DLQWriter = &kafkago.Writer{
	Addr:     kafkago.TCP("localhost:9092"),
	Topic:    "order-created-dlq",
	Balancer: &kafkago.LeastBytes{},
}

func PublishToDLQ(message []byte) error {
	return DLQWriter.WriteMessages(
		context.Background(),
		kafkago.Message{
			Value: message,
		},
	)
}