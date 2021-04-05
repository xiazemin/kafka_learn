package main

import (
	"fmt"

	message "kafka_learn/exp2/proto/kafka/proto"

	"github.com/golang/protobuf/proto"

	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

func main() {

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost",
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		panic(err)
	}

	c.SubscribeTopics([]string{"test", "^aRegex.*[Tt]opic"}, nil)

	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			data := message.Kafka{}
			if msg != nil {
				fmt.Println(proto.Unmarshal(msg.Value, &data))
				fmt.Printf("Consumer error: %v (%v)  %#v\n", err, data, msg.Key)
			}
			fmt.Printf("Message on %#v,%s: %s\n", err, msg.TopicPartition, string(msg.Key))
		} else {
			// The client will automatically try to recover from all errors.
			fmt.Printf("Consumer error: %v (%v) \n", err, msg)
		}
	}

	c.Close()
}
