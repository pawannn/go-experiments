package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/IBM/sarama"
)

type Log struct {
	Id      string `json:"id"`
	LogType string `json:"log_type"`
	LogData string `json:"log_data"`
}

func main() {
	brokers := []string{"localhost:9092"}
	logHandler, err := NewLogsHandler(brokers)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/log", logHandler.handleLogs)
	http.ListenAndServe(":8080", nil)
}

type LogsHandler struct {
	producer sarama.SyncProducer
}

func NewLogsHandler(brokers []string) (*LogsHandler, error) {
	config := sarama.NewConfig()
	config.Producer.Retry.Max = 5
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		return nil, err
	}

	return &LogsHandler{producer}, nil
}

func (l *LogsHandler) handleLogs(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("method not allowed"))
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid log request"))
		return
	}

	logs := []Log{}

	if err := json.Unmarshal(body, &logs); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid log data"))
		return
	}

	if err := l.PushToQueue(logs); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("unable to process logs at the movement"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("logs processed successfully"))
}

func (l *LogsHandler) PushToQueue(logs []Log) error {
	topic := "log"
	var messages []*sarama.ProducerMessage

	for _, log := range logs {
		logInBytes, err := json.Marshal(log)
		if err != nil {
			return err
		}

		kafkaMsg := &sarama.ProducerMessage{
			Topic: topic,
			Value: sarama.StringEncoder(logInBytes),
		}

		messages = append(messages, kafkaMsg)
	}

	err := l.producer.SendMessages(messages)
	if err != nil {
		return err
	}

	fmt.Printf("Batch sent %d\n", len(messages))
	return nil
}
