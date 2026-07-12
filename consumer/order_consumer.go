package consumer

import (
	"context"
	"encoding/json"
	"fmt"
	"orderflow/events"
	"orderflow/kafka"
	"orderflow/repositories"
	"time"

	kafkago "github.com/segmentio/kafka-go"
)

var Reader = kafkago.NewReader(kafkago.ReaderConfig{
	Brokers: []string{"localhost:9092"},
	Topic:   "order-created",
	GroupID: "order-consumer-group",
})

func StartConsumer() {

	fmt.Println(" Kafka Consumer Started")

	for {

		message, err := Reader.ReadMessage(context.Background())
		var event events.OrderCreatedEvent
		err = json.Unmarshal(message.Value, &event)

		if err != nil {
			fmt.Println("Invalid message. Sending to DLQ...")

			if err := kafka.PublishToDLQ(message.Value); err != nil {
				fmt.Println("DLQ Publish Failed:", err)
			}
			continue
		}

		fmt.Printf("Processing Order %d\n", event.OrderID)

		time.Sleep(5 * time.Second)

		repositories.UpdateOrderStatus(event.OrderID, "COMPLETED")
		if err != nil {

			fmt.Println("Processing failed. Sending to DLQ...")
			data, _ := json.Marshal(event)
			if err := kafka.PublishToDLQ(data); err != nil {
				fmt.Println("DLQ Publish Failed:", err)
			}
			continue

		}

		fmt.Printf("Order %d Completed\n", event.OrderID)
	}
}
