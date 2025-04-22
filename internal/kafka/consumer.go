package kafka

import (
	"log"

	"github.com/IBM/sarama"
)

func ConsumeMessages(brokerList []string, topic string) error {
	consumer, err := sarama.NewConsumer(brokerList, nil)
	if err != nil {
		return err
	}
	defer consumer.Close()

	partitionConsumer, err := consumer.ConsumePartition(topic, 0, sarama.OffsetNewest)
	if err != nil {
		return err
	}
	defer partitionConsumer.Close()

	log.Printf("Listening for messages on topic: %s\n", topic)

	for msg := range partitionConsumer.Messages() {
		log.Printf("Received message: %s\n", string(msg.Value))
	}

	return nil
}
