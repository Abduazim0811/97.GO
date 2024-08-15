package handlers

import (
	"encoding/json"
	"net/http"

	"97.GO/producer/internal/kafka"
	"97.GO/producer/internal/models"
)

func CreateOrderHandler(w http.ResponseWriter, r *http.Request) {
	var order models.Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	order.Status = "new"

	if err := kafka.SendOrderToKafka(order); err != nil {
		http.Error(w, "Failed to send order", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Order received"))
}
