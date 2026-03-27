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
}

type OrdersConsumer struct {
	group        sarama.ConsumerGroup
	errorHandler *ErrorHandler
	topics       []string
	maxRetry     int
	delay        time.Duration
}

func main() {
	brokers := []string{"localhost:9092"}
	groupID := "order-group"
	topics := []string{"order"}

	orderConsumer, err := newOrdersConsumer(brokers, groupID, topics, 0)
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
			"order_retry_30s",
			[]string{"order_error_30s"},
			time.Duration(time.Second*30),
		)

		if err != nil {
			log.Fatal(err)
		}

		if err := retryConsumer.start(ctx); err != nil {
			log.Fatal(err)
		}
	}()

	go func() {
		retryConsumer, err := newOrdersConsumer(
			brokers,
			"order_error_60s",
			[]string{"order_error_60s"},
			time.Duration(time.Second*60),
		)

		if err != nil {
			log.Fatal(err)
		}

		if err := retryConsumer.start(ctx); err != nil {
			log.Fatal(err)
		}
	}()

	go func() {
		retryConsumer, err := newOrdersConsumer(
			brokers,
			"order_error_5m",
			[]string{"order_error_5m"},
			time.Duration(time.Minute*5),
		)

		if err != nil {
			log.Fatal(err)
		}

		if err := retryConsumer.start(ctx); err != nil {
			log.Fatal(err)
		}
	}()

	go func() {
		retryConsumer, err := newOrdersConsumer(
			brokers,
			"order_error_10m",
			[]string{"order_error_10m"},
			time.Duration(time.Minute*10),
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

func newOrdersConsumer(brokers []string, groupID string, topics []string, delay time.Duration) (*OrdersConsumer, error) {
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
		delay:        delay,
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
	time.Sleep(time.Duration(oC.delay))

	order := new(Order)

	if err := json.Unmarshal(msg.Value, order); err != nil {
		return err
	}

	fmt.Println("Processing order:", order.OrderID, "retry:", order.RetryCount)

	err := ProcessOrder(order)
	if err != nil {
		order.RetryCount++

		fmt.Println("Failed order:", order.OrderID, "retry:", order.RetryCount)

		msgInBytes, err := json.Marshal(order)
		if err != nil {
			return err
		}

		if order.RetryCount >= oC.maxRetry {
			fmt.Println("Sending to DLQ:", order.OrderID)
			return oC.errorHandler.SendToDLQ(msgInBytes)
		}

		return oC.errorHandler.SendToRetry(msgInBytes, order.RetryCount)
	}

	fmt.Println("Success:", order.OrderID)
	return nil
}

func ProcessOrder(order *Order) error {
	if order.RetryCount == 3 {
		return nil
	}

	return fmt.Errorf("Some error occoured")
}
