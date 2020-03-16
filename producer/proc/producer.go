package proc

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"time"

	"dba-hell/rmq"
	"dba-hell/rmq/consts"
	"dba-hell/rmq/types"
)

type Procucer struct {
	mx      *sync.Mutex
	wg      *sync.WaitGroup
	rabbit  *rmq.RabbitMQ
	counter uint64
}

// NewProducer init Producer struct
func NewProducer(rabbit *rmq.RabbitMQ) *Procucer {
	return &Procucer{
		mx:      new(sync.Mutex),
		wg:      new(sync.WaitGroup),
		rabbit:  rabbit,
		counter: 0,
	}
}

// Start function start process generate message and send to RabbitMQ
func (o *Procucer) Start() {
	for i := 0; i < 4; i++ {
		o.wg.Add(1)
		go o.process(i)
	}

	o.wg.Wait()
}

func (o *Procucer) process(i int) {
	defer o.wg.Done()
	for {
		o.mx.Lock()
		o.counter++
		body := &types.Message{
			Id:        o.counter,
			Name:      "TODO: Create name",
			Number:    i,
			Body:      fmt.Sprintf("Process number %v", i),
			Timestamp: time.Now().UTC().Unix(),
		}
		o.mx.Unlock()

		b, err := json.Marshal(body)
		if err != nil {
			log.Println(err)
			continue
		}

		err = o.rabbit.Publish(consts.PRODUCER_EXCHANGE, consts.ROUTING_KEY, b)
		if err != nil {
			log.Println(err)
			continue
		}

		time.Sleep(time.Millisecond * 10)
	}
}
