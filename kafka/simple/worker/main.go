package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/IBM/sarama"
)

func main() {
	topic := "order"
	orderCount := 0

	worker, err := ConnectConsumer([]string{"localhost:9092"})
	if err != nil {
		log.Fatal(err)
	}

	consumer, err := worker.ConsumePartition(topic, 0, sarama.OffsetOldest)

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGTERM, syscall.SIGINT)
	done := make(chan struct{})

	go func() {
		for {
			select {
			case err := <-consumer.Errors():
				fmt.Println(err.Error())
			case msg := <-consumer.Messages():
				fmt.Printf("Order count %d, topic: %s, message %s", orderCount, topic, string(msg.Value))
			case <-sigchan:
				done <- struct{}{}
			}
		}
	}()

	<-done
	fmt.Println("No of orders processed : ", orderCount)
	if err := worker.Close(); err != nil {
		log.Fatal(err)
	}
}

func ConnectConsumer(brokers []string) (sarama.Consumer, error) {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	return sarama.NewConsumer(brokers, config)
}
