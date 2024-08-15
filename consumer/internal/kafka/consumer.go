package kafka

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"97.GO/consumer/internal/models"
	"github.com/twmb/franz-go/pkg/kgo"
)

// SetupKafkaConsumer initializes Kafka consumer
func SetupKafkaConsumer() *kgo.Client {
	client, err := kgo.NewClient(
		kgo.SeedBrokers("localhost:9092"),
		kgo.ConsumeTopics("order_events"),
		kgo.WithLogger(kgo.BasicLogger(os.Stdout, kgo.LogLevelInfo, nil)), // `os.Stdout` io.Writer interfeysiga mos keladi
	)
	if err != nil {
		log.Fatalf("Failed to create Kafka consumer: %v", err)
	}
	return client
}

func ConsumeMessages(consumer *kgo.Client) {
	ctx := context.Background()

	for {
		fetches := consumer.PollFetches(ctx)
		if fetches.IsClientClosed() {
			break
		}

		fetches.EachRecord(func(record *kgo.Record) {
			processMessage(record.Value)
		})
	}
}

func processMessage(value []byte) {
	var order models.Order
	if err := json.Unmarshal(value, &order); err != nil {
		log.Printf("Failed to unmarshal message: %v", err)
		return
	}

	if order.Status == "new" {
		fmt.Printf("New Order processed: ID=%s, Amount=%d\n", order.ID, order.Amount)
	}
}
