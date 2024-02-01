package mq

import (
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

func CreateMQConnection(url string) (*amqp.Connection, error) {
	conn, err := amqp.Dial(url)
	if err == nil {
		fmt.Println("RABBITMQ CONNECTED")
	}
	return conn, err
}
