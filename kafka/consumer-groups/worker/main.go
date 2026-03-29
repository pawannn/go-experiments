package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/IBM/sarama"
)

type Order struct {
	CustomerID string `json:"customer_id"`
	OrderID    string `json:"order_id"`
	RetryCount int    `json:"retry_count"`
	NextRetry  int64  `json:"next_retry"`
}

type OrderConsumer struct {
	group        sarama.ConsumerGroup
	errorHandler *ErrorHandler
	MaxRetry     int
}

func main() {
	brokers := []string{"localhost:9092"}
	consumerGroupID := "order-group"

	ordersGroup, err := NewOrderConsumer(brokers, consumerGroupID)
	if err != nil {
		log.Fatal(err)
	}

	topics := []string{"order", ordersGroup.errorHandler.retryTopic}
	err = ordersGroup.Start(context.Background(), topics)

	if err != nil {
		log.Fatal(err)
	}
}

func NewOrderConsumer(brokers []string, consumerGroupID string) (*OrderConsumer, error) {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	group, err := sarama.NewConsumerGroup(brokers, consumerGroupID, config)
	if err != nil {
		return nil, err
	}

	errorHandler, err := NewErrorHandler(brokers, "order")
	if err != nil {
		return nil, err
	}

	return &OrderConsumer{
		group:        group,
		errorHandler: errorHandler,
		MaxRetry:     5,
	}, nil
}

func (o *OrderConsumer) Start(ctx context.Context, topics []string) error {
	orderHandler := &ConsumerGroupHandler{
		Handler: o.Handler,
	}

	for {
		err := o.group.Consume(ctx, topics, orderHandler)
		if err != nil {
			log.Println("unable to consume message: ", err.Error())
			continue
		}

		if ctx.Err() != nil {
			return ctx.Err()
		}
	}
}

func (o *OrderConsumer) Handler(msg *sarama.ConsumerMessage) error {
	order := new(Order)
	if err := json.Unmarshal(msg.Value, order); err != nil {
		return fmt.Errorf("Unable to consume message : %v", err.Error())
	}

	now := time.Now().Unix()
	if order.NextRetry > now {
		fmt.Println("The message has to wait some more time")
		o.errorHandler.PushToRetry(msg.Value)
		return nil
	}

	err := processOrder(order)
	if err != nil {
		order.RetryCount++
		order.NextRetry = time.Now().Add(time.Second * 10).Unix()
		orderInBytes, err := json.Marshal(order)
		if err != nil {
			return err
		}

		if order.RetryCount >= o.MaxRetry {
			fmt.Println("pushing to DLQ")
			o.errorHandler.PushToDLQ(orderInBytes)
			return nil
		}

		fmt.Println("pushing to Retry")
		o.errorHandler.PushToRetry(orderInBytes)
		return nil
	}

	return nil
}

func processOrder(order *Order) error {
	fmt.Printf("recieved orderID : %s from customerID : %s\n", order.OrderID, order.CustomerID)
	return fmt.Errorf("invalid order")
}
