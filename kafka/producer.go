package kafka

import (
	"context"

	kafkago "github.com/segmentio/kafka-go"
)

var Writer = &kafkago.Writer{
	Addr:     kafkago.TCP("localhost:9092"),
	Topic:    "order-created",
	Balancer: &kafkago.LeastBytes{},
}

func Publish(message []byte) error {
	return Writer.WriteMessages(
		context.Background(),
		kafkago.Message{
			Value: message,
		},
	)
}