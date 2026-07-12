package kafka

import (
	"context"
	"fmt"

	kafkago "github.com/segmentio/kafka-go"
)

var Writer = &kafkago.Writer{
	Addr:     kafkago.TCP("localhost:9092"),
	Topic:    "order-created",
	Balancer: &kafkago.LeastBytes{},
}

func Publish(message []byte) error {

	fmt.Println("Kafka Broker: localhost:9092")
	fmt.Println("Publishing:", string(message))

	err := Writer.WriteMessages(
		context.Background(),
		kafkago.Message{
			Value: message,
		},
	)

	if err != nil {
		fmt.Println("Kafka Error:", err)
		return err
	}

	fmt.Println("Message Published Successfully")

	return nil
}