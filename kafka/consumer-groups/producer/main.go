package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/IBM/sarama"
)

type Order struct {
	CustomerID string `json:"customer_id"`
	OrderID    string `json:"order_id"`
}

type OrderProducer struct {
	producer sarama.SyncProducer
}

func main() {
	brokers := []string{"localhost:9092"}

	config := sarama.NewConfig()
	config.Producer.Retry.Max = 5
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		log.Fatal(err.Error())
	}

	oP := &OrderProducer{
		producer: producer,
	}

	http.HandleFunc("/order", oP.handlerOrder)
	http.ListenAndServe(":8080", nil)
}

func (o *OrderProducer) handlerOrder(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed"))
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid order request"))
		return
	}

	order := new(Order)
	if err := json.Unmarshal(body, order); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid order details"))
		return
	}

	// TODO: Do operations on order

	err = o.pushOrderToQueue(order)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("unable to process order at the movement"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Order placed successfully"))
}

func (o *OrderProducer) pushOrderToQueue(order *Order) error {
	orderInBytes, err := json.Marshal(order)
	if err != nil {
		return err
	}

	kafkaMsg := &sarama.ProducerMessage{
		Topic: "order",
		Value: sarama.StringEncoder(orderInBytes),
	}

	partition, offset, err := o.producer.SendMessage(kafkaMsg)
	if err != nil {
		return err
	}

	fmt.Printf("order stored in topic(order)/partition(%d)/offset(%d)\n", partition, offset)

	return nil
}
