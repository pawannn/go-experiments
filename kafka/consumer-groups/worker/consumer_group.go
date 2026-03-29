package main

import (
	"log"

	"github.com/IBM/sarama"
)

type ConsumerGroupHandler struct {
	Handler func(msg *sarama.ConsumerMessage) error
}

func (c *ConsumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error {
	return nil
}

func (c *ConsumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error {
	return nil
}

func (c *ConsumerGroupHandler) ConsumeClaim(sessions sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		if err := c.Handler(msg); err != nil {
			log.Println("Error consuming message : ", err.Error())
			continue
		}

		sessions.MarkMessage(msg, "")
	}

	return nil
}
