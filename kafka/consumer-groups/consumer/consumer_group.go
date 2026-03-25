package main

import (
	"fmt"
	"log"

	"github.com/IBM/sarama"
)

type ConsumerGroupHandler struct {
	Handler func(*sarama.ConsumerMessage) error
}

func (h *ConsumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error {
	return nil
}

func (h *ConsumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error {
	return nil
}

func (h *ConsumerGroupHandler) ConsumeClaim(
	session sarama.ConsumerGroupSession,
	claim sarama.ConsumerGroupClaim,
) error {

	for msg := range claim.Messages() {
		if err := h.Handler(msg); err != nil {
			fmt.Println(err.Error())
			log.Println("Unable to process message : ", string(msg.Key))
			continue
		}

		session.MarkMessage(msg, "")
	}

	return nil
}
