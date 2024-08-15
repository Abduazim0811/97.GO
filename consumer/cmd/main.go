package main

import (
	"log"

	"97.GO/consumer/internal/kafka"
)

func main() {
	consumer := kafka.SetupKafkaConsumer()
	defer consumer.Close()

	log.Println("Starting consumer...")
	kafka.ConsumeMessages(consumer)
}
