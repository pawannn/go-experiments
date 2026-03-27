package main

import "github.com/IBM/sarama"

// Producer is added here for retry and DLQ management

type ErrorHandler struct {
	producer   sarama.SyncProducer
	errorTopic string
	dlqTopic   string
}

func NewErroHandler(brokers []string, topic string) (*ErrorHandler, error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		return nil, err
	}

	return &ErrorHandler{
		producer:   producer,
		errorTopic: topic + "_error",
		dlqTopic:   topic + "_dlq",
	}, nil
}

func (e *ErrorHandler) SendToRetry(msg []byte) error {
	producerMessage := &sarama.ProducerMessage{
		Topic: e.errorTopic,
		Value: sarama.StringEncoder(msg),
	}

	_, _, err := e.producer.SendMessage(producerMessage)
	return err
}

func (e *ErrorHandler) SendToDLQ(msg []byte) error {
	producerMessage := &sarama.ProducerMessage{
		Topic: e.dlqTopic,
		Value: sarama.StringEncoder(msg),
	}

	_, _, err := e.producer.SendMessage(producerMessage)
	return err
}
