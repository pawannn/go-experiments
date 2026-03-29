package main

import (
	"fmt"

	"github.com/IBM/sarama"
)

type ErrorHandler struct {
	producer   sarama.AsyncProducer
	retryTopic string
	dlqTopic   string
}

func NewErrorHandler(brokers []string, consumerTopic string) (*ErrorHandler, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = false
	config.Producer.Return.Errors = true
	config.Producer.Retry.Max = 3

	producer, err := sarama.NewAsyncProducer(brokers, config)
	if err != nil {
		return nil, err
	}

	go func() {
		for err := range producer.Errors() {
			fmt.Printf("Error in %s topic : %s", err.Msg.Topic, err.Err.Error())
		}
	}()

	return &ErrorHandler{
		producer:   producer,
		retryTopic: consumerTopic + "_retry",
		dlqTopic:   consumerTopic + "_dlq",
	}, nil
}

func (e *ErrorHandler) PushToRetry(msg []byte) {
	kafkaMsg := &sarama.ProducerMessage{
		Topic: e.retryTopic,
		Value: sarama.StringEncoder(msg),
	}

	e.producer.Input() <- kafkaMsg
}

func (e *ErrorHandler) PushToDLQ(msg []byte) {
	kafkaMsg := &sarama.ProducerMessage{
		Topic: e.dlqTopic,
		Value: sarama.StringEncoder(msg),
	}

	e.producer.Input() <- kafkaMsg
}
