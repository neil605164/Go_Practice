package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

const (
	kafkaURL = "kafka:9092"
	topic    = "neil_test"
)

func newKafkaWriter(kafkaURL, topic string) *kafka.Writer {
	return &kafka.Writer{
		Addr:  kafka.TCP(kafkaURL),
		Topic: topic,
		// GroupID: "test", // 有设置 GroupID 没有设置 like single worker ,GroupID like fanout
		Balancer:    &kafka.LeastBytes{},
		ErrorLogger: kafka.LoggerFunc(errorLog), // 當下無法正常寫入 kafka 時的異常紀錄
		Logger:      kafka.LoggerFunc(logFn),    // 紀錄每一次寫入 kafka 的做紀錄
	}
}

// logFn ====> error writing messages to neil_test (partition 0): kafka.(*Client).Produce: fetch request error: topic partition has no leader (topic="neil_test" partition=0)
func errorLog(msg string, a ...interface{}) {
	log.Printf("logFn ====> "+msg, a...)

	for idx, arg := range a {
		log.Printf("Argument %d: %v", idx, arg)
	}
}

// logFn ----> writing 1 messages to neil_test (partition: 0)
func logFn(msg string, a ...interface{}) {
	// 使用传递的 msg 和 a 参数打印日志
	log.Printf("logFn ----> "+msg, a...)

	// 逐个打印传入的参数
	for idx, arg := range a {
		log.Printf("Argument %d: %v", idx, arg)
	}
}

func main() {

	writer := newKafkaWriter(kafkaURL, topic)
	defer writer.Close()

	fmt.Println("start producing ... !!")
	for i := 0; ; i++ {
		key := fmt.Sprintf("Key-%d", i)
		msg := kafka.Message{
			Value: []byte(fmt.Sprintf(`{"index:" %d, "timestamp": %v}`, i, time.Now())),
		}
		err := writer.WriteMessages(context.Background(), msg)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("produced", key)
		}
		time.Sleep(1 * time.Second)
	}
}
