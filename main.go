package main

import (
	"fmt"

	datafeeder "github.com/maccam912/gotrader-datafeeder"
	nats "github.com/nats-io/nats.go"
)

func main() {

	nc, _ := nats.Connect(nats.DefaultURL)

	nc.Subscribe("prices", func(m *nats.Msg) {
		fmt.Printf("Received a message: %s\n", string(m.Data))
	})

	datafeeder.RandomDataProducer()
}
