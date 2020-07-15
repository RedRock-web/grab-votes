// @program: grab-votes 
// @author: edte
// @create: 2020-07-15 22:27
package service

import (
	"encoding/json"
	"log"
	"time"

	"grab-votes/model"
)

type Order struct {
	Oid  int
	time int
	Uid  int
	Mid  int
}

var OrderCh = make(chan string, 100000)

var NewOrderQueueName = func() string {
	return "oders"
}

var OrderQueueName = NewOrderQueueName()

func GrabOrder(o Order) {
	q := NewRabbitMQSimple(OrderQueueName)

	b, err := json.Marshal(o)
	if err != nil {
		log.Printf("failed marshal order to json:%s", err)
		return
	}

	q.PublishSimple(string(b))
}

func order() {
	q := NewRabbitMQSimple(OrderQueueName)
	msgs := q.ConsumeSimple()
	var o Order

	for _, msg := range msgs {
		json.Unmarshal([]byte(msg), &o)
		model.AddOrder(model.Order{
			Oid:  0,
			Time: time.Now().Unix(),
			Uid:  0,
			Mid:  0,
		})
		RemoveMovie(o.Mid)
	}
}
