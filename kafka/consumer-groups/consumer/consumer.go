package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/IBM/sarama"
)

type Order struct {
	CustomerName string `json:"customer_name"`
	OrderID      string `json:"order_id"`
}

type OrdersConsumer struct {
	group  sarama.ConsumerGroup
	topics []string
}

func main() {
	brokers := []string{"localhost:9092"}
	groupID := "order-group"
	topics := []string{"order"}

	orderConsumer, err := newOrdersConsumer(brokers, groupID, topics)
	if err != nil {
		log.Fatal(err)
	}

	err = orderConsumer.start(context.Background())
	if err != nil {
		log.Fatal(err)
	}

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

	return &OrdersConsumer{
		group:  consumerGroup,
		topics: topics,
	}, nil
}

func (oC *OrdersConsumer) start(ctx context.Context) error {
	handlers := &ConsumerGroupHandler{
		Handler: oC.Handler,
	}

	for {
		err := oC.group.Consume(ctx, oC.topics, handlers)
		if err != nil {
			log.Panicln("Unable to consumer message : ", err.Error())
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

	fmt.Println("Recieved a new order : ", order.OrderID, " by customer : ", order.CustomerName)

	return nil
}
