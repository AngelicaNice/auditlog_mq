package server

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	audit "github.com/AngelicaNice/auditlog_mq/pkg/domain"
	amqp "github.com/rabbitmq/amqp091-go"
	log "github.com/sirupsen/logrus"
)

type Audit interface {
	Log(ctx context.Context, item audit.LogItem) error
}

type AuditServer struct {
	amqpConn     *amqp.Connection
	qname        string
	auditService Audit
}

func NewAuditServer(amqpConn *amqp.Connection, qname string, auditService Audit) *AuditServer {
	return &AuditServer{
		amqpConn:     amqpConn,
		qname:        qname,
		auditService: auditService,
	}
}

func (s *AuditServer) Run() {
	fmt.Println("TRY TO START AUDITSERVER")

	ch, err := s.amqpConn.Channel()
	if err != nil {
		log.WithField("rabbitmq", "failed to open a channel").Fatal(err)
	}

	fmt.Println("CHANNEL IS OPENED")

	defer ch.Close()

	q, err := ch.QueueDeclare(
		s.qname,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal("failed to declare a queue")
	}

	fmt.Println("QUEUE IS DECLARED")

	msgs, err := ch.Consume(
		q.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal("failed to register a consumer")
	}

	fmt.Println("WAITING FOR MESSAGES")

	forever := make(chan bool)

	for msg := range msgs {
		fmt.Println("MESSAGE WAS RECEIVED")
		//if item, err := deserialize(msg.Body); err == nil {
		//	fmt.Printf("%+v", item)
		fmt.Printf("%+v", msg)
		//	s.auditService.Log(context.Background(), item)
		//}
	}

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

func deserialize(b []byte) (audit.LogItem, error) {
	var msg audit.LogItem
	buf := bytes.NewBuffer(b)
	decoder := json.NewDecoder(buf)
	err := decoder.Decode(&msg)
	return msg, err
}
