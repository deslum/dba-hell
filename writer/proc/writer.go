package proc

import (
	"database/sql"
	"encoding/json"
	"github.com/streadway/amqp"
	"log"
	"sync"

	"dba-hell/rmq/types"
)

type Writer struct {
	db           *sql.DB
	wg           *sync.WaitGroup
	deliveryChan <-chan amqp.Delivery
	counter      int64
}

func NewWriter(deliveryChan <-chan amqp.Delivery, db *sql.DB) *Writer {
	return &Writer{
		db:           db,
		wg:           new(sync.WaitGroup),
		deliveryChan: deliveryChan,
		counter:      0,
	}
}

func (o *Writer) Start() {
	for i := 0; i < 4; i++ {
		o.wg.Add(1)
		go o.process()
	}

	o.wg.Wait()
}

func (o *Writer) process() {
	defer o.wg.Done()
	sqlStatement := `INSERT INTO "dba_test.threads" (id, body, timestamp) VALUES ($1, $2, to_timestamp($3))`
	txCount := 0
	tx, err := o.db.Begin()
	if err != nil {
		log.Println(err)
	}

	for {
		for msg := range o.deliveryChan {
			var message types.Message
			if err := json.Unmarshal(msg.Body, &message); err != nil {
				log.Println(err)
				continue
			}

			if txCount >= 30 {
				if err := tx.Commit(); err != nil {
					log.Println(err)
					_ = tx.Rollback()
				}

				txCount = 0
				tx, err = o.db.Begin()
				if err != nil {
					log.Println(err)
					continue
				}
			}

			txCount++

			_, err = o.db.Exec(sqlStatement, message.Id, message.Body, message.Timestamp)
			if err != nil {
				log.Println(err)
				continue
			}

			if err = msg.Ack(false); err != nil {
				log.Println(err)
				continue
			}
		}
	}
}
