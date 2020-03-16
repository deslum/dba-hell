package main

import (
	"github.com/streadway/amqp"
	"log"

	"dba-hell/producer/proc"
	"dba-hell/rmq"
	"dba-hell/rmq/consts"
)

func main() {
	publisher := rmq.NewRabbitMQ(consts.RMQ_URI)
	if err := publisher.InitPublisher(); err != nil {
		log.Println(err)
		return
	}

	if err := publisher.ExchangeDeclare(consts.PRODUCER_EXCHANGE, amqp.ExchangeDirect); err != nil {
		log.Println(err)
		return
	}

	publisher.QueueDeclare(consts.STATISTIC_QUEUE)

	if err := publisher.QueueBind(
		consts.STATISTIC_QUEUE,
		consts.ROUTING_KEY,
		consts.PRODUCER_EXCHANGE,
	); err != nil {
		log.Println(err)
		return
	}

	generator := proc.NewProducer(publisher)
	generator.Start()
}
