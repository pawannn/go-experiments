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
	CustomerName string `json:"customer_name"`
	OrderType    string `json:"order_type"`
}

func main() {
	http.HandleFunc("/order", placeOrder)
	http.ListenAndServe(":8080", nil)
}

func placeOrder(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("incorrect method for the given enpoint"))
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("unable to read body"))
		return
	}

	order := new(Order)
	if err := json.Unmarshal(body, order); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("unable to de-serialize the payload"))
		return
	}

	orderInBytes, err := json.Marshal(order)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Unable to serialize the payload"))
		return
	}

	if err := pushOrderToQueue("orders", orderInBytes); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("unable to push data to queue"))
		return
	}

	var response map[string]string = map[string]string{
		"message": "order placed successfully",
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func pushOrderToQueue(topic string, msg []byte) error {
	brokers := []string{"localhost:9092"}
	// create connection to broker
	producer, err := ConnectBrokers(brokers)
	if err != nil {
		return fmt.Errorf("Unable to connect to brokers : %s", err.Error())
	}
	defer producer.Close()

	producerMessage := &sarama.ProducerMessage{
		Topic: "order",
		Value: sarama.StringEncoder(msg),
	}

	partition, offset, err := producer.SendMessage(producerMessage)
	if err != nil {
		return err
	}

	log.Printf("message stored in topic(%s)/partition(%d)/offset(%d)", topic, partition, offset)
	return nil
}

func ConnectBrokers(brokers []string) (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5

	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		return nil, err
	}

	return producer, nil
}
