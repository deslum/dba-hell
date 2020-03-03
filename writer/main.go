package main

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq"

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

		connStr := "postgres://dba-test:dba-test@postgres/dba_test?sslmode=disable"
		db, err := sql.Open("postgres", connStr)
		if err != nil {
			log.Println(err)
			continue
		}

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

		generator := proc.NewWriter(deliveryChan, db)
		generator.Start()
	}

}
