package main

import (
	"context"
	"fmt"
	"log"

	"github.com/twmb/franz-go/pkg/kgo"
)
func createTopicWithPartitions(ctx context.Context)
func main(){
	client, err := kgo.NewClient(
		kgo.SeedBrokers("localhost:9092"),
		kgo.AllowAutoTopicCreation(),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	ctx := context.Background()
	err = client.Ping(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// record := kgo.Record{
	// 	Topic: "franz-topic",
	// 	Value: []byte("Hello, Kafka!!!"),
	// }
	for {
		record := kgo.Record{
		Key: []byte("my-key"),
		Topic:	"franz-topic",
		Value: []byte("Hello, kafka with key"),
	}

	err = client.ProduceSync(ctx, &record).FirstErr()
	if err != nil {
		fmt.Printf("Failed to produce message: %v\n", err)
	}else{
		fmt.Println("Message produced successfully")
	}
	}
	
}