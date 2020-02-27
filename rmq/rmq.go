package rmq

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

type RabbitMQ struct {
	uri      string
	producer *amqp.Channel
}

func NewRabbitMQ(amqpURI string) *RabbitMQ {
	return &RabbitMQ{uri: amqpURI}
}

func (o *RabbitMQ) InitPublisher() error {
	log.Printf("dialing %q", o.uri)
	connection, err := amqp.Dial(o.uri)
	if err != nil {
		return fmt.Errorf("dial: %s", err)
	}

	log.Printf("got Connection, getting Channel")
	o.producer, err = connection.Channel()
	if err != nil {
		return fmt.Errorf("channel: %s", err)
	}

	return nil
}

func (o *RabbitMQ) ExchangeDeclare(exchange string, exchangeType string) error {
	if err := o.producer.ExchangeDeclare(
		exchange,     // name
		exchangeType, // type
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // noWait
		nil,          // arguments
	); err != nil {
		return fmt.Errorf("exchange Declare: %s", err)
	}

	return nil
}

func (o *RabbitMQ) QueueDeclare(name string) {
	o.producer.QueueDeclare(name, // name of the queue
		true,                     // durable
		false,                    // delete when usused
		false,                    // exclusive
		false,                    // noWait
		nil,
	)
}

func (o *RabbitMQ) QueueBind(name string, key string, exchange string) error {
	if err := o.producer.QueueBind(
		name,     // name of the queue
		key,      // bindingKey
		exchange, // sourceExchange
		false,    // noWait
		nil,      // arguments
	); err != nil {
		return fmt.Errorf("Queue Bind: %s", err)
	}

	return nil
}

func (o *RabbitMQ) Publish(exchange, routingKey string, body []byte) error {
	if err := o.producer.Publish(
		exchange,   // publish to an exchange
		routingKey, // routing to 0 or more queues
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			Headers:         amqp.Table{},
			ContentType:     "text/plain",
			ContentEncoding: "",
			Body:            body,
			DeliveryMode:    amqp.Transient, // 1=non-persistent, 2=persistent
			Priority:        0,              // 0-9
		},
	); err != nil {
		return fmt.Errorf("Exchange Publish: %s", err)
	}

	return nil
}

func (o *RabbitMQ) Consume(name string, appTag string) (<-chan amqp.Delivery, error) {
	deliveries, err := o.producer.Consume(
		name,   // name
		appTag, // consumerTag,
		false,  // noAck
		false,  // exclusive
		false,  // noLocal
		false,  // noWait
		nil,    // arguments
	)
	if err != nil {
		return nil, fmt.Errorf("queue Consume: %s", err)
	}

	return deliveries, nil
}
