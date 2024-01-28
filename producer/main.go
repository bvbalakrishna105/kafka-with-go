package main

import (
	"fmt"
	"log"

	"github.com/Shopify/sarama"
)

func main() {
	// Set the Kafka broker address
	brokerList := []string{"192.168.49.1:9092"}

	// Create a new producer configuration
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true

	// Create a new producer using the configuration
	producer, err := sarama.NewSyncProducer(brokerList, config)
	if err != nil {
		log.Fatalf("Error creating producer: %v", err)
	}

	// Create a new message to be sent
	message := &sarama.ProducerMessage{
		Topic: "test-topic",
		Value: sarama.StringEncoder("Hello, Kafka!"),
	}

	// Send the message to the Kafka topic
	partition, offset, err := producer.SendMessage(message)
	if err != nil {
		log.Fatalf("Failed to send message: %v", err)
	}

	fmt.Printf("Message sent to partition %d at offset %d\n", partition, offset)

	// Close the producer
	producer.Close()
}
