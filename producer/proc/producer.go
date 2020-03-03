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

func NewProducer(rabbit *rmq.RabbitMQ) *Procucer {
	return &Procucer{
		mx:      new(sync.Mutex),
		wg:      new(sync.WaitGroup),
		rabbit:  rabbit,
		counter: 0,
	}
}

func (o *Procucer) Start() {
	for i := 0; i < 10; i++ {
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
			Body:      fmt.Sprintf("Process number %v", i),
			Timestamp: time.Now().UTC().Unix(),
		}
		log.Println(body)
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

		time.Sleep(time.Millisecond * 50)
	}
}
