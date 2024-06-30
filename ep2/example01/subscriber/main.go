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

	nc.Subscribe("intros", func(msg *nats.Msg) {
		fmt.Println("i got message:%s\n", string(msg.Data))
	})
	time.Sleep(1 * time.Hour)

}
