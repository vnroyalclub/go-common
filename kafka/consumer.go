/*
   封装kafka comsumer group
*/

package kafka

import (
	"context"
	"fmt"
	"time"

	"github.com/Shopify/sarama"
)

type Handler struct {
	Run  func(msg []byte, args ...interface{})
	Args []interface{}
}

type Consumer struct {
	client  sarama.ConsumerGroup
	brokers []string
	groupId string
}

func NewConsumer(addr []string, groupId string) (consumer Consumer, err error) {

	config := sarama.NewConfig()
	config.Version = sarama.V2_2_0_0
	config.Consumer.Return.Errors = true
	c, err := sarama.NewConsumerGroup(addr, groupId, config)
	if err != nil {
		return
	}
	consumer.client = c
	consumer.groupId = groupId
	consumer.brokers = addr
	return
}

func (c Consumer) Consume(topicHandler map[string]Handler) (err error) {
	if len(topicHandler) == 0 {
		err = fmt.Errorf("topic handler is empty")
		return
	}

	topics := make([]string, 0)

	go func() {
		for err := range c.client.Errors() {
			fmt.Println("consume group:",c.groupId,"brokers:",c.brokers,"err:", err)
		}
	}()

	ctx := context.Background()
	for topic, _ := range topicHandler {
		topics = append(topics, topic)
	}
	for {
		groupHandler := consumeHandler{
			topicHandler: topicHandler,
		}
		err = c.client.Consume(ctx, topics, groupHandler)
		if err != nil {
			fmt.Println("consume group:",c.groupId,"brokers:",c.brokers,"err:", err ,"retry again after 3s")
			<-time.After(time.Second * 3)
		}
	}

	return
}

func (c Consumer) Close() error {
	return c.client.Close()
}

type consumeHandler struct {
	topicHandler map[string]Handler
}

func (consumeHandler) Setup(_ sarama.ConsumerGroupSession) error { return nil }

func (consumeHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }

func (p consumeHandler) ConsumeClaim(session sarama.ConsumerGroupSession,
	claim sarama.ConsumerGroupClaim) error {

	for msg := range claim.Messages() {
		hander := p.topicHandler[msg.Topic]
		if hander.Run != nil {
			hander.Run(msg.Value, hander.Args...)
		}
		session.MarkMessage(msg, "")
	}
	return nil
}
