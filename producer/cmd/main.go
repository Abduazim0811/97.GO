package main

import (
	"log"
	"net/http"

	handlers "97.GO/producer/internal/handler"
)

func main() {
	http.HandleFunc("/create_order", handlers.CreateOrderHandler)
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
