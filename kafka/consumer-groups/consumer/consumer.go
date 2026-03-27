package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/IBM/sarama"
)

type Order struct {
	CustomerName string `json:"customer_name"`
	OrderID      string `json:"order_id"`
	RetryCount   int    `json:"retry_count"`
	ProcessAt    int64
}

type OrdersConsumer struct {
	group        sarama.ConsumerGroup
	errorHandler *ErrorHandler
	topics       []string
	maxRetry     int
}

func main() {
	brokers := []string{"localhost:9092"}
	groupID := "order-group"
	topics := []string{"order"}

	orderConsumer, err := newOrdersConsumer(brokers, groupID, topics)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	go func() {
		if err := orderConsumer.start(ctx); err != nil {
			log.Fatal(err)
		}
	}()

	go func() {
		retryConsumer, err := newOrdersConsumer(
			brokers,
			"order-retry-group",
			[]string{"order_error"},
		)
		if err != nil {
			log.Fatal(err)
		}

		if err := retryConsumer.start(ctx); err != nil {
			log.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan
}

func newOrdersConsumer(brokers []string, groupID string, topics []string) (*OrdersConsumer, error) {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	consumerGroup, err := sarama.NewConsumerGroup(brokers, groupID, config)
	if err != nil {
		return nil, err
	}

	errHanlder, err := NewErroHandler(brokers, "order")
	if err != nil {
		return nil, err
	}

	return &OrdersConsumer{
		group:        consumerGroup,
		topics:       topics,
		errorHandler: errHanlder,
		maxRetry:     3,
	}, nil
}

func (oC *OrdersConsumer) start(ctx context.Context) error {
	handlers := &ConsumerGroupHandler{
		Handler: oC.Handler,
	}

	for {
		err := oC.group.Consume(ctx, oC.topics, handlers)
		if err != nil {
			log.Println("Unable to consumer message : ", err.Error())
			continue
		}

		if ctx.Err() != nil {
			return ctx.Err()
		}
	}
}

func (oC *OrdersConsumer) Handler(msg *sarama.ConsumerMessage) error {
	order := new(Order)

	if err := json.Unmarshal(msg.Value, order); err != nil {
		return err
	}

	now := time.Now().Unix()

	if order.ProcessAt > 0 && order.ProcessAt > now {
		fmt.Println("Not ready yet, requeueing:", order.OrderID)

		msgInBytes, _ := json.Marshal(order)

		return oC.errorHandler.SendToRetry(msgInBytes)
	}
	fmt.Println("Processing order:", order.OrderID, "retry:", order.RetryCount)

	err := ProcessOrder(order)
	if err != nil {
		order.RetryCount++

		order.ProcessAt = time.Now().Add(30 * time.Second).Unix()

		fmt.Println("Failed order:", order.OrderID, "retry:", order.RetryCount)

		msgInBytes, err := json.Marshal(order)
		if err != nil {
			return err
		}

		if order.RetryCount > oC.maxRetry {
			fmt.Println("Sending to DLQ:", order.OrderID)
			return oC.errorHandler.SendToDLQ(msgInBytes)
		}

		return oC.errorHandler.SendToRetry(msgInBytes)
	}

	fmt.Println("Success:", order.OrderID)
	return nil
}

func ProcessOrder(order *Order) error {
	return fmt.Errorf("Some error occoured")
}
