package main

import (
	"fmt"
	"log"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	startConsumer()
}

func startConsumer() {
	fmt.Println("starting consumer...")

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":    "localhost:9092",
		"group.id":             "foo",
		"default.topic.config": kafka.ConfigMap{"auto.offset.reset": "smallest"},
	})

	if err != nil {
		log.Fatalf("unable to proceed -> %v", err)
	}

	fmt.Printf("consumer -> %v", c)

	err = c.SubscribeTopics([]string{"test"}, nil)

	startProducer()

	run := true
	for run == true {
		ev := c.Poll(3000)
		switch e := ev.(type) {
		case *kafka.Message:
			fmt.Printf("%% Message on %s:\n%s\n", e.TopicPartition, string(e.Value))
		case kafka.Error:
			fmt.Fprintf(os.Stderr, "%% Error: %v\n", e)
			run = false
		default:
			fmt.Printf("Ignoring -> %v\n", e)
		}
	}
}

func startProducer() {
	fmt.Println("starting producer...")

	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers":    "localhost:9092",
		"client.id":            "foo producer",
		"default.topic.config": kafka.ConfigMap{"acks": "all"},
	})

	if err != nil {
		log.Fatalf("unable to proceed -> %v", err)
	}

	fmt.Printf("producer -> %v", p)
	topic := "test"
	value := "Hello Go!"

	dchan := make(chan kafka.Event, 10000)
	defer close(dchan)

	err = p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(value),
	}, dchan)

	e := <-dchan
	m := e.(*kafka.Message)

	if m.TopicPartition.Error != nil {
		fmt.Printf("Delivery failed: %v\n", m.TopicPartition.Error)
	} else {
		fmt.Printf("Delivered message to topic %s [%d] at offset %v\n",
			*m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
	}
}
