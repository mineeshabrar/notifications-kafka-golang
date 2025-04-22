package kafka

import (
	"log"

	"github.com/IBM/sarama"
)

func ProduceMessage(brokerList []string, topic string, message string) error {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer(brokerList, config)
	if err != nil {
		return err
	}
	defer producer.Close()

	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	}

	_, _, err = producer.SendMessage(msg)
	if err != nil {
		return err
	}

	log.Printf("Message sent to topic %s: %s\n", topic, message)
	return nil
}
