package main

import (
	"fmt"
	message "kafka_learn/exp2/proto/kafka/proto"

	"github.com/golang/protobuf/proto"

	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

func main() {

	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost"})
	if err != nil {
		panic(err)
	}

	defer p.Close()

	// Delivery report handler for produced messages
	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
				}
			}
		}
	}()

	// Produce messages to topic (asynchronously)
	topic := "test"
	msg := &message.Kafka{
		MessageId:      1234567,
		MessageContent: "test proto kafka",
	}
	data, _ := proto.Marshal(msg)
	p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(data),
	}, nil)

	// Wait for message deliveries before shutting down
	p.Flush(15 * 1000)
}
