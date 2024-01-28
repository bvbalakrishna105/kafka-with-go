package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"

	"github.com/Shopify/sarama"
)

func main() {
	// Set the Kafka broker address
	brokerList := []string{"192.168.49.1:9092"}

	// Create a new consumer configuration
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	// Create a new consumer using the configuration
	consumer, err := sarama.NewConsumer(brokerList, config)
	if err != nil {
		log.Fatalf("Error creating consumer: %v", err)
	}

	// Subscribe to the Kafka topic
	topic := "test-topic"
	partitionConsumer, err := consumer.ConsumePartition(topic, 0, sarama.OffsetOldest)
	if err != nil {
		log.Fatalf("Error creating partition consumer: %v", err)
	}

	// Set up a goroutine to handle incoming messages
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case msg := <-partitionConsumer.Messages():
				fmt.Printf("Received message: %s\n", string(msg.Value))
			case err := <-partitionConsumer.Errors():
				log.Printf("Error: %v\n", err)
			}
		}
	}()

	// Wait for a signal to exit
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	<-sig

	// Close the consumer and wait for the goroutine to finish
	partitionConsumer.Close()
	consumer.Close()
	wg.Wait()
}
