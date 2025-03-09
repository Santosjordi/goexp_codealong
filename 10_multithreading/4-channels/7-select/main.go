package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type Message struct {
	id  int64
	Msg string
}

func main() {
	c1 := make(chan Message)
	c2 := make(chan Message)
	var i int64 = 0

	go func() {
		for {
			time.Sleep(time.Second * 2)
			atomic.AddInt64(&i, 1)
			msg := Message{id: i, Msg: "Hello from RabbitMQ"}
			c1 <- msg
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 1)
			atomic.AddInt64(&i, 1)
			msg := Message{id: i, Msg: "Hello from Kafka"}
			c2 <- msg
		}
	}()
	// for i := 0; i < 30; i++ {
	for {
		select {
		case msg := <-c1: //rabbitmq
			fmt.Printf("received from rabbitMq: %v - id: %d \n", msg.Msg, msg.id)

		case msg := <-c2: // kafka
			fmt.Printf("received from Kafka: %v - id: %d \n", msg.Msg, msg.id)

		case <-time.After(time.Second * 3):
			println("timeout")

			// default:
			// 	println("no activity")
		}
	}
}
