package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"

	"dba-hell/rmq"
	"dba-hell/rmq/consts"
	"dba-hell/writer/proc"
)

const (
	_RMQ_APP_NAME = "writer"
)

func main() {
	connStr := "postgres://dba-test:dba-test@postgres/dba_test?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Println(err)
		return
	}

	publisher := rmq.NewRabbitMQ(consts.RMQ_URI)
	if err := publisher.InitPublisher(); err != nil {
		log.Println(err)
		return
	}

	deliveryChan, err := publisher.Consume(consts.STATISTIC_QUEUE, _RMQ_APP_NAME)
	if err != nil {
		log.Println(err)
		return
	}

	generator := proc.NewWriter(deliveryChan, db)
	generator.Start()

}
