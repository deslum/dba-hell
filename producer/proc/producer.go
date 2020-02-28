package proc

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"sync/atomic"
	"time"

	"dba-hell/rmq"
	"dba-hell/rmq/consts"
	"dba-hell/rmq/types"
)

type Procucer struct {
	wg      *sync.WaitGroup
	rabbit  *rmq.RabbitMQ
	counter uint64
}

func NewProducer(rabbit *rmq.RabbitMQ) *Procucer {
	return &Procucer{
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
		body := &types.Message{
			Id:        atomic.LoadUint64(&o.counter),
			Body:      fmt.Sprintf("Process number %v", i),
			Timestamp: time.Now().UTC().Unix(),
		}

		log.Println(body)

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

		atomic.AddUint64(&o.counter, 1)
		time.Sleep(time.Millisecond * 50)
	}
}
