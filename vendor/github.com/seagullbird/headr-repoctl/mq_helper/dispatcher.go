package mq_helper

import (
	"github.com/streadway/amqp"
	"log"
	"encoding/json"
	"github.com/seagullbird/headr-repoctl/config"
)

type Dispatcher interface {
	DispatchMessage(message interface{}) (err error)
}

type AMQPDispatcher struct {
	channel       *amqp.Channel
	queueName     string
	mandatorySend bool
}

func (d *AMQPDispatcher) DispatchMessage(message interface{}) (err error) {
	log.Printf("Dispatching message to queue %s\n", d.queueName)
	body, err := json.Marshal(message)
	if err == nil {
		err = d.channel.Publish(
			"",              // exchange
			d.queueName,     // routing key
			d.mandatorySend, // mandatory
			false,           // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(body),
			})
		if err != nil {
			log.Printf("Failed to dispatch message: %s\n", err)
		}
	} else {
		log.Printf("Failed to marshal message %v (%s)\n", message, err)
	}
	return
}

func NewDispatcher(queueName string) Dispatcher {
	uri := amqp.URI{
		Scheme:   "amqp",
		Host:     config.MQSERVERNAME,
		Port:     5672,
		Username: "user",
		Password: "kQS5MZHEFC",
		Vhost:    "/",
	}
	conn, err := amqp.Dial(uri.String())
	FailOnError(err, "Failed to connect to RabbitMQ")

	ch, err := conn.Channel()
	FailOnError(err, "Failed to open a channel")

	q, err := ch.QueueDeclare(
		queueName, 			// name
		false,		// durable
		false,	// delete when unused
		false,		// exclusive
		false,		// no-wait
		nil,			// arguments
	)
	FailOnError(err, "Failed to declare a queue")
	return &AMQPDispatcher{
		channel: ch,
		queueName: q.Name,
		mandatorySend: false,
	}
}