package main

import (
	"bytes"
	"fmt"
	"log"
	"time"

	"github.com/streadway/amqp"
)

// AggregatorOption Queue 基礎項目限制
type AggregatorOption struct {
	Timer       time.Duration // 倒數計時時間
	UpperLimmit int           // 累積的上限數
}

// receving from queue
func main() {
	arg := AggregatorOption{
		Timer:       time.Second * 60,
		UpperLimmit: 10,
	}

	// 處理 MQ 事項
	if err := MQHandle(arg); err != nil {
		fmt.Println("====>", err)
	}
}

// connectMQ 連線到 MQ
func connectMQ() *amqp.Channel {
	// MQ Connect
	conn, err := amqp.Dial("amqp://guest:guest@127.0.0.1:5672/")
	if err != nil {
		log.Fatalf("%s: %s", "Failed to connect to RabbitMQ", err)
	}
	// defer conn.Close()

	// open a MQ channel
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("%s: %s", "Failed to open a channel", err)
	}

	// defer ch.Close()

	return ch
}

// MQHandle 接收 MQ 資訊 + 執行任務
func MQHandle(arg AggregatorOption) error {
	// connect to MQ
	ch := connectMQ()

	// create a new exchange if exchange is not exist, if exchange not exist, is not allow to bind the queue, you'll geet mistake
	err := ch.ExchangeDeclare(
		"logs",   // name
		"fanout", // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)

	if err != nil {
		log.Fatalf("%s: %s", "Failed to declare a exchange", err)
	}

	// create a new queue
	q, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)

	if err != nil {
		log.Fatalf("%s: %s", "Failed to declare a queue", err)
		return err
	}

	// binding exchange and queue
	err = ch.QueueBind(
		q.Name, // queue name
		"",     // routing key
		"logs", // exchange
		false,
		nil,
	)

	if err != nil {
		log.Fatalf("%s: %s", "Failed to bind a queue", err)
	}

	// receive queue
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)

	if err != nil {
		log.Fatalf("%s: %s", "Failed to consume a message", err)
		return err
	}

	// make a new channel
	eventQueue := make(chan amqp.Delivery, 1)

	go func(arg AggregatorOption) {

		// definition task list
		taskList := []amqp.Delivery{}

		timer := time.NewTimer(arg.Timer)

		defer timer.Stop()

		for {
			select {
			case msg := <-eventQueue:
				// 塞入資料
				taskList = append(taskList, msg)

				fmt.Println("===> I'm in queue, now the data length is", len(taskList))

				if len(taskList) == 1 {
					// 如果是第一筆資料，初始化倒數時間
					timer.Reset(arg.Timer)
				}

				// 如果未滿上限數量，繼續塞入資料並等待
				if len(taskList) < arg.UpperLimmit {
					break
				}

				// 執行任務
				if err := DoingTask(taskList); err != nil {
					// 錯誤處理
					fmt.Println("===>", err.Error())
					return
				}

				// clear array list
				taskList = []amqp.Delivery{}
			case <-timer.C:
				fmt.Println("-------->", time.Now())

				// 任務清單有東西才執行，否則重新計算
				if len(taskList) > 0 {

					// 執行任務
					if err := DoingTask(taskList); err != nil {
						// 錯誤處理
						fmt.Println("===>", err.Error())
						return
					}

					// clear array list
					taskList = []amqp.Delivery{}
				}

				timer.Reset(arg.Timer)
			}
		}
	}(arg)

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	for d := range msgs {
		eventQueue <- d
	}
	return nil
}

// 執行 MQ 要做的任務
func DoingTask(list []amqp.Delivery) error {

	for k := range list {
		fmt.Printf("position %d value is %v", k, string(list[k].Body))
		dotCount := bytes.Count(list[k].Body, []byte("."))
		t := time.Duration(dotCount)
		time.Sleep(t * time.Second)
		fmt.Println("Done")

		// _ = list[k].Ack(false)
	}

	return nil
}
