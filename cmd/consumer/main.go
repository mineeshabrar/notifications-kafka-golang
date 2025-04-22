package main

import (
	"fmt"
	"os"

	"go-kafka-notifications/internal/kafka"
)

func main() {
	brokers := []string{"localhost:9092"}
	topic := "notifications"

	err := kafka.ConsumeMessages(brokers, topic)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error consuming messages: %v\n", err)
		os.Exit(1)
	}
}
