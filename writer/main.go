package main

import (
	"log"
	"time"

	"dba-hell/rmq"
	"dba-hell/rmq/consts"
	"dba-hell/writer/proc"
)

const (
	_RMQ_APP_NAME = "writer"
)

func main() {
	for {
		time.Sleep(time.Millisecond * 5000)

		publisher := rmq.NewRabbitMQ("amqp://rabbitmq:rabbitmq@rabbit1:5672/")
		if err := publisher.InitPublisher(); err != nil {
			log.Println(err)
			continue
		}

		deliveryChan, err := publisher.Consume(consts.STATISTIC_QUEUE, _RMQ_APP_NAME)
		if err != nil {
			log.Println(err)
			continue
		}

		generator := proc.NewWriter(deliveryChan)
		generator.Start()
	}

}
