package proc

import (
	"encoding/json"
	"log"
	"sync"

	"github.com/streadway/amqp"

	"dba-hell/rmq/types"
)

type Writer struct {
	wg           *sync.WaitGroup
	deliveryChan <-chan amqp.Delivery
	counter      int64
}

func NewWriter(deliveryChan <-chan amqp.Delivery) *Writer {
	return &Writer{
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
	for {
		for msg := range o.deliveryChan {
			var message types.Message
			err := json.Unmarshal(msg.Body, &message)
			if err != nil {
				log.Println(err)
				continue
			}

			err = msg.Ack(false)
			if err != nil {
				log.Println(err)
				continue
			}
		}
	}
}
