package main

import (
	"dba-hell/producer/proc"
	"dba-hell/rmq/consts"
	"github.com/streadway/amqp"
	"log"
	"time"

	"dba-hell/rmq"
)

func main() {
	for {
		time.Sleep(time.Millisecond * 5000)

		publisher := rmq.NewRabbitMQ("amqp://rabbitmq:rabbitmq@rabbit1:5672/")
		if err := publisher.InitPublisher(); err != nil {
			log.Println(err)
			continue
		}

		if err := publisher.ExchangeDeclare(consts.PRODUCER_EXCHANGE, amqp.ExchangeDirect); err != nil {
			log.Println(err)
			continue
		}

		publisher.QueueDeclare(consts.STATISTIC_QUEUE)

		if err := publisher.QueueBind(
			consts.STATISTIC_QUEUE,
			consts.ROUTING_KEY,
			consts.PRODUCER_EXCHANGE,
		); err != nil {
			log.Println(err)
			continue
		}

		generator := proc.NewProducer(publisher)
		generator.Start()

	}
}
