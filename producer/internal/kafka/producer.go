package kafka

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"97.GO/producer/internal/models"
	"github.com/twmb/franz-go/pkg/kgo"
)

var producer *kgo.Client

func init() {
	producer = setupKafkaProducer()
}

func setupKafkaProducer() *kgo.Client {
	client, err := kgo.NewClient(
		kgo.SeedBrokers("localhost:9092"),
		kgo.WithLogger(kgo.BasicLogger(os.Stdout, kgo.LogLevelInfo, nil)), 
	)
	if err != nil {
		log.Fatalf("Failed to create Kafka producer: %v", err)
	}
	return client
}

func SendOrderToKafka(order models.Order) error {
	orderBytes, err := json.Marshal(order)
	if err != nil {
		return fmt.Errorf("failed to marshal order: %w", err)
	}

	record := &kgo.Record{
		Topic: "order_events",
		Value: orderBytes,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	producer.Produce(ctx, record, func(record *kgo.Record, err error) {
		if err != nil {
			log.Printf("Failed to send message: %v", err)
			return
		}
		fmt.Println("Message sent successfully")
	})

	producer.Flush(ctx)
	return nil
}
