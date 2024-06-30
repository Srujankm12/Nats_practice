package main

import (
	"fmt"
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {
	nc, err := nats.Connect(nats.DefaultURL)

	if err != nil {
		log.Fatalf("cant connect to nats: %v", err)
	}
	defer nc.Close()
	for {
		nc.Publish("intros", []byte("hello"))
		fmt.Println("i sent message")
		time.Sleep(1 * time.Second)
	}
}
