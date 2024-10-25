package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/segmentio/kafka-go"
)

const (
	kafkaURL = "kafka:9092"
	topic    = "neil_test"
)

func getKafkaReader(kafkaURL, topic string) *kafka.Reader {
	brokers := strings.Split(kafkaURL, ",")
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers: brokers,
		Topic:   topic,
		// GroupID: "test", // 有设置 GroupID 没有设置 like single worker ,GroupID like fanout
		MinBytes: 10e3, // 10KB 最小累積處理的資料量
		MaxBytes: 10e6, // 10MB 最大累積處理的資料量
	})
}

func main() {

	reader := getKafkaReader(kafkaURL, topic)

	defer reader.Close()

	fmt.Println("start consuming ... !!")
	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("message at topic: %v partition: %v offset: %v value: %s\n", m.Topic, m.Partition, m.Offset, string(m.Value))
	}

}
