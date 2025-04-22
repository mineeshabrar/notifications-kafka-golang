package main

import (
	"fmt"
	"os"

	"go-kafka-notifications/internal/kafka"
)

func main() {
	brokers := []string{"localhost:9092"}
	topic := "notifications"
	message := "Hello from Go Kafka Producer!"

	err := kafka.ProduceMessage(brokers, topic, message)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error producing message: %v\n", err)
		os.Exit(1)
	}
}
