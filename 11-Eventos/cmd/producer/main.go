// cmd/producer/main.go

package main

import (
	"curso-go/events/pkg/rabbitmq"
	"encoding/json"
	"strconv"
	"time"
)

func main() {
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	// Simulating a producer
	// The producer sends a message every second
	i := 0
	for {
		i++
		jsonBody, err := json.Marshal(map[string]string{
			"order_id": strconv.Itoa(i),
			"product":  "Camiseta",
			"quantity": "2",
		})
		if err != nil {
			panic(err)
		}

		rabbitmq.Publish(ch, "amq.direct", jsonBody)
		time.Sleep(1 * time.Second)
	}
}
