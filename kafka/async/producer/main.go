package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/IBM/sarama"
)

type LogProducer struct {
	producer sarama.AsyncProducer
}

type Log struct {
	LogID      string `json:"log_id"`
	LogDetails string `json:"log_details"`
}

func main() {
	brokers := []string{"localhost:9092"}

	config := sarama.NewConfig()
	config.Producer.Return.Errors = true
	config.Producer.Return.Successes = false
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5

	producer, err := sarama.NewAsyncProducer(brokers, config)
	if err != nil {
		log.Fatal(err)
	}
	defer producer.Close()

	lP := LogProducer{
		producer: producer,
	}

	go func() {
		for err := range lP.producer.Errors() {
			log.Printf("failed to send message: %v\n", err.Err)

			if err.Msg != nil {
				log.Println("failed message : ", err.Msg)
			}
		}
	}()

	http.HandleFunc("/log", lP.handlerLog)
	http.ListenAndServe(":8080", nil)
}

func (o *LogProducer) handlerLog(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("method not allowed"))
		return
	}

	log := new(Log)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid log request"))
		return
	}

	if err := json.Unmarshal(body, log); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid log details"))
		return
	}

	// perform operations on log

	logInBytes, err := json.Marshal(log)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("unable to structure the log"))
		return
	}

	o.PushToQueue("log", logInBytes)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("log processed successfully"))
}

func (o *LogProducer) PushToQueue(topic string, msg []byte) {
	kafkaMsg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(msg),
	}

	select {
	case o.producer.Input() <- kafkaMsg:
	default:
		log.Println("producer buffer full, dropping message")
	}
}
