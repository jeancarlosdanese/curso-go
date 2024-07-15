// cmd/consumer/main.go

package main

import (
	"curso-go/events/pkg/rabbitmq"
	"encoding/json"
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	msgs := make(chan amqp.Delivery)

	go rabbitmq.Consumer(ch, msgs, "orders")

	for msg := range msgs {
		var jsonMsg map[string]interface{} // Usando um mapa para simplicidade
		if err := json.Unmarshal(msg.Body, &jsonMsg); err != nil {
			log.Println("Error unmarshaling JSON:", err)
			continue
		}

		prettyJSON, err := json.MarshalIndent(jsonMsg, "", "    ")
		if err != nil {
			log.Println("Error marshaling JSON:", err)
			continue
		}

		fmt.Println(string(prettyJSON))
		msg.Ack(false)
	}
}
