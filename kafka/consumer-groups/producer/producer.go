package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/IBM/sarama"
)

type Order struct {
	CustomerName string `json:"customer_name"`
	OrderID      string `json:"order_id"`
}

func main() {
	http.HandleFunc("/order", HandlerOrder)

	fmt.Println("server started listening at 8080...")
	http.ListenAndServe(":8080", nil)
}

func HandlerOrder(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed"))
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid request"))
		return
	}

	order := new(Order)
	if err := json.Unmarshal(body, order); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid order request"))
		return
	}

	// validating or other things in order

	orderInBytes, err := json.Marshal(order)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("unable to place order at the movement"))
		return
	}

	if err := PushOrderToQueue("order", orderInBytes); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("unable to process order at the movement"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("order processed successfully"))
}

func PushOrderToQueue(topic string, msg []byte) error {
	brokers := []string{"localhost:9092"}
	producer, err := ConnectToBroker(brokers)
	if err != nil {
		return err
	}

	kafkaMsg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(msg),
	}

	partition, offset, err := producer.SendMessage(kafkaMsg)
	if err != nil {
		return nil
	}

	fmt.Println("message stored in topic ", topic, " partition ", partition, " offset ", offset)
	return nil
}

func ConnectToBroker(brokers []string) (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true

	return sarama.NewSyncProducer(brokers, config)
}
